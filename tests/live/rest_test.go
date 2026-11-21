//go:build live

package live_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/prettylittlelies/cry/build"
	"github.com/prettylittlelies/cry/fingerprint"
	"github.com/prettylittlelies/cry/properties"
	"github.com/prettylittlelies/cry/rest"
	"github.com/prettylittlelies/cry/transport"
)

func TestExperimentsUnauthenticated(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
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
	c := rest.New("", tr, rest.Identity{Session: session, Host: host, Chrome: r.MacOS.Chrome})
	out, err := rest.Call[struct {
		Fingerprint string `json:"fingerprint"`
	}](ctx, c, &rest.Request{Route: rest.NewRoute(http.MethodGet, "/experiments"), NoAuth: true})
	if err != nil {
		t.Fatal(err)
	}
	if out.Fingerprint == "" {
		t.Fatal("no fingerprint returned")
	}
	t.Logf("fingerprint issued (%d chars), ua=%s", len(out.Fingerprint), session.UserAgent())
}
