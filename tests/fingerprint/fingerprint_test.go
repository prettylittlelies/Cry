package fingerprint_test

import (
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
