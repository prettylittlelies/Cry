package state

import (
	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
)

func (s *State) applySocial(event any) {
	switch e := event.(type) {
	case *discord.RelationshipAddEvent:
		s.putRelationship(e.AsRelationship())
	case *discord.RelationshipUpdateEvent:
		if r, ok := s.relationships[e.ID]; ok {
			remarshal(e, r)
		}
	case *discord.RelationshipRemoveEvent:
		delete(s.relationships, e.ID)
	case *discord.PresenceUpdateEvent:
		p := *e
		s.presences[e.User.ID] = &p
	}
}

func (s *State) putRelationship(r discord.Relationship) {
	if id, err := snowflake.Parse(r.ID); err == nil {
		s.relationships[id] = &r
		if r.User.ID != 0 {
			u := r.User
			s.users[u.ID] = &u
		}
	}
}

func (s *State) putPresence(p discord.Presence) {
	var e discord.PresenceUpdateEvent
	remarshal(p, &e)
	s.presences[p.User.ID] = &e
}
