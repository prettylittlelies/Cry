package gateway

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/prettylittlelies/cry/gateway/codec"
	"github.com/prettylittlelies/cry/transport/websocket"
)

type session struct {
	g           *Gateway
	conn        *websocket.Conn
	ctx         context.Context
	interval    time.Duration
	acked       bool
	lastBeat    time.Time
	established bool
	closeCode   int
}

func (s *session) write(op Opcode, d any) error {
	msgOp, b, err := codec.Encode(s.g.cfg.Encoding, outbound{Op: op, D: d})
	if err != nil {
		return err
	}
	return s.conn.Write(msgOp, b)
}

func (s *session) heartbeat() error {
	g := s.g
	g.mu.Lock()
	payload := heartbeat{Seq: g.seq, QoS: g.qos.consume()}
	g.mu.Unlock()
	s.lastBeat = time.Now()
	return s.write(OpQoSHeartbeat, payload)
}

func (s *session) loop(frames <-chan frame) error {
	s.closeCode = 1000
	helloTimer := time.NewTimer(helloTimeout)
	defer helloTimer.Stop()
	var (
		first  *time.Timer
		ticker *time.Ticker
		firstC <-chan time.Time
		tickC  <-chan time.Time
	)
	stop := func() {
		if first != nil {
			first.Stop()
		}
		if ticker != nil {
			ticker.Stop()
		}
	}
	defer stop()
	for {
		select {
		case <-s.ctx.Done():
			return s.ctx.Err()
		case <-helloTimer.C:
			s.g.clearResumeURL()
			return &reconnect{cause: errors.New("gateway: did not receive HELLO in time")}
		case <-firstC:
			firstC = nil
			ticker = time.NewTicker(s.interval)
			tickC = ticker.C
			s.acked = true
			if err := s.beatOrTimeout(); err != nil {
				return err
			}
		case <-tickC:
			if err := s.beatOrTimeout(); err != nil {
				return err
			}
		case <-s.g.wake:
			if s.established {
				if err := s.heartbeat(); err != nil {
					return s.lost(err)
				}
			}
		case o := <-s.g.send:
			if err := s.write(o.Op, o.D); err != nil {
				return s.lost(err)
			}
		case f, ok := <-frames:
			if !ok {
				return s.lost(websocket.ErrClosed)
			}
			if f.err != nil {
				return s.lost(f.err)
			}
			if f.p.S != nil {
				s.g.mu.Lock()
				s.g.seq = *f.p.S
				s.g.mu.Unlock()
			}
			switch f.p.Op {
			case OpHello:
				helloTimer.Stop()
				var h hello
				if err := json.Unmarshal(f.p.D, &h); err != nil || h.HeartbeatInterval <= 0 {
					return &reconnect{cause: errors.New("gateway: invalid HELLO")}
				}
				s.interval = time.Duration(h.HeartbeatInterval) * time.Millisecond
				stop()
				first = time.NewTimer(jitter(s.interval))
				firstC = first.C
				if err := s.resumeOrIdentify(); err != nil {
					return s.lost(err)
				}
			case OpHeartbeat:
				if err := s.heartbeat(); err != nil {
					return s.lost(err)
				}
				if ticker != nil {
					ticker.Reset(s.interval)
				}
			case OpHeartbeatAck:
				s.acked = true
				s.g.mu.Lock()
				s.g.lastAck = time.Now()
				s.g.mu.Unlock()
			case OpReconnect:
				s.closeCode = 4000
				return &reconnect{immediate: true, established: s.established, cause: errors.New("gateway: server requested reconnect")}
			case OpInvalidSession:
				var resumable bool
				json.Unmarshal(f.p.D, &resumable)
				var err error
				if resumable {
					err = s.resumeOrIdentify()
				} else {
					err = s.identify()
				}
				if err != nil {
					return s.lost(err)
				}
			case OpDispatch:
				s.dispatch(f.p)
			}
			if s.interval > 0 && ticker != nil && time.Since(s.lastBeat) > s.interval+heartbeatGrace {
				if err := s.heartbeat(); err != nil {
					return s.lost(err)
				}
			}
		}
	}
}

