package backoff

import (
	"math/rand/v2"
	"time"
)

const DefaultMin = 500 * time.Millisecond

type Backoff struct {
	Min     time.Duration
	Max     time.Duration
	Jitter  bool
	current time.Duration
	fails   int
}

func New(min, max time.Duration) *Backoff {
	if min <= 0 {
		min = DefaultMin
	}
	if max <= 0 {
		max = 10 * min
	}
	return &Backoff{Min: min, Max: max, Jitter: true, current: min}
}

func Default() *Backoff {
	return New(DefaultMin, 0)
}

func (b *Backoff) Fails() int {
	return b.fails
}

func (b *Backoff) Succeed() {
	b.fails = 0
	b.current = b.Min
}

func (b *Backoff) Fail(floor time.Duration) time.Duration {
	b.fails++
	step := 2 * float64(b.current)
	if b.Jitter {
		step *= rand.Float64()
	}
	b.current = min(b.current+time.Duration(step), b.Max)
	return max(b.current, floor)
}
