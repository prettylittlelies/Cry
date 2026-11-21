package rest

import (
	"context"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	bucketIdle = 5 * time.Minute
	pruneEvery = time.Minute
)

type bucket struct {
	mu        sync.Mutex
	limit     int
	remaining int
	resetAt   time.Time
	usedAt    time.Time
}

func (b *bucket) idle(now time.Time) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return now.Sub(b.usedAt) >= bucketIdle && !now.Before(b.resetAt)
}

func (b *bucket) acquire(ctx context.Context, now func() time.Time) error {
	for {
		b.mu.Lock()
		t := now()
		b.usedAt = t
		if !b.resetAt.IsZero() && !t.Before(b.resetAt) {
			b.remaining, b.resetAt = b.limit, time.Time{}
		}
		if b.limit == 0 || b.remaining > 0 {
			if b.limit > 0 {
				b.remaining--
			}
			b.mu.Unlock()
			return nil
		}
		wait := b.resetAt.Sub(t)
		b.mu.Unlock()
		if err := sleep(ctx, wait); err != nil {
			return err
		}
	}
}

func (b *bucket) update(h http.Header, now time.Time) {
	limit, err := strconv.Atoi(h.Get("X-RateLimit-Limit"))
	if err != nil {
		return
	}
	remaining, _ := strconv.Atoi(h.Get("X-RateLimit-Remaining"))
	after, _ := strconv.ParseFloat(h.Get("X-RateLimit-Reset-After"), 64)
	b.mu.Lock()
	b.limit, b.remaining = limit, remaining
	b.resetAt = now.Add(seconds(after))
	b.mu.Unlock()
}

func (b *bucket) exhaust(until time.Time) {
	b.mu.Lock()
	if b.limit == 0 {
		b.limit = 1
	}
	b.remaining = 0
	if until.After(b.resetAt) {
		b.resetAt = until
	}
	b.mu.Unlock()
}

type limiter struct {
	mu      sync.Mutex
	hashes  map[string]string
	buckets map[string]*bucket
	global  time.Time
	pruned  time.Time
	now     func() time.Time
}

func newLimiter() *limiter {
	return &limiter{hashes: make(map[string]string), buckets: make(map[string]*bucket), now: time.Now}
}

func (l *limiter) bucket(r Route) *bucket {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.prune()
	id := r.Key()
	if h, ok := l.hashes[id]; ok {
		id = h
	}
	id += ":" + r.Major()
	b, ok := l.buckets[id]
	if !ok {
		b = &bucket{usedAt: l.now()}
		l.buckets[id] = b
	}
	return b
}

func (l *limiter) prune() {
	now := l.now()
	if now.Sub(l.pruned) < pruneEvery {
		return
	}
	l.pruned = now
	for id, b := range l.buckets {
		if b.idle(now) {
			delete(l.buckets, id)
		}
	}
}

func (l *limiter) size() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return len(l.buckets)
}

func (l *limiter) learn(r Route, b *bucket, hash string) {
	if hash == "" {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.hashes[r.Key()] == hash {
		return
	}
	l.hashes[r.Key()] = hash
	l.buckets[hash+":"+r.Major()] = b
}

func (l *limiter) waitGlobal(ctx context.Context) error {
	l.mu.Lock()
	wait := l.global.Sub(l.now())
	l.mu.Unlock()
	return sleep(ctx, wait)
}

func (l *limiter) blockGlobal(d time.Duration) {
	l.mu.Lock()
	if until := l.now().Add(d); until.After(l.global) {
		l.global = until
	}
	l.mu.Unlock()
}

func seconds(f float64) time.Duration {
	return time.Duration(f * float64(time.Second))
}

func sleep(ctx context.Context, d time.Duration) error {
	if d <= 0 {
		return ctx.Err()
	}
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-t.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
