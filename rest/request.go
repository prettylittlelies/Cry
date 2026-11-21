package rest

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
)

type Request struct {
	Route       Route
	Query       url.Values
	JSON        any
	Form        *Form
	Body        []byte
	ContentType string
	Reason      string
	Context     any
	Header      http.Header
	NoAuth      bool
}

type Response struct {
	Status int
	Header http.Header
	Body   []byte
}

func (r *Request) payload() ([]byte, string, error) {
	if r.JSON != nil {
		b, err := json.Marshal(r.JSON)
		return b, "application/json", err
	}
	if r.Form != nil {
		b, ct := r.Form.Encode()
		return b, ct, nil
	}
	return r.Body, r.ContentType, nil
}

func (r *Request) contextHeader() (string, error) {
	if r.Context == nil {
		return "", nil
	}
	b, err := json.Marshal(r.Context)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func Call[T any](ctx context.Context, c *Client, req *Request) (T, error) {
	var out T
	resp, err := c.Do(ctx, req)
	if err != nil {
		return out, err
	}
	if len(resp.Body) == 0 || resp.Status == http.StatusNoContent {
		return out, nil
	}
	if err := json.Unmarshal(resp.Body, &out); err != nil {
		return out, &DecodeError{Route: req.Route.Key(), Err: err}
	}
	return out, nil
}

type DecodeError struct {
	Route string
	Err   error
}

func (e *DecodeError) Error() string {
	return "rest: decode " + e.Route + ": " + e.Err.Error()
}

func (e *DecodeError) Unwrap() error {
	return e.Err
}
