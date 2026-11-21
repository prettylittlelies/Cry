package converting

import (
	"context"
	"regexp"
	"strconv"

	"github.com/prettylittlelies/cry"
	"github.com/prettylittlelies/cry/snowflake"
)

var (
	idPattern           = regexp.MustCompile(`^([0-9]{15,20})$`)
	userMentionPattern  = regexp.MustCompile(`^<@!?([0-9]{15,20})>$`)
	roleMentionPattern  = regexp.MustCompile(`^<@&([0-9]{15,20})>$`)
	channelPattern      = regexp.MustCompile(`^<#([0-9]{15,20})>$`)
	anyMentionPattern   = regexp.MustCompile(`^<(?:@[!&]?|#)([0-9]{15,20})>$`)
	emojiPattern        = regexp.MustCompile(`^<a?:[a-zA-Z0-9_~]{1,32}:([0-9]{15,20})>$`)
	partialEmojiPattern = regexp.MustCompile(`^<(a?):([a-zA-Z0-9_~]{1,32}):([0-9]{15,20})>$`)
)

type Scope struct {
	Context context.Context
	Client  *cry.Client
	GuildID snowflake.ID
}

func (s Scope) guilds() []snowflake.ID {
	if s.GuildID != 0 {
		return []snowflake.ID{s.GuildID}
	}
	return s.Client.State().Guilds()
}

func Object(argument string) (snowflake.ID, bool) {
	return matchID(argument, anyMentionPattern)
}

func matchID(argument string, mention *regexp.Regexp) (snowflake.ID, bool) {
	m := idPattern.FindStringSubmatch(argument)
	if m == nil && mention != nil {
		m = mention.FindStringSubmatch(argument)
	}
	if m == nil {
		return 0, false
	}
	return parseID(m[len(m)-1])
}

func parseID(s string) (snowflake.ID, bool) {
	n, err := strconv.ParseUint(s, 10, 64)
	return snowflake.ID(n), err == nil
}

func findFirst[T any](items []T, pred func(*T) bool) (T, bool) {
	for i := range items {
		if pred(&items[i]) {
			return items[i], true
		}
	}
	var zero T
	return zero, false
}

func splitTag(argument string) (name, discriminator string, ok bool) {
	if n := len(argument); n > 5 && argument[n-5] == '#' {
		return argument[:n-5], argument[n-4:], true
	}
	return argument, "", false
}

func equalPtr(p *string, s string) bool {
	return p != nil && *p == s
}
