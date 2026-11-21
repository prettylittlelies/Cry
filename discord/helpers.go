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

func (u *User) Partial() PartialUser {
	return PartialUser{
		ID:                   u.ID,
		Username:             u.Username,
		Discriminator:        u.Discriminator,
		GlobalName:           u.GlobalName,
		Avatar:               u.Avatar,
		AvatarDecorationData: u.AvatarDecorationData,
		Collectibles:         u.Collectibles,
		DisplayNameStyles:    u.DisplayNameStyles,
		PrimaryGuild:         u.PrimaryGuild,
		Bot:                  u.Bot,
		System:               u.System,
	}
}

func (t ChannelType) IsThread() bool {
	return t == ChannelTypeNewsThread || t == ChannelTypePublicThread || t == ChannelTypePrivateThread
}

func (t ChannelType) Messageable() bool {
	switch t {
	case ChannelTypeGuildText, ChannelTypeDM, ChannelTypeGuildVoice, ChannelTypeGroupDM, ChannelTypeGuildNews, ChannelTypeGuildStageVoice:
		return true
	}
	return t.IsThread()
}
