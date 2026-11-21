package testkit

import (
	"net/http"
	"testing"

	"github.com/prettylittlelies/cry/build"
	"github.com/prettylittlelies/cry/fingerprint"
	"github.com/prettylittlelies/cry/properties"
	"github.com/prettylittlelies/cry/rest"
	"github.com/prettylittlelies/cry/transport"
)

func version(t testing.TB, s string) build.Version {
	t.Helper()
	v, err := build.ParseVersion(s)
	if err != nil {
		t.Fatal(err)
	}
	return v
}

func Identity(t testing.TB) rest.Identity {
	t.Helper()
	chrome := version(t, "152.0.7977.169")
	host := properties.MacOSHost()
	d := properties.NewDesktop(host, properties.DesktopRelease{
		Channel:  build.Stable,
		Host:     version(t, "0.0.421"),
		Electron: version(t, "44.10.0"),
		Chrome:   chrome,
		Number:   678818,
	})
	return rest.Identity{Session: properties.NewSession(d), Host: host, Chrome: chrome}
}

func RESTClient(t testing.TB, h http.HandlerFunc, opts ...rest.Option) (*rest.Client, string) {
	t.Helper()
	srv := NewH2Server(t, h)
	tr := transport.New(fingerprint.Electron44(), transport.WithRootCAs(srv.Pool))
	t.Cleanup(tr.CloseIdleConnections)
	return rest.New("token", tr, Identity(t), append([]rest.Option{rest.WithBaseURL(srv.URL)}, opts...)...), srv.URL
}
