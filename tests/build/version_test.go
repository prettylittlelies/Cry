package build_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/prettylittlelies/cry/build"
)

func errorsAs[T error](err error, target *T) bool {
	return errors.As(err, target)
}

func TestVersionJSON(t *testing.T) {
	for _, in := range []string{`[1,0,9269]`, `"1.0.9269"`} {
		var v build.Version
		if err := json.Unmarshal([]byte(in), &v); err != nil {
			t.Fatalf("%s: %v", in, err)
		}
		if v.String() != "1.0.9269" {
			t.Errorf("%s: got %s", in, v)
		}
	}
}

func TestVersionCompare(t *testing.T) {
	a, _ := build.ParseVersion("44.10.2")
	b, _ := build.ParseVersion("44.11.5")
	if a.Compare(b) >= 0 || a.Major() != 44 {
		t.Errorf("compare %s %s", a, b)
	}
}

func TestVersionInvalid(t *testing.T) {
	_, err := build.ParseVersion("1.x")
	var pe *build.ParseError
	if !errorsAs(err, &pe) {
		t.Fatalf("want *build.ParseError, got %v", err)
	}
}

func TestAndroidUserAgent(t *testing.T) {
	m, err := build.ParseMobileVersion("280.2")
	if err != nil {
		t.Fatal(err)
	}
	if got := m.AndroidUserAgent(build.TrackAlpha); got != "Discord-Android/280202;RNA" {
		t.Errorf("got %s", got)
	}
}
