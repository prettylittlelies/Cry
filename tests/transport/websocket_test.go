package transport_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/gobwas/ws"
	"github.com/prettylittlelies/cry/fingerprint"
	"github.com/prettylittlelies/cry/tests/testkit"
	"github.com/prettylittlelies/cry/transport/websocket"
)

var browser = websocket.Browser{
	UserAgent:      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) discord/0.0.421 Chrome/152.0.7977.169 Electron/44.10.0 Safari/537.36",
	Origin:         "https://discord.com",
	AcceptLanguage: "en-US",
}

func dial(t *testing.T, srv *testkit.WSServer, jar http.CookieJar) (*websocket.Conn, *testkit.WSPeer) {
	t.Helper()
	d := websocket.NewDialer(fingerprint.Chromium152HTTP1)
	d.RootCAs = srv.Pool
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	c, err := websocket.Dial(ctx, srv.URL+"?encoding=etf&v=9&compress=zstd-stream", websocket.Config{Dialer: d, Browser: browser, Jar: jar})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { c.Close(1000, "") })
	return c, <-srv.Peers
}

func captured(t *testing.T, name string) []string {
	t.Helper()
	line, _, _ := bytes.Cut(testkit.Fixture(t, "fingerprint", name), []byte("\n"))
	var rec struct {
		Headers []string `json:"headers"`
	}
	if err := json.Unmarshal(line, &rec); err != nil {
		t.Fatal(err)
	}
	return rec.Headers
}

func headerNames(lines []string) []string {
	var out []string
	for _, l := range lines[1:] {
		k, _, _ := strings.Cut(l, ": ")
		out = append(out, k)
	}
	return out
}

func TestHandshakeOrderMatchesChromium(t *testing.T) {
	srv := testkit.NewWSServer(t)
	_, peer := dial(t, srv, nil)
	want := captured(t, "electron44-websocket.jsonl")
	if got := headerNames(peer.Request); !slices.Equal(got, headerNames(want)) {
		t.Fatalf("got  %v\nwant %v", got, headerNames(want))
	}
	if peer.Request[0] != "GET /?encoding=etf&v=9&compress=zstd-stream HTTP/1.1" {
		t.Fatalf("request line %q", peer.Request[0])
	}
}

func TestHandshakeCookieRoundTrip(t *testing.T) {
	srv := testkit.NewWSServer(t)
	srv.Extra = []string{"Set-Cookie: __cf_bm=abc; HttpOnly; SameSite=None; Partitioned; Secure; Path=/"}
	jar, _ := cookiejar.New(nil)
	first, _ := dial(t, srv, jar)
	first.Close(1000, "")
	_, peer := dial(t, srv, jar)
	want := headerNames(captured(t, "electron44-websocket-cookie.jsonl"))
	if got := headerNames(peer.Request); !slices.Equal(got, want) {
		t.Fatalf("got  %v\nwant %v", got, want)
	}
	if !slices.Contains(peer.Request, "Cookie: __cf_bm=abc") {
		t.Fatalf("cookie not replayed: %v", peer.Request)
	}
}

func TestPingFragmentsAndClose(t *testing.T) {
	srv := testkit.NewWSServer(t)
	c, peer := dial(t, srv, nil)
	go func() {
		peer.Send(ws.NewPingFrame([]byte("hb")))
		peer.Send(ws.Frame{Header: ws.Header{OpCode: ws.OpBinary, Length: 3}, Payload: []byte("dis")})
		peer.Send(ws.Frame{Header: ws.Header{Fin: true, OpCode: ws.OpContinuation, Length: 4}, Payload: []byte("cord")})
		peer.Send(ws.NewCloseFrame(ws.NewCloseFrameBody(4004, "Authentication failed.")))
	}()
	op, msg, err := c.Read()
	if err != nil || op != websocket.Binary || string(msg) != "discord" {
		t.Fatalf("op=%v msg=%q err=%v", op, msg, err)
	}
	pong, err := peer.Receive()
	if err != nil || pong.Header.OpCode != ws.OpPong || string(pong.Payload) != "hb" || !pong.Header.Masked {
		t.Fatalf("pong %+v err=%v", pong.Header, err)
	}
	_, _, err = c.Read()
	var closeErr *websocket.CloseError
	if !errors.As(err, &closeErr) || closeErr.Code != 4004 || closeErr.Reason != "Authentication failed." {
		t.Fatalf("got %v", err)
	}
	if err := c.Write(websocket.Binary, []byte("late")); !errors.Is(err, websocket.ErrClosed) {
		t.Fatalf("write after close: %v", err)
	}
}

func TestClientFramesMasked(t *testing.T) {
	srv := testkit.NewWSServer(t)
	c, peer := dial(t, srv, nil)
	payload := bytes.Repeat([]byte{0x83}, 70000)
	if err := c.Write(websocket.Binary, payload); err != nil {
		t.Fatal(err)
	}
	f, err := peer.Receive()
	if err != nil || !f.Header.Masked || !f.Header.Fin || !bytes.Equal(f.Payload, payload) {
		t.Fatalf("frame %+v err=%v", f.Header, err)
	}
}

func TestRejectsAcceptedExtension(t *testing.T) {
	srv := testkit.NewWSServer(t)
	srv.Extra = []string{"Sec-WebSocket-Extensions: permessage-deflate"}
	d := websocket.NewDialer(fingerprint.Chromium152HTTP1)
	d.RootCAs = srv.Pool
	_, err := websocket.Dial(context.Background(), srv.URL, websocket.Config{Dialer: d, Browser: browser})
	var herr *websocket.HandshakeError
	if !errors.As(err, &herr) || !strings.Contains(herr.Reason, "extension") {
		t.Fatalf("got %v", err)
	}
}
