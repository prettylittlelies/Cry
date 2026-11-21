package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	ErrUnauthorized = errors.New("rest: unauthorized")
	ErrForbidden    = errors.New("rest: forbidden")
	ErrNotFound     = errors.New("rest: not found")
	ErrServer       = errors.New("rest: discord server error")
)

type HTTPError struct {
	Method  string
	URL     string
	Status  int
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Errors  json.RawMessage `json:"errors"`
	Body    []byte
}

func (e *HTTPError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("rest: %s %s: %d %s (code %d)", e.Method, e.URL, e.Status, e.Message, e.Code)
	}
	return fmt.Sprintf("rest: %s %s: %d %s", e.Method, e.URL, e.Status, http.StatusText(e.Status))
}

func (e *HTTPError) Is(target error) bool {
	switch target {
	case ErrUnauthorized:
		return e.Status == http.StatusUnauthorized
	case ErrForbidden:
		return e.Status == http.StatusForbidden
	case ErrNotFound:
		return e.Status == http.StatusNotFound
	case ErrServer:
		return e.Status >= 500
	}
	return false
}

func newHTTPError(method, url string, status int, body []byte) *HTTPError {
	e := &HTTPError{Method: method, URL: url, Status: status, Body: body}
	_ = json.Unmarshal(body, e)
	return e
}

type RateLimitedError struct {
	RetryAfter time.Duration
	Global     bool
	Scope      string
	Cloudflare bool
}

func (e *RateLimitedError) Error() string {
	return fmt.Sprintf("rest: rate limited for %s (global=%v scope=%q cloudflare=%v)", e.RetryAfter, e.Global, e.Scope, e.Cloudflare)
}

type CloudflareBanError struct {
	Code string
}

func (e *CloudflareBanError) Error() string {
	return "rest: cloudflare ban (code " + e.Code + ")"
}

type CaptchaRequiredError struct {
	*HTTPError
	Captcha *Captcha
}

func (e *CaptchaRequiredError) Error() string {
	return fmt.Sprintf("rest: captcha required (%s)", e.Captcha.Service)
}

func (e *CaptchaRequiredError) Unwrap() error {
	return e.HTTPError
}
