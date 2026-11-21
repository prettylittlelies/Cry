package fingerprint_test

import (
	"net/http"
	"slices"
	"testing"

	"github.com/prettylittlelies/cry/fingerprint"
)

func TestParseUrgency(t *testing.T) {
	cases := map[string]fingerprint.Urgency{
		"u=1, i": 1,
		"u=4, i": 4,
		"i":      fingerprint.DefaultUrgency,
		"":       fingerprint.DefaultUrgency,
		"u=9":    fingerprint.DefaultUrgency,
	}
	for in, want := range cases {
		if got := fingerprint.ParseUrgency(in); got != want {
			t.Errorf("%q: got %d, want %d", in, got, want)
		}
	}
}

func TestChromiumWeights(t *testing.T) {
	h2 := fingerprint.Chromium152H2()
	for u, want := range map[fingerprint.Urgency]int{0: 256, 1: 220, 3: 147, 4: 110} {
		p := h2.Priority(u)
		if int(p.Weight)+1 != want || !p.Exclusive || p.StreamDep != 0 {
			t.Errorf("u=%d: got %+v, want weight %d", u, p, want)
		}
	}
}

func headerNames(o fingerprint.HeaderOrder, keys ...string) []string {
	h := http.Header{}
	for _, k := range keys {
		h.Set(k, "v")
	}
	var out []string
	for _, f := range o.Fields(h) {
		out = append(out, f.Name)
	}
	return out
}

func TestDesktopHeaderOrder(t *testing.T) {
	got := headerNames(fingerprint.DiscordDesktopOrder(),
		"Priority", "User-Agent", "X-Zeta", "Authorization", "Sec-Ch-Ua-Platform", "X-Alpha", "Connection",
		"X-Super-Properties", "Accept", "Content-Type", "X-Context-Properties", "Accept-Language", "X-Discord-Locale", "Cookie")
	want := []string{
		"sec-ch-ua-platform", "content-type", "x-alpha", "x-zeta", "x-context-properties",
		"authorization", "x-super-properties", "accept-language", "x-discord-locale",
		"user-agent", "accept", "cookie", "priority",
	}
	if !slices.Equal(got, want) {
		t.Fatalf("got  %v\nwant %v", got, want)
	}
}

func TestWebHeaderOrder(t *testing.T) {
	got := headerNames(fingerprint.DiscordWebOrder(), "Accept-Language", "Authorization", "Accept-Encoding", "X-Discord-Locale")
	want := []string{"authorization", "x-discord-locale", "accept-encoding", "accept-language"}
	if !slices.Equal(got, want) {
		t.Fatalf("got  %v\nwant %v", got, want)
	}
}
