package build

import (
	"net/url"

	"github.com/google/uuid"
)

type Endpoints struct {
	App             func(Channel) string
	Updates         string
	API             string
	ElectronIndex   string
	PlayStore       string
	AndroidVersions string
	ITunesLookup    string
}

func DefaultEndpoints() Endpoints {
	return Endpoints{
		App:             func(c Channel) string { return "https://" + c.Subdomain() + "discord.com/app" },
		Updates:         "https://updates.discord.com",
		API:             "https://discord.com/api",
		ElectronIndex:   "https://releases.electronjs.org/releases.json",
		PlayStore:       "https://play.google.com/store/apps/details?id=com.discord&hl=en",
		AndroidVersions: "https://dl.discordapp.net/apps/android/versions.json",
		ITunesLookup:    "https://itunes.apple.com/lookup?bundleId=com.hammerandchisel.discord",
	}
}

func (e Endpoints) windowsManifest(c Channel) string {
	q := url.Values{
		"install_id": {uuid.NewString()},
		"channel":    {string(c)},
		"platform":   {"win"},
		"arch":       {"x64"},
	}
	return e.Updates + "/distributions/app/manifests/latest?" + q.Encode()
}

func (e Endpoints) macUpdate(c Channel) string {
	return e.API + "/updates/" + string(c) + "?platform=osx"
}
