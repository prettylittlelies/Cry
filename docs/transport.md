# transport

```go
import "github.com/prettylittlelies/cry/transport"
```

## Index

- [type GoAwayError](#type-goawayerror)
- [type Option](#type-option)
- [type StreamError](#type-streamerror)
- [type Transport](#type-transport)

## Variables

```go
var (
	ErrNotH2         = errors.New("transport: server did not negotiate h2")
	ErrUnsupported   = errors.New("transport: only https is supported")
	ErrConnClosed    = errors.New("transport: connection closed")
	ErrMissingStatus = errors.New("transport: response missing :status")
)
```

## Types

### type GoAwayError

```go
type GoAwayError struct {
	LastStreamID uint32
	Code         http2.ErrCode
	Debug        string
}
```

#### func (*GoAwayError) Error

```go
func (e *GoAwayError) Error() string
```

### type Option

```go
type Option func(*Transport)
```

#### func WithNetDialer

```go
func WithNetDialer(d fingerprint.ContextDialer) Option
```

#### func WithRootCAs

```go
func WithRootCAs(pool *x509.CertPool) Option
```

### type StreamError

```go
type StreamError struct {
	StreamID uint32
	Code     http2.ErrCode
}
```

#### func (*StreamError) Error

```go
func (e *StreamError) Error() string
```

### type Transport

```go
type Transport struct {
	// contains filtered or unexported fields
}
```

#### func New

```go
func New(p fingerprint.Profile, opts ...Option) *Transport
```

#### func (*Transport) CloseIdleConnections

```go
func (t *Transport) CloseIdleConnections()
```

#### func (*Transport) Profile

```go
func (t *Transport) Profile() fingerprint.Profile
```

#### func (*Transport) RoundTrip

```go
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error)
```

