//go:build live

package live_test

import (
	"context"
	"encoding/json"
	"net/http/cookiejar"
	"testing"
	"time"

	"github.com/prettylittlelies/cry/fingerprint"
	"github.com/prettylittlelies/cry/gateway/codec"
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
	r, err := codec.NewReader(c, codec.ETF, codec.ZstdStream)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	raw, err := r.Next()
	if err != nil {
		t.Fatal(err)
	}
	var hello struct {
		Op int `json:"op"`
		D  struct {
			HeartbeatInterval int      `json:"heartbeat_interval"`
			Trace             []string `json:"_trace"`
		} `json:"d"`
	}
	if err := json.Unmarshal(raw, &hello); err != nil {
		t.Fatal(err)
	}
	if hello.Op != 10 || hello.D.HeartbeatInterval == 0 {
		t.Fatalf("unexpected hello %s", raw)
	}
	t.Logf("HELLO: %s", raw)
}
