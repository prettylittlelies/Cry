package state

import (
	"encoding/json"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
)

func (s *State) Apply(event any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	switch e := event.(type) {
	case *discord.ReadyEvent:
		s.ready(e)
	case *discord.ReadySupplementalEvent:
		s.readySupplemental(e)
	case *discord.UserUpdateEvent:
		s.me = e.AsUser()
	default:
		s.applyGuild(event)
		s.applyChannel(event)
		s.applyMessage(event)
		s.applySocial(event)
	}
}

func remarshal(src, dst any) {
	raw, err := json.Marshal(src)
	if err == nil {
		_ = json.Unmarshal(raw, dst)
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
		s.putRelationship(e.Relationships[i])
	}
}

func (s *State) readySupplemental(e *discord.ReadySupplementalEvent) {
	for _, p := range e.MergedPresences.Friends {
		s.putPresence(p)
	}
	for i, presences := range e.MergedPresences.Guilds {
		for _, p := range presences {
			if p.GuildID == 0 && i < len(e.Guilds) {
				p.GuildID = e.Guilds[i].ID
			}
			s.putPresence(p)
		}
	}
	for i, g := range e.Guilds {
		for _, v := range g.VoiceStates {
			s.putVoiceState(g.ID, v)
		}
		if i < len(e.MergedMembers) {
			for _, raw := range e.MergedMembers[i] {
				var m discord.GuildMember
				if json.Unmarshal(raw, &m) == nil {
					s.putMember(g.ID, m)
				}
			}
		}
	}
	for _, c := range e.LazyPrivateChannels {
		s.putChannel(c)
	}
}

func (s *State) putChannel(c discord.Channel) {
	s.channels[c.ID] = &c
}

func (s *State) applyChannel(event any) {
	switch e := event.(type) {
	case *discord.Channel:
		s.putChannel(*e)
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
	case *discord.ThreadListSyncEvent:
		for _, c := range e.Threads {
			c.GuildID = e.GuildID
			s.putChannel(c)
		}
	case *discord.ChannelRecipientAddEvent:
		if c, ok := s.channels[e.ChannelID]; ok {
			c.Recipients = append(c.Recipients, e.User)
		}
	case *discord.ChannelRecipientRemoveEvent:
		if c, ok := s.channels[e.ChannelID]; ok {
			c.Recipients = deleteWhere(c.Recipients, func(u discord.PartialUser) bool { return u.ID == e.User.ID })
		}
	}
}

func deleteWhere[T any](items []T, match func(T) bool) []T {
	out := items[:0]
	for _, v := range items {
		if !match(v) {
			out = append(out, v)
		}
	}
	return out
}

func idOf(id *snowflake.ID) snowflake.ID {
	if id == nil {
		return 0
	}
	return *id
}
