package properties

import (
	"strings"

	"github.com/prettylittlelies/cry/build"
)

type Platform string

const (
	Darwin Platform = "darwin"
	Win32  Platform = "win32"
	Linux  Platform = "linux"
)

func (p Platform) OS() string {
	switch p {
	case Darwin:
		return "Mac OS X"
	case Win32:
		return "Windows"
	case Linux:
		return "Linux"
	}
	return string(p)
}

func (p Platform) ClientHint() string {
	switch p {
	case Darwin:
		return "macOS"
	case Win32:
		return "Windows"
	}
	return "Linux"
}

func (p Platform) userAgentOS() string {
	switch p {
	case Darwin:
		return "Macintosh; Intel Mac OS X 10_15_7"
	case Win32:
		return "Windows NT 10.0; Win64; x64"
	}
	return "X11; Linux x86_64"
}

type Arch string

const (
	X64   Arch = "x64"
	ARM64 Arch = "arm64"
	IA32  Arch = "ia32"
)

type LinuxSession struct {
	WindowManager      string
	Distro             string
	RuntimeEnvironment string
	DisplayServer      string
}

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

func MacOSHost() Host {
	return Host{
		Platform:      Darwin,
		Release:       "25.1.0",
		SystemVersion: "26.1.0",
		Arch:          ARM64,
		AppArch:       ARM64,
		Locale:        "en-US",
		Languages:     []string{"en-US", "en"},
		Timezone:      "America/New_York",
	}
}

func WindowsHost() Host {
	return Host{
		Platform:      Win32,
		Release:       "10.0.26100",
		SystemVersion: "10.0.26100",
		Arch:          X64,
		AppArch:       X64,
		BuildRevision: "6899",
		Locale:        "en-US",
		Languages:     []string{"en-US", "en"},
		Timezone:      "America/New_York",
	}
}

func (h Host) sdkVersion() string {
	parts := strings.Split(h.Release, ".")
	switch {
	case h.Platform == Darwin && len(parts) > 0:
		return parts[0]
	case h.Platform == Win32 && len(parts) > 2:
		return parts[2]
	}
	return ""
}

func (h Host) ElectronUserAgent(host build.Version, chrome, electron build.Version) string {
	return "Mozilla/5.0 (" + h.Platform.userAgentOS() + ") AppleWebKit/537.36 (KHTML, like Gecko) discord/" +
		host.String() + " Chrome/" + chrome.String() + " Electron/" + electron.String() + " Safari/537.36"
}

func (h Host) AcceptLanguage() string {
	var b strings.Builder
	q := 10
	for i, l := range h.Languages {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(l)
		if q < 10 {
			b.WriteString(";q=0.")
			b.WriteByte('0' + byte(q))
		}
		q = max(q-1, 1)
	}
	return b.String()
}
