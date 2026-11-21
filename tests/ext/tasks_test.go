package ext_test

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"github.com/prettylittlelies/cry/ext/tasks"
)

func TestCountAndHooks(t *testing.T) {
	var runs, before, after atomic.Int32
	l := tasks.Every(5*time.Millisecond, func(context.Context) error { runs.Add(1); return nil }).
		Count(3).
		Before(func(context.Context) { before.Add(1) }).
		After(func(context.Context) { after.Add(1) })
	if err := l.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
	if err := l.Start(context.Background()); !errors.Is(err, tasks.ErrRunning) {
		t.Fatalf("second start: %v", err)
	}
	l.Wait()
	if runs.Load() != 3 || before.Load() != 1 || after.Load() != 1 || l.Running() || l.Iteration() != 3 {
		t.Fatalf("runs=%d before=%d after=%d running=%v iteration=%d", runs.Load(), before.Load(), after.Load(), l.Running(), l.Iteration())
	}
}

func TestStopCancelRestart(t *testing.T) {
	var runs atomic.Int32
	l := tasks.Every(time.Hour, func(context.Context) error { runs.Add(1); return nil })
	l.Start(context.Background())
	time.Sleep(20 * time.Millisecond)
	l.Cancel()
	l.Wait()
	if runs.Load() != 1 || l.Running() {
		t.Fatalf("cancel: runs=%d running=%v", runs.Load(), l.Running())
	}
	if err := l.Restart(context.Background()); err != nil {
		t.Fatal(err)
	}
	time.Sleep(20 * time.Millisecond)
	if runs.Load() != 2 || !l.Running() {
		t.Fatalf("restart: runs=%d running=%v", runs.Load(), l.Running())
	}
	l.Stop()
	l.SetInterval(time.Millisecond)
	l.Wait()
	if l.Running() {
		t.Fatal("stop should end the loop after the current iteration")
	}
}

func TestSetIntervalWakesSleep(t *testing.T) {
	var runs atomic.Int32
	l := tasks.Every(time.Hour, func(context.Context) error { runs.Add(1); return nil }).Count(2)
	l.Start(context.Background())
	time.Sleep(20 * time.Millisecond)
	l.SetInterval(5 * time.Millisecond)
	done := make(chan struct{})
	go func() { l.Wait(); close(done) }()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("SetInterval did not shorten the pending sleep")
	}
	if runs.Load() != 2 {
		t.Fatalf("runs=%d", runs.Load())
	}
}

func TestErrors(t *testing.T) {
	boom := errors.New("boom")
	var handled atomic.Int32
	l := tasks.Every(time.Millisecond, func(context.Context) error { return boom }).
		Count(3).
		OnError(func(_ context.Context, err error) {
			if errors.Is(err, boom) {
				handled.Add(1)
			}
		})
	l.Start(context.Background())
	l.Wait()
	if handled.Load() != 3 {
		t.Fatalf("handled %d", handled.Load())
	}
	var runs atomic.Int32
	unhandled := tasks.Every(time.Millisecond, func(context.Context) error { runs.Add(1); return boom })
	unhandled.Start(context.Background())
	unhandled.Wait()
	if runs.Load() != 1 {
		t.Fatalf("unhandled error should stop the loop, ran %d times", runs.Load())
	}
}

func TestTimeOfDaySchedule(t *testing.T) {
	var fired atomic.Int32
	soon := time.Now().UTC().Add(1500 * time.Millisecond)
	l := tasks.At([]tasks.TimeOfDay{
		{Hour: (soon.Hour() + 12) % 24, Minute: soon.Minute(), Second: soon.Second()},
		{Hour: soon.Hour(), Minute: soon.Minute(), Second: soon.Second()},
	}, func(context.Context) error { fired.Add(1); return nil }).Count(2)
	l.Start(context.Background())
	time.Sleep(50 * time.Millisecond)
	if fired.Load() != 0 {
		t.Fatal("time-based loop must wait for its first scheduled time")
	}
	if next := l.Next(); next.Sub(soon) > time.Second || soon.Sub(next) > time.Second {
		t.Fatalf("next run %v, expected near %v", next, soon)
	}
	l.Cancel()
	l.Wait()
}
