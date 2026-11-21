package discord

import (
	"time"

	"github.com/prettylittlelies/cry/snowflake"
)

type ChannelType int

const (
	ChannelGuildText          ChannelType = 0
	ChannelDM                 ChannelType = 1
	ChannelGuildVoice         ChannelType = 2
	ChannelGroupDM            ChannelType = 3
	ChannelGuildCategory      ChannelType = 4
	ChannelGuildNews          ChannelType = 5
	ChannelGuildStore         ChannelType = 6
	ChannelGuildNewsThread    ChannelType = 10
	ChannelGuildPublicThread  ChannelType = 11
	ChannelGuildPrivateThread ChannelType = 12
	ChannelGuildStageVoice    ChannelType = 13
	ChannelGuildDirectory     ChannelType = 14
	ChannelGuildForum         ChannelType = 15
	ChannelGuildMedia         ChannelType = 16
)

func (t ChannelType) IsThread() bool {
	return t == ChannelGuildNewsThread || t == ChannelGuildPublicThread || t == ChannelGuildPrivateThread
}

func (t ChannelType) IsPrivate() bool {
	return t == ChannelDM || t == ChannelGroupDM
}

type ThreadMetadata struct {
	Archived            bool       `json:"archived"`
	AutoArchiveDuration int        `json:"auto_archive_duration"`
	ArchiveTimestamp    *time.Time `json:"archive_timestamp"`
	Locked              bool       `json:"locked"`
	Invitable           *bool      `json:"invitable,omitempty"`
	CreateTimestamp     *time.Time `json:"create_timestamp,omitempty"`
}

type ForumTag struct {
	ID        snowflake.ID  `json:"id"`
	Name      string        `json:"name"`
	Moderated bool          `json:"moderated"`
	EmojiID   *snowflake.ID `json:"emoji_id"`
	EmojiName *string       `json:"emoji_name"`
}

type Channel struct {
	ID                   snowflake.ID    `json:"id"`
	Type                 ChannelType     `json:"type"`
	GuildID              *snowflake.ID   `json:"guild_id,omitempty"`
	Position             int             `json:"position,omitempty"`
	PermissionOverwrites []Overwrite     `json:"permission_overwrites,omitempty"`
	Name                 *string         `json:"name,omitempty"`
	Topic                *string         `json:"topic,omitempty"`
	NSFW                 bool            `json:"nsfw,omitempty"`
	LastMessageID        *snowflake.ID   `json:"last_message_id,omitempty"`
	LastPinTimestamp     *time.Time      `json:"last_pin_timestamp,omitempty"`
	Bitrate              int             `json:"bitrate,omitempty"`
	UserLimit            int             `json:"user_limit,omitempty"`
	RateLimitPerUser     int             `json:"rate_limit_per_user,omitempty"`
	Recipients           []User          `json:"recipients,omitempty"`
	RecipientIDs         []snowflake.ID  `json:"recipient_ids,omitempty"`
	Icon                 *string         `json:"icon,omitempty"`
	OwnerID              *snowflake.ID   `json:"owner_id,omitempty"`
	ParentID             *snowflake.ID   `json:"parent_id,omitempty"`
	RTCRegion            *string         `json:"rtc_region,omitempty"`
	ThreadMetadata       *ThreadMetadata `json:"thread_metadata,omitempty"`
	MessageCount         int             `json:"message_count,omitempty"`
	MemberCount          int             `json:"member_count,omitempty"`
	Flags                int             `json:"flags,omitempty"`
	AvailableTags        []ForumTag      `json:"available_tags,omitempty"`
	AppliedTags          []snowflake.ID  `json:"applied_tags,omitempty"`
	IsMessageRequest     bool            `json:"is_message_request,omitempty"`
	IsSpam               bool            `json:"is_spam,omitempty"`
}

func (c *Channel) Mention() string {
	return "<#" + c.ID.String() + ">"
}

func (c *Channel) Participants() []snowflake.ID {
	if len(c.RecipientIDs) > 0 {
		return c.RecipientIDs
	}
	ids := make([]snowflake.ID, len(c.Recipients))
	for i, u := range c.Recipients {
		ids[i] = u.ID
	}
	return ids
}
