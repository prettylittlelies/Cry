# properties

```go
import "github.com/prettylittlelies/cry/properties"
```

## Index

- [func LaunchSignature](#func-launchsignature)
- [func MobileLaunchSignature](#func-mobilelaunchsignature)
- [type AppState](#type-appstate)
- [type Arch](#type-arch)
- [type ClientMod](#type-clientmod)
- [type Desktop](#type-desktop)
- [type DesktopRelease](#type-desktoprelease)
- [type Host](#type-host)
- [type LinuxSession](#type-linuxsession)
- [type Platform](#type-platform)
- [type Session](#type-session)

## Constants

```go
const HeartbeatRotation = 30 * time.Minute
```

## Functions

### func LaunchSignature

```go
func LaunchSignature() string
```

### func MobileLaunchSignature

```go
func MobileLaunchSignature(now time.Time) string
```

## Types

### type AppState

```go
type AppState string
```

```go
const (
	Focused    AppState = "focused"
	Unfocused  AppState = "unfocused"
	Active     AppState = "active"
	Inactive   AppState = "inactive"
	Background AppState = "background"
)
```

### type Arch

```go
type Arch string
```

```go
const (
	X64   Arch = "x64"
	ARM64 Arch = "arm64"
	IA32  Arch = "ia32"
)
```

### type ClientMod

```go
type ClientMod uint8
```

```go
const (
	JQuery        ClientMod = 119
	BetterDiscord ClientMod = 108
	Rambox        ClientMod = 100
	Vendetta      ClientMod = 91
	Vencord       ClientMod = 84
	Replugged     ClientMod = 75
	Legcord       ClientMod = 61
	Dorion        ClientMod = 55
	GoofCord      ClientMod = 48
	OpenAsar      ClientMod = 38
	Shelter       ClientMod = 24
	Moonlight     ClientMod = 11
)
```

#### func (ClientMod) In

```go
func (m ClientMod) In(signature string) bool
```

### type Desktop

```go
type Desktop struct {
	OS                       string              `json:"os"`
	Browser                  string              `json:"browser"`
	ReleaseChannel           build.Channel       `json:"release_channel"`
	ClientVersion            string              `json:"client_version"`
	OSVersion                string              `json:"os_version"`
	OSVersionIsRelease       bool                `json:"os_version_is_release"`
	OSSDKVersion             string              `json:"os_sdk_version,omitempty"`
	OSArch                   Arch                `json:"os_arch"`
	AppArch                  Arch                `json:"app_arch"`
	SystemLocale             string              `json:"system_locale"`
	HasClientMods            bool                `json:"has_client_mods"`
	ClientLaunchID           string              `json:"client_launch_id"`
	BrowserUserAgent         string              `json:"browser_user_agent"`
	BrowserVersion           string              `json:"browser_version"`
	WindowManager            string              `json:"window_manager,omitempty"`
	Distro                   string              `json:"distro,omitempty"`
	RuntimeEnvironment       string              `json:"runtime_environment,omitempty"`
	DisplayServer            string              `json:"display_server,omitempty"`
	OSBuildRevision          string              `json:"os_build_revision,omitempty"`
	ClientBuildNumber        build.Number        `json:"client_build_number"`
	NativeBuildNumber        *build.NativeNumber `json:"native_build_number,omitempty"`
	ClientEventSource        *string             `json:"client_event_source"`
	InstallSource            string              `json:"install_source,omitempty"`
	LaunchSignature          string              `json:"launch_signature,omitempty"`
	ClientHeartbeatSessionID string              `json:"client_heartbeat_session_id,omitempty"`
	ClientAppState           AppState            `json:"client_app_state,omitempty"`
}
```

#### func NewDesktop

```go
func NewDesktop(h Host, r DesktopRelease) *Desktop
```

### type DesktopRelease

```go
type DesktopRelease struct {
	Channel  build.Channel
	Host     build.Version
	Electron build.Version
	Chrome   build.Version
	Number   build.Number
	Native   build.NativeNumber
}
```

#### func MacOSRelease

```go
func MacOSRelease(r build.Release) DesktopRelease
```

### type Host

```go
type Host struct {
	Platform      Platform
	Release       string
	SystemVersion string
	Arch          Arch
	AppArch       Arch
	BuildRevision string
	Locale        string
	Languages     []string
	Timezone      string
	Linux         *LinuxSession
}
```

#### func MacOSHost

```go
func MacOSHost() Host
```

#### func WindowsHost

```go
func WindowsHost() Host
```

#### func (Host) AcceptLanguage

```go
func (h Host) AcceptLanguage() string
```

#### func (Host) ElectronUserAgent

```go
func (h Host) ElectronUserAgent(host build.Version, chrome, electron build.Version) string
```

### type LinuxSession

```go
type LinuxSession struct {
	WindowManager      string
	Distro             string
	RuntimeEnvironment string
	DisplayServer      string
}
```

### type Platform

```go
type Platform string
```

```go
const (
	Darwin Platform = "darwin"
	Win32  Platform = "win32"
	Linux  Platform = "linux"
)
```

#### func (Platform) ClientHint

```go
func (p Platform) ClientHint() string
```

#### func (Platform) OS

```go
func (p Platform) OS() string
```

### type Session

```go
type Session struct {
	// contains filtered or unexported fields
}
```

#### func NewSession

```go
func NewSession(d *Desktop) *Session
```

#### func (*Session) Encoded

```go
func (s *Session) Encoded() string
```

#### func (*Session) Properties

```go
func (s *Session) Properties() Desktop
```

#### func (*Session) SetAppState

```go
func (s *Session) SetAppState(state AppState)
```

#### func (*Session) UserAgent

```go
func (s *Session) UserAgent() string
```

