package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/publicsuffix"
)

const (
	DefaultBaseURL = "https://discord.com/api/v9"
	maxAttempts    = 5
)

var (
	retryStatuses     = map[int]struct{}{502: {}, 504: {}, 507: {}, 522: {}, 523: {}, 524: {}}
	cloudflareCodeExp = regexp.MustCompile(`error code: (\d+)`)
)

type Option func(*Client)

func WithBaseURL(u string) Option {
	return func(c *Client) { c.base = strings.TrimSuffix(u, "/") }
}

func WithCaptcha(h CaptchaHandler) Option {
	return func(c *Client) { c.captcha = h }
}

func WithMaxRetryAfter(d time.Duration) Option {
	return func(c *Client) { c.maxRetryAfter = d }
}

func WithReferer(r string) Option {
	return func(c *Client) { c.browser.referer = r }
}

func WithDebugOptions(opts ...string) Option {
	return func(c *Client) { c.browser.debug = strings.Join(opts, ",") }
}

func WithClock(now func() time.Time) Option {
	return func(c *Client) { c.limiter.now = now }
}

func WithInstallationID(id string) Option {
	return func(c *Client) { c.installation = id }
}

type Client struct {
	http          *http.Client
	browser       *browser
	limiter       *limiter
	token         string
	installation  string
	base          string
	captcha       CaptchaHandler
	maxRetryAfter time.Duration
}

func New(token string, rt http.RoundTripper, id Identity, opts ...Option) *Client {
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	c := &Client{
		http:    &http.Client{Transport: rt, Jar: jar},
		browser: newBrowser(id),
		limiter: newLimiter(),
		token:   token,
		base:    DefaultBaseURL,
	}
	for _, o := range opts {
		o(c)
	}
	return c
}

type Stats struct {
	Buckets int
}

func (c *Client) Stats() Stats {
	return Stats{Buckets: c.limiter.size()}
}

func (c *Client) SetToken(token string) {
	c.token = token
}

func (c *Client) Do(ctx context.Context, req *Request) (*Response, error) {
	body, contentType, err := req.payload()
	if err != nil {
		return nil, err
	}
	contextHeader, err := req.contextHeader()
	if err != nil {
		return nil, err
	}
	query := newQuery(req)
	extra := req.Header.Clone()
	if extra == nil {
		extra = http.Header{}
	}
	b := c.limiter.bucket(req.Route)
	failed := 0
	var last *Response
	for attempt := range maxAttempts {
		if err := c.limiter.waitGlobal(ctx); err != nil {
			return nil, err
		}
		if err := b.acquire(ctx, c.limiter.now); err != nil {
			return nil, err
		}
		h := c.compose(req, contentType, contextHeader, extra, failed)
		resp, err := c.send(ctx, req.Route.Method, c.base+req.Route.Path()+query.encode(), body, h)
		if err != nil {
			if attempt < maxAttempts-1 && retryableNetErr(err) && ctx.Err() == nil {
				failed++
				if err := sleep(ctx, backoff(attempt)); err != nil {
					return nil, err
				}
				continue
			}
			return nil, err
		}
		last = resp
		c.limiter.learn(req.Route, b, resp.Header.Get("X-RateLimit-Bucket"))
		if resp.Status != http.StatusTooManyRequests {
			b.update(resp.Header, c.limiter.now())
		}
		switch {
		case resp.Status == http.StatusAccepted && retryLater(resp.Body) > 0:
			query.bump()
			if err := sleep(ctx, retryLater(resp.Body)); err != nil {
				return nil, err
			}
			continue
		case resp.Status >= 200 && resp.Status < 300:
			return resp, nil
		case resp.Status == http.StatusTooManyRequests:
			rl, err := c.rateLimited(resp)
			if err != nil {
				return nil, err
			}
			if rl.Global {
				c.limiter.blockGlobal(rl.RetryAfter)
			} else {
				b.exhaust(c.limiter.now().Add(rl.RetryAfter))
			}
			if err := sleep(ctx, rl.RetryAfter); err != nil {
				return nil, err
			}
			continue
		case isRetryStatus(resp.Status):
			failed++
			if err := sleep(ctx, backoff(attempt)); err != nil {
				return nil, err
			}
			continue
		}
		herr := newHTTPError(req.Route.Method, req.Route.Path(), resp.Status, resp.Body)
		captcha := parseCaptcha(resp.Status, resp.Body)
		if captcha == nil {
			return nil, herr
		}
		cerr := &CaptchaRequiredError{HTTPError: herr, Captcha: captcha}
		if c.captcha == nil || attempt == maxAttempts-1 {
			return nil, cerr
		}
		key, err := c.captcha.Solve(ctx, captcha)
		if err != nil {
			return nil, err
		}
		captcha.apply(extra, key)
	}
	return nil, newHTTPError(req.Route.Method, req.Route.Path(), last.Status, last.Body)
}

