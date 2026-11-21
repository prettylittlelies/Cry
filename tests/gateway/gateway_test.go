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

	"github.com/prettylittlelies/cry/build"
	"github.com/prettylittlelies/cry/gateway"
	"github.com/prettylittlelies/cry/properties"
	"github.com/prettylittlelies/cry/tests/testkit"
	"github.com/prettylittlelies/cry/transport/websocket"
)

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

func TestFastConnectIdentifyThenHeartbeat(t *testing.T) {
	h := start(t, true)
	p := testkit.AcceptGatewayPeer(t, h.srv)
	p.Hello(150)
	id := p.Expect(gateway.OpIdentify)
	if want := []string{"op", "d"}; !slices.Equal(id.Keys, want) {
		t.Fatalf("frame keys %v", id.Keys)
	}
	if got, want := testkit.KeysOf(t, id.D), []string{"token", "capabilities", "properties", "client_state"}; !slices.Equal(got, want) {
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
	p.Send(map[string]any{"op": 0, "s": 1, "t": "READY", "d": map[string]any{"session_id": "abc", "resume_gateway_url": h.srv.URL}})
	if e := <-h.events; e.Type != "READY" || e.Seq != 1 {
		t.Fatalf("event %+v", e)
	}
	ts := p.Expect(gateway.OpUpdateTimeSpentSessionID)
	if got := testkit.KeysOf(t, ts.D); !slices.Equal(got, []string{"initialization_timestamp", "session_id", "client_launch_id"}) {
		t.Fatalf("time spent keys %v", got)
	}
	hb := p.Expect(gateway.OpQoSHeartbeat)
	if got := testkit.KeysOf(t, hb.D); !slices.Equal(got, []string{"seq", "qos"}) {
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
	p := testkit.AcceptGatewayPeer(t, h.srv)
	p.Hello(10000)
	p.Expect(gateway.OpIdentify)
	p.Send(map[string]any{"op": 0, "s": 7, "t": "READY", "d": map[string]any{"session_id": "abc", "resume_gateway_url": h.srv.URL}})
	<-h.events
	p.Send(map[string]any{"op": 7, "d": nil})
	next := testkit.AcceptGatewayPeer(t, h.srv)
	next.Hello(10000)
	r := next.Expect(gateway.OpResume)
	if got := testkit.KeysOf(t, r.D); !slices.Equal(got, []string{"token", "session_id", "seq"}) {
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
	p := testkit.AcceptGatewayPeer(t, h.srv)
	p.Hello(10000)
	p.Expect(gateway.OpIdentify)
	derived := base64.StdEncoding.EncodeToString([]byte{0x0a, 0x02, 0xca, 0xfe})
	p.Send(map[string]any{"op": 0, "s": 1, "t": "READY", "d": map[string]any{"session_id": "abc", "qos_token": derived}})
	<-h.events
	p.Send(map[string]any{"op": 9, "d": false})
	id := p.Expect(gateway.OpIdentify)
	if got, want := testkit.KeysOf(t, id.D), []string{"token", "capabilities", "properties", "presence", "compress", "client_state", "qos_token"}; !slices.Equal(got, want) {
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
	p := testkit.AcceptGatewayPeer(t, h.srv)
	p.Hello(10000)
	p.Expect(gateway.OpIdentify)
	p.Close(4004)
	select {
	case err := <-h.done:
		if !errors.Is(err, gateway.ErrAuthenticationFailed) {
			t.Fatalf("got %v", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("gateway kept running after 4004")
	}
}
