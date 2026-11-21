package cry

import (
	"log/slog"

	"github.com/prettylittlelies/cry/build"
	"github.com/prettylittlelies/cry/properties"
	"github.com/prettylittlelies/cry/rest"
)

const DefaultMessageCache = 1000

type config struct {
	host        properties.Host
	channel     build.Channel
	snapshot    *build.Snapshot
	captcha     rest.CaptchaHandler
	messages    int
	fastConnect bool
	logger      *slog.Logger
	endpoint    string
	restBase    string
}

func defaults() config {
	return config{
		host:        properties.MacOSHost(),
		channel:     build.Stable,
		messages:    DefaultMessageCache,
		fastConnect: true,
		logger:      slog.New(slog.DiscardHandler),
	}
}

type Option func(*config)

func WithHost(h properties.Host) Option {
	return func(c *config) { c.host = h }
}

func WithChannel(ch build.Channel) Option {
	return func(c *config) { c.channel = ch }
}

func WithSnapshot(s *build.Snapshot) Option {
	return func(c *config) { c.snapshot = s }
}

func WithCaptcha(h rest.CaptchaHandler) Option {
	return func(c *config) { c.captcha = h }
}

func WithMessageCache(n int) Option {
	return func(c *config) { c.messages = n }
}

func WithFastConnect(enabled bool) Option {
	return func(c *config) { c.fastConnect = enabled }
}

func WithLogger(l *slog.Logger) Option {
	return func(c *config) { c.logger = l }
}

func WithGatewayEndpoint(url string) Option {
	return func(c *config) { c.endpoint = url }
}

func WithAPIBase(url string) Option {
	return func(c *config) { c.restBase = url }
}
