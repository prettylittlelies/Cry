package state

import (
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
