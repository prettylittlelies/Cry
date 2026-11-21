package converting

import (
	"strings"

	"github.com/prettylittlelies/cry/discord"
)

func (s Scope) User(argument string) (discord.PartialUser, bool) {
	st := s.Client.State()
	if id, ok := matchID(argument, userMentionPattern); ok {
		if u, ok := st.User(id); ok {
			return u, true
		}
		u, err := s.Client.API().GetUser(s.Context, id)
		if err != nil {
			return discord.PartialUser{}, false
		}
		return u.Partial(), true
	}
	arg := strings.TrimPrefix(argument, "@")
	if name, disc, ok := splitTag(arg); ok {
		for _, g := range st.Guilds() {
			if m, ok := findFirst(st.Members(g), func(m *discord.GuildMember) bool {
				return m.User.Username == name && m.User.Discriminator == disc
			}); ok {
				return m.User, true
			}
		}
	}
	return st.UserByName(arg)
}

func (s Scope) Member(argument string) (discord.GuildMember, bool) {
	st := s.Client.State()
	id, byID := matchID(argument, userMentionPattern)
	for _, g := range s.guilds() {
		if byID {
			if m, ok := st.Member(g, id); ok {
				return m, true
			}
		} else if m, ok := memberByName(st.Members(g), argument); ok {
			return m, true
		}
	}
	if !byID || s.GuildID == 0 {
		return discord.GuildMember{}, false
	}
	m, err := s.Client.API().GetGuildMember(s.Context, s.GuildID, id)
	if err != nil {
		return discord.GuildMember{}, false
	}
	return *m, true
}

func memberByName(members []discord.GuildMember, argument string) (discord.GuildMember, bool) {
	if name, disc, ok := splitTag(argument); ok {
		if m, ok := findFirst(members, func(m *discord.GuildMember) bool {
			return m.User.Username == name && m.User.Discriminator == disc
		}); ok {
			return m, true
		}
	}
	return findFirst(members, func(m *discord.GuildMember) bool {
		return m.User.Username == argument || equalPtr(m.User.GlobalName, argument) || equalPtr(m.Nick, argument)
	})
}
