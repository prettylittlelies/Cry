package fingerprint_test

import (
	"strings"
	"testing"

	"github.com/prettylittlelies/cry/fingerprint"
)

func TestSecCHUAMatchesElectron44(t *testing.T) {
	want := ""
	for _, l := range reference(t, "electron44-renderer.json").HTTP2.Frames[2].Headers {
		if v, ok := strings.CutPrefix(l, "sec-ch-ua: "); ok {
			want = v
		}
	}
	if got := fingerprint.SecCHUA(fingerprint.Brands(152, "")); got != want {
		t.Fatalf("got  %s\nwant %s", got, want)
	}
}

func TestChromeBrandsIncludeProduct(t *testing.T) {
	got := fingerprint.SecCHUA(fingerprint.Brands(152, "Google Chrome"))
	for _, b := range []string{`"Chromium";v="152"`, `"Google Chrome";v="152"`, `"Not?A_Brand";v="24"`} {
		if !strings.Contains(got, b) {
			t.Fatalf("%s missing %s", got, b)
		}
	}
}
