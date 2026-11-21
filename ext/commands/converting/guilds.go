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
