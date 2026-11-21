# websocket

```go
import "github.com/prettylittlelies/cry/transport/websocket"
```

## Index

- [func NewDialer](#func-newdialer)
- [type Browser](#type-browser)
- [type CloseError](#type-closeerror)
- [type Config](#type-config)
- [type Conn](#type-conn)
- [type HandshakeError](#type-handshakeerror)
- [type Opcode](#type-opcode)

## Constants

```go
const (
	Text   = ws.OpText
	Binary = ws.OpBinary
)
```

```go
const (
	DefaultMaxMessage = 64 << 20
)
```

## Variables

```go
var (
	ErrClosed          = errors.New("websocket: connection closed")
	ErrProtocol        = errors.New("websocket: protocol violation")
	ErrMessageTooLarge = errors.New("websocket: message exceeds limit")
)
```

## Functions

### func NewDialer

```go
func NewDialer(hello fingerprint.ClientHello) *fingerprint.Dialer
```

## Types

### type Browser

```go
type Browser struct {
	UserAgent      string
	Origin         string
	AcceptLanguage string
}
```

### type CloseError

```go
type CloseError struct {
	Code   int
	Reason string
}
```

#### func (*CloseError) Error

```go
func (e *CloseError) Error() string
```

### type Config

```go
type Config struct {
	Dialer     *fingerprint.Dialer
	Browser    Browser
	Jar        http.CookieJar
	MaxMessage int
}
```

### type Conn

```go
type Conn struct {
	Response *http.Response
	// contains filtered or unexported fields
}
```

#### func Dial

```go
func Dial(ctx context.Context, rawURL string, cfg Config) (*Conn, error)
```

#### func (*Conn) Close

```go
func (c *Conn) Close(code int, reason string) error
```

#### func (*Conn) Read

```go
func (c *Conn) Read() (Opcode, []byte, error)
```

#### func (*Conn) Write

```go
func (c *Conn) Write(op Opcode, payload []byte) error
```

### type HandshakeError

```go
type HandshakeError struct {
	Status int
	Reason string
	Body   []byte
}
```

#### func (*HandshakeError) Error

```go
func (e *HandshakeError) Error() string
```

### type Opcode

```go
type Opcode = ws.OpCode
```

