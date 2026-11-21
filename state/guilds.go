package state

import (
	"slices"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
)

func (s *State) applyGuild(event any) {
	switch e := event.(type) {
	case *discord.GuildCreateEvent:
		s.putGuild(e.AsGatewayGuild(), nil)
	case *discord.GuildUpdateEvent:
		if g, ok := s.guilds[e.ID]; ok {
			remarshal(e, &g.Properties)
		}
	case *discord.GuildDeleteEvent:
		s.deleteGuild(e.ID, e.Unavailable)
	case *discord.GuildRoleCreateEvent:
		s.putRole(e.GuildID, e.Role)
	case *discord.GuildRoleUpdateEvent:
		s.putRole(e.GuildID, e.Role)
	case *discord.GuildRoleDeleteEvent:
		s.deleteRole(e.GuildID, e.RoleID)
	case *discord.GuildEmojisUpdateEvent:
		if g, ok := s.guilds[e.GuildID]; ok {
			g.Emojis = e.Emojis
		}
	case *discord.GuildStickersUpdateEvent:
		if g, ok := s.guilds[e.GuildID]; ok {
			g.Stickers = e.Stickers
		}
	case *discord.GuildScheduledEventCreateEvent:
		s.putScheduledEvent(e.AsGuildScheduledEvent())
	case *discord.GuildScheduledEventUpdateEvent:
		s.putScheduledEvent(e.AsGuildScheduledEvent())
	case *discord.GuildScheduledEventDeleteEvent:
		if g, ok := s.guilds[e.GuildID]; ok {
			g.GuildScheduledEvents = deleteWhere(g.GuildScheduledEvents, func(x discord.GuildScheduledEvent) bool { return x.ID == e.ID })
		}
	case *discord.GuildMemberAddEvent:
		s.putMember(e.GuildID, e.AsGuildMember())
	case *discord.GuildMemberUpdateEvent:
		var m discord.GuildMember
		remarshal(e, &m)
		s.putMember(e.GuildID, m)
	case *discord.GuildMemberRemoveEvent:
		delete(s.members[e.GuildID], e.User.ID)
	case *discord.GuildMembersChunkEvent:
		for _, m := range e.Members {
			s.putMember(e.GuildID, m)
		}
		for _, p := range e.Presences {
			p.GuildID = e.GuildID
			s.putPresence(p)
		}
	case *discord.VoiceStateUpdateEvent:
		s.putVoiceState(idOf(e.GuildID), e.AsVoiceState())
	case *discord.PassiveUpdateV2Event:
		for _, v := range e.UpdatedVoiceStates {
			s.putVoiceState(e.GuildID, v)
		}
		for _, user := range e.RemovedVoiceStates {
			delete(s.voice[e.GuildID], user)
		}
		for _, m := range e.UpdatedMembers {
			s.putMember(e.GuildID, m)
		}
	}
}

func (s *State) putGuild(g discord.GatewayGuild, merged []discord.GuildMember) {
	for _, c := range append(g.Channels, g.Threads...) {
		c.GuildID = g.ID
		s.putChannel(c)
	}
	for _, m := range append(g.Members, merged...) {
		s.putMember(g.ID, m)
	}
	for _, v := range g.VoiceStates {
		s.putVoiceState(g.ID, v)
	}
	s.guilds[g.ID] = &g
}

func (s *State) deleteGuild(id snowflake.ID, unavailable bool) {
	if g, ok := s.guilds[id]; ok && unavailable {
		g.Unavailable = true
		return
	}
	delete(s.guilds, id)
	delete(s.members, id)
	delete(s.voice, id)
	for cid, c := range s.channels {
		if c.GuildID == id {
			delete(s.channels, cid)
		}
	}
}

func (s *State) putRole(guild snowflake.ID, r discord.Role) {
	g, ok := s.guilds[guild]
	if !ok {
		return
	}
	if i := slices.IndexFunc(g.Roles, func(o discord.Role) bool { return o.ID == r.ID }); i >= 0 {
		g.Roles[i] = r
		return
	}
	g.Roles = append(g.Roles, r)
}

func (s *State) deleteRole(guild, role snowflake.ID) {
	if g, ok := s.guilds[guild]; ok {
		g.Roles = deleteWhere(g.Roles, func(r discord.Role) bool { return r.ID == role })
	}
	for _, m := range s.members[guild] {
		m.Roles = deleteWhere(m.Roles, func(id snowflake.ID) bool { return id == role })
	}
}

func (s *State) putMember(guild snowflake.ID, m discord.GuildMember) {
	if s.members[guild] == nil {
		s.members[guild] = map[snowflake.ID]*discord.GuildMember{}
	}
	id := m.User.ID
	if id == 0 {
		id = m.UserID
	}
	if m.User.ID != 0 {
		u := m.User
		s.users[u.ID] = &u
	}
	s.members[guild][id] = &m
}

func (s *State) putVoiceState(guild snowflake.ID, v discord.VoiceState) {
	if v.ChannelID == nil {
		delete(s.voice[guild], v.UserID)
		return
	}
	if s.voice[guild] == nil {
		s.voice[guild] = map[snowflake.ID]*discord.VoiceState{}
	}
	s.voice[guild][v.UserID] = &v
}

func (s *State) putScheduledEvent(e discord.GuildScheduledEvent) {
	g, ok := s.guilds[e.GuildID]
	if !ok {
		return
	}
	if i := slices.IndexFunc(g.GuildScheduledEvents, func(x discord.GuildScheduledEvent) bool { return x.ID == e.ID }); i >= 0 {
		g.GuildScheduledEvents[i] = e
		return
	}
	g.GuildScheduledEvents = append(g.GuildScheduledEvents, e)
}
