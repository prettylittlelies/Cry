package voice

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/prettylittlelies/cry"
	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/mls"
	"github.com/prettylittlelies/cry/snowflake"
	"github.com/prettylittlelies/cry/transport/websocket"
	"github.com/prettylittlelies/cry/voice/dave"
)

const (
	frameDuration   = 20 * time.Millisecond
	keepaliveEvery  = 5 * time.Second
	trailingSilence = 5
	udpPingMagic    = 0x1337CAFE
)

var (
	ErrNotReady = errors.New("voice: connection not ready")
	ErrClosed   = errors.New("voice: connection closed")
)

var silenceFrame = []byte{0xF8, 0xFF, 0xFE}

type Options struct {
	SelfMute         bool
	SelfDeaf         bool
	DisableDAVE      bool
	Signer           *mls.SignaturePrivateKey
	PreferredRegions []string
	Video            MediaType
	OnOpus           func(user snowflake.ID, frame []byte)
}

type FrameSource interface {
	Next() ([]byte, error)
}

type Conn struct {
	client  *cry.Client
	guild   snowflake.ID
	channel snowflake.ID
	self    snowflake.ID
	opts    Options

	sessionID string
	token     string
	endpoint  string

	ws         atomic.Pointer[socket]
	udp        atomic.Pointer[udpLink]
	ssrc       uint32
	crypt      atomic.Pointer[transport]
	rtp        rtpSender
	video      *videoSender
	videoFPS   atomic.Int32
	videoCodec atomic.Pointer[string]
	handshake  chan error
	resumed    chan struct{}

	usersMu sync.RWMutex
	users   map[uint32]snowflake.ID

	e2ee *e2ee

	ctx    context.Context
	cancel context.CancelFunc
	done   chan error
	err    atomic.Pointer[error]
}

func Join(ctx context.Context, c *cry.Client, guild, channel snowflake.ID, opts Options) (*Conn, error) {
	self := c.State().Me().ID
	v := &Conn{client: c, guild: guild, channel: channel, self: self, opts: opts, users: map[uint32]snowflake.ID{}, done: make(chan error, 1)}
	if opts.Video != "" {
		v.video = &videoSender{}
	}
	if !opts.DisableDAVE {
		session, err := dave.NewSession(self, opts.Signer)
		if err != nil {
			return nil, err
		}
		v.e2ee = newE2EE(v, session)
	}
	states := make(chan *discord.VoiceStateUpdateEvent, 1)
	servers := make(chan *discord.VoiceServerUpdateEvent, 1)
	removeState := cry.On(c, func(_ context.Context, e *discord.VoiceStateUpdateEvent) {
		if e.UserID == self && e.ChannelID != nil && *e.ChannelID == channel {
			select {
			case states <- e:
			default:
			}
		}
	})
	defer removeState()
	removeServer := cry.On(c, func(_ context.Context, e *discord.VoiceServerUpdateEvent) {
		if guild != 0 && e.GuildID != nil && *e.GuildID == guild || guild == 0 && e.GuildID == nil && e.ChannelID == channel {
			select {
			case servers <- e:
			default:
			}
		}
	})
	defer removeServer()
	if err := c.Gateway().UpdateVoiceState(ctx, v.voiceState(&channel)); err != nil {
		return nil, err
	}
	var state *discord.VoiceStateUpdateEvent
	var server *discord.VoiceServerUpdateEvent
	for state == nil || server == nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case state = <-states:
		case server = <-servers:
		}
	}
	if server.Endpoint == nil {
		return nil, fmt.Errorf("voice: server has no endpoint yet")
	}
	v.sessionID, v.token, v.endpoint = state.SessionID, server.Token, *server.Endpoint
	v.ctx, v.cancel = context.WithCancel(context.Background())
	if err := v.connect(ctx, false); err != nil {
		v.cancel()
		return nil, err
	}
	go v.run()
	return v, nil
}

func (v *Conn) voiceState(channel *snowflake.ID) discord.UpdateVoiceState {
	var guild *snowflake.ID
	if v.guild != 0 {
		guild = &v.guild
	}
	return discord.UpdateVoiceState{GuildID: guild, ChannelID: channel, SelfMute: v.opts.SelfMute, SelfDeaf: v.opts.SelfDeaf, PreferredRegions: v.opts.PreferredRegions}
}

func (v *Conn) serverID() snowflake.ID {
	if v.guild != 0 {
		return v.guild
	}
	return v.channel
}

func (v *Conn) socketConfig() websocket.Config {
	g := v.client.Gateway()
	return websocket.Config{Dialer: g.Dialer(), Browser: g.Browser(), Jar: g.Jar()}
}

