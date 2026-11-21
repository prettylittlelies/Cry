package cry

import (
	"context"
	"encoding/json"
	"reflect"
	"sync"
)

type RawEvent struct {
	Type string
	Seq  int64
	Data json.RawMessage
}

type listener struct {
	id   uint64
	call func(context.Context, any)
}

type registry struct {
	mu        sync.RWMutex
	next      uint64
	listeners map[reflect.Type][]listener
}

func newRegistry() *registry {
	return &registry{listeners: map[reflect.Type][]listener{}}
}

func (r *registry) add(t reflect.Type, call func(context.Context, any)) func() {
	r.mu.Lock()
	r.next++
	id := r.next
	r.listeners[t] = append(r.listeners[t], listener{id: id, call: call})
	r.mu.Unlock()
	var once sync.Once
	return func() {
		once.Do(func() {
			r.mu.Lock()
			defer r.mu.Unlock()
			ls := r.listeners[t]
			for i, l := range ls {
				if l.id == id {
					r.listeners[t] = append(ls[:i:i], ls[i+1:]...)
					return
				}
			}
		})
	}
}

func (r *registry) emit(ctx context.Context, v any) {
	r.mu.RLock()
	ls := append([]listener(nil), r.listeners[reflect.TypeOf(v)]...)
	r.mu.RUnlock()
	for _, l := range ls {
		go l.call(ctx, v)
	}
}

func On[T any](c *Client, fn func(context.Context, *T)) (remove func()) {
	return c.events.add(reflect.TypeFor[*T](), func(ctx context.Context, v any) {
		fn(ctx, v.(*T))
	})
}

func WaitFor[T any](ctx context.Context, c *Client, check func(*T) bool) (*T, error) {
	found := make(chan *T, 1)
	remove := On(c, func(_ context.Context, v *T) {
		if check == nil || check(v) {
			select {
			case found <- v:
			default:
			}
		}
	})
	defer remove()
	select {
	case v := <-found:
		return v, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (c *Client) Dispatch(ctx context.Context, v any) {
	c.events.emit(ctx, v)
}
