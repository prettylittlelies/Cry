package voice

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/prettylittlelies/cry/transport/websocket"
)

var ErrHeartbeatTimeout = errors.New("voice: heartbeat not acknowledged")

type handler interface {
	onJSON(op Opcode, d json.RawMessage) error
	onBinary(op uint8, payload []byte) error
}

type socket struct {
	conn     *websocket.Conn
	seq      atomic.Int64
	acked    atomic.Bool
	writeMu  sync.Mutex
	interval chan time.Duration
}

func dialSocket(ctx context.Context, endpoint string, cfg websocket.Config) (*socket, error) {
	url := fmt.Sprintf("wss://%s/?v=%d", endpoint, GatewayVersion)
	conn, err := websocket.Dial(ctx, url, cfg)
	if err != nil {
		return nil, err
	}
	s := &socket{conn: conn, interval: make(chan time.Duration, 1)}
	s.seq.Store(-1)
	return s, nil
}

func (s *socket) send(op Opcode, d any) error {
	b, err := json.Marshal(outbound{Op: op, D: d})
	if err != nil {
		return err
	}
	s.writeMu.Lock()
	defer s.writeMu.Unlock()
	return s.conn.Write(websocket.Text, b)
}

func (s *socket) sendBinary(op uint8, payload []byte) error {
	s.writeMu.Lock()
	defer s.writeMu.Unlock()
	return s.conn.Write(websocket.Binary, append([]byte{op}, payload...))
}

func (s *socket) close() { s.conn.Close(1000, "") }

func (s *socket) read(h handler) error {
	for {
		kind, data, err := s.conn.Read()
		if err != nil {
			return err
		}
		if kind == websocket.Binary {
			if len(data) < 3 {
				continue
			}
			s.seq.Store(int64(binary.BigEndian.Uint16(data)))
			if err := h.onBinary(data[2], data[3:]); err != nil {
				return err
			}
			continue
		}
		var p packet
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
		if p.Seq != nil {
			s.seq.Store(int64(*p.Seq))
		}
		switch p.Op {
		case OpHello:
			var hl hello
			if err := json.Unmarshal(p.D, &hl); err != nil {
				return err
			}
			s.interval <- time.Duration(hl.Interval * float64(time.Millisecond))
		case OpHeartbeatAck:
			s.acked.Store(true)
		case OpHeartbeat:
			if err := s.beat(); err != nil {
				return err
			}
		default:
			if err := h.onJSON(p.Op, p.D); err != nil {
				return err
			}
		}
	}
}

func (s *socket) beat() error {
	return s.send(OpHeartbeat, heartbeat{T: time.Now().UnixMilli(), SeqAck: int(s.seq.Load())})
}

func (s *socket) heartbeat(ctx context.Context) error {
	var interval time.Duration
	select {
	case <-ctx.Done():
		return ctx.Err()
	case interval = <-s.interval:
	}
	s.acked.Store(true)
	t := time.NewTicker(interval)
	defer t.Stop()
	for {
		if !s.acked.Swap(false) {
			return ErrHeartbeatTimeout
		}
		if err := s.beat(); err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.C:
		}
	}
}
