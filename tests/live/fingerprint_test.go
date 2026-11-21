//go:build live

package live_test

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/prettylittlelies/cry/fingerprint"
	"github.com/prettylittlelies/cry/tests/testkit"
	"golang.org/x/net/http2"
)

type peet struct {
	TLS struct {
		JA4       string `json:"ja4"`
		PeetPrint string `json:"peetprint"`
	} `json:"tls"`
}

func capture(t *testing.T, d *fingerprint.Dialer) peet {
	t.Helper()
	tr := &http2.Transport{DialTLSContext: func(ctx context.Context, network, addr string, _ *tls.Config) (net.Conn, error) {
		return d.DialTLS(ctx, network, addr)
	}}
	defer tr.CloseIdleConnections()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://tls.peet.ws/api/all", nil)
	resp, err := tr.RoundTrip(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	var p peet
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		t.Fatal(err)
	}
	return p
}

func reference(t *testing.T, name string) peet {
	t.Helper()
	var p peet
	if err := json.Unmarshal(testkit.Fixture(t, "fingerprint", name), &p); err != nil {
		t.Fatal(err)
	}
	return p
}

func TestChromium152ClientHello(t *testing.T) {
	d := fingerprint.NewDialer(fingerprint.Chromium152)
	fresh, resumed := capture(t, d), capture(t, d)
	if want := reference(t, "electron44-main.json").TLS.JA4; fresh.TLS.JA4 != want {
		t.Errorf("fresh ja4: got %s, want %s", fresh.TLS.JA4, want)
	}
	if want := reference(t, "electron44-renderer.json").TLS.JA4; resumed.TLS.JA4 != want {
		t.Errorf("resumed ja4: got %s, want %s", resumed.TLS.JA4, want)
	}
	if want := reference(t, "electron44-renderer.json").TLS.PeetPrint; resumed.TLS.PeetPrint != want {
		t.Errorf("resumed peetprint:\n got %s\nwant %s", resumed.TLS.PeetPrint, want)
	}
}
