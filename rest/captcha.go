package rest

import (
	"context"
	"encoding/json"
	"net/http"
)

type CaptchaService string

const (
	HCaptcha            CaptchaService = "hcaptcha"
	ReCaptcha           CaptchaService = "recaptcha"
	ReCaptchaEnterprise CaptchaService = "recaptcha_enterprise"
)

type Captcha struct {
	Errors    []string       `json:"captcha_key"`
	Service   CaptchaService `json:"captcha_service"`
	Sitekey   string         `json:"captcha_sitekey"`
	SessionID string         `json:"captcha_session_id"`
	RqData    string         `json:"captcha_rqdata"`
	RqToken   string         `json:"captcha_rqtoken"`
	Invisible bool           `json:"should_serve_invisible"`
}

type CaptchaHandler interface {
	Solve(ctx context.Context, c *Captcha) (string, error)
}

type CaptchaFunc func(ctx context.Context, c *Captcha) (string, error)

func (f CaptchaFunc) Solve(ctx context.Context, c *Captcha) (string, error) {
	return f(ctx, c)
}

func parseCaptcha(status int, body []byte) *Captcha {
	if status != http.StatusBadRequest {
		return nil
	}
	var probe struct {
		Key *json.RawMessage `json:"captcha_key"`
	}
	if json.Unmarshal(body, &probe) != nil || probe.Key == nil {
		return nil
	}
	c := new(Captcha)
	if json.Unmarshal(body, c) != nil {
		return nil
	}
	return c
}

func (c *Captcha) apply(h http.Header, key string) {
	h.Set("X-Captcha-Key", key)
	if c.SessionID != "" {
		h.Set("X-Captcha-Session-Id", c.SessionID)
	}
	if c.RqToken != "" {
		h.Set("X-Captcha-Rqtoken", c.RqToken)
	}
}
