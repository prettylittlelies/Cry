# voice

```go
import "github.com/prettylittlelies/cry/voice"
```

## Index

- [type Conn](#type-conn)
- [type FrameSource](#type-framesource)
- [type IVF](#type-ivf)
- [type IVFHeader](#type-ivfheader)
- [type MediaType](#type-mediatype)
- [type OggOpus](#type-oggopus)
- [type OggWriter](#type-oggwriter)
- [type Opcode](#type-opcode)
- [type Options](#type-options)
- [type SpeakingFlags](#type-speakingflags)
- [type VideoOptions](#type-videooptions)

## Constants

```go
const (
	ModeAES256GCM         = "aead_aes256_gcm_rtpsize"
	ModeXChaCha20Poly1305 = "aead_xchacha20_poly1305_rtpsize"
)
```

```go
const GatewayVersion = 8
```

## Variables

```go
var (
	ErrNotReady = errors.New("voice: connection not ready")
	ErrClosed   = errors.New("voice: connection closed")
)
```

```go
var (
	ErrNoMode       = errors.New("voice: no supported transport encryption mode")
	ErrShortPacket  = errors.New("voice: packet too short")
	ErrTransportTag = errors.New("voice: transport decryption failed")
)
```

```go
var (
	ErrNoVideo    = errors.New("voice: video was not negotiated on this connection")
	ErrVideoCodec = errors.New("voice: negotiated video codec is not VP8")
)
```

```go
var ErrDiscovery = errors.New("voice: ip discovery failed")
```

```go
var ErrHeartbeatTimeout = errors.New("voice: heartbeat not acknowledged")
```

```go
var ErrNotIVF = errors.New("voice: not an ivf stream")
```

```go
var ErrNotOgg = errors.New("voice: not an ogg stream")
```

## Types

### type Conn

```go
type Conn struct {
	// contains filtered or unexported fields
}
```

#### func Join

```go
func Join(ctx context.Context, c *cry.Client, guild, channel snowflake.ID, opts Options) (*Conn, error)
```

#### func (*Conn) Close

```go
func (v *Conn) Close(ctx context.Context) error
```

#### func (*Conn) Done

```go
func (v *Conn) Done() <-chan struct{}
```

#### func (*Conn) Encrypted

```go
func (v *Conn) Encrypted() bool
```

#### func (*Conn) EpochAuthenticator

```go
func (v *Conn) EpochAuthenticator() []byte
```

#### func (*Conn) Err

```go
func (v *Conn) Err() error
```

#### func (*Conn) Play

```go
func (v *Conn) Play(ctx context.Context, src FrameSource) error
```

#### func (*Conn) PlayIVF

```go
func (v *Conn) PlayIVF(ctx context.Context, r io.Reader, opts VideoOptions) error
```

#### func (*Conn) SetVideo

```go
func (v *Conn) SetVideo(active bool, opts VideoOptions) error
```

#### func (*Conn) Speaking

```go
func (v *Conn) Speaking(flags SpeakingFlags) error
```

#### func (*Conn) StartStream

```go
func (v *Conn) StartStream(ctx context.Context, opts Options) (*Conn, error)
```

#### func (*Conn) StreamKey

```go
func (v *Conn) StreamKey() string
```

#### func (*Conn) VideoCodec

```go
func (v *Conn) VideoCodec() string
```

#### func (*Conn) WriteOpus

```go
func (v *Conn) WriteOpus(frame []byte) error
```

#### func (*Conn) WriteVP8

```go
func (v *Conn) WriteVP8(frame []byte) error
```

### type FrameSource

```go
type FrameSource interface {
	Next() ([]byte, error)
}
```

### type IVF

```go
type IVF struct {
	Header IVFHeader
	// contains filtered or unexported fields
}
```

#### func NewIVF

```go
func NewIVF(r io.Reader) (*IVF, error)
```

#### func (*IVF) Next

```go
func (v *IVF) Next() ([]byte, error)
```

### type IVFHeader

```go
type IVFHeader struct {
	FourCC    string
	Width     int
	Height    int
	Rate      uint32
	Scale     uint32
	FrameRate float64
}
```

### type MediaType

```go
type MediaType string
```

```go
const (
	Camera MediaType = "video"
	Screen MediaType = "screen"
)
```

### type OggOpus

```go
type OggOpus struct {
	// contains filtered or unexported fields
}
```

#### func NewOggOpus

```go
func NewOggOpus(r io.Reader) *OggOpus
```

#### func (*OggOpus) Next

```go
func (o *OggOpus) Next() ([]byte, error)
```

### type OggWriter

```go
type OggWriter struct {
	// contains filtered or unexported fields
}
```

#### func NewOggWriter

```go
func NewOggWriter(w io.Writer) *OggWriter
```

#### func (*OggWriter) Close

```go
func (o *OggWriter) Close() error
```

#### func (*OggWriter) WritePacket

```go
func (o *OggWriter) WritePacket(packet []byte) error
```

### type Opcode

```go
type Opcode int
```

```go
const (
	OpIdentify           Opcode = 0
	OpSelectProtocol     Opcode = 1
	OpReady              Opcode = 2
	OpHeartbeat          Opcode = 3
	OpSessionDescription Opcode = 4
	OpSpeaking           Opcode = 5
	OpHeartbeatAck       Opcode = 6
	OpResume             Opcode = 7
	OpHello              Opcode = 8
	OpResumed            Opcode = 9
	OpClientsConnect     Opcode = 11
	OpVideo              Opcode = 12
	OpClientDisconnect   Opcode = 13
	OpSessionUpdate      Opcode = 14
	OpClientFlags        Opcode = 18
	OpClientPlatform     Opcode = 20
)
```

### type Options

```go
type Options struct {
	SelfMute         bool
	SelfDeaf         bool
	DisableDAVE      bool
	Signer           *mls.SignaturePrivateKey
	PreferredRegions []string
	Video            MediaType
	OnOpus           func(user snowflake.ID, frame []byte)
}
```

### type SpeakingFlags

```go
type SpeakingFlags int
```

```go
const (
	Voice      SpeakingFlags = 1 << 0
	Soundshare SpeakingFlags = 1 << 1
	Priority   SpeakingFlags = 1 << 2
)
```

### type VideoOptions

```go
type VideoOptions struct {
	Width   int
	Height  int
	FPS     int
	Bitrate int
}
```

