package transport

import (
	"bytes"
	"io"
	"sync"
)

type body struct {
	mu       sync.Mutex
	cond     *sync.Cond
	buf      bytes.Buffer
	err      error
	consumed func(int)
	closed   func()
}

func newBody(consumed func(int), closed func()) *body {
	b := &body{consumed: consumed, closed: closed}
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (b *body) write(p []byte) {
	b.mu.Lock()
	if b.err == nil {
		b.buf.Write(p)
	}
	b.mu.Unlock()
	b.cond.Broadcast()
}

func (b *body) fail(err error) {
	b.mu.Lock()
	if b.err == nil {
		b.err = err
	}
	b.mu.Unlock()
	b.cond.Broadcast()
}

func (b *body) Read(p []byte) (int, error) {
	b.mu.Lock()
	for b.buf.Len() == 0 && b.err == nil {
		b.cond.Wait()
	}
	if b.buf.Len() == 0 {
		err := b.err
		b.mu.Unlock()
		return 0, err
	}
	n, _ := b.buf.Read(p)
	b.mu.Unlock()
	b.consumed(n)
	return n, nil
}

func (b *body) Close() error {
	b.mu.Lock()
	open := b.err == nil
	if open {
		b.err = io.ErrClosedPipe
	}
	b.buf.Reset()
	b.mu.Unlock()
	b.cond.Broadcast()
	if open {
		b.closed()
	}
	return nil
}
