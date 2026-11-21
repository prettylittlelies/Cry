# codec

```go
import "github.com/prettylittlelies/cry/gateway/codec"
```

## Index

- [func Encode](#func-encode)
- [type Compression](#type-compression)
- [type Encoding](#type-encoding)
- [type Reader](#type-reader)
- [type Source](#type-source)
- [type UnsupportedError](#type-unsupportederror)

## Functions

### func Encode

```go
func Encode(enc Encoding, v any) (websocket.Opcode, []byte, error)
```

## Types

### type Compression

```go
type Compression string
```

```go
const (
	None       Compression = ""
	ZstdStream Compression = "zstd-stream"
	ZlibStream Compression = "zlib-stream"
)
```

### type Encoding

```go
type Encoding string
```

```go
const (
	ETF  Encoding = "etf"
	JSON Encoding = "json"
)
```

### type Reader

```go
type Reader struct {
	// contains filtered or unexported fields
}
```

#### func NewReader

```go
func NewReader(src Source, enc Encoding, comp Compression) (*Reader, error)
```

#### func (*Reader) Close

```go
func (r *Reader) Close()
```

#### func (*Reader) Next

```go
func (r *Reader) Next() (json.RawMessage, error)
```

### type Source

```go
type Source interface {
	Read() (websocket.Opcode, []byte, error)
}
```

### type UnsupportedError

```go
type UnsupportedError struct {
	Encoding    Encoding
	Compression Compression
}
```

#### func (*UnsupportedError) Error

```go
func (e *UnsupportedError) Error() string
```

