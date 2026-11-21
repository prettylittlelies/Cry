package gateway_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"slices"
	"testing"
	"time"

	"github.com/gobwas/ws"
	"github.com/klauspost/compress/zstd"
	"github.com/prettylittlelies/cry/build"
	"github.com/prettylittlelies/cry/gateway"
	"github.com/prettylittlelies/cry/internal/etf"
	"github.com/prettylittlelies/cry/properties"
	"github.com/prettylittlelies/cry/tests/testkit"
	"github.com/prettylittlelies/cry/transport/websocket"
)

type peer struct {
	t   *testing.T
	ws  *testkit.WSPeer
	buf bytes.Buffer
	zw  *zstd.Encoder
}

func newPeer(t *testing.T, srv *testkit.WSServer) *peer {
	t.Helper()
	select {
	case p := <-srv.Peers:
		g := &peer{t: t, ws: p}
		g.zw, _ = zstd.NewWriter(&g.buf)
		return g
	case <-time.After(5 * time.Second):
		t.Fatal("client never connected")
	}
	return nil
}

func (p *peer) send(v any) {
	p.t.Helper()
	b, err := etf.Marshal(v)
	if err != nil {
		p.t.Fatal(err)
	}
	p.zw.Write(b)
	p.zw.Flush()
	msg := bytes.Clone(p.buf.Bytes())
	p.buf.Reset()
	if err := p.ws.Send(ws.NewBinaryFrame(msg)); err != nil {
		p.t.Fatal(err)
	}
}

type packet struct {
	Op   int             `json:"op"`
	D    json.RawMessage `json:"d"`
	raw  json.RawMessage
	keys []string
}

func keysOf(t *testing.T, raw json.RawMessage) []string {
	t.Helper()
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.Token()
	var out []string
	for dec.More() {
		k, _ := dec.Token()
		out = append(out, k.(string))
		var skip json.RawMessage
		dec.Decode(&skip)
	}
	return out
}

func (p *peer) recv() packet {
	p.t.Helper()
	p.ws.Conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	f, err := p.ws.Receive()
	if err != nil {
		p.t.Fatal(err)
	}
	if f.Header.OpCode != ws.OpBinary {
		p.t.Fatalf("expected binary frame, got %v", f.Header.OpCode)
	}
	raw, err := etf.ToJSON(f.Payload)
	if err != nil {
		p.t.Fatal(err)
	}
	var pk packet
	json.Unmarshal(raw, &pk)
	pk.raw, pk.keys = raw, keysOf(p.t, raw)
	return pk
}

func (p *peer) expect(op int) packet {
	p.t.Helper()
	for {
		pk := p.recv()
		if pk.Op == op {
			return pk
		}
		if pk.Op != int(gateway.OpQoSHeartbeat) && pk.Op != int(gateway.OpUpdateTimeSpentSessionID) {
			p.t.Fatalf("expected op %d, got %s", op, pk.raw)
		}
	}
}

func (p *peer) close(code int) {
	p.ws.Send(ws.NewCloseFrame(ws.NewCloseFrameBody(ws.StatusCode(code), "")))
	p.ws.Conn.Close()
}

func session(t *testing.T) *properties.Session {
	t.Helper()
	v := func(s string) build.Version { out, _ := build.ParseVersion(s); return out }
	return properties.NewSession(properties.NewDesktop(properties.MacOSHost(), properties.DesktopRelease{
		Channel: build.Stable, Host: v("0.0.421"), Electron: v("44.10.0"), Chrome: v("152.0.7977.169"), Number: 678818,
	}))
}

type harness struct {
	srv    *testkit.WSServer
	gw     *gateway.Gateway
	events chan gateway.Event
	done   chan error
	cancel context.CancelFunc
}

func start(t *testing.T, fast bool) *harness {
	t.Helper()
	srv := testkit.NewWSServer(t)
	h := &harness{srv: srv, events: make(chan gateway.Event, 32), done: make(chan error, 1)}
	h.gw = gateway.New(gateway.Config{
		Token:       "token",
		Session:     session(t),
		Endpoint:    srv.URL,
		FastConnect: fast,
		Browser:     websocket.Browser{UserAgent: "ua", Origin: "https://discord.com", AcceptLanguage: "en-US"},
		Handler:     func(e gateway.Event) { h.events <- e },
	})
	h.gw.Dialer().RootCAs = srv.Pool
	ctx, cancel := context.WithCancel(context.Background())
	h.cancel = cancel
	t.Cleanup(cancel)
	go func() { h.done <- h.gw.Run(ctx) }()
	return h
}

func hello(interval int) map[string]any {
	return map[string]any{"op": 10, "d": map[string]any{"heartbeat_interval": interval}}
}

