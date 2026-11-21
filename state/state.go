package state

import (
	"sync"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
)

type State struct {
	mu            sync.RWMutex
	me            discord.User
	sessionID     string
	users         map[snowflake.ID]*discord.PartialUser
	guilds        map[snowflake.ID]*discord.GatewayGuild
	channels      map[snowflake.ID]*discord.Channel
	members       map[snowflake.ID]map[snowflake.ID]*discord.GuildMember
	relationships map[snowflake.ID]*discord.Relationship
	presences     map[snowflake.ID]*discord.PresenceUpdateEvent
	messages      *messageCache
}

func New(messageLimit int) *State {
	return &State{
		users:         map[snowflake.ID]*discord.PartialUser{},
		guilds:        map[snowflake.ID]*discord.GatewayGuild{},
		channels:      map[snowflake.ID]*discord.Channel{},
		members:       map[snowflake.ID]map[snowflake.ID]*discord.GuildMember{},
		relationships: map[snowflake.ID]*discord.Relationship{},
		presences:     map[snowflake.ID]*discord.PresenceUpdateEvent{},
		messages:      newMessageCache(messageLimit),
	}
}
