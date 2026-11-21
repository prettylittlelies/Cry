package converting

import (
	"regexp"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
)

var cleanPattern = regexp.MustCompile(`<(@[!&]?|#)([0-9]{15,20})>`)

type CleanOptions struct {
	FixChannelMentions bool
	UseNicknames       bool
	EscapeMarkdown     bool
	RemoveMarkdown     bool
}

func DefaultCleanOptions() CleanOptions {
	return CleanOptions{UseNicknames: true}
}

func (s Scope) Clean(opts CleanOptions, mentions []discord.PartialUser, argument string) string {
	result := cleanPattern.ReplaceAllStringFunc(argument, func(m string) string {
		sub := cleanPattern.FindStringSubmatch(m)
		id, _ := parseID(sub[2])
		switch sub[1] {
		case "#":
			return s.cleanChannel(opts, id)
		case "@&":
			return s.cleanRole(id)
		}
		return s.cleanMember(opts, mentions, id)
	})
	switch {
	case opts.EscapeMarkdown:
		result = discord.EscapeMarkdown(result, false, true)
	case opts.RemoveMarkdown:
		result = discord.RemoveMarkdown(result, true)
	}
	return discord.EscapeMentions(result)
}

func (s Scope) cleanMember(opts CleanOptions, mentions []discord.PartialUser, id snowflake.ID) string {
	st := s.Client.State()
	user, ok := findFirst(mentions, func(u *discord.PartialUser) bool { return u.ID == id })
	if s.GuildID == 0 {
		if !ok {
			user, ok = st.User(id)
		}
		if !ok {
			return "@deleted-user"
		}
		return "@" + user.DisplayName()
	}
	member, isMember := st.Member(s.GuildID, id)
	switch {
	case !ok && !isMember:
		return "@deleted-user"
	case !ok:
		user = member.User
	}
	if !opts.UseNicknames {
		return "@" + user.Username
	}
	if isMember && member.Nick != nil {
		return "@" + *member.Nick
	}
	return "@" + user.DisplayName()
}

func (s Scope) cleanRole(id snowflake.ID) string {
	if g, ok := s.Client.State().Guild(s.GuildID); ok {
		if r, ok := findFirst(g.Roles, func(r *discord.Role) bool { return r.ID == id }); ok {
			return "@" + r.Name
		}
	}
	return "@deleted-role"
}

func (s Scope) cleanChannel(opts CleanOptions, id snowflake.ID) string {
	if !opts.FixChannelMentions || s.GuildID == 0 {
		return "<#" + id.String() + ">"
	}
	if c, ok := s.Client.State().Channel(id); ok && c.GuildID == s.GuildID && c.Name != nil {
		return "#" + *c.Name
	}
	return "#deleted-channel"
}
