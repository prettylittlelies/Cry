package converting

import (
	"regexp"
	"strconv"
	"time"

	"github.com/prettylittlelies/cry/discord"
)

var (
	timestampPattern = regexp.MustCompile(`^<t:(-?\d+)(?::[tTdDfFsSR])?>`)
	eventLinkPattern = regexp.MustCompile(`(?i)^https?://(?:(ptb|canary|www)\.)?discord\.com/events/([0-9]{15,20})/([0-9]{15,20})$`)
)

func ParseTimestamp(argument string) (time.Time, bool) {
	m := timestampPattern.FindStringSubmatch(argument)
	if m == nil {
		return time.Time{}, false
	}
	n, err := strconv.ParseInt(m[1], 10, 64)
	return time.Unix(n, 0).UTC(), err == nil
}

func (s Scope) Sticker(argument string) (discord.Sticker, bool) {
	st := s.Client.State()
	if id, ok := matchID(argument, nil); ok {
		for _, g := range st.AllGuilds() {
			if sticker, ok := findFirst(g.Stickers, func(x *discord.Sticker) bool { return x.ID == id }); ok {
				return sticker, true
			}
		}
		return discord.Sticker{}, false
	}
	byName := func(x *discord.Sticker) bool { return x.Name == argument }
	if g, ok := st.Guild(s.GuildID); ok {
		if sticker, ok := findFirst(g.Stickers, byName); ok {
			return sticker, true
		}
	}
	for _, g := range st.AllGuilds() {
		if sticker, ok := findFirst(g.Stickers, byName); ok {
			return sticker, true
		}
	}
	return discord.Sticker{}, false
}

func (s Scope) ScheduledEvent(argument string) (discord.GuildScheduledEvent, bool) {
	st := s.Client.State()
	search := func(pred func(*discord.GuildScheduledEvent) bool) (discord.GuildScheduledEvent, bool) {
		for _, g := range s.guilds() {
			guild, _ := st.Guild(g)
			if e, ok := findFirst(guild.GuildScheduledEvents, pred); ok {
				return e, true
			}
		}
		return discord.GuildScheduledEvent{}, false
	}
	if id, ok := matchID(argument, nil); ok {
		return search(func(e *discord.GuildScheduledEvent) bool { return e.ID == id })
	}
	if m := eventLinkPattern.FindStringSubmatch(argument); m != nil {
		guildID, _ := parseID(m[2])
		eventID, _ := parseID(m[3])
		g, ok := st.Guild(guildID)
		if !ok {
			return discord.GuildScheduledEvent{}, false
		}
		return findFirst(g.GuildScheduledEvents, func(e *discord.GuildScheduledEvent) bool { return e.ID == eventID })
	}
	return search(func(e *discord.GuildScheduledEvent) bool { return e.Name == argument })
}
