package discord

import (
	"encoding/json"

	"github.com/prettylittlelies/cry/snowflake"
)

type Versioned[T any] struct {
	Entries []T  `json:"entries"`
	Partial bool `json:"partial"`
	Version int  `json:"version"`
}

type Ready struct {
	V                 int                        `json:"v"`
	User              ClientUser                 `json:"user"`
	Users             []User                     `json:"users"`
	Guilds            []Guild                    `json:"guilds"`
	MergedMembers     [][]Member                 `json:"merged_members"`
	PrivateChannels   []Channel                  `json:"private_channels"`
	Relationships     []Relationship             `json:"relationships"`
	Sessions          []Session                  `json:"sessions"`
	ReadState         Versioned[ReadState]       `json:"read_state"`
	UserGuildSettings Versioned[json.RawMessage] `json:"user_guild_settings"`
	UserSettingsProto string                     `json:"user_settings_proto"`
	SessionID         string                     `json:"session_id"`
	SessionType       string                     `json:"session_type"`
	ResumeGatewayURL  string                     `json:"resume_gateway_url"`
	AnalyticsToken    string                     `json:"analytics_token"`
	CountryCode       string                     `json:"country_code"`
	ConnectedAccounts []json.RawMessage          `json:"connected_accounts"`
	Experiments       json.RawMessage            `json:"experiments"`
	GuildExperiments  json.RawMessage            `json:"guild_experiments"`
}

type MergedPresences struct {
	Friends []Presence   `json:"friends"`
	Guilds  [][]Presence `json:"guilds"`
}

type SupplementalGuild struct {
	ID          snowflake.ID      `json:"id"`
	VoiceStates []json.RawMessage `json:"voice_states"`
}

type ReadySupplemental struct {
	Guilds              []SupplementalGuild `json:"guilds"`
	MergedMembers       [][]Member          `json:"merged_members"`
	MergedPresences     MergedPresences     `json:"merged_presences"`
	LazyPrivateChannels []Channel           `json:"lazy_private_channels"`
}
