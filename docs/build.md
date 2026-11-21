# build

```go
import "github.com/prettylittlelies/cry/build"
```

## Index

- [type Channel](#type-channel)
- [type Endpoints](#type-endpoints)
- [type ExtractError](#type-extracterror)
- [type MacOS](#type-macos)
- [type Mobile](#type-mobile)
- [type MobileTrack](#type-mobiletrack)
- [type MobileVersion](#type-mobileversion)
- [type NativeNumber](#type-nativenumber)
- [type Number](#type-number)
- [type Option](#type-option)
- [type ParseError](#type-parseerror)
- [type Release](#type-release)
- [type Scraper](#type-scraper)
- [type Snapshot](#type-snapshot)
- [type Source](#type-source)
- [type SourceError](#type-sourceerror)
- [type Version](#type-version)
- [type Windows](#type-windows)

## Types

### type Channel

```go
type Channel string
```

```go
const (
	Stable Channel = "stable"
	PTB    Channel = "ptb"
	Canary Channel = "canary"
)
```

#### func Channels

```go
func Channels() []Channel
```

#### func (Channel) Subdomain

```go
func (c Channel) Subdomain() string
```

#### func (Channel) Valid

```go
func (c Channel) Valid() bool
```

### type Endpoints

```go
type Endpoints struct {
	App             func(Channel) string
	Updates         string
	API             string
	ElectronIndex   string
	PlayStore       string
	AndroidVersions string
	ITunesLookup    string
}
```

#### func DefaultEndpoints

```go
func DefaultEndpoints() Endpoints
```

### type ExtractError

```go
type ExtractError struct {
	Source string
	Field  string
}
```

#### func (*ExtractError) Error

```go
func (e *ExtractError) Error() string
```

### type MacOS

```go
type MacOS struct {
	Host     Version
	Electron Version
	Chrome   Version
}
```

### type Mobile

```go
type Mobile struct {
	Android        MobileVersion
	AndroidMinimum uint32
	IOS            MobileVersion
	IOSMinimumOS   Version
}
```

### type MobileTrack

```go
type MobileTrack uint8
```

```go
const (
	TrackStable MobileTrack = iota
	TrackBeta
	TrackAlpha
	TrackDevelopment
)
```

### type MobileVersion

```go
type MobileVersion struct {
	Major uint32
	Minor uint32
}
```

#### func ParseMobileVersion

```go
func ParseMobileVersion(s string) (MobileVersion, error)
```

#### func (MobileVersion) AndroidCode

```go
func (m MobileVersion) AndroidCode(t MobileTrack) uint32
```

#### func (MobileVersion) AndroidUserAgent

```go
func (m MobileVersion) AndroidUserAgent(t MobileTrack) string
```

#### func (MobileVersion) String

```go
func (m MobileVersion) String() string
```

### type NativeNumber

```go
type NativeNumber uint32
```

#### func (NativeNumber) String

```go
func (n NativeNumber) String() string
```

### type Number

```go
type Number uint32
```

#### func (Number) String

```go
func (n Number) String() string
```

### type Option

```go
type Option func(*Scraper)
```

#### func WithChannels

```go
func WithChannels(c ...Channel) Option
```

#### func WithDoer

```go
func WithDoer(d fetch.Doer) Option
```

#### func WithEndpoints

```go
func WithEndpoints(e Endpoints) Option
```

### type ParseError

```go
type ParseError struct {
	Kind  string
	Input string
	Err   error
}
```

#### func (*ParseError) Error

```go
func (e *ParseError) Error() string
```

#### func (*ParseError) Unwrap

```go
func (e *ParseError) Unwrap() error
```

### type Release

```go
type Release struct {
	Channel Channel
	Number  Number
	Windows Windows
	MacOS   MacOS
}
```

### type Scraper

```go
type Scraper struct {
	// contains filtered or unexported fields
}
```

#### func NewScraper

```go
func NewScraper(opts ...Option) *Scraper
```

#### func (*Scraper) ChromeFor

```go
func (s *Scraper) ChromeFor(ctx context.Context, electron Version) (Version, error)
```

#### func (*Scraper) Fetch

```go
func (s *Scraper) Fetch(ctx context.Context) (*Snapshot, error)
```

### type Snapshot

```go
type Snapshot struct {
	FetchedAt time.Time
	Releases  map[Channel]Release
	Mobile    Mobile
}
```

#### func (*Snapshot) Release

```go
func (s *Snapshot) Release(c Channel) (Release, bool)
```

### type Source

```go
type Source interface {
	Fetch(ctx context.Context) (*Snapshot, error)
}
```

### type SourceError

```go
type SourceError struct {
	Source  string
	Channel Channel
	Err     error
}
```

#### func (*SourceError) Error

```go
func (e *SourceError) Error() string
```

#### func (*SourceError) Unwrap

```go
func (e *SourceError) Unwrap() error
```

### type Version

```go
type Version []uint32
```

#### func ParseVersion

```go
func ParseVersion(s string) (Version, error)
```

#### func (Version) Compare

```go
func (v Version) Compare(o Version) int
```

#### func (Version) Major

```go
func (v Version) Major() uint32
```

#### func (Version) MarshalText

```go
func (v Version) MarshalText() ([]byte, error)
```

#### func (Version) Segment

```go
func (v Version) Segment(i int) uint32
```

#### func (Version) String

```go
func (v Version) String() string
```

#### func (*Version) UnmarshalJSON

```go
func (v *Version) UnmarshalJSON(b []byte) error
```

#### func (*Version) UnmarshalText

```go
func (v *Version) UnmarshalText(b []byte) error
```

### type Windows

```go
type Windows struct {
	Host   Version
	Native NativeNumber
}
```

