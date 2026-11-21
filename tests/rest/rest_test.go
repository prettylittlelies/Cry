package rest_test

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sync/atomic"
	"testing"
	"time"

	"github.com/prettylittlelies/cry/build"
	"github.com/prettylittlelies/cry/fingerprint"
	"github.com/prettylittlelies/cry/properties"
	"github.com/prettylittlelies/cry/rest"
	"github.com/prettylittlelies/cry/tests/testkit"
	"github.com/prettylittlelies/cry/transport"
)

func identity(t *testing.T) rest.Identity {
	t.Helper()
	chrome, _ := build.ParseVersion("152.0.7977.169")
	electron, _ := build.ParseVersion("44.10.0")
	host, _ := build.ParseVersion("0.0.421")
	h := properties.MacOSHost()
	d := properties.NewDesktop(h, properties.DesktopRelease{Channel: build.Stable, Host: host, Electron: electron, Chrome: chrome, Number: 678818})
	return rest.Identity{Session: properties.NewSession(d), Host: h, Chrome: chrome}
}

func clientURL(t *testing.T, h http.HandlerFunc, opts ...rest.Option) (*rest.Client, string) {
	t.Helper()
	srv := testkit.NewH2Server(t, h)
	tr := transport.New(fingerprint.Electron44(), transport.WithRootCAs(srv.Pool))
	t.Cleanup(tr.CloseIdleConnections)
	return rest.New("token", tr, identity(t), append([]rest.Option{rest.WithBaseURL(srv.URL)}, opts...)...), srv.URL
}

func client(t *testing.T, h http.HandlerFunc, opts ...rest.Option) *rest.Client {
	t.Helper()
	c, _ := clientURL(t, h, opts...)
	return c
}

func me() *rest.Request {
	return &rest.Request{Route: rest.NewRoute(http.MethodGet, "/users/@me")}
}

func TestBrowserHeaders(t *testing.T) {
	seen := make(chan http.Header, 2)
	c := client(t, func(w http.ResponseWriter, r *http.Request) {
		seen <- r.Header.Clone()
		io.WriteString(w, `{}`)
	})
	if _, err := c.Do(context.Background(), me()); err != nil {
		t.Fatal(err)
	}
	post := &rest.Request{Route: rest.NewRoute(http.MethodPost, "/channels/{channel_id}/messages", rest.P("channel_id", "1")), JSON: map[string]string{"content": "x"}}
	if _, err := c.Do(context.Background(), post); err != nil {
		t.Fatal(err)
	}
	get, sent := <-seen, <-seen
	raw, err := base64.StdEncoding.DecodeString(get.Get("X-Super-Properties"))
	if err != nil || !json.Valid(raw) {
		t.Fatalf("super properties %q", get.Get("X-Super-Properties"))
	}
	checks := map[string]string{
		"Authorization":      "token",
		"Sec-Fetch-Site":     "same-origin",
		"Sec-Fetch-Mode":     "cors",
		"Referer":            rest.DefaultReferer,
		"X-Discord-Locale":   "en-US",
		"X-Discord-Timezone": "America/New_York",
		"Sec-Ch-Ua-Platform": `"macOS"`,
		"Sec-Ch-Ua-Mobile":   "?0",
		"Accept-Language":    "en-US,en;q=0.9",
	}
	for k, want := range checks {
		if got := get.Get(k); got != want {
			t.Errorf("%s: got %q, want %q", k, got, want)
		}
	}
	if get.Get("Origin") != "" || sent.Get("Origin") != rest.DefaultOrigin {
		t.Errorf("origin get=%q post=%q", get.Get("Origin"), sent.Get("Origin"))
	}
	if sent.Get("Content-Type") != "application/json" {
		t.Errorf("content-type %q", sent.Get("Content-Type"))
	}
}

func TestRouteRateLimitRetry(t *testing.T) {
	var calls atomic.Int32
	c := client(t, func(w http.ResponseWriter, r *http.Request) {
		if calls.Add(1) == 1 {
			w.Header().Set("Via", "1.1 google")
			w.WriteHeader(http.StatusTooManyRequests)
			io.WriteString(w, `{"retry_after":0.05,"global":false}`)
			return
		}
		io.WriteString(w, `{"id":"1"}`)
	})
	start := time.Now()
	out, err := rest.Call[struct {
		ID string `json:"id"`
	}](context.Background(), c, me())
	if err != nil || out.ID != "1" || calls.Load() != 2 || time.Since(start) < 50*time.Millisecond {
		t.Fatalf("out=%+v err=%v calls=%d", out, err, calls.Load())
	}
}

func TestAcceptedRetryBumpsAttempts(t *testing.T) {
	var calls atomic.Int32
	c := client(t, func(w http.ResponseWriter, r *http.Request) {
		if calls.Add(1) == 1 {
			w.WriteHeader(http.StatusAccepted)
			io.WriteString(w, `{"code":110000,"retry_after":0.01}`)
			return
		}
		io.WriteString(w, `{"attempts":"`+r.URL.Query().Get("attempts")+`"}`)
	})
	out, err := rest.Call[map[string]string](context.Background(), c, me())
	if err != nil || out["attempts"] != "1" {
		t.Fatalf("out=%v err=%v", out, err)
	}
}

