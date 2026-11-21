# rest

```go
import "github.com/prettylittlelies/cry/rest"
```

## Index

- [func Call](#func-call)
- [func Values](#func-values)
- [type Captcha](#type-captcha)
- [type CaptchaFunc](#type-captchafunc)
- [type CaptchaHandler](#type-captchahandler)
- [type CaptchaRequiredError](#type-captcharequirederror)
- [type CaptchaService](#type-captchaservice)
- [type Client](#type-client)
- [type CloudAttachment](#type-cloudattachment)
- [type CloudFile](#type-cloudfile)
- [type CloudflareBanError](#type-cloudflarebanerror)
- [type DecodeError](#type-decodeerror)
- [type External](#type-external)
- [type File](#type-file)
- [type Form](#type-form)
- [type HTTPError](#type-httperror)
- [type Identity](#type-identity)
- [type MessageAttachment](#type-messageattachment)
- [type Option](#type-option)
- [type Param](#type-param)
- [type ParamValue](#type-paramvalue)
- [type Part](#type-part)
- [type PreflightError](#type-preflighterror)
- [type RateLimitedError](#type-ratelimitederror)
- [type RefreshedURL](#type-refreshedurl)
- [type Request](#type-request)
- [type Response](#type-response)
- [type Route](#type-route)
- [type Stats](#type-stats)
- [type Upload](#type-upload)
- [type UploadError](#type-uploaderror)

## Constants

```go
const (
	DefaultOrigin  = "https://discord.com"
	DefaultReferer = DefaultOrigin + "/channels/@me"
)
```

```go
const (
	DefaultBaseURL = "https://discord.com/api/v9"
)
```

## Variables

```go
var (
	ErrUnauthorized = errors.New("rest: unauthorized")
	ErrForbidden    = errors.New("rest: forbidden")
	ErrNotFound     = errors.New("rest: not found")
	ErrServer       = errors.New("rest: discord server error")
)
```

```go
var ErrEmptyFile = errors.New("rest: cannot upload an empty file")
```

## Functions

### func Call

```go
func Call[T any](ctx context.Context, c *Client, req *Request) (T, error)
```

### func Values

```go
func Values(v any) url.Values
```

## Types

### type Captcha

```go
type Captcha struct {
	Errors    []string       `json:"captcha_key"`
	Service   CaptchaService `json:"captcha_service"`
	Sitekey   string         `json:"captcha_sitekey"`
	SessionID string         `json:"captcha_session_id"`
	RqData    string         `json:"captcha_rqdata"`
	RqToken   string         `json:"captcha_rqtoken"`
	Invisible bool           `json:"should_serve_invisible"`
}
```

### type CaptchaFunc

```go
type CaptchaFunc func(ctx context.Context, c *Captcha) (string, error)
```

#### func (CaptchaFunc) Solve

```go
func (f CaptchaFunc) Solve(ctx context.Context, c *Captcha) (string, error)
```

### type CaptchaHandler

```go
type CaptchaHandler interface {
	Solve(ctx context.Context, c *Captcha) (string, error)
}
```

### type CaptchaRequiredError

```go
type CaptchaRequiredError struct {
	*HTTPError
	Captcha *Captcha
}
```

#### func (*CaptchaRequiredError) Error

```go
func (e *CaptchaRequiredError) Error() string
```

#### func (*CaptchaRequiredError) Unwrap

```go
func (e *CaptchaRequiredError) Unwrap() error
```

### type CaptchaService

```go
type CaptchaService string
```

```go
const (
	HCaptcha            CaptchaService = "hcaptcha"
	ReCaptcha           CaptchaService = "recaptcha"
	ReCaptchaEnterprise CaptchaService = "recaptcha_enterprise"
)
```

### type Client

```go
type Client struct {
	// contains filtered or unexported fields
}
```

#### func New

```go
func New(token string, rt http.RoundTripper, id Identity, opts ...Option) *Client
```

#### func (*Client) DeleteUpload

```go
func (c *Client) DeleteUpload(ctx context.Context, uploadedFilename string) error
```

#### func (*Client) Do

```go
func (c *Client) Do(ctx context.Context, req *Request) (*Response, error)
```

#### func (*Client) External

```go
func (c *Client) External(ctx context.Context, e *External) (*Response, error)
```

#### func (*Client) RefreshAttachmentURLs

```go
func (c *Client) RefreshAttachmentURLs(ctx context.Context, urls ...string) ([]RefreshedURL, error)
```

#### func (*Client) SetToken

```go
func (c *Client) SetToken(token string)
```

#### func (*Client) Stats

```go
func (c *Client) Stats() Stats
```

#### func (*Client) UploadAttachment

```go
func (c *Client) UploadAttachment(ctx context.Context, channel snowflake.ID, f CloudFile) (*Upload, error)
```

### type CloudAttachment

```go
type CloudAttachment struct {
	ID             string `json:"id"`
	UploadURL      string `json:"upload_url"`
	UploadFilename string `json:"upload_filename"`
}
```

### type CloudFile

```go
type CloudFile struct {
	Name        string
	ContentType string
	Data        []byte
	Description string
	Spoiler     bool
}
```

### type CloudflareBanError

```go
type CloudflareBanError struct {
	Code string
}
```

#### func (*CloudflareBanError) Error

```go
func (e *CloudflareBanError) Error() string
```

### type DecodeError

```go
type DecodeError struct {
	Route string
	Err   error
}
```

#### func (*DecodeError) Error

```go
func (e *DecodeError) Error() string
```

#### func (*DecodeError) Unwrap

```go
func (e *DecodeError) Unwrap() error
```

### type External

```go
type External struct {
	Method string
	URL    string
	Header http.Header
	Body   []byte
}
```

### type File

```go
type File struct {
	Name        string
	ContentType string
	Data        []byte
}
```

### type Form

```go
type Form struct {
	Parts []Part
}
```

#### func MessageForm

```go
func MessageForm(payload any, files ...File) (*Form, error)
```

#### func (*Form) Encode

```go
func (f *Form) Encode() ([]byte, string)
```

### type HTTPError

```go
type HTTPError struct {
	Method  string
	URL     string
	Status  int
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Errors  json.RawMessage `json:"errors"`
	Body    []byte
}
```

#### func (*HTTPError) Error

```go
func (e *HTTPError) Error() string
```

#### func (*HTTPError) Is

```go
func (e *HTTPError) Is(target error) bool
```

### type Identity

```go
type Identity struct {
	Session *properties.Session
	Host    properties.Host
	Chrome  build.Version
	Mobile  bool
}
```

### type MessageAttachment

```go
type MessageAttachment struct {
	ID               string `json:"id"`
	Description      string `json:"description,omitempty"`
	Filename         string `json:"filename"`
	UploadedFilename string `json:"uploaded_filename"`
	IsSpoiler        bool   `json:"is_spoiler,omitempty"`
}
```

### type Option

```go
type Option func(*Client)
```

#### func WithBaseURL

```go
func WithBaseURL(u string) Option
```

#### func WithCaptcha

```go
func WithCaptcha(h CaptchaHandler) Option
```

#### func WithClock

```go
func WithClock(now func() time.Time) Option
```

#### func WithDebugOptions

```go
func WithDebugOptions(opts ...string) Option
```

#### func WithInstallationID

```go
func WithInstallationID(id string) Option
```

#### func WithMaxRetryAfter

```go
func WithMaxRetryAfter(d time.Duration) Option
```

#### func WithReferer

```go
func WithReferer(r string) Option
```

### type Param

```go
type Param struct {
	Name  string
	Value string
	Raw   bool
}
```

#### func P

```go
func P[T ParamValue](name string, value T) Param
```

#### func Raw

```go
func Raw(name, value string) Param
```

### type ParamValue

```go
type ParamValue interface {
	~string | ~uint64
}
```

### type Part

```go
type Part struct {
	Name  string
	File  *File
	Value string
}
```

### type PreflightError

```go
type PreflightError struct {
	URL    string
	Status int
	Origin string
}
```

#### func (*PreflightError) Error

```go
func (e *PreflightError) Error() string
```

### type RateLimitedError

```go
type RateLimitedError struct {
	RetryAfter time.Duration
	Global     bool
	Scope      string
	Cloudflare bool
}
```

#### func (*RateLimitedError) Error

```go
func (e *RateLimitedError) Error() string
```

### type RefreshedURL

```go
type RefreshedURL struct {
	Original  string `json:"original"`
	Refreshed string `json:"refreshed"`
}
```

### type Request

```go
type Request struct {
	Route       Route
	Query       url.Values
	JSON        any
	Form        *Form
	Body        []byte
	ContentType string
	Reason      string
	Context     any
	Header      http.Header
	NoAuth      bool
}
```

### type Response

```go
type Response struct {
	Status int
	Header http.Header
	Body   []byte
}
```

### type Route

```go
type Route struct {
	Method   string
	Template string
	Params   []Param
}
```

#### func NewRoute

```go
func NewRoute(method, template string, params ...Param) Route
```

#### func (Route) Key

```go
func (r Route) Key() string
```

#### func (Route) Major

```go
func (r Route) Major() string
```

#### func (Route) Path

```go
func (r Route) Path() string
```

### type Stats

```go
type Stats struct {
	Buckets int
}
```

### type Upload

```go
type Upload struct {
	File             CloudFile
	UploadedFilename string
}
```

#### func (Upload) Attachment

```go
func (u Upload) Attachment(index int) MessageAttachment
```

### type UploadError

```go
type UploadError struct {
	Filename string
	Attempts int
	Status   int
	Err      error
}
```

#### func (*UploadError) Error

```go
func (e *UploadError) Error() string
```

#### func (*UploadError) Unwrap

```go
func (e *UploadError) Unwrap() error
```