func (v *Conn) connect(ctx context.Context, resuming bool) error {
	ws, err := dialSocket(ctx, v.endpoint, v.socketConfig())
	if err != nil {
		return err
	}
	if old := v.ws.Load(); old != nil {
		ws.seq.Store(old.seq.Load())
	}
	v.handshake, v.resumed = make(chan error, 1), make(chan struct{}, 1)
	v.ws.Store(ws)
	go func() { v.done <- ws.read(v) }()
	go ws.heartbeat(v.ctx)
	if resuming {
		if err := ws.send(OpResume, resume{ServerID: v.serverID(), ChannelID: v.channel, SessionID: v.sessionID, Token: v.token, SeqAck: int(ws.seq.Load())}); err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-v.resumed:
			return nil
		}
	}
	maxDAVE := 0
	if v.e2ee != nil {
		maxDAVE = dave.ProtocolVersion
	}
	id := identify{ServerID: v.serverID(), ChannelID: v.channel, UserID: v.self, SessionID: v.sessionID, Token: v.token, MaxDAVEProtocol: maxDAVE}
	if v.video != nil {
		id.Video = true
		id.Streams = []streamInfo{{Type: string(v.opts.Video), RID: primaryRID, Quality: 100}}
	}
	if err := ws.send(OpIdentify, id); err != nil {
		return err
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-v.handshake:
		return err
	case err := <-v.done:
		return fmt.Errorf("voice: handshake: %w", err)
	}
}

func (v *Conn) run() {
	go v.receive()
	go v.keepalive()
	for {
		select {
		case <-v.ctx.Done():
			return
		case err := <-v.done:
			if v.ctx.Err() != nil {
				return
			}
			var ce *websocket.CloseError
			if errors.As(err, &ce) && ce.Code >= 4000 && ce.Code != 4015 {
				v.fail(err)
				return
			}
			v.ws.Load().close()
			ctx, cancel := context.WithTimeout(v.ctx, 10*time.Second)
			err = v.connect(ctx, true)
			cancel()
			if err != nil {
				v.fail(err)
				return
			}
		}
	}
}

func (v *Conn) fail(err error) {
	v.err.CompareAndSwap(nil, &err)
	v.cancel()
	v.closeTransports()
}

func (v *Conn) Err() error {
	if err := v.err.Load(); err != nil {
		return *err
	}
	return nil
}

func (v *Conn) closeTransports() {
	if ws := v.ws.Load(); ws != nil {
		ws.close()
	}
	if u := v.udp.Load(); u != nil {
		u.close()
	}
}

func (v *Conn) onJSON(op Opcode, d json.RawMessage) error {
	switch op {
	case OpReady:
		var r ready
		if err := json.Unmarshal(d, &r); err != nil {
			return err
		}
		go func() {
			if err := v.selectProtocol(r); err != nil {
				v.handshake <- err
			}
		}()
	case OpSessionDescription:
		var sd sessionDescription
		if err := json.Unmarshal(d, &sd); err != nil {
			return err
		}
		key := make([]byte, len(sd.SecretKey))
		for i, b := range sd.SecretKey {
			key[i] = byte(b)
		}
		t, err := newTransport(sd.Mode, key)
		if err != nil {
			return err
		}
		v.crypt.Store(t)
		if sd.VideoCodec != "" {
			v.videoCodec.Store(&sd.VideoCodec)
		}
		if v.e2ee != nil {
			if err := v.e2ee.start(sd.DAVEProtocol); err != nil {
				return err
			}
		}
		select {
		case v.handshake <- nil:
		default:
		}
	case OpResumed:
		select {
		case v.resumed <- struct{}{}:
		default:
		}
	case OpSpeaking:
		var s speaking
		if err := json.Unmarshal(d, &s); err != nil {
			return err
		}
		v.usersMu.Lock()
		v.users[s.SSRC] = s.UserID
		v.usersMu.Unlock()
	case OpClientsConnect:
		var c clientsConnect
		if err := json.Unmarshal(d, &c); err != nil {
			return err
		}
		if v.e2ee != nil {
			v.e2ee.connected(c.UserIDs...)
		}
	case OpClientDisconnect:
		var c clientDisconnect
		if err := json.Unmarshal(d, &c); err != nil {
			return err
		}
		v.usersMu.Lock()
		for ssrc, u := range v.users {
			if u == c.UserID {
				delete(v.users, ssrc)
			}
		}
		v.usersMu.Unlock()
		if v.e2ee != nil {
			v.e2ee.disconnected(c.UserID)
		}
	default:
		if v.e2ee != nil {
			return v.e2ee.onJSON(op, d)
		}
	}
	return nil
}

func (v *Conn) onBinary(op uint8, payload []byte) error {
	if v.e2ee == nil {
		return nil
	}
	return v.e2ee.onBinary(dave.Opcode(op), payload)
}

