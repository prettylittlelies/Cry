package gateway

import (
	"context"
	"slices"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
)

type MemberRequest struct {
	GuildIDs  []snowflake.ID `json:"guild_id"`
	Query     *string        `json:"query"`
	Limit     *int           `json:"limit"`
	Presences bool           `json:"presences"`
	UserIDs   []snowflake.ID `json:"user_ids"`
	Nonce     string         `json:"nonce,omitempty"`
}

type GuildSubscription struct {
	Typing            *bool                     `json:"typing,omitempty"`
	Threads           *bool                     `json:"threads,omitempty"`
	Activities        *bool                     `json:"activities,omitempty"`
	MemberUpdates     *bool                     `json:"member_updates,omitempty"`
	Members           []snowflake.ID            `json:"members,omitempty"`
	Channels          map[snowflake.ID][][2]int `json:"channels,omitempty"`
	ThreadMemberLists []snowflake.ID            `json:"thread_member_lists,omitempty"`
}

type guildSubscriptionsBulk struct {
	Subscriptions map[snowflake.ID]GuildSubscription `json:"subscriptions"`
}

func (g *Gateway) Presence() Presence {
	g.mu.Lock()
	defer g.mu.Unlock()
	p := g.presence
	p.Activities = slices.Clone(p.Activities)
	return p
}

func (g *Gateway) SetPresence(ctx context.Context, p Presence) error {
	if p.Activities == nil {
		p.Activities = []discord.Activity{}
	}
	if p.Status == "" {
		p.Status = discord.StatusTypeOnline
	}
	g.mu.Lock()
	g.presence = p
	g.mu.Unlock()
	return g.Send(ctx, OpPresenceUpdate, p)
}

func (g *Gateway) RequestMembers(ctx context.Context, r MemberRequest) error {
	return g.Send(ctx, OpRequestGuildMembers, r)
}

func (g *Gateway) SubscribeGuilds(ctx context.Context, subscriptions map[snowflake.ID]GuildSubscription) error {
	return g.Send(ctx, OpGuildSubscriptionsBulk, guildSubscriptionsBulk{Subscriptions: subscriptions})
}

func (g *Gateway) SearchRecentMembers(ctx context.Context, r discord.SearchRecentMembers) error {
	return g.Send(ctx, OpSearchRecentMembers, r)
}

func (g *Gateway) RequestLastMessages(ctx context.Context, r discord.RequestLastMessages) error {
	return g.Send(ctx, OpRequestLastMessages, r)
}

func (g *Gateway) RequestChannelStatuses(ctx context.Context, guild snowflake.ID) error {
	return g.Send(ctx, OpRequestChannelStatuses, discord.RequestChannelStatuses{GuildID: guild})
}

func (g *Gateway) RequestChannelMemberCount(ctx context.Context, guild, channel snowflake.ID) error {
	return g.Send(ctx, OpRequestChannelMemberCount, discord.RequestChannelMemberCount{GuildID: guild, ChannelID: channel})
}

func (g *Gateway) RequestChannelInfo(ctx context.Context, r discord.RequestChannelInfo) error {
	return g.Send(ctx, OpRequestChannelInfo, r)
}

func (g *Gateway) RequestSoundboardSounds(ctx context.Context, guilds ...snowflake.ID) error {
	return g.Send(ctx, OpRequestSoundboardSounds, discord.RequestSoundboardSounds{GuildIDs: guilds})
}
