package discord

import (
	"slices"
	"time"

	"github.com/prettylittlelies/cry/snowflake"
)

type MemberFlags int

type Member struct {
	User                       *User          `json:"user,omitempty"`
	UserID                     snowflake.ID   `json:"user_id,omitempty"`
	GuildID                    snowflake.ID   `json:"guild_id,omitempty"`
	Nick                       *string        `json:"nick"`
	Avatar                     *string        `json:"avatar"`
	Banner                     *string        `json:"banner,omitempty"`
	Roles                      []snowflake.ID `json:"roles"`
	JoinedAt                   *time.Time     `json:"joined_at"`
	PremiumSince               *time.Time     `json:"premium_since"`
	Deaf                       bool           `json:"deaf"`
	Mute                       bool           `json:"mute"`
	Pending                    bool           `json:"pending,omitempty"`
	Flags                      MemberFlags    `json:"flags"`
	CommunicationDisabledUntil *time.Time     `json:"communication_disabled_until"`
}

func (m *Member) ID() snowflake.ID {
	if m.User != nil {
		return m.User.ID
	}
	return m.UserID
}

func (m *Member) HasRole(id snowflake.ID) bool {
	return slices.Contains(m.Roles, id)
}
