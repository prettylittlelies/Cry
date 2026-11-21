package properties_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"slices"
	"strings"
	"testing"

	"github.com/prettylittlelies/cry/build"
	"github.com/prettylittlelies/cry/properties"
)

func release(t *testing.T) properties.DesktopRelease {
	t.Helper()
	v := func(s string) build.Version {
		out, err := build.ParseVersion(s)
		if err != nil {
			t.Fatal(err)
		}
		return out
	}
	return properties.DesktopRelease{
		Channel:  build.Stable,
		Host:     v("0.0.421"),
		Electron: v("44.10.0"),
		Chrome:   v("152.0.7977.169"),
		Number:   678818,
	}
}

func keys(t *testing.T, raw []byte) []string {
	t.Helper()
	dec := json.NewDecoder(bytes.NewReader(raw))
	if _, err := dec.Token(); err != nil {
		t.Fatal(err)
	}
	var out []string
	for dec.More() {
		tok, err := dec.Token()
		if err != nil {
			t.Fatal(err)
		}
		out = append(out, tok.(string))
		var skip json.RawMessage
		if err := dec.Decode(&skip); err != nil {
			t.Fatal(err)
		}
	}
	return out
}

func decode(t *testing.T, s *properties.Session) []byte {
	t.Helper()
	raw, err := base64.StdEncoding.DecodeString(s.Encoded())
	if err != nil {
		t.Fatal(err)
	}
	return raw
}

func TestMacOSKeyOrder(t *testing.T) {
	s := properties.NewSession(properties.NewDesktop(properties.MacOSHost(), release(t)))
	want := []string{
		"os", "browser", "release_channel", "client_version", "os_version", "os_version_is_release",
		"os_sdk_version", "os_arch", "app_arch", "system_locale", "has_client_mods", "client_launch_id",
		"browser_user_agent", "browser_version", "client_build_number", "client_event_source",
		"launch_signature", "client_heartbeat_session_id", "client_app_state",
	}
	if got := keys(t, decode(t, s)); !slices.Equal(got, want) {
		t.Fatalf("got  %v\nwant %v", got, want)
	}
}

func TestWindowsKeyOrder(t *testing.T) {
	r := release(t)
	r.Native = 98981
	s := properties.NewSession(properties.NewDesktop(properties.WindowsHost(), r))
	want := []string{
		"os", "browser", "release_channel", "client_version", "os_version", "os_version_is_release",
		"os_sdk_version", "os_arch", "app_arch", "system_locale", "has_client_mods", "client_launch_id",
		"browser_user_agent", "browser_version", "os_build_revision", "client_build_number",
		"native_build_number", "client_event_source", "launch_signature", "client_heartbeat_session_id",
		"client_app_state",
	}
	if got := keys(t, decode(t, s)); !slices.Equal(got, want) {
		t.Fatalf("got  %v\nwant %v", got, want)
	}
}

func TestDesktopValues(t *testing.T) {
	d := properties.NewDesktop(properties.MacOSHost(), release(t))
	if d.OS != "Mac OS X" || d.OSSDKVersion != "25" || d.BrowserVersion != "44.10.0" {
		t.Fatalf("mac props %+v", d)
	}
	wantUA := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) discord/0.0.421 Chrome/152.0.7977.169 Electron/44.10.0 Safari/537.36"
	if d.BrowserUserAgent != wantUA {
		t.Fatalf("ua %s", d.BrowserUserAgent)
	}
	if w := properties.NewDesktop(properties.WindowsHost(), release(t)); w.OSSDKVersion != "26100" || w.OS != "Windows" {
		t.Fatalf("windows props %+v", w)
	}
	if !strings.Contains(string(decodeDesktop(t, d)), `"client_event_source":null`) {
		t.Fatal("client_event_source must serialize as null")
	}
}

func decodeDesktop(t *testing.T, d *properties.Desktop) []byte {
	t.Helper()
	raw, err := json.Marshal(d)
	if err != nil {
		t.Fatal(err)
	}
	return raw
}

func TestSessionAppState(t *testing.T) {
	s := properties.NewSession(properties.NewDesktop(properties.MacOSHost(), release(t)))
	before := s.Encoded()
	s.SetAppState(properties.Unfocused)
	if s.Encoded() == before || s.Properties().ClientAppState != properties.Unfocused {
		t.Fatal("app state change must re-encode")
	}
}

func TestLaunchSignatureClearsModBits(t *testing.T) {
	mods := []properties.ClientMod{
		properties.JQuery, properties.BetterDiscord, properties.Rambox, properties.Vendetta,
		properties.Vencord, properties.Replugged, properties.Legcord, properties.Dorion,
		properties.GoofCord, properties.OpenAsar, properties.Shelter, properties.Moonlight,
	}
	for range 256 {
		sig := properties.LaunchSignature()
		for _, m := range mods {
			if m.In(sig) {
				t.Fatalf("%s carries mod bit %d", sig, m)
			}
		}
	}
	if !properties.BetterDiscord.In("ffffffff-ffff-ffff-ffff-ffffffffffff") {
		t.Fatal("bit lookup broken")
	}
}

func TestAcceptLanguage(t *testing.T) {
	h := properties.MacOSHost()
	h.Languages = []string{"en-US", "en", "ro", "de", "fr", "es", "it", "pt", "nl", "pl", "sv", "fi"}
	want := "en-US,en;q=0.9,ro;q=0.8,de;q=0.7,fr;q=0.6,es;q=0.5,it;q=0.4,pt;q=0.3,nl;q=0.2,pl;q=0.1,sv;q=0.1,fi;q=0.1"
	if got := h.AcceptLanguage(); got != want {
		t.Fatalf("got  %s\nwant %s", got, want)
	}
}
