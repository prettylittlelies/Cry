package state

import (
	"encoding/json"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
)

type messageCache struct {
	limit int
	order []snowflake.ID
	head  int
	byID  map[snowflake.ID]discord.MessageCreateEvent
}

func newMessageCache(limit int) *messageCache {
	return &messageCache{limit: limit, byID: map[snowflake.ID]discord.MessageCreateEvent{}}
}

func (c *messageCache) put(m discord.MessageCreateEvent) {
	if c.limit <= 0 {
		return
	}
	if _, ok := c.byID[m.ID]; ok {
		c.byID[m.ID] = m
		return
	}
	if len(c.order) < c.limit {
		c.order = append(c.order, m.ID)
	} else {
		delete(c.byID, c.order[c.head])
		c.order[c.head] = m.ID
		c.head = (c.head + 1) % c.limit
	}
	c.byID[m.ID] = m
}

func (c *messageCache) get(id snowflake.ID) (discord.MessageCreateEvent, bool) {
	m, ok := c.byID[id]
	return m, ok
}

func (c *messageCache) remove(id snowflake.ID) {
	delete(c.byID, id)
}

func (c *messageCache) edit(id snowflake.ID, fn func(*discord.MessageCreateEvent)) {
	if m, ok := c.byID[id]; ok {
		fn(&m)
		c.byID[id] = m
	}
}

func (s *State) applyMessage(event any) {
	switch e := event.(type) {
	case *discord.MessageCreateEvent:
		s.messages.put(*e)
	case *discord.MessageUpdateEvent:
		s.messages.edit(e.ID, func(m *discord.MessageCreateEvent) { remarshal(e, m) })
	case *discord.MessageDeleteEvent:
		s.messages.remove(e.ID)
	case *discord.MessageDeleteBulkEvent:
		for _, id := range e.IDs {
			s.messages.remove(id)
		}
	case *discord.MessageReactionAddEvent:
		burst := e.Type == discord.ReactionTypeBurst
		mine := e.UserID == s.me.ID
		s.messages.edit(e.MessageID, func(m *discord.MessageCreateEvent) { m.Reactions = addReaction(m.Reactions, e.Emoji, burst, mine) })
	case *discord.MessageReactionRemoveEvent:
		var emoji discord.Emoji
		if json.Unmarshal(e.Emoji, &emoji) != nil {
			return
		}
		burst := e.Type == discord.ReactionTypeBurst
		mine := e.UserID == s.me.ID
		s.messages.edit(e.MessageID, func(m *discord.MessageCreateEvent) { m.Reactions = removeReaction(m.Reactions, emoji, burst, mine) })
	case *discord.MessageReactionRemoveEmojiEvent:
		s.messages.edit(e.MessageID, func(m *discord.MessageCreateEvent) {
			m.Reactions = deleteWhere(m.Reactions, func(r discord.Reaction) bool { return sameEmoji(r.Emoji, e.Emoji) })
		})
	case *discord.MessageReactionRemoveAllEvent:
		s.messages.edit(e.MessageID, func(m *discord.MessageCreateEvent) { m.Reactions = nil })
	}
}
