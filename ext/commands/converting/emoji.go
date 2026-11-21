package converting

import (
	"slices"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
)

type PartialEmoji struct {
	ID       snowflake.ID
	Name     string
	Animated bool
}

func ParsePartialEmoji(argument string) (PartialEmoji, bool) {
	m := partialEmojiPattern.FindStringSubmatch(argument)
	if m == nil {
		return PartialEmoji{}, false
	}
	id, _ := parseID(m[3])
	return PartialEmoji{ID: id, Name: m[2], Animated: m[1] == "a"}, true
}

func (s Scope) Emoji(argument string) (discord.Emoji, bool) {
	pred := func(e *discord.Emoji) bool { return e.Name == argument }
	if id, ok := matchID(argument, emojiPattern); ok {
		pred = func(e *discord.Emoji) bool { return e.ID != nil && *e.ID == id }
	}
	guilds := s.Client.State().AllGuilds()
	if i := slices.IndexFunc(guilds, func(g discord.GatewayGuild) bool { return g.ID == s.GuildID }); i > 0 {
		guilds[0], guilds[i] = guilds[i], guilds[0]
	}
	for _, g := range guilds {
		if e, ok := findFirst(g.Emojis, pred); ok {
			return e, true
		}
	}
	return discord.Emoji{}, false
}
