package discord

import (
	"encoding/json"
	"time"

	"github.com/prettylittlelies/cry/snowflake"
)

type VerificationLevel int

type NSFWLevel int

type PremiumTier int

type GuildFeature string

type Guild struct {
	ID                          snowflake.ID      `json:"id"`
	Name                        string            `json:"name"`
	Icon                        *string           `json:"icon"`
	Splash                      *string           `json:"splash"`
	DiscoverySplash             *string           `json:"discovery_splash"`
	Banner                      *string           `json:"banner"`
	Description                 *string           `json:"description"`
	OwnerID                     snowflake.ID      `json:"owner_id"`
	AFKChannelID                *snowflake.ID     `json:"afk_channel_id"`
	AFKTimeout                  int               `json:"afk_timeout"`
	SystemChannelID             *snowflake.ID     `json:"system_channel_id"`
	SystemChannelFlags          int               `json:"system_channel_flags"`
	RulesChannelID              *snowflake.ID     `json:"rules_channel_id"`
	PublicUpdatesChannelID      *snowflake.ID     `json:"public_updates_channel_id"`
	VerificationLevel           VerificationLevel `json:"verification_level"`
	DefaultMessageNotifications int               `json:"default_message_notifications"`
	ExplicitContentFilter       int               `json:"explicit_content_filter"`
	MFALevel                    int               `json:"mfa_level"`
	NSFWLevel                   NSFWLevel         `json:"nsfw_level"`
	PremiumTier                 PremiumTier       `json:"premium_tier"`
	PremiumSubscriptionCount    int               `json:"premium_subscription_count"`
	PreferredLocale             string            `json:"preferred_locale"`
	VanityURLCode               *string           `json:"vanity_url_code"`
	Features                    []GuildFeature    `json:"features"`
	MaxMembers                  int               `json:"max_members,omitempty"`
	MemberCount                 int               `json:"member_count,omitempty"`
	JoinedAt                    *time.Time        `json:"joined_at,omitempty"`
	Large                       bool              `json:"large,omitempty"`
	Lazy                        bool              `json:"lazy,omitempty"`
	Unavailable                 bool              `json:"unavailable,omitempty"`
	Version                     int64             `json:"version,omitempty"`
	Channels                    []Channel         `json:"channels,omitempty"`
	Threads                     []Channel         `json:"threads,omitempty"`
	Roles                       []Role            `json:"roles,omitempty"`
	Emojis                      []Emoji           `json:"emojis,omitempty"`
	Stickers                    []Sticker         `json:"stickers,omitempty"`
	Members                     []Member          `json:"members,omitempty"`
}

type guildAlias Guild

func (g *Guild) UnmarshalJSON(b []byte) error {
	var wire struct {
		guildAlias
		Properties json.RawMessage `json:"properties"`
	}
	if err := json.Unmarshal(b, &wire); err != nil {
		return err
	}
	*g = Guild(wire.guildAlias)
	if len(wire.Properties) == 0 || string(wire.Properties) == "null" {
		return nil
	}
	outer := *g
	if err := json.Unmarshal(wire.Properties, (*guildAlias)(g)); err != nil {
		return err
	}
	g.ID = outer.ID
	g.MemberCount = outer.MemberCount
	g.JoinedAt = outer.JoinedAt
	g.Large, g.Lazy, g.Unavailable, g.Version = outer.Large, outer.Lazy, outer.Unavailable, outer.Version
	g.PremiumSubscriptionCount = max(g.PremiumSubscriptionCount, outer.PremiumSubscriptionCount)
	g.Channels, g.Threads, g.Roles = outer.Channels, outer.Threads, outer.Roles
	g.Emojis, g.Stickers, g.Members = outer.Emojis, outer.Stickers, outer.Members
	return nil
}

func (g *Guild) Role(id snowflake.ID) *Role {
	for i := range g.Roles {
		if g.Roles[i].ID == id {
			return &g.Roles[i]
		}
	}
	return nil
}
