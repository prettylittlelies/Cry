package discord

import "github.com/prettylittlelies/cry/snowflake"

func displayName(global *string, username string) string {
	if global != nil && *global != "" {
		return *global
	}
	return username
}

func (u *User) DisplayName() string {
	return displayName(u.GlobalName, u.Username)
}

func (u *User) Mention() string {
	return "<@" + u.ID.String() + ">"
}

func (u *PartialUser) DisplayName() string {
	return displayName(u.GlobalName, u.Username)
}

func (u *PartialUser) Mention() string {
	return "<@" + u.ID.String() + ">"
}

func (c *Channel) Mention() string {
	return "<#" + c.ID.String() + ">"
}

func (r *Role) Mention() string {
	return "<@&" + r.ID.String() + ">"
}

func (m *Message) JumpURL(guildID snowflake.ID) string {
	guild := "@me"
	if guildID != 0 {
		guild = guildID.String()
	}
	return "https://discord.com/channels/" + guild + "/" + m.ChannelID.String() + "/" + m.ID.String()
}
