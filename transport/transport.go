package transport

import (
	"context"
	"crypto/x509"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/prettylittlelies/cry/fingerprint"
)

const dialTimeout = 30 * time.Second

type Transport struct {
	profile fingerprint.Profile
	dialer  *fingerprint.Dialer

	mu      sync.Mutex
	conns   map[string]*conn
	pending map[string]*dial
}

type dial struct {
	done chan struct{}
	conn *conn
	err  error
}

type Option func(*Transport)

func WithNetDialer(d fingerprint.ContextDialer) Option {
	return func(t *Transport) { t.dialer.Net = d }
}

func WithRootCAs(pool *x509.CertPool) Option {
	return func(t *Transport) { t.dialer.RootCAs = pool }
}

var _ http.RoundTripper = (*Transport)(nil)

func New(p fingerprint.Profile, opts ...Option) *Transport {
	t := &Transport{
		profile: p,
		dialer:  fingerprint.NewDialer(p.Hello),
		conns:   make(map[string]*conn),
		pending: make(map[string]*dial),
	}
	for _, o := range opts {
		o(t)
	}
	return t
}

func (t *Transport) Profile() fingerprint.Profile {
	return t.profile
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.Scheme != "https" {
		return nil, ErrUnsupported
	}
	addr := authority(req.URL.Host, req.URL.Port())
	c, err := t.conn(req.Context(), addr)
	if err != nil {
		return nil, err
	}
	resp, err := c.roundTrip(req)
	if err != nil && !c.usable() {
		t.evict(addr, c)
	}
	return resp, err
}

func (t *Transport) CloseIdleConnections() {
	t.mu.Lock()
	conns := t.conns
	t.conns = make(map[string]*conn)
	t.mu.Unlock()
	for _, c := range conns {
		c.close()
	}
}

func (t *Transport) conn(ctx context.Context, addr string) (*conn, error) {
	t.mu.Lock()
	if c, ok := t.conns[addr]; ok && c.usable() {
		t.mu.Unlock()
		return c, nil
	}
	d, inflight := t.pending[addr]
	if !inflight {
		d = &dial{done: make(chan struct{})}
		t.pending[addr] = d
	}
	t.mu.Unlock()

	if !inflight {
		dctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), dialTimeout)
		d.conn, d.err = t.open(dctx, addr)
		cancel()
		t.mu.Lock()
		delete(t.pending, addr)
		if d.err == nil {
			t.conns[addr] = d.conn
		}
		t.mu.Unlock()
		close(d.done)
	}
	select {
	case <-d.done:
		return d.conn, d.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (t *Transport) open(ctx context.Context, addr string) (*conn, error) {
	tc, err := t.dialer.DialTLS(ctx, "tcp", addr)
	if err != nil {
		return nil, err
	}
	if tc.ConnectionState().NegotiatedProtocol != "h2" {
		tc.Close()
		return nil, ErrNotH2
	}
	c, err := newConn(tc, t.profile)
	if err != nil {
		tc.Close()
		return nil, err
	}
	return c, nil
}

func (t *Transport) evict(addr string, c *conn) {
	t.mu.Lock()
	if t.conns[addr] == c {
		delete(t.conns, addr)
	}
	t.mu.Unlock()
}

func authority(host, port string) string {
	if port != "" {
		return host
	}
	return net.JoinHostPort(host, "443")
}
