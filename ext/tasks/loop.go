package tasks

import (
	"context"
	"errors"
	"sync"
	"time"
)

var ErrRunning = errors.New("tasks: loop already running")

type Func func(ctx context.Context) error

type Loop struct {
	fn      Func
	before  func(ctx context.Context)
	after   func(ctx context.Context)
	onError func(ctx context.Context, err error)
	count   int
	now     func() time.Time

	mu        sync.Mutex
	schedule  schedule
	running   bool
	stopping  bool
	cancel    context.CancelFunc
	done      chan struct{}
	iteration int
	nextAt    time.Time
	reset     chan struct{}
}

func Every(d time.Duration, fn Func) *Loop {
	return newLoop(interval(d), fn)
}

func At(times []TimeOfDay, fn Func) *Loop {
	return newLoop(daily(times), fn)
}

func newLoop(s schedule, fn Func) *Loop {
	return &Loop{fn: fn, schedule: s, now: time.Now, reset: make(chan struct{}, 1)}
}

func (l *Loop) Count(n int) *Loop {
	l.count = n
	return l
}

func (l *Loop) Before(fn func(ctx context.Context)) *Loop {
	l.before = fn
	return l
}

func (l *Loop) After(fn func(ctx context.Context)) *Loop {
	l.after = fn
	return l
}

func (l *Loop) OnError(fn func(ctx context.Context, err error)) *Loop {
	l.onError = fn
	return l
}

func (l *Loop) Start(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.running {
		return ErrRunning
	}
	ctx, cancel := context.WithCancel(ctx)
	l.running, l.stopping, l.iteration = true, false, 0
	l.cancel, l.done = cancel, make(chan struct{})
	go l.run(ctx, l.done)
	return nil
}

func (l *Loop) Stop() {
	l.mu.Lock()
	l.stopping = true
	l.mu.Unlock()
}

func (l *Loop) Cancel() {
	l.mu.Lock()
	cancel := l.cancel
	l.mu.Unlock()
	if cancel != nil {
		cancel()
	}
}

func (l *Loop) Wait() {
	l.mu.Lock()
	done := l.done
	l.mu.Unlock()
	if done != nil {
		<-done
	}
}

func (l *Loop) Restart(ctx context.Context) error {
	l.Cancel()
	l.Wait()
	return l.Start(ctx)
}

func (l *Loop) Running() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.running
}

func (l *Loop) Iteration() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.iteration
}

func (l *Loop) Next() time.Time {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.nextAt
}

func (l *Loop) SetInterval(d time.Duration) {
	l.mu.Lock()
	l.schedule = interval(d)
	l.mu.Unlock()
	select {
	case l.reset <- struct{}{}:
	default:
	}
}

func (l *Loop) run(ctx context.Context, done chan struct{}) {
	defer func() {
		if l.after != nil {
			l.after(context.WithoutCancel(ctx))
		}
		l.mu.Lock()
		l.running, l.cancel = false, nil
		l.mu.Unlock()
		close(done)
	}()
	if l.before != nil {
		l.before(ctx)
	}
	if _, timed := l.schedule.(daily); timed {
		l.mu.Lock()
		l.nextAt = l.schedule.next(l.now())
		wait := time.Until(l.nextAt)
		l.mu.Unlock()
		if !l.sleep(ctx, wait) {
			return
		}
	}
	for {
		if err := l.fn(ctx); err != nil && ctx.Err() == nil {
			if l.onError == nil {
				return
			}
			l.onError(ctx, err)
		}
		l.mu.Lock()
		l.iteration++
		finished := l.stopping || (l.count > 0 && l.iteration >= l.count)
		l.nextAt = l.schedule.next(l.now())
		wait := time.Until(l.nextAt)
		l.mu.Unlock()
		if finished || ctx.Err() != nil {
			return
		}
		if !l.sleep(ctx, wait) {
			return
		}
	}
}

func (l *Loop) sleep(ctx context.Context, d time.Duration) bool {
	for {
		t := time.NewTimer(d)
		select {
		case <-t.C:
			return true
		case <-ctx.Done():
			t.Stop()
			return false
		case <-l.reset:
			t.Stop()
			l.mu.Lock()
			l.nextAt = l.schedule.next(l.now())
			d = time.Until(l.nextAt)
			l.mu.Unlock()
		}
	}
}
