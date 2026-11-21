package fingerprint_test

import (
	"encoding/json"
	"testing"

	"github.com/prettylittlelies/cry/tests/testkit"
)

type capture struct {
	HTTP2 struct {
		Frames []struct {
			Type    string   `json:"frame_type"`
			Headers []string `json:"headers"`
		} `json:"sent_frames"`
	} `json:"http2"`
}

func reference(t *testing.T, name string) capture {
	t.Helper()
	var c capture
	if err := json.Unmarshal(testkit.Fixture(t, "fingerprint", name), &c); err != nil {
		t.Fatal(err)
	}
	return c
}
