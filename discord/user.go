package discord

import "github.com/prettylittlelies/cry/snowflake"

type UserFlags uint64

type PremiumType int

const (
	PremiumNone PremiumType = iota
	PremiumClassic
	PremiumNitro
	PremiumBasic
)

type AvatarDecoration struct {
	Asset     string       `json:"asset"`
	SkuID     snowflake.ID `json:"sku_id"`
	ExpiresAt *int64       `json:"expires_at"`
}

type PrimaryGuild struct {
	IdentityEnabled bool          `json:"identity_enabled"`
	IdentityGuildID *snowflake.ID `json:"identity_guild_id"`
	Tag             *string       `json:"tag"`
	Badge           *string       `json:"badge"`
}

type User struct {
	ID                   snowflake.ID      `json:"id"`
	Username             string            `json:"username"`
	Discriminator        string            `json:"discriminator"`
	GlobalName           *string           `json:"global_name"`
	Avatar               *string           `json:"avatar"`
	AvatarDecorationData *AvatarDecoration `json:"avatar_decoration_data"`
	Banner               *string           `json:"banner,omitempty"`
	AccentColor          *int              `json:"accent_color,omitempty"`
	Bot                  bool              `json:"bot,omitempty"`
	System               bool              `json:"system,omitempty"`
	PublicFlags          UserFlags         `json:"public_flags,omitempty"`
	PrimaryGuild         *PrimaryGuild     `json:"primary_guild,omitempty"`
}

func (u *User) DisplayName() string {
	if u.GlobalName != nil && *u.GlobalName != "" {
		return *u.GlobalName
	}
	return u.Username
}

func (u *User) Mention() string {
	return "<@" + u.ID.String() + ">"
}

type ClientUser struct {
	User
	Email          *string     `json:"email"`
	Phone          *string     `json:"phone"`
	Verified       bool        `json:"verified"`
	MFAEnabled     bool        `json:"mfa_enabled"`
	NSFWAllowed    *bool       `json:"nsfw_allowed"`
	Premium        bool        `json:"premium"`
	PremiumType    PremiumType `json:"premium_type"`
	Flags          UserFlags   `json:"flags"`
	PurchasedFlags int         `json:"purchased_flags"`
	Bio            string      `json:"bio"`
	Pronouns       string      `json:"pronouns"`
	Desktop        bool        `json:"desktop"`
	Mobile         bool        `json:"mobile"`
}