func TestFastConnectIdentifyThenHeartbeat(t *testing.T) {
	h := start(t, true)
	p := newPeer(t, h.srv)
	p.send(hello(150))
	id := p.expect(int(gateway.OpIdentify))
	if want := []string{"op", "d"}; !slices.Equal(id.keys, want) {
		t.Fatalf("frame keys %v", id.keys)
	}
	if got, want := keysOf(t, id.D), []string{"token", "capabilities", "properties", "client_state"}; !slices.Equal(got, want) {
		t.Fatalf("identify keys %v", got)
	}
	var d struct {
		Token        string         `json:"token"`
		Capabilities int64          `json:"capabilities"`
		Properties   map[string]any `json:"properties"`
		ClientState  map[string]any `json:"client_state"`
	}
	json.Unmarshal(id.D, &d)
	if d.Capabilities != 1767421 || d.Properties["is_fast_connect"] != true || d.Properties["os"] != "Mac OS X" {
		t.Fatalf("identify %s", id.D)
	}
	for _, absent := range []string{"launch_signature", "client_heartbeat_session_id", "client_app_state"} {
		if _, ok := d.Properties[absent]; ok {
			t.Fatalf("fast connect properties carry %s", absent)
		}
	}
	p.send(map[string]any{"op": 0, "s": 1, "t": "READY", "d": map[string]any{"session_id": "abc", "resume_gateway_url": h.srv.URL}})
	if e := <-h.events; e.Type != "READY" || e.Seq != 1 {
		t.Fatalf("event %+v", e)
	}
	ts := p.expect(int(gateway.OpUpdateTimeSpentSessionID))
	if got := keysOf(t, ts.D); !slices.Equal(got, []string{"initialization_timestamp", "session_id", "client_launch_id"}) {
		t.Fatalf("time spent keys %v", got)
	}
	hb := p.expect(int(gateway.OpQoSHeartbeat))
	if got := keysOf(t, hb.D); !slices.Equal(got, []string{"seq", "qos"}) {
		t.Fatalf("heartbeat keys %v", got)
	}
	var beat struct {
		Seq int64 `json:"seq"`
		QoS struct {
			Active  bool     `json:"active"`
			Ver     int      `json:"ver"`
			Reasons []string `json:"reasons"`
		} `json:"qos"`
	}
	json.Unmarshal(hb.D, &beat)
	if beat.Seq != 1 || !beat.QoS.Active || beat.QoS.Ver != 31 || !slices.Equal(beat.QoS.Reasons, []string{"foregrounded"}) {
		t.Fatalf("heartbeat %s", hb.D)
	}
}

func TestReconnectResumes(t *testing.T) {
	h := start(t, true)
	p := newPeer(t, h.srv)
	p.send(hello(10000))
	p.expect(int(gateway.OpIdentify))
	p.send(map[string]any{"op": 0, "s": 7, "t": "READY", "d": map[string]any{"session_id": "abc", "resume_gateway_url": h.srv.URL}})
	<-h.events
	p.send(map[string]any{"op": 7, "d": nil})
	next := newPeer(t, h.srv)
	next.send(hello(10000))
	r := next.expect(int(gateway.OpResume))
	if got := keysOf(t, r.D); !slices.Equal(got, []string{"token", "session_id", "seq"}) {
		t.Fatalf("resume keys %v", got)
	}
	var res struct {
		Token     string `json:"token"`
		SessionID string `json:"session_id"`
		Seq       int64  `json:"seq"`
	}
	json.Unmarshal(r.D, &res)
	if res.Token != "token" || res.SessionID != "abc" || res.Seq != 7 {
		t.Fatalf("resume %s", r.D)
	}
}

func TestInvalidSessionFullIdentify(t *testing.T) {
	h := start(t, true)
	p := newPeer(t, h.srv)
	p.send(hello(10000))
	p.expect(int(gateway.OpIdentify))
	derived := base64.StdEncoding.EncodeToString([]byte{0x0a, 0x02, 0xca, 0xfe})
	p.send(map[string]any{"op": 0, "s": 1, "t": "READY", "d": map[string]any{"session_id": "abc", "qos_token": derived}})
	<-h.events
	p.send(map[string]any{"op": 9, "d": false})
	id := p.expect(int(gateway.OpIdentify))
	if got, want := keysOf(t, id.D), []string{"token", "capabilities", "properties", "presence", "compress", "client_state", "qos_token"}; !slices.Equal(got, want) {
		t.Fatalf("identify keys %v", got)
	}
	var d struct {
		Properties map[string]any  `json:"properties"`
		Presence   json.RawMessage `json:"presence"`
		Compress   bool            `json:"compress"`
		QoSToken   string          `json:"qos_token"`
	}
	json.Unmarshal(id.D, &d)
	if d.Properties["is_fast_connect"] != false || d.Properties["gateway_connect_reasons"] != "COLD_START" || d.Properties["launch_signature"] == nil {
		t.Fatalf("properties %v", d.Properties)
	}
	if string(d.Presence) != `{"status":"online","since":0,"activities":[],"afk":false}` {
		t.Fatalf("presence %s", d.Presence)
	}
	token, _ := base64.StdEncoding.DecodeString(d.QoSToken)
	if want := []byte{0x0a, 0x02, 0x08, 0x01, 0x12, 0x04, 0x0a, 0x02, 0xca, 0xfe}; !bytes.Equal(token, want) {
		t.Fatalf("qos token % x", token)
	}
}

func TestAuthenticationFailureStops(t *testing.T) {
	h := start(t, true)
	p := newPeer(t, h.srv)
	p.send(hello(10000))
	p.expect(int(gateway.OpIdentify))
	p.close(4004)
	select {
	case err := <-h.done:
		if !errors.Is(err, gateway.ErrAuthenticationFailed) {
			t.Fatalf("got %v", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("gateway kept running after 4004")
	}
}
