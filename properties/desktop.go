package properties

import (
	"github.com/google/uuid"
	"github.com/prettylittlelies/cry/build"
)

type AppState string

const (
	Focused    AppState = "focused"
	Unfocused  AppState = "unfocused"
	Active     AppState = "active"
	Inactive   AppState = "inactive"
	Background AppState = "background"
)

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

type DesktopRelease struct {
	Channel  build.Channel
	Host     build.Version
	Electron build.Version
	Chrome   build.Version
	Number   build.Number
	Native   build.NativeNumber
}

func MacOSRelease(r build.Release) DesktopRelease {
	return DesktopRelease{Channel: r.Channel, Host: r.MacOS.Host, Electron: r.MacOS.Electron, Chrome: r.MacOS.Chrome, Number: r.Number}
}

func NewDesktop(h Host, r DesktopRelease) *Desktop {
	d := &Desktop{
		OS:                 h.Platform.OS(),
		Browser:            "Discord Client",
		ReleaseChannel:     r.Channel,
		ClientVersion:      r.Host.String(),
		OSVersion:          h.SystemVersion,
		OSVersionIsRelease: true,
		OSSDKVersion:       h.sdkVersion(),
		OSArch:             h.Arch,
		AppArch:            h.AppArch,
		SystemLocale:       h.Locale,
		ClientLaunchID:     uuid.NewString(),
		BrowserUserAgent:   h.ElectronUserAgent(r.Host, r.Chrome, r.Electron),
		BrowserVersion:     r.Electron.String(),
		ClientBuildNumber:  r.Number,
		LaunchSignature:    LaunchSignature(),
	}
	if h.Platform == Win32 {
		d.OSBuildRevision = h.BuildRevision
	}
	if h.Platform == Linux && h.Linux != nil {
		d.WindowManager = h.Linux.WindowManager
		d.Distro = h.Linux.Distro
		d.RuntimeEnvironment = h.Linux.RuntimeEnvironment
		d.DisplayServer = h.Linux.DisplayServer
	}
	if r.Native != 0 {
		native := r.Native
		d.NativeBuildNumber = &native
	}
	return d
}
