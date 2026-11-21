package converting

import "github.com/prettylittlelies/cry/discord"

func (s Scope) Role(argument string) (discord.Role, bool) {
	g, ok := s.Client.State().Guild(s.GuildID)
	if !ok {
		return discord.Role{}, false
	}
	pred := func(r *discord.Role) bool { return r.Name == argument }
	if id, ok := matchID(argument, roleMentionPattern); ok {
		pred = func(r *discord.Role) bool { return r.ID == id }
	}
	return findFirst(g.Roles, pred)
}

func (s Scope) Guild(argument string) (discord.GatewayGuild, bool) {
	st := s.Client.State()
	if id, ok := matchID(argument, nil); ok {
		if g, ok := st.Guild(id); ok {
			return g, true
		}
	}
	return findFirst(st.AllGuilds(), func(g *discord.GatewayGuild) bool { return g.Properties.Name == argument })
}

func (s Scope) Channel(argument string) (discord.Channel, bool) {
	st := s.Client.State()
	if id, ok := matchID(argument, channelPattern); ok {
		c, ok := st.Channel(id)
		return c, ok && (s.GuildID == 0 || c.GuildID == s.GuildID)
	}
	for _, g := range s.guilds() {
		if c, ok := findFirst(st.GuildChannels(g), func(c *discord.Channel) bool { return equalPtr(c.Name, argument) }); ok {
			return c, true
		}
	}
	return discord.Channel{}, false
}
