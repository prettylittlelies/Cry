# fingerprint

```go
import "github.com/prettylittlelies/cry/fingerprint"
```

## Index

- [func BlinkOrder](#func-blinkorder)
- [func Chromium152](#func-chromium152)
- [func Chromium152HTTP1](#func-chromium152http1)
- [func SecCHUA](#func-secchua)
- [type Brand](#type-brand)
- [type ClientHello](#type-clienthello)
- [type ContextDialer](#type-contextdialer)
- [type Dialer](#type-dialer)
- [type H2](#type-h2)
- [type HeaderPolicy](#type-headerpolicy)
- [type Profile](#type-profile)
- [type Pseudo](#type-pseudo)
- [type Urgency](#type-urgency)

## Constants

```go
const (
	MLDSA44 utls.SignatureScheme = 0x0904
	MLDSA65 utls.SignatureScheme = 0x0905
	MLDSA87 utls.SignatureScheme = 0x0906
)
```

```go
const Wildcard = "*"
```

## Variables

```go
var (
	ALPNHTTP2 = []string{"h2", "http/1.1"}
	ALPNHTTP1 = []string{"http/1.1"}
)
```

## Functions

### func BlinkOrder

```go
func BlinkOrder(names ...[]string) []string
```

### func Chromium152

```go
func Chromium152() *utls.ClientHelloSpec
```

### func Chromium152HTTP1

```go
func Chromium152HTTP1() *utls.ClientHelloSpec
```

### func SecCHUA

```go
func SecCHUA(brands []Brand) string
```

## Types

### type Brand

```go
type Brand struct {
	Name    string
	Version string
}
```

#### func Brands

```go
func Brands(major uint32, product string) []Brand
```

### type ClientHello

```go
type ClientHello func() *utls.ClientHelloSpec
```

### type ContextDialer

```go
type ContextDialer interface {
	DialContext(ctx context.Context, network, addr string) (net.Conn, error)
}
```

### type Dialer

```go
type Dialer struct {
	Hello    ClientHello
	ALPN     []string
	Sessions utls.ClientSessionCache
	Net      ContextDialer
	RootCAs  *x509.CertPool
}
```

#### func NewDialer

```go
func NewDialer(hello ClientHello) *Dialer
```

#### func (*Dialer) DialTLS

```go
func (d *Dialer) DialTLS(ctx context.Context, network, addr string) (*utls.UConn, error)
```

### type H2

```go
type H2 struct {
	Settings       []http2.Setting
	ConnectionFlow uint32
	PseudoOrder    []Pseudo
	Weights        []uint16
	Exclusive      bool
}
```

#### func Chromium152H2

```go
func Chromium152H2() H2
```

#### func (H2) Priority

```go
func (h H2) Priority(u Urgency) http2.PriorityParam
```

#### func (H2) Setting

```go
func (h H2) Setting(id http2.SettingID) (uint32, bool)
```

### type HeaderPolicy

```go
type HeaderPolicy struct {
	Insertion []string
	Leading   []string
	Hints     []string
	Trailing  []string
	Preflight []string
}
```

#### func DiscordDesktopPolicy

```go
func DiscordDesktopPolicy() HeaderPolicy
```

#### func DiscordWebPolicy

```go
func DiscordWebPolicy() HeaderPolicy
```

#### func (HeaderPolicy) Arrange

```go
func (p HeaderPolicy) Arrange(inserted []string) []string
```

#### func (HeaderPolicy) Fields

```go
func (p HeaderPolicy) Fields(method string, h http.Header) []hpack.HeaderField
```

### type Profile

```go
type Profile struct {
	Name    string
	Hello   ClientHello
	H2      H2
	Headers HeaderPolicy
}
```

#### func Electron44

```go
func Electron44() Profile
```

### type Pseudo

```go
type Pseudo string
```

```go
const (
	Method    Pseudo = ":method"
	Authority Pseudo = ":authority"
	Scheme    Pseudo = ":scheme"
	Path      Pseudo = ":path"
)
```

### type Urgency

```go
type Urgency uint8
```

```go
const DefaultUrgency Urgency = 3
```

#### func ParseUrgency

```go
func ParseUrgency(priority string) Urgency
```

