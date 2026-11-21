package state

import (
	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
)

func (s *State) Apply(event any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	switch e := event.(type) {
	case *discord.ReadyEvent:
		s.ready(e)
	case *discord.GuildCreateEvent:
		s.putGuild(e.AsGatewayGuild(), nil)
	case *discord.GuildDeleteEvent:
		s.deleteGuild(e.ID, e.Unavailable)
	case *discord.ChannelCreateEvent:
		s.putChannel(e.AsChannel())
	case *discord.ChannelUpdateEvent:
		s.putChannel(e.AsChannel())
	case *discord.ChannelDeleteEvent:
		delete(s.channels, e.ID)
	case *discord.ThreadCreateEvent:
		s.putChannel(e.AsChannel())
	case *discord.ThreadUpdateEvent:
		s.putChannel(e.AsChannel())
	case *discord.ThreadDeleteEvent:
		delete(s.channels, e.ID)
	case *discord.GuildMemberAddEvent:
		s.putMember(e.GuildID, e.AsGuildMember())
	case *discord.GuildMemberRemoveEvent:
		delete(s.members[e.GuildID], e.User.ID)
	case *discord.RelationshipAddEvent:
		if id, err := snowflake.Parse(e.ID); err == nil {
			r := e.AsRelationship()
			s.relationships[id] = &r
		}
	case *discord.RelationshipRemoveEvent:
		delete(s.relationships, e.ID)
	case *discord.PresenceUpdateEvent:
		s.presences[e.User.ID] = e
	case *discord.UserUpdateEvent:
		s.me = e.AsUser()
	case *discord.MessageCreateEvent:
		s.messages.put(*e)
	case *discord.MessageDeleteEvent:
		s.messages.remove(e.ID)
	}
}

func (s *State) ready(e *discord.ReadyEvent) {
	s.me = e.User
	s.sessionID = e.SessionID
	for i := range e.Users {
		s.users[e.Users[i].ID] = &e.Users[i]
	}
	for i, g := range e.Guilds {
		var merged []discord.GuildMember
		if i < len(e.MergedMembers) {
			merged = e.MergedMembers[i]
		}
		s.putGuild(g, merged)
	}
	for _, c := range e.PrivateChannels {
		s.putChannel(c)
	}
	for i := range e.Relationships {
		r := &e.Relationships[i]
		if id, err := snowflake.Parse(r.ID); err == nil {
			s.relationships[id] = r
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
	s.guilds[g.ID] = &g
}

func (s *State) deleteGuild(id snowflake.ID, unavailable bool) {
	if g, ok := s.guilds[id]; ok && unavailable {
		g.Unavailable = true
		return
	}
	delete(s.guilds, id)
	delete(s.members, id)
	for cid, c := range s.channels {
		if c.GuildID == id {
			delete(s.channels, cid)
		}
	}
}

func (s *State) putChannel(c discord.Channel) {
	s.channels[c.ID] = &c
}

func (s *State) putMember(guild snowflake.ID, m discord.GuildMember) {
	if s.members[guild] == nil {
		s.members[guild] = map[snowflake.ID]*discord.GuildMember{}
	}
	id := m.User.ID
	if id == 0 {
		id = m.UserID
	}
	s.members[guild][id] = &m
}
