package transport

import (
	"errors"
	"fmt"

	"golang.org/x/net/http2"
)

var (
	ErrNotH2         = errors.New("transport: server did not negotiate h2")
	ErrUnsupported   = errors.New("transport: only https is supported")
	ErrConnClosed    = errors.New("transport: connection closed")
	ErrMissingStatus = errors.New("transport: response missing :status")
)

type StreamError struct {
	StreamID uint32
	Code     http2.ErrCode
}

func (e *StreamError) Error() string {
	return fmt.Sprintf("transport: stream %d reset: %v", e.StreamID, e.Code)
}

type GoAwayError struct {
	LastStreamID uint32
	Code         http2.ErrCode
	Debug        string
}

func (e *GoAwayError) Error() string {
	return fmt.Sprintf("transport: goaway %v (last stream %d) %s", e.Code, e.LastStreamID, e.Debug)
}
