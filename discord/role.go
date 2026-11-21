package discord

import "github.com/prettylittlelies/cry/snowflake"

type RoleColors struct {
	Primary   int  `json:"primary_color"`
	Secondary *int `json:"secondary_color"`
	Tertiary  *int `json:"tertiary_color"`
}

type RoleTags struct {
	BotID                 *snowflake.ID `json:"bot_id,omitempty"`
	IntegrationID         *snowflake.ID `json:"integration_id,omitempty"`
	SubscriptionListingID *snowflake.ID `json:"subscription_listing_id,omitempty"`
}

type Role struct {
	ID           snowflake.ID `json:"id"`
	Name         string       `json:"name"`
	Description  *string      `json:"description"`
	Color        int          `json:"color"`
	Colors       *RoleColors  `json:"colors,omitempty"`
	Hoist        bool         `json:"hoist"`
	Icon         *string      `json:"icon"`
	UnicodeEmoji *string      `json:"unicode_emoji"`
	Position     int          `json:"position"`
	Permissions  Permissions  `json:"permissions"`
	Managed      bool         `json:"managed"`
	Mentionable  bool         `json:"mentionable"`
	Flags        int          `json:"flags"`
	Tags         *RoleTags    `json:"tags,omitempty"`
}

func (r *Role) Mention() string {
	return "<@&" + r.ID.String() + ">"
}
