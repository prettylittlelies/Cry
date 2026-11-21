package discord

import (
	"encoding/json"
	"time"

	"github.com/prettylittlelies/cry/snowflake"
)

type MessageType int

const (
	MessageDefault MessageType = 0
	MessageReply   MessageType = 19
)

type MessageFlags int

const (
	MessageCrossposted           MessageFlags = 1 << 0
	MessageSuppressEmbeds        MessageFlags = 1 << 2
	MessageEphemeral             MessageFlags = 1 << 6
	MessageSuppressNotifications MessageFlags = 1 << 12
	MessageIsVoiceMessage        MessageFlags = 1 << 13
)

type Attachment struct {
	ID           snowflake.ID `json:"id"`
	Filename     string       `json:"filename"`
	Title        *string      `json:"title,omitempty"`
	Description  *string      `json:"description,omitempty"`
	ContentType  *string      `json:"content_type,omitempty"`
	Size         int          `json:"size"`
	URL          string       `json:"url"`
	ProxyURL     string       `json:"proxy_url"`
	Height       *int         `json:"height,omitempty"`
	Width        *int         `json:"width,omitempty"`
	Spoiler      bool         `json:"spoiler,omitempty"`
	DurationSecs *float64     `json:"duration_secs,omitempty"`
	Waveform     *string      `json:"waveform,omitempty"`
	Flags        int          `json:"flags,omitempty"`
}

type EmbedMedia struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type Embed struct {
	Type        string     `json:"type,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	URL         string     `json:"url,omitempty"`
	Timestamp   *time.Time `json:"timestamp,omitempty"`
	Color       int        `json:"color,omitempty"`
	Footer      *struct {
		Text    string `json:"text"`
		IconURL string `json:"icon_url,omitempty"`
	} `json:"footer,omitempty"`
	Image     *EmbedMedia `json:"image,omitempty"`
	Thumbnail *EmbedMedia `json:"thumbnail,omitempty"`
	Video     *EmbedMedia `json:"video,omitempty"`
	Provider  *struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"provider,omitempty"`
	Author *struct {
		Name    string `json:"name"`
		URL     string `json:"url,omitempty"`
		IconURL string `json:"icon_url,omitempty"`
	} `json:"author,omitempty"`
	Fields []EmbedField `json:"fields,omitempty"`
}

type ReactionCounts struct {
	Burst  int `json:"burst"`
	Normal int `json:"normal"`
}

type Reaction struct {
	Count        int            `json:"count"`
	CountDetails ReactionCounts `json:"count_details"`
	Me           bool           `json:"me"`
	MeBurst      bool           `json:"me_burst"`
	Emoji        Emoji          `json:"emoji"`
	BurstColors  []string       `json:"burst_colors,omitempty"`
}

type MessageReference struct {
	Type            int           `json:"type,omitempty"`
	MessageID       *snowflake.ID `json:"message_id,omitempty"`
	ChannelID       *snowflake.ID `json:"channel_id,omitempty"`
	GuildID         *snowflake.ID `json:"guild_id,omitempty"`
	FailIfNotExists *bool         `json:"fail_if_not_exists,omitempty"`
}

type StickerItem struct {
	ID         snowflake.ID  `json:"id"`
	Name       string        `json:"name"`
	FormatType StickerFormat `json:"format_type"`
}

type Message struct {
	ID                snowflake.ID      `json:"id"`
	ChannelID         snowflake.ID      `json:"channel_id"`
	GuildID           *snowflake.ID     `json:"guild_id,omitempty"`
	Author            User              `json:"author"`
	Member            *Member           `json:"member,omitempty"`
	Content           string            `json:"content"`
	Timestamp         time.Time         `json:"timestamp"`
	EditedTimestamp   *time.Time        `json:"edited_timestamp"`
	TTS               bool              `json:"tts"`
	MentionEveryone   bool              `json:"mention_everyone"`
	Mentions          []User            `json:"mentions"`
	MentionRoles      []snowflake.ID    `json:"mention_roles"`
	Attachments       []Attachment      `json:"attachments"`
	Embeds            []Embed           `json:"embeds"`
	Reactions         []Reaction        `json:"reactions,omitempty"`
	Nonce             json.RawMessage   `json:"nonce,omitempty"`
	Pinned            bool              `json:"pinned"`
	WebhookID         *snowflake.ID     `json:"webhook_id,omitempty"`
	Type              MessageType       `json:"type"`
	Flags             MessageFlags      `json:"flags,omitempty"`
	MessageReference  *MessageReference `json:"message_reference,omitempty"`
	ReferencedMessage *Message          `json:"referenced_message,omitempty"`
	StickerItems      []StickerItem     `json:"sticker_items,omitempty"`
	Components        json.RawMessage   `json:"components,omitempty"`
	Poll              json.RawMessage   `json:"poll,omitempty"`
}

func (m *Message) JumpURL() string {
	guild := "@me"
	if m.GuildID != nil {
		guild = m.GuildID.String()
	}
	return "https://discord.com/channels/" + guild + "/" + m.ChannelID.String() + "/" + m.ID.String()
}