func (c *Client) compose(req *Request, contentType, contextHeader string, extra http.Header, failed int) http.Header {
	h := http.Header{}
	if contentType != "" {
		h.Set("Content-Type", contentType)
	}
	for k, v := range extra {
		h[k] = v
	}
	if req.Reason != "" {
		h.Set("X-Audit-Log-Reason", encodeURIComponent(req.Reason))
	}
	if contextHeader != "" {
		h.Set("X-Context-Properties", contextHeader)
	}
	if failed > 0 {
		h.Set("X-Failed-Requests", strconv.Itoa(failed))
	}
	if !req.NoAuth && c.token != "" {
		h.Set("Authorization", c.token)
	}
	if c.installation != "" {
		h.Set("X-Installation-ID", c.installation)
	}
	c.browser.headers(req.Route.Method, h)
	return h
}

func (c *Client) send(ctx context.Context, method, target string, body []byte, h http.Header) (*Response, error) {
	var r io.Reader
	if body != nil {
		r = bytes.NewReader(body)
	}
	hr, err := http.NewRequestWithContext(ctx, method, target, r)
	if err != nil {
		return nil, err
	}
	hr.Header = h
	resp, err := c.http.Do(hr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &Response{Status: resp.StatusCode, Header: resp.Header, Body: data}, nil
}

func (c *Client) rateLimited(resp *Response) (*RateLimitedError, error) {
	rl := &RateLimitedError{
		Scope:      resp.Header.Get("X-RateLimit-Scope"),
		Global:     resp.Header.Get("X-RateLimit-Global") == "true",
		Cloudflare: resp.Header.Get("Via") == "",
	}
	headerAfter, hasHeader := strconv.ParseFloat(resp.Header.Get("Retry-After"), 64)
	var body struct {
		RetryAfter float64 `json:"retry_after"`
		Global     bool    `json:"global"`
	}
	if json.Unmarshal(resp.Body, &body) == nil {
		rl.RetryAfter = seconds(body.RetryAfter)
		rl.Global = rl.Global || body.Global
	} else if hasHeader == nil && headerAfter > 0 {
		rl.RetryAfter = seconds(headerAfter)
	} else {
		code := "unknown"
		if m := cloudflareCodeExp.FindSubmatch(resp.Body); m != nil {
			code = string(m[1])
		}
		return nil, &CloudflareBanError{Code: code}
	}
	if c.maxRetryAfter > 0 && rl.RetryAfter > c.maxRetryAfter {
		return nil, rl
	}
	return rl, nil
}

func retryLater(body []byte) time.Duration {
	var v struct {
		Code       int      `json:"code"`
		RetryAfter *float64 `json:"retry_after"`
	}
	if json.Unmarshal(body, &v) != nil || v.Code < 110000 || v.Code > 119999 {
		return 0
	}
	if v.RetryAfter == nil || *v.RetryAfter == 0 {
		return 5 * time.Second
	}
	return seconds(*v.RetryAfter)
}

func isRetryStatus(s int) bool {
	_, ok := retryStatuses[s]
	return ok
}

func backoff(attempt int) time.Duration {
	return time.Duration(1+2*attempt) * time.Second
}
