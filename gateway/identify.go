package gateway

import (
	"slices"
	"strings"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/properties"
)

const ColdStart = "COLD_START"

type Presence struct {
	Status     discord.StatusType `json:"status"`
	Since      int64              `json:"since"`
	Activities []discord.Activity `json:"activities"`
	AFK        bool               `json:"afk"`
}

func DefaultPresence() Presence {
	return Presence{Status: discord.StatusTypeOnline, Activities: []discord.Activity{}}
}

type ClientState struct {
	GuildVersions map[string]int64 `json:"guild_versions"`
}

type fastConnectProperties struct {
	properties.Desktop
	IsFastConnect  bool   `json:"is_fast_connect"`
	InstallationID string `json:"installation_id,omitempty"`
}

type identifyProperties struct {
	properties.Desktop
	IsFastConnect         bool   `json:"is_fast_connect"`
	GatewayConnectReasons string `json:"gateway_connect_reasons"`
	InstallationID        string `json:"installation_id,omitempty"`
}

type fastIdentify struct {
	Token        string                `json:"token"`
	Capabilities Capabilities          `json:"capabilities"`
	Properties   fastConnectProperties `json:"properties"`
	ClientState  ClientState           `json:"client_state"`
}

type identify struct {
	Token        string             `json:"token"`
	Capabilities Capabilities       `json:"capabilities"`
	Properties   identifyProperties `json:"properties"`
	Presence     Presence           `json:"presence"`
	Compress     bool               `json:"compress"`
	ClientState  ClientState        `json:"client_state"`
	QoSToken     string             `json:"qos_token"`
}

func freshClientState() ClientState {
	return ClientState{GuildVersions: map[string]int64{}}
}

func fastConnectDesktop(d properties.Desktop) properties.Desktop {
	d.LaunchSignature = ""
	d.ClientHeartbeatSessionID = ""
	d.ClientAppState = ""
	return d
}

func connectReasons(reasons []string) string {
	if len(reasons) == 0 {
		return "NO_REASONS"
	}
	sorted := slices.Clone(reasons)
	slices.Sort(sorted)
	return strings.Join(sorted, ",")
}