func (s *session) beatOrTimeout() error {
	if !s.acked {
		s.closeCode = 4000
		return &reconnect{established: s.established, cause: errors.New("gateway: heartbeat ack timeout")}
	}
	s.acked = false
	if err := s.heartbeat(); err != nil {
		return s.lost(err)
	}
	return nil
}

func (s *session) lost(err error) error {
	var ce *websocket.CloseError
	if errors.As(err, &ce) && CloseCode(ce.Code) == CloseAuthenticationFailed {
		return ErrAuthenticationFailed
	}
	return &reconnect{established: s.established, cause: err}
}

func (s *session) resumeOrIdentify() error {
	g := s.g
	g.mu.Lock()
	canResume := g.sessionID != "" && (g.lastAck.IsZero() || time.Since(g.lastAck) <= resumeWindow)
	r := resume{Token: g.cfg.Token, SessionID: g.sessionID, Seq: g.seq}
	g.mu.Unlock()
	var err error
	if canResume {
		err = s.write(OpResume, r)
	} else {
		err = s.identify()
	}
	g.mu.Lock()
	g.lastAck = time.Now()
	g.mu.Unlock()
	return err
}

func (s *session) identify() error {
	g := s.g
	g.mu.Lock()
	g.seq, g.sessionID = 0, ""
	fast := g.identified == 0 && g.cfg.FastConnect
	g.identified++
	reasons := ""
	if !fast {
		reasons = connectReasons(g.reasons)
		g.reasons = nil
	}
	active := g.qos.active()
	derived := g.derivedQoS
	presence := g.presence
	g.mu.Unlock()
	props := g.cfg.Session.Properties()
	if fast {
		return s.write(OpIdentify, fastIdentify{
			Token:        g.cfg.Token,
			Capabilities: g.cfg.Capabilities,
			Properties:   fastConnectProperties{Desktop: fastConnectDesktop(props), IsFastConnect: true, InstallationID: g.cfg.InstallationID},
			ClientState:  freshClientState(),
		})
	}
	return s.write(OpIdentify, identify{
		Token:        g.cfg.Token,
		Capabilities: g.cfg.Capabilities,
		Properties:   identifyProperties{Desktop: props, GatewayConnectReasons: reasons, InstallationID: g.cfg.InstallationID},
		Presence:     presence,
		Compress:     false,
		ClientState:  freshClientState(),
		QoSToken:     qosToken(active, derived),
	})
}

func (s *session) dispatch(p inbound) {
	g := s.g
	switch p.T {
	case "READY":
		var r ready
		json.Unmarshal(p.D, &r)
		g.mu.Lock()
		g.sessionID = r.SessionID
		g.resumeURL = trimSlash(r.ResumeGatewayURL)
		if r.QoSToken != "" {
			g.derivedQoS = r.QoSToken
		}
		g.connected = true
		g.mu.Unlock()
		s.established = true
	case "READY_SUPPLEMENTAL", "RESUMED":
		g.mu.Lock()
		g.connected = true
		g.mu.Unlock()
		s.established = true
	}
	seq := int64(0)
	if p.S != nil {
		seq = *p.S
	}
	g.cfg.Handler(Event{Type: p.T, Seq: seq, Data: p.D})
	if p.T == "READY" || p.T == "RESUMED" {
		s.reportTimeSpent()
	}
}

func (s *session) reportTimeSpent() {
	g := s.g
	g.mu.Lock()
	if !g.qos.active() {
		g.mu.Unlock()
		return
	}
	g.timeSpent.touch(time.Now())
	ts := g.timeSpent
	g.timeSpent.sent = ts.id
	g.mu.Unlock()
	if ts.sent == ts.id {
		return
	}
	launch := g.cfg.Session.Properties().ClientLaunchID
	if err := s.write(OpUpdateTimeSpentSessionID, timeSpentSession{InitializationTimestamp: ts.created.UnixMilli(), SessionID: ts.id, ClientLaunchID: launch}); err == nil {
		s.heartbeat()
	}
}
