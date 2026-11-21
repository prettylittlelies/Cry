package converting

import (
	"regexp"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
)

var channelLinkPattern = regexp.MustCompile(`^https?://(?:(?:ptb|canary|www)\.)?discord(?:app)?\.com/channels/(?:[0-9]{15,20}|@me)/([0-9]{15,20})(?:/(?:[0-9]{15,20})/?)?$`)

func matchChannel(argument string) (snowflake.ID, bool) {
	if id, ok := matchID(argument, channelPattern); ok {
		return id, true
	}
	if m := channelLinkPattern.FindStringSubmatch(argument); m != nil {
		return parseID(m[1])
	}
	return 0, false
}

func (s Scope) resolveChannel(argument string, accept func(*discord.Channel) bool) (discord.Channel, bool) {
	st := s.Client.State()
	if id, ok := matchChannel(argument); ok {
		c, ok := st.Channel(id)
		return c, ok && c.GuildID != 0 && (s.GuildID == 0 || c.GuildID == s.GuildID) && accept(&c)
	}
	for _, g := range s.guilds() {
		if c, ok := findFirst(st.GuildChannels(g), func(c *discord.Channel) bool { return equalPtr(c.Name, argument) && accept(c) }); ok {
			return c, true
		}
	}
	return discord.Channel{}, false
}

func (s Scope) Channel(argument string) (discord.Channel, bool) {
	return s.resolveChannel(argument, func(c *discord.Channel) bool { return !c.Type.IsThread() })
}

func (s Scope) ChannelOfType(argument string, types ...discord.ChannelType) (discord.Channel, bool) {
	return s.resolveChannel(argument, func(c *discord.Channel) bool {
		for _, t := range types {
			if c.Type == t {
				return true
			}
		}
		return false
	})
}

type Thread struct{ discord.Channel }

func (s Scope) Thread(argument string) (Thread, bool) {
	if s.GuildID == 0 {
		return Thread{}, false
	}
	c, ok := s.resolveChannel(argument, func(c *discord.Channel) bool { return c.Type.IsThread() })
	return Thread{c}, ok
}
