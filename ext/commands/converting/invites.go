package converting

import (
	"net/url"
	"path"
	"regexp"

	"github.com/prettylittlelies/cry/discord"
)

var (
	inviteURLPattern  = regexp.MustCompile(`^(?:https?://)?discord(?:\.gg|(?:app)?\.com/invite)/[^/]+`)
	inviteCodePattern = regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`)
)

type ResolvedInvite struct {
	Code    string
	EventID string
}

func ResolveInvite(argument string) (ResolvedInvite, bool) {
	if !inviteURLPattern.MatchString(argument) {
		return ResolvedInvite{Code: argument}, inviteCodePattern.MatchString(argument)
	}
	raw := argument
	if u, err := url.Parse(argument); err == nil && u.Scheme == "" {
		raw = "https://" + argument
	}
	u, err := url.Parse(raw)
	if err != nil {
		return ResolvedInvite{}, false
	}
	return ResolvedInvite{Code: path.Base(u.Path), EventID: u.Query().Get("event")}, true
}

func (s Scope) Invite(argument string) (discord.Invite, bool) {
	resolved, ok := ResolveInvite(argument)
	if !ok {
		return discord.Invite{}, false
	}
	query := discord.GetInviteQuery{WithCounts: true}
	if resolved.EventID != "" {
		query.GuildScheduledEventID, ok = parseID(resolved.EventID)
		if !ok {
			return discord.Invite{}, false
		}
	}
	invite, err := s.Client.API().GetInvite(s.Context, resolved.Code, query)
	if err != nil {
		return discord.Invite{}, false
	}
	return *invite, true
}
