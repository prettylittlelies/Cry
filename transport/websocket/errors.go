package websocket

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrClosed          = errors.New("websocket: connection closed")
	ErrProtocol        = errors.New("websocket: protocol violation")
	ErrMessageTooLarge = errors.New("websocket: message exceeds limit")
)

type HandshakeError struct {
	Status int
	Reason string
	Body   []byte
}

func (e *HandshakeError) Error() string {
	if e.Reason != "" {
		return fmt.Sprintf("websocket: handshake failed (%d): %s", e.Status, e.Reason)
	}
	return fmt.Sprintf("websocket: handshake failed: %d %s", e.Status, http.StatusText(e.Status))
}

type CloseError struct {
	Code   int
	Reason string
}

func (e *CloseError) Error() string {
	return fmt.Sprintf("websocket: closed with %d %q", e.Code, e.Reason)
}
