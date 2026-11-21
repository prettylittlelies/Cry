package cry

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"sync/atomic"

	"github.com/prettylittlelies/cry/build"
	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/fingerprint"
	"github.com/prettylittlelies/cry/gateway"
	"github.com/prettylittlelies/cry/properties"
	"github.com/prettylittlelies/cry/rest"
	"github.com/prettylittlelies/cry/state"
	"github.com/prettylittlelies/cry/transport"
	"github.com/prettylittlelies/cry/transport/websocket"
	"golang.org/x/net/publicsuffix"
)

var ErrUnknownChannel = errors.New("cry: release not present in build snapshot")

type Client struct {
	cfg       config
	transport *transport.Transport
	session   *properties.Session
	rest      *rest.Client
	gateway   *gateway.Gateway
	state     *state.State
	events    *registry
	runCtx    atomic.Pointer[context.Context]
}

func New(ctx context.Context, token string, opts ...Option) (*Client, error) {
	cfg := defaults()
	for _, o := range opts {
		o(&cfg)
	}
	tr := transport.New(fingerprint.Electron44())
	snap := cfg.snapshot
	if snap == nil {
		var err error
		snap, err = build.NewScraper(build.WithDoer(&http.Client{Transport: tr}), build.WithChannels(cfg.channel)).Fetch(ctx)
		if err != nil {
			return nil, err
		}
	}
	release, ok := snap.Release(cfg.channel)
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrUnknownChannel, cfg.channel)
	}
	session := properties.NewSession(properties.NewDesktop(cfg.host, properties.MacOSRelease(release)))
	c := &Client{
		cfg:       cfg,
		transport: tr,
		session:   session,
		state:     state.New(cfg.messages),
		events:    newRegistry(),
	}
	restOpts := []rest.Option{}
	if cfg.captcha != nil {
		restOpts = append(restOpts, rest.WithCaptcha(cfg.captcha))
	}
	if cfg.restBase != "" {
		restOpts = append(restOpts, rest.WithBaseURL(cfg.restBase))
	}
	c.rest = rest.New(token, tr, rest.Identity{Session: session, Host: cfg.host, Chrome: release.MacOS.Chrome}, restOpts...)
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	c.gateway = gateway.New(gateway.Config{
		Token:       token,
		Session:     session,
		Jar:         jar,
		Endpoint:    cfg.endpoint,
		FastConnect: cfg.fastConnect,
		Logger:      cfg.logger,
		Browser:     websocket.Browser{UserAgent: session.UserAgent(), Origin: rest.DefaultOrigin, AcceptLanguage: cfg.host.Locale},
		Handler:     c.handle,
	})
	return c, nil
}

func (c *Client) REST() *rest.Client {
	return c.rest
}

func (c *Client) State() *state.State {
	return c.state
}

func (c *Client) Gateway() *gateway.Gateway {
	return c.gateway
}

func (c *Client) Session() *properties.Session {
	return c.session
}

func (c *Client) Run(ctx context.Context) error {
	c.runCtx.Store(&ctx)
	defer c.transport.CloseIdleConnections()
	return c.gateway.Run(ctx)
}

func (c *Client) context() context.Context {
	if p := c.runCtx.Load(); p != nil {
		return *p
	}
	return context.Background()
}

func (c *Client) handle(e gateway.Event) {
	ctx := c.context()
	newEvent, ok := discord.Events[e.Type]
	if !ok {
		c.events.emit(ctx, &RawEvent{Type: e.Type, Seq: e.Seq, Data: e.Data})
		return
	}
	v := newEvent()
	if err := json.Unmarshal(e.Data, v); err != nil {
		c.cfg.logger.Warn("event decode failed", "type", e.Type, "err", err)
		c.events.emit(ctx, &RawEvent{Type: e.Type, Seq: e.Seq, Data: e.Data})
		return
	}
	c.state.Apply(v)
	c.events.emit(ctx, v)
}
