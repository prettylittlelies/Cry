package fetch

import (
	"fmt"
	"net/http"
)

type StatusError struct {
	URL  string
	Code int
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("fetch: %s: %d %s", e.URL, e.Code, http.StatusText(e.Code))
}

type DecodeError struct {
	URL string
	Err error
}

func (e *DecodeError) Error() string {
	return fmt.Sprintf("fetch: decode %s: %v", e.URL, e.Err)
}

func (e *DecodeError) Unwrap() error {
	return e.Err
}