func TestServerErrorRetryCountsFailures(t *testing.T) {
	if testing.Short() {
		t.Skip("waits for backoff")
	}
	var calls atomic.Int32
	c := client(t, func(w http.ResponseWriter, r *http.Request) {
		if calls.Add(1) == 1 {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		io.WriteString(w, `{"failed":"`+r.Header.Get("X-Failed-Requests")+`"}`)
	})
	out, err := rest.Call[map[string]string](context.Background(), c, me())
	if err != nil || out["failed"] != "1" {
		t.Fatalf("out=%v err=%v", out, err)
	}
}

const captchaBody = `{"captcha_key":["captcha-required"],"captcha_sitekey":"site","captcha_service":"hcaptcha","captcha_session_id":"sess","captcha_rqdata":"data","captcha_rqtoken":"rq"}`

func TestCaptchaSolvedAndRetried(t *testing.T) {
	var got http.Header
	c := client(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Captcha-Key") == "" {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, captchaBody)
			return
		}
		got = r.Header.Clone()
		io.WriteString(w, `{}`)
	}, rest.WithCaptcha(rest.CaptchaFunc(func(_ context.Context, cap *rest.Captcha) (string, error) {
		if cap.Sitekey != "site" || cap.RqData != "data" || cap.Service != rest.HCaptcha {
			t.Errorf("captcha %+v", cap)
		}
		return "solved", nil
	})))
	if _, err := c.Do(context.Background(), me()); err != nil {
		t.Fatal(err)
	}
	if got.Get("X-Captcha-Key") != "solved" || got.Get("X-Captcha-Session-Id") != "sess" || got.Get("X-Captcha-Rqtoken") != "rq" {
		t.Fatalf("captcha headers %v", got)
	}
}

func TestCaptchaWithoutHandler(t *testing.T) {
	c := client(t, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, captchaBody)
	})
	_, err := c.Do(context.Background(), me())
	var cerr *rest.CaptchaRequiredError
	if !errors.As(err, &cerr) || cerr.Captcha.SessionID != "sess" {
		t.Fatalf("got %v", err)
	}
}

func TestNotFound(t *testing.T) {
	c := client(t, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, `{"message":"Unknown Channel","code":10003}`)
	})
	_, err := c.Do(context.Background(), me())
	var herr *rest.HTTPError
	if !errors.Is(err, rest.ErrNotFound) || !errors.As(err, &herr) || herr.Code != 10003 || herr.Message != "Unknown Channel" {
		t.Fatalf("got %v", err)
	}
}

func TestCloudflareBan(t *testing.T) {
	c := client(t, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		io.WriteString(w, `<html>error code: 1015</html>`)
	})
	_, err := c.Do(context.Background(), me())
	var ban *rest.CloudflareBanError
	if !errors.As(err, &ban) || ban.Code != "1015" {
		t.Fatalf("got %v", err)
	}
}

func TestMaxRetryAfter(t *testing.T) {
	c := client(t, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		io.WriteString(w, `{"retry_after":600,"global":true}`)
	}, rest.WithMaxRetryAfter(time.Second))
	_, err := c.Do(context.Background(), me())
	var rl *rest.RateLimitedError
	if !errors.As(err, &rl) || !rl.Global || rl.RetryAfter != 600*time.Second {
		t.Fatalf("got %v", err)
	}
}

func TestCookiesPersist(t *testing.T) {
	var calls atomic.Int32
	var cookie string
	c := client(t, func(w http.ResponseWriter, r *http.Request) {
		if calls.Add(1) == 1 {
			http.SetCookie(w, &http.Cookie{Name: "__dcfduid", Value: "abc", Path: "/", Secure: true})
		} else {
			cookie = r.Header.Get("Cookie")
		}
		io.WriteString(w, `{}`)
	})
	for range 2 {
		if _, err := c.Do(context.Background(), me()); err != nil {
			t.Fatal(err)
		}
	}
	if cookie != "__dcfduid=abc" {
		t.Fatalf("cookie %q", cookie)
	}
}

func TestAuditLogReasonEncoding(t *testing.T) {
	var reason string
	c := client(t, func(w http.ResponseWriter, r *http.Request) {
		reason = r.Header.Get("X-Audit-Log-Reason")
		io.WriteString(w, `{}`)
	})
	req := me()
	req.Reason = "spam & raid (bye!) ü"
	if _, err := c.Do(context.Background(), req); err != nil {
		t.Fatal(err)
	}
	if want := "spam%20%26%20raid%20(bye!)%20%C3%BC"; reason != want {
		t.Fatalf("got %q, want %q", reason, want)
	}
}
