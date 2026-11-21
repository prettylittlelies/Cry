//go:build live

package live_test

import (
	"context"
	"net/http/cookiejar"
	"testing"
	"time"

	"github.com/prettylittlelies/cry/fingerprint"
	"github.com/prettylittlelies/cry/transport/websocket"
)

func TestGatewayHandshake(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	jar, _ := cookiejar.New(nil)
	c, err := websocket.Dial(ctx, "wss://gateway.discord.gg/?encoding=etf&v=9&compress=zstd-stream", websocket.Config{
		Dialer: websocket.NewDialer(fingerprint.Chromium152HTTP1),
		Browser: websocket.Browser{
			UserAgent:      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) discord/0.0.421 Chrome/152.0.7977.169 Electron/44.10.0 Safari/537.36",
			Origin:         "https://discord.com",
			AcceptLanguage: "en-US",
		},
		Jar: jar,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close(1000, "")
	op, msg, err := c.Read()
	if err != nil {
		t.Fatal(err)
	}
	if op != websocket.Binary || len(msg) < 4 || msg[0] != 0x28 || msg[1] != 0xb5 || msg[2] != 0x2f || msg[3] != 0xfd {
		t.Fatalf("expected a zstd frame, got op=%v % x", op, msg[:min(len(msg), 8)])
	}
	t.Logf("hello frame: %d zstd bytes, cookies set: %d", len(msg), len(c.Response.Cookies()))
}
