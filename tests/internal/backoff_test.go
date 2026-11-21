package internal_test

import (
	"testing"
	"time"

	"github.com/prettylittlelies/cry/internal/backoff"
)

func TestBackoffMatchesDiscord(t *testing.T) {
	b := backoff.Default()
	if b.Min != 500*time.Millisecond || b.Max != 5*time.Second {
		t.Fatalf("defaults min=%s max=%s", b.Min, b.Max)
	}
	b.Jitter = false
	want := []time.Duration{1500 * time.Millisecond, 4500 * time.Millisecond, 5 * time.Second, 5 * time.Second}
	for i, w := range want {
		if got := b.Fail(0); got != w {
			t.Fatalf("fail %d: got %s, want %s", i, got, w)
		}
	}
	if got := b.Fail(9 * time.Second); got != 9*time.Second {
		t.Fatalf("floor not honored: %s", got)
	}
	b.Succeed()
	if b.Fails() != 0 || b.Fail(0) != 1500*time.Millisecond {
		t.Fatal("succeed must reset")
	}
}

func TestBackoffJitterBounded(t *testing.T) {
	b := backoff.Default()
	for range 100 {
		if d := b.Fail(0); d < b.Min || d > b.Max {
			t.Fatalf("out of range %s", d)
		}
	}
}