func (v *Conn) selectProtocol(r ready) error {
	mode, err := chooseMode(r.Modes)
	if err != nil {
		return err
	}
	link, err := dialUDP(v.ctx, r.IP, r.Port)
	if err != nil {
		return err
	}
	if old := v.udp.Swap(link); old != nil {
		old.close()
	}
	v.ssrc, v.rtp.ssrc = r.SSRC, r.SSRC
	if v.video != nil {
		for _, s := range r.Streams {
			if s.RID == primaryRID || v.video.ssrc == 0 {
				v.video.ssrc, v.video.rtx = s.SSRC, s.RTXSSRC
			}
		}
		if v.video.rtx == 0 && v.video.ssrc != 0 {
			v.video.rtx = v.video.ssrc + 1
		}
	}
	codecs := []codec{{Name: "opus", Type: "audio", Priority: 1000, PayloadType: opusPayloadType}}
	if v.video != nil {
		codecs = append(codecs, codec{Name: "VP8", Type: "video", Priority: 1000, PayloadType: vp8PayloadType, RTXPayloadType: vp8RTXPayloadType})
	}
	ip, port, err := link.discover(r.SSRC)
	if err != nil {
		return err
	}
	return v.ws.Load().send(OpSelectProtocol, selectProtocol{
		Protocol:        "udp",
		Data:            protocolData{Address: ip, Port: port, Mode: mode},
		Codecs:          codecs,
		RTCConnectionID: uuid.NewString(),
	})
}

func (v *Conn) keepalive() {
	t := time.NewTicker(keepaliveEvery)
	defer t.Stop()
	var seq uint32
	for {
		select {
		case <-v.ctx.Done():
			return
		case <-t.C:
			seq++
			ping := binary.BigEndian.AppendUint32(binary.BigEndian.AppendUint32(nil, udpPingMagic), seq)
			if u := v.udp.Load(); u != nil {
				u.write(ping)
				v.report(u)
			}
		}
	}
}

func (v *Conn) report(u *udpLink) {
	t := v.crypt.Load()
	if t == nil {
		return
	}
	snapshots := []func() (uint32, senderStats){v.rtp.snapshot}
	if v.video != nil {
		snapshots = append(snapshots, v.video.snapshot)
	}
	for _, snap := range snapshots {
		if ssrc, stats := snap(); stats.packets > 0 {
			header, body := senderReport(ssrc, time.Now(), stats)
			u.write(t.seal(header, body))
		}
	}
}

func (v *Conn) Speaking(flags SpeakingFlags) error {
	ws := v.ws.Load()
	if ws == nil {
		return ErrNotReady
	}
	return ws.send(OpSpeaking, speaking{Speaking: flags, SSRC: v.ssrc})
}

func (v *Conn) WriteOpus(frame []byte) error {
	t, u := v.crypt.Load(), v.udp.Load()
	if t == nil || u == nil {
		return ErrNotReady
	}
	if v.ctx.Err() != nil {
		return ErrClosed
	}
	payload := frame
	if v.e2ee != nil && !dave.IsSilence(frame) {
		var err error
		if payload, err = v.e2ee.protect(frame); err != nil {
			return err
		}
	}
	return u.write(t.seal(v.rtp.next(len(payload)), payload))
}

func (v *Conn) Play(ctx context.Context, src FrameSource) error {
	if err := v.Speaking(Voice); err != nil {
		return err
	}
	defer v.Speaking(0)
	t := time.NewTicker(frameDuration)
	defer t.Stop()
	for {
		frame, err := src.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-v.ctx.Done():
			return ErrClosed
		case <-t.C:
		}
		if err := v.WriteOpus(frame); err != nil {
			return err
		}
	}
	for range trailingSilence {
		<-t.C
		if err := v.WriteOpus(silenceFrame); err != nil {
			return err
		}
	}
	return nil
}

func (v *Conn) receive() {
	buf := make([]byte, 2048)
	for v.ctx.Err() == nil {
		u := v.udp.Load()
		if u == nil {
			time.Sleep(frameDuration)
			continue
		}
		n, err := u.read(buf)
		if err != nil {
			continue
		}
		pkt, ok := parseRTP(buf[:n])
		t := v.crypt.Load()
		if !ok || pkt.payloadType != opusPayloadType || t == nil || v.opts.OnOpus == nil {
			continue
		}
		_, payload, err := t.open(buf[:n])
		if err != nil {
			continue
		}
		v.usersMu.RLock()
		user, known := v.users[pkt.ssrc]
		v.usersMu.RUnlock()
		if !known {
			continue
		}
		frame := append([]byte(nil), payload...)
		if v.e2ee != nil && !dave.IsSilence(frame) {
			if frame, err = v.e2ee.unprotect(user, frame); err != nil {
				continue
			}
		}
		v.opts.OnOpus(user, frame)
	}
}

func (v *Conn) EpochAuthenticator() []byte {
	if v.e2ee == nil {
		return nil
	}
	return v.e2ee.session.EpochAuthenticator()
}

func (v *Conn) Encrypted() bool {
	return v.e2ee != nil && v.e2ee.active()
}

func (v *Conn) Done() <-chan struct{} { return v.ctx.Done() }

func (v *Conn) Close(ctx context.Context) error {
	v.cancel()
	v.closeTransports()
	return v.client.Gateway().UpdateVoiceState(ctx, v.voiceState(nil))
}
