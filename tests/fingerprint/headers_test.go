package fingerprint_test

import (
	"encoding/json"
	"net/http"
	"slices"
	"testing"

	"github.com/prettylittlelies/cry/fingerprint"
	"github.com/prettylittlelies/cry/tests/testkit"
)

type blinkCase struct {
	Set  []string `json:"set"`
	Sent []string `json:"sent"`
}

func TestArrangeMatchesChromium(t *testing.T) {
	var cases []blinkCase
	if err := json.Unmarshal(testkit.Fixture(t, "fingerprint", "blink-orders.json"), &cases); err != nil {
		t.Fatal(err)
	}
	policy := fingerprint.DiscordWebPolicy()
	for i, c := range cases {
		inserted := slices.Clone(c.Set)
		if !slices.Contains(inserted, "content-type") {
			inserted = append(inserted, "content-type")
		}
		for _, n := range c.Sent {
			if !slices.Contains(inserted, n) {
				inserted = append(inserted, n)
			}
		}
		if got := policy.Arrange(inserted); !slices.Equal(got, c.Sent) {
			t.Errorf("case %d %v\n got  %v\n want %v", i, c.Set, got, c.Sent)
		}
	}
	if len(cases) < 200 {
		t.Fatalf("fixture has only %d cases", len(cases))
	}
}

func names(p fingerprint.HeaderPolicy, h http.Header) []string {
	var out []string
	for _, f := range p.Fields(h) {
		out = append(out, f.Name)
	}
	return out
}

func header(keys ...string) http.Header {
	h := http.Header{}
	for _, k := range keys {
		h.Set(k, "v")
	}
	return h
}

func TestDesktopAcceptLanguageIsPageHeader(t *testing.T) {
	h := header("Authorization", "Accept-Language", "User-Agent", "Accept-Encoding", "Priority")
	web := names(fingerprint.DiscordWebPolicy(), h)
	desktop := names(fingerprint.DiscordDesktopPolicy(), h)
	if web[len(web)-2] != "accept-language" {
		t.Errorf("web keeps accept-language trailing: %v", web)
	}
	if desktop[len(desktop)-1] != "priority" || desktop[len(desktop)-2] != "accept-encoding" {
		t.Errorf("desktop moves accept-language into page map: %v", desktop)
	}
}

func TestFieldsDropsHopByHop(t *testing.T) {
	got := names(fingerprint.DiscordWebPolicy(), header("Connection", "Te", "Authorization", "Content-Length"))
	if !slices.Equal(got, []string{"content-length", "authorization"}) {
		t.Fatalf("got %v", got)
	}
}
