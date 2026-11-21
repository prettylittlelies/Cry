//go:build live

package live_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/prettylittlelies/cry/build"
	"github.com/prettylittlelies/cry/fingerprint"
	"github.com/prettylittlelies/cry/gateway"
	"github.com/prettylittlelies/cry/properties"
	"github.com/prettylittlelies/cry/transport"
	"github.com/prettylittlelies/cry/transport/websocket"
)

func token(t *testing.T) string {
	t.Helper()
	if v := os.Getenv("CRY_TOKEN"); v != "" {
		return v
	}
	_, file, _, _ := runtime.Caller(0)
	raw, err := os.ReadFile(filepath.Join(filepath.Dir(file), ".token"))
	if err != nil {
		t.Skip("no token: set CRY_TOKEN or create tests/live/.token")
	}
	return strings.TrimSpace(string(raw))
}

func TestGatewayReady(t *testing.T) {
	tok := token(t)
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()
	tr := transport.New(fingerprint.Electron44())
	defer tr.CloseIdleConnections()
	snap, err := build.NewScraper(build.WithDoer(&http.Client{Transport: tr}), build.WithChannels(build.Stable)).Fetch(ctx)
	if err != nil {
		t.Fatal(err)
	}
	r, _ := snap.Release(build.Stable)
	host := properties.MacOSHost()
	session := properties.NewSession(properties.NewDesktop(host, properties.MacOSRelease(r)))
	jar, _ := cookiejar.New(nil)
	ready := make(chan gateway.Event, 1)
	gw := gateway.New(gateway.Config{
		Token:       tok,
		Session:     session,
		FastConnect: true,
		Jar:         jar,
		Browser:     websocket.Browser{UserAgent: session.UserAgent(), Origin: "https://discord.com", AcceptLanguage: host.Locale},
		Handler: func(e gateway.Event) {
			if e.Type == "READY" {
				select {
				case ready <- e:
				default:
				}
			}
		},
	})
	errc := make(chan error, 1)
	go func() { errc <- gw.Run(ctx) }()
	select {
	case e := <-ready:
		var d struct {
			User struct {
				ID       string `json:"id"`
				Username string `json:"username"`
			} `json:"user"`
			SessionID string `json:"session_id"`
			Guilds    []any  `json:"guilds"`
			QoSToken  string `json:"qos_token"`
		}
		if err := json.Unmarshal(e.Data, &d); err != nil {
			t.Fatal(err)
		}
		t.Logf("READY: session=%s guilds=%d qos_token=%v ready_bytes=%d", d.SessionID, len(d.Guilds), d.QoSToken != "", len(e.Data))
		if d.SessionID == "" || d.User.ID == "" {
			t.Fatal("READY missing session or user")
		}
	case err := <-errc:
		t.Fatalf("gateway stopped: %v", err)
	case <-ctx.Done():
		t.Fatal("no READY before timeout")
	}
}
