package cry

import (
	"context"
	"time"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/gateway"
	"github.com/prettylittlelies/cry/snowflake"
)

type Guild struct {
	client *Client
	ID     snowflake.ID
}

type Member struct {
	client  *Client
	GuildID snowflake.ID
	UserID  snowflake.ID
}

func (c *Client) Guild(id snowflake.ID) Guild {
	return Guild{client: c, ID: id}
}

func (c *Client) Member(guild, user snowflake.ID) Member {
	return Member{client: c, GuildID: guild, UserID: user}
}

func (g Guild) Cached() (discord.GatewayGuild, bool) {
	return g.client.state.Guild(g.ID)
}

func (g Guild) Fetch(ctx context.Context) (*discord.Guild, error) {
	return g.client.api.GetGuild(ctx, g.ID, discord.GetGuildQuery{WithCounts: true})
}

func (g Guild) Member(user snowflake.ID) Member {
	return g.client.Member(g.ID, user)
}

func (g Guild) Members() []discord.GuildMember {
	return g.client.state.Members(g.ID)
}

func (g Guild) Roles(ctx context.Context) ([]discord.Role, error) {
	return g.client.api.ListGuildRoles(ctx, g.ID)
}

func (g Guild) Leave(ctx context.Context) error {
	return g.client.api.LeaveGuild(ctx, g.ID, discord.LeaveGuildParams{})
}

func (g Guild) SetNick(ctx context.Context, nick string) error {
	_, err := g.client.api.ModifyCurrentGuildMember(ctx, g.ID, discord.ModifyCurrentGuildMemberParams{Nick: &nick})
	return err
}

func (g Guild) QueryMembers(ctx context.Context, query string, limit int, presences bool) ([]discord.GuildMember, error) {
	return g.requestMembers(ctx, gateway.MemberRequest{Query: &query, Limit: &limit, Presences: presences})
}

func (g Guild) FetchMembers(ctx context.Context, users ...snowflake.ID) ([]discord.GuildMember, error) {
	return g.requestMembers(ctx, gateway.MemberRequest{UserIDs: users})
}

func (g Guild) requestMembers(ctx context.Context, r gateway.MemberRequest) ([]discord.GuildMember, error) {
	r.GuildIDs = []snowflake.ID{g.ID}
	r.Nonce = snowflake.FromTime(time.Now()).String()
	chunks := make(chan *discord.GuildMembersChunkEvent, 8)
	remove := On(g.client, func(_ context.Context, e *discord.GuildMembersChunkEvent) {
		if e.Nonce == r.Nonce {
			chunks <- e
		}
	})
	defer remove()
	if err := g.client.gateway.RequestMembers(ctx, r); err != nil {
		return nil, err
	}
	var out []discord.GuildMember
	for received := 0; ; {
		select {
		case <-ctx.Done():
			return out, ctx.Err()
		case e := <-chunks:
			out = append(out, e.Members...)
			if received++; int64(received) >= e.ChunkCount {
				return out, nil
			}
		}
	}
}

func (m Member) Guild() Guild {
	return m.client.Guild(m.GuildID)
}

func (m Member) User() User {
	return m.client.User(m.UserID)
}

func (m Member) Cached() (discord.GuildMember, bool) {
	return m.client.state.Member(m.GuildID, m.UserID)
}

func (m Member) Fetch(ctx context.Context) (*discord.GuildMember, error) {
	return m.client.api.GetGuildMember(ctx, m.GuildID, m.UserID)
}

func (m Member) Edit(ctx context.Context, params discord.ModifyGuildMemberParams) error {
	_, err := m.client.api.ModifyGuildMember(ctx, m.GuildID, m.UserID, params)
	return err
}

func (m Member) AddRoles(ctx context.Context, roles ...snowflake.ID) error {
	for _, r := range roles {
		if err := m.client.api.AddGuildMemberRole(ctx, m.GuildID, m.UserID, r); err != nil {
			return err
		}
	}
	return nil
}

func (m Member) RemoveRoles(ctx context.Context, roles ...snowflake.ID) error {
	for _, r := range roles {
		if err := m.client.api.RemoveGuildMemberRole(ctx, m.GuildID, m.UserID, r); err != nil {
			return err
		}
	}
	return nil
}

func (m Member) Timeout(ctx context.Context, until time.Time) error {
	return m.Edit(ctx, discord.ModifyGuildMemberParams{CommunicationDisabledUntil: &until})
}

func (m Member) Kick(ctx context.Context) error {
	return m.client.api.RemoveGuildMember(ctx, m.GuildID, m.UserID)
}

func (m Member) Ban(ctx context.Context, deleteMessages time.Duration) error {
	return m.client.api.CreateGuildBan(ctx, m.GuildID, m.UserID, discord.CreateGuildBanParams{DeleteMessageSeconds: int64(deleteMessages / time.Second)})
}

func (m Member) Unban(ctx context.Context) error {
	return m.client.api.DeleteGuildBan(ctx, m.GuildID, m.UserID)
}
