package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/prettylittlelies/cry/rest"
)

type API struct {
	rest *rest.Client
}

func New(c *rest.Client) *API {
	return &API{rest: c}
}

type Option func(*rest.Request)

func WithReason(reason string) Option {
	return func(r *rest.Request) { r.Reason = reason }
}

func WithContextProperties(props any) Option {
	return func(r *rest.Request) { r.Context = props }
}

func WithHeader(key, value string) Option {
	return func(r *rest.Request) {
		if r.Header == nil {
			r.Header = http.Header{}
		}
		r.Header.Set(key, value)
	}
}

func WithoutAuth() Option {
	return func(r *rest.Request) { r.NoAuth = true }
}

func prepare(req *rest.Request, opts []Option) *rest.Request {
	for _, o := range opts {
		o(req)
	}
	return req
}

func call[T any](ctx context.Context, a *API, req *rest.Request, opts []Option) (*T, error) {
	out, err := rest.Call[T](ctx, a.rest, prepare(req, opts))
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func list[T any](ctx context.Context, a *API, req *rest.Request, opts []Option) ([]T, error) {
	return rest.Call[[]T](ctx, a.rest, prepare(req, opts))
}

func raw(ctx context.Context, a *API, req *rest.Request, opts []Option) (json.RawMessage, error) {
	resp, err := a.rest.Do(ctx, prepare(req, opts))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func none(ctx context.Context, a *API, req *rest.Request, opts []Option) error {
	_, err := a.rest.Do(ctx, prepare(req, opts))
	return err
}
