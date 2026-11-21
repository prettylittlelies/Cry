package fetch

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	Doer   Doer
	Header http.Header
}

func New(d Doer, header http.Header) Client {
	if d == nil {
		d = http.DefaultClient
	}
	return Client{Doer: d, Header: header}
}

func (c Client) Get(ctx context.Context, url string, extra http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range c.Header {
		req.Header[k] = v
	}
	for k, v := range extra {
		req.Header[k] = v
	}
	resp, err := c.Doer.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		resp.Body.Close()
		return nil, &StatusError{URL: url, Code: resp.StatusCode}
	}
	return resp, nil
}

func (c Client) Bytes(ctx context.Context, url string) ([]byte, error) {
	resp, err := c.Get(ctx, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func JSON[T any](ctx context.Context, c Client, url string) (T, error) {
	var out T
	resp, err := c.Get(ctx, url, nil)
	if err != nil {
		return out, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return out, &DecodeError{URL: url, Err: err}
	}
	return out, nil
}
