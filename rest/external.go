package rest

import (
	"context"
	"fmt"
	"net/http"
)

type External struct {
	Method string
	URL    string
	Header http.Header
	Body   []byte
}

type PreflightError struct {
	URL    string
	Status int
	Origin string
}

func (e *PreflightError) Error() string {
	return fmt.Sprintf("rest: cors preflight for %s failed (status %d, allow-origin %q)", e.URL, e.Status, e.Origin)
}

func (c *Client) External(ctx context.Context, e *External) (*Response, error) {
	page := e.Header.Clone()
	if page == nil {
		page = http.Header{}
	}
	unsafe := unsafeHeaders(page)
	key := c.browser.origin + " " + e.URL
	if needsPreflight(e.Method, unsafe) && !c.cors.covers(key, e.Method, unsafe) {
		if err := c.preflight(ctx, e.URL, e.Method, unsafe, key); err != nil {
			return nil, err
		}
	}
	c.browser.crossSite(page)
	return c.sendWith(ctx, c.anonymous, e.Method, e.URL, e.Body, page)
}

func (c *Client) preflight(ctx context.Context, target, method string, unsafe []string, key string) error {
	resp, err := c.sendWith(ctx, c.anonymous, http.MethodOptions, target, nil, c.browser.preflight(method, unsafe))
	if err != nil {
		return err
	}
	allow := resp.Header.Get("Access-Control-Allow-Origin")
	if resp.Status < 200 || resp.Status > 299 || (allow != "*" && allow != c.browser.origin) {
		return &PreflightError{URL: target, Status: resp.Status, Origin: allow}
	}
	c.cors.store(key, resp.Header)
	return nil
}
