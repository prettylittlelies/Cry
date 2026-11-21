package fingerprint_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"net/http"
	"slices"
	"strings"
	"testing"

	"github.com/prettylittlelies/cry/fingerprint"
	"github.com/prettylittlelies/cry/tests/testkit"
)

type labRecord struct {
	Method  string   `json:"method"`
	Headers []string `json:"headers"`
}

func labRecords(t *testing.T, name string) []labRecord {
	t.Helper()
	var out []labRecord
	sc := bufio.NewScanner(bytes.NewReader(testkit.Fixture(t, "fingerprint", name)))
	sc.Buffer(nil, 1<<20)
	for sc.Scan() {
		var r labRecord
		if err := json.Unmarshal(sc.Bytes(), &r); err != nil {
			t.Fatal(err)
		}
		out = append(out, r)
	}
	return out
}

func split(lines []string) (http.Header, []string) {
	h := http.Header{}
	var names []string
	for _, l := range lines {
		k, v, _ := strings.Cut(l, ": ")
		if strings.HasPrefix(k, ":") {
			continue
		}
		h.Add(k, v)
		names = append(names, k)
	}
	return h, names
}

func TestUploadOrdersMatchChromium(t *testing.T) {
	records := labRecords(t, "electron44-upload.jsonl")
	if len(records) != 4 {
		t.Fatalf("fixture has %d records", len(records))
	}
	policy := fingerprint.DiscordWebPolicy()
	for i, r := range records {
		h, want := split(r.Headers)
		var got []string
		for _, f := range policy.Fields(r.Method, h) {
			got = append(got, f.Name)
		}
		if !slices.Equal(got, want) {
			t.Errorf("record %d %s\n got  %v\n want %v", i, r.Method, got, want)
		}
	}
}
