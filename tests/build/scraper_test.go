package build_test

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/prettylittlelies/cry/build"
	"github.com/prettylittlelies/cry/tests/testkit"
)

const macApp = "Discord.app/Contents/Frameworks/Electron Framework.framework/Versions/A/Resources/Info.plist"

func macPackage(t *testing.T) []byte {
	t.Helper()
	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)
	for _, name := range []string{"Discord.app/Contents/Info.plist", macApp} {
		w, err := zw.Create(name)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := w.Write(testkit.Fixture(t, "build", "Info.plist")); err != nil {
			t.Fatal(err)
		}
	}
	if err := zw.Close(); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func newScraper(t *testing.T) *build.Scraper {
	t.Helper()
	srv := testkit.NewServer(t)
	srv.Serve("/app", "text/html", testkit.Fixture(t, "build", "app.html"))
	srv.Serve("/distributions/app/manifests/latest", "application/json", testkit.Fixture(t, "build", "manifest.json"))
	srv.Serve("/api/updates/stable", "application/json", fmt.Appendf(nil, string(testkit.Fixture(t, "build", "macupdate.json")), srv.URL))
	srv.Serve("/Discord.zip", "application/zip", macPackage(t))
	srv.Serve("/electron.json", "application/json", testkit.Fixture(t, "build", "electron.json"))
	srv.Serve("/play", "text/html", testkit.Fixture(t, "build", "play.html"))
	srv.Serve("/android.json", "application/json", testkit.Fixture(t, "build", "androidversions.json"))
	srv.Serve("/itunes", "application/json", testkit.Fixture(t, "build", "itunes.json"))
	return build.NewScraper(
		build.WithDoer(srv.Client()),
		build.WithChannels(build.Stable),
		build.WithEndpoints(build.Endpoints{
			App:             func(build.Channel) string { return srv.URL + "/app" },
			Updates:         srv.URL,
			API:             srv.URL + "/api",
			ElectronIndex:   srv.URL + "/electron.json",
			PlayStore:       srv.URL + "/play",
			AndroidVersions: srv.URL + "/android.json",
			ITunesLookup:    srv.URL + "/itunes",
		}),
	)
}

func TestScraperFetch(t *testing.T) {
	snap, err := newScraper(t).Fetch(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	r, ok := snap.Release(build.Stable)
	if !ok {
		t.Fatal("stable release missing")
	}
	cases := []struct {
		name      string
		got, want string
	}{
		{"number", r.Number.String(), "678818"},
		{"windows host", r.Windows.Host.String(), "1.0.9269"},
		{"windows native", r.Windows.Native.String(), "98981"},
		{"macos host", r.MacOS.Host.String(), "0.0.421"},
		{"electron", r.MacOS.Electron.String(), "44.10.2"},
		{"chrome", r.MacOS.Chrome.String(), "152.0.7977.169"},
		{"android", snap.Mobile.Android.String(), "354.12"},
		{"android minimum", fmt.Sprint(snap.Mobile.AndroidMinimum), "388"},
		{"ios", snap.Mobile.IOS.String(), "354.0"},
		{"ios minimum os", snap.Mobile.IOSMinimumOS.String(), "16.0"},
	}
	for _, c := range cases {
		if c.got != c.want {
			t.Errorf("%s: got %q, want %q", c.name, c.got, c.want)
		}
	}
}

func TestScraperSourceError(t *testing.T) {
	srv := testkit.NewServer(t)
	srv.Handle("/", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusServiceUnavailable) })
	s := build.NewScraper(build.WithDoer(srv.Client()), build.WithChannels(build.Stable), build.WithEndpoints(build.Endpoints{
		App:     func(build.Channel) string { return srv.URL },
		Updates: srv.URL, API: srv.URL, ElectronIndex: srv.URL, PlayStore: srv.URL, AndroidVersions: srv.URL, ITunesLookup: srv.URL,
	}))
	_, err := s.Fetch(context.Background())
	var se *build.SourceError
	if !errorsAs(err, &se) {
		t.Fatalf("want *build.SourceError, got %v", err)
	}
}
