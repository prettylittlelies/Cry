package discord

import "github.com/prettylittlelies/cry/snowflake"

type Emoji struct {
	ID            *snowflake.ID  `json:"id"`
	Name          *string        `json:"name"`
	Roles         []snowflake.ID `json:"roles,omitempty"`
	User          *User          `json:"user,omitempty"`
	RequireColons bool           `json:"require_colons,omitempty"`
	Managed       bool           `json:"managed,omitempty"`
	Animated      bool           `json:"animated,omitempty"`
	Available     bool           `json:"available,omitempty"`
}

func (e *Emoji) APIName() string {
	name := ""
	if e.Name != nil {
		name = *e.Name
	}
	if e.ID == nil {
		return name
	}
	return name + ":" + e.ID.String()
}

func (e *Emoji) String() string {
	if e.ID == nil {
		return e.APIName()
	}
	prefix := "<:"
	if e.Animated {
		prefix = "<a:"
	}
	return prefix + e.APIName() + ">"
}

type StickerFormat int

const (
	StickerPNG StickerFormat = iota + 1
	StickerAPNG
	StickerLottie
	StickerGIF
)

type Sticker struct {
	ID          snowflake.ID  `json:"id"`
	PackID      *snowflake.ID `json:"pack_id,omitempty"`
	Name        string        `json:"name"`
	Description *string       `json:"description"`
	Tags        string        `json:"tags"`
	Type        int           `json:"type"`
	FormatType  StickerFormat `json:"format_type"`
	Available   bool          `json:"available,omitempty"`
	GuildID     *snowflake.ID `json:"guild_id,omitempty"`
	User        *User         `json:"user,omitempty"`
	SortValue   *int          `json:"sort_value,omitempty"`
}
