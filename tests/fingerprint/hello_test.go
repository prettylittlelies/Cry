package fingerprint_test

import (
	"bytes"
	"encoding/json"
	"net"
	"slices"
	"testing"

	"github.com/prettylittlelies/cry/fingerprint"
	"github.com/prettylittlelies/cry/tests/testkit"
	utls "github.com/refraction-networking/utls"
)

var deterministic = []uint16{5, 10, 11, 13, 16, 17613, 18, 23, 27, 35, 43, 45, 65281}

func capturedHello(t *testing.T, name string) *testkit.Hello {
	t.Helper()
	line, _, _ := bytes.Cut(testkit.Fixture(t, "fingerprint", name), []byte("\n"))
	var rec struct {
		Hello []byte `json:"client_hello"`
	}
	if err := json.Unmarshal(line, &rec); err != nil {
		t.Fatal(err)
	}
	h, err := testkit.ParseHello(rec.Hello)
	if err != nil {
		t.Fatal(err)
	}
	return h
}

func builtHello(t *testing.T, spec fingerprint.ClientHello, alpn []string) *testkit.Hello {
	t.Helper()
	client, server := net.Pipe()
	defer client.Close()
	defer server.Close()
	u := utls.UClient(client, &utls.Config{ServerName: "gateway.discord.gg", NextProtos: alpn, OmitEmptyPsk: true}, utls.HelloCustom)
	if err := u.ApplyPreset(spec()); err != nil {
		t.Fatal(err)
	}
	if err := u.BuildHandshakeState(); err != nil {
		t.Fatal(err)
	}
	raw := u.HandshakeState.Hello.Raw
	record := append([]byte{0x16, 0x03, 0x01, byte(len(raw) >> 8), byte(len(raw))}, raw...)
	h, err := testkit.ParseHello(record)
	if err != nil {
		t.Fatal(err)
	}
	return h
}

func compareHello(t *testing.T, got, want *testkit.Hello) {
	t.Helper()
	if !slices.Equal(got.Ciphers, want.Ciphers) {
		t.Errorf("ciphers\n got  %v\n want %v", got.Ciphers, want.Ciphers)
	}
	if !slices.Equal(got.ExtensionSet(), want.ExtensionSet()) {
		t.Errorf("extensions\n got  %v\n want %v", got.ExtensionSet(), want.ExtensionSet())
	}
	for _, id := range deterministic {
		if !bytes.Equal(got.Data[id], want.Data[id]) {
			t.Errorf("extension %d\n got  %x\n want %x", id, got.Data[id], want.Data[id])
		}
	}
}

func TestWebSocketHelloMatchesChromium(t *testing.T) {
	compareHello(t, builtHello(t, fingerprint.Chromium152HTTP1, fingerprint.ALPNHTTP1), capturedHello(t, "electron44-websocket.jsonl"))
}

func TestHTTP2HelloMatchesChromium(t *testing.T) {
	compareHello(t, builtHello(t, fingerprint.Chromium152, fingerprint.ALPNHTTP2), capturedHello(t, "electron44-upload.jsonl"))
}
