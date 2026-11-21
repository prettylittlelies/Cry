package gateway

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/prettylittlelies/cry/fingerprint"
	"github.com/prettylittlelies/cry/gateway/codec"
	"github.com/prettylittlelies/cry/internal/backoff"
	"github.com/prettylittlelies/cry/properties"
	"github.com/prettylittlelies/cry/transport/websocket"
)

const (
	DefaultEndpoint = "wss://gateway.discord.gg"
	APIVersion      = 9
	helloTimeout    = 30 * time.Second
	resumeWindow    = 3 * time.Minute
	heartbeatGrace  = 5 * time.Second
	maxSessionFails = 4
)

var ErrAuthenticationFailed = errors.New("gateway: authentication failed")

type Config struct {
	Token          string
	Session        *properties.Session
	Browser        websocket.Browser
	Hello          fingerprint.ClientHello
	Jar            http.CookieJar
	Endpoint       string
	Encoding       codec.Encoding
	Compression    codec.Compression
	Capabilities   Capabilities
	Presence       Presence
	InstallationID string
	FastConnect    bool
	Handler        func(Event)
	Logger         *slog.Logger
}

func (c *Config) defaults() {
	if c.Endpoint == "" {
		c.Endpoint = DefaultEndpoint
	}
	if c.Encoding == "" {
		c.Encoding = codec.ETF
	}
	if c.Compression == codec.None && c.Encoding == codec.ETF {
		c.Compression = codec.ZstdStream
	}
	if c.Capabilities == 0 {
		c.Capabilities = DefaultCapabilities
	}
	if c.Presence.Status == "" {
		c.Presence = DefaultPresence()
	}
	if c.Hello == nil {
		c.Hello = fingerprint.Chromium152HTTP1
	}
	if c.Handler == nil {
		c.Handler = func(Event) {}
	}
	if c.Logger == nil {
		c.Logger = slog.New(slog.DiscardHandler)
	}
}

type Gateway struct {
	cfg    Config
	dialer *fingerprint.Dialer
	send   chan outbound

	mu         sync.Mutex
	seq        int64
	sessionID  string
	resumeURL  string
	derivedQoS string
	lastAck    time.Time
	identified int
	reasons    []string
	qos        qosState
	connected  bool
	wake       chan struct{}
	timeSpent  timeSpent
	presence   Presence
}

func New(cfg Config) *Gateway {
	cfg.defaults()
	g := &Gateway{
		cfg:     cfg,
		dialer:  websocket.NewDialer(cfg.Hello),
		send:    make(chan outbound, 64),
		reasons: []string{ColdStart},
		wake:    make(chan struct{}, 1),
	}
	g.presence = cfg.Presence
	g.qos.change(initialQoS(cfg.Session.Properties().ClientAppState))
	return g
}

func initialQoS(state properties.AppState) QoS {
	if state == properties.Focused {
		return NewQoS(Foregrounded)
	}
	return NewQoS()
}

func (g *Gateway) Dialer() *fingerprint.Dialer {
	return g.dialer
}

func (g *Gateway) Send(ctx context.Context, op Opcode, d any) error {
	select {
	case g.send <- outbound{Op: op, D: d}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (g *Gateway) SetActivity(q QoS) {
	g.mu.Lock()
	wake := g.qos.change(q) && g.connected
	g.mu.Unlock()
	if wake {
		select {
		case g.wake <- struct{}{}:
		default:
		}
	}
}

func (g *Gateway) url() string {
	g.mu.Lock()
	base := g.resumeURL
	g.mu.Unlock()
	if base == "" {
		base = g.cfg.Endpoint
	}
	u, err := url.Parse(base)
	if err != nil {
		u, _ = url.Parse(g.cfg.Endpoint)
	}
	if u.Path == "" {
		u.Path = "/"
	}
	raw := "encoding=" + string(g.cfg.Encoding) + "&v=" + strconv.Itoa(APIVersion)
	if g.cfg.Compression != codec.None {
		raw += "&compress=" + url.QueryEscape(string(g.cfg.Compression))
	}
	u.RawQuery = raw
	return u.String()
}

func (g *Gateway) Run(ctx context.Context) error {
	wait := backoff.New(time.Second, time.Minute)
	for {
		err := g.connect(ctx)
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if errors.Is(err, ErrAuthenticationFailed) {
			return err
		}
		var re *reconnect
		if errors.As(err, &re) && re.immediate {
			continue
		}
		if errors.As(err, &re) && re.established {
			wait.Succeed()
		}
		delay := wait.Fail(0)
		if wait.Fails() > maxSessionFails {
			g.resetSession()
		}
		g.cfg.Logger.Info("gateway closed", "err", err, "retry_in", delay)
		if err := sleep(ctx, delay); err != nil {
			return err
		}
	}
}

func (g *Gateway) resetSession() {
	g.mu.Lock()
	g.sessionID, g.seq = "", 0
	g.mu.Unlock()
}

type reconnect struct {
	immediate   bool
	established bool
	cause       error
}

func (r *reconnect) Error() string {
	return fmt.Sprintf("gateway: reconnect (immediate=%v): %v", r.immediate, r.cause)
}

func (r *reconnect) Unwrap() error {
	return r.cause
}

type frame struct {
	p   inbound
	err error
}

func (g *Gateway) connect(ctx context.Context) error {
	dctx, cancel := context.WithTimeout(ctx, helloTimeout)
	conn, err := websocket.Dial(dctx, g.url(), websocket.Config{Dialer: g.dialer, Browser: g.cfg.Browser, Jar: g.cfg.Jar})
	cancel()
	if err != nil {
		g.clearResumeURL()
		return err
	}
	reader, err := codec.NewReader(conn, g.cfg.Encoding, g.cfg.Compression)
	if err != nil {
		conn.Close(1000, "")
		return err
	}
	frames := make(chan frame, 16)
	go func() {
		defer close(frames)
		for {
			raw, err := reader.Next()
			if err != nil {
				frames <- frame{err: err}
				return
			}
			var p inbound
			if err := json.Unmarshal(raw, &p); err != nil {
				frames <- frame{err: err}
				return
			}
			frames <- frame{p: p}
		}
	}()
	s := &session{g: g, conn: conn, ctx: ctx}
	err = s.loop(frames)
	conn.Close(s.closeCode, "")
	for range frames {
	}
	reader.Close()
	g.mu.Lock()
	g.connected = false
	g.mu.Unlock()
	return err
}

func (g *Gateway) clearResumeURL() {
	g.mu.Lock()
	g.resumeURL = ""
	g.mu.Unlock()
}

func sleep(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-t.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func jitter(d time.Duration) time.Duration {
	return time.Duration(rand.Int64N(int64(d)))
}

func trimSlash(s string) string {
	return strings.TrimSuffix(s, "/")
}
