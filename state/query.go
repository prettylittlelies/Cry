package state

import (
	"maps"
	"slices"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
)

func lookup[T any](s *State, m map[snowflake.ID]*T, id snowflake.ID) (T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := m[id]
	if !ok {
		var zero T
		return zero, false
	}
	return *v, true
}

func (s *State) Me() discord.User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.me
}

func (s *State) SessionID() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.sessionID
}

func (s *State) User(id snowflake.ID) (discord.PartialUser, bool) {
	return lookup(s, s.users, id)
}

func (s *State) Guild(id snowflake.ID) (discord.GatewayGuild, bool) {
	return lookup(s, s.guilds, id)
}

func (s *State) Channel(id snowflake.ID) (discord.Channel, bool) {
	return lookup(s, s.channels, id)
}

func (s *State) Relationship(user snowflake.ID) (discord.Relationship, bool) {
	return lookup(s, s.relationships, user)
}

func (s *State) Presence(user snowflake.ID) (discord.PresenceUpdateEvent, bool) {
	return lookup(s, s.presences, user)
}

func (s *State) Guilds() []snowflake.ID {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return slices.Collect(maps.Keys(s.guilds))
}

func (s *State) Member(guild, user snowflake.ID) (discord.GuildMember, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	m, ok := s.members[guild][user]
	if !ok {
		return discord.GuildMember{}, false
	}
	return *m, true
}

func (s *State) Message(id snowflake.ID) (discord.MessageCreateEvent, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.messages.get(id)
}
