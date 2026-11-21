//go:build live

package live_test

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/prettylittlelies/cry/fingerprint"
	"github.com/prettylittlelies/cry/tests/testkit"
	"github.com/prettylittlelies/cry/transport"
)

type peetPriority struct {
	Weight    int `json:"weight"`
	DependsOn int `json:"depends_on"`
	Exclusive int `json:"exclusive"`
}

type peetFrame struct {
	Type     string        `json:"frame_type"`
	Headers  []string      `json:"headers"`
	Priority *peetPriority `json:"priority"`
}

type peet struct {
	TLS struct {
		JA4       string `json:"ja4"`
		PeetPrint string `json:"peetprint"`
	} `json:"tls"`
	HTTP2 struct {
		Akamai string      `json:"akamai_fingerprint"`
		Frames []peetFrame `json:"sent_frames"`
	} `json:"http2"`
}

func (p peet) headers() peetFrame {
	for _, f := range slices.Backward(p.HTTP2.Frames) {
		if f.Type == "HEADERS" {
			return f
		}
	}
	return peetFrame{}
}

func names(lines []string) []string {
	out := make([]string, len(lines))
	for i, l := range lines {
		out[i], _, _ = strings.Cut(l, ": ")
	}
	return out
}

func reference(t *testing.T, name string) peet {
	t.Helper()
	var p peet
	if err := json.Unmarshal(testkit.Fixture(t, "fingerprint", name), &p); err != nil {
		t.Fatal(err)
	}
	return p
}

func capture(t *testing.T, tr *transport.Transport, header http.Header) peet {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://tls.peet.ws/api/all", nil)
	req.Header = header.Clone()
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

func rendererHeaders(ref peet) http.Header {
	h := http.Header{}
	for _, l := range ref.headers().Headers {
		k, v, _ := strings.Cut(l, ": ")
		if !strings.HasPrefix(k, ":") {
			h.Add(k, v)
		}
	}
	return h
}

func TestElectron44Fingerprint(t *testing.T) {
	fresh, renderer := reference(t, "electron44-main.json"), reference(t, "electron44-renderer.json")
	tr := transport.New(fingerprint.Electron44())
	header := rendererHeaders(renderer)

	first := capture(t, tr, header)
	if first.TLS.JA4 != fresh.TLS.JA4 {
		t.Errorf("fresh ja4: got %s, want %s", first.TLS.JA4, fresh.TLS.JA4)
	}
	if first.HTTP2.Akamai != renderer.HTTP2.Akamai {
		t.Errorf("akamai: got %s, want %s", first.HTTP2.Akamai, renderer.HTTP2.Akamai)
	}
	got, want := first.headers(), renderer.headers()
	if g, w := names(got.Headers), names(want.Headers); !slices.Equal(g, w) {
		t.Errorf("header order:\n got %v\nwant %v", g, w)
	}
	if *got.Priority != *want.Priority {
		t.Errorf("priority: got %+v, want %+v", *got.Priority, *want.Priority)
	}

	tr.CloseIdleConnections()
	resumed := capture(t, tr, header)
	if resumed.TLS.JA4 != renderer.TLS.JA4 {
		t.Errorf("resumed ja4: got %s, want %s", resumed.TLS.JA4, renderer.TLS.JA4)
	}
	if resumed.TLS.PeetPrint != renderer.TLS.PeetPrint {
		t.Errorf("resumed peetprint:\n got %s\nwant %s", resumed.TLS.PeetPrint, renderer.TLS.PeetPrint)
	}
}
