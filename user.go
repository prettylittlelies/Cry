package cry

import (
	"context"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
)

type User struct {
	client *Client
	ID     snowflake.ID
}

func (c *Client) User(id snowflake.ID) User {
	return User{client: c, ID: id}
}

func (u User) Cached() (discord.PartialUser, bool) {
	return u.client.state.User(u.ID)
}

func (u User) Fetch(ctx context.Context) (*discord.User, error) {
	return u.client.api.GetUser(ctx, u.ID)
}

func (u User) Profile(ctx context.Context, query discord.GetUserProfileQuery) (*discord.GetUserProfileResponse, error) {
	return u.client.api.GetUserProfile(ctx, u.ID, query)
}

func (u User) DM(ctx context.Context) (Channel, error) {
	if id, ok := u.client.state.DirectMessage(u.ID); ok {
		return u.client.Channel(id), nil
	}
	ch, err := u.client.api.CreatePrivateChannel(ctx, discord.CreatePrivateChannelParams{Recipients: []snowflake.ID{u.ID}})
	if err != nil {
		return Channel{}, err
	}
	u.client.state.Apply(ch)
	return u.client.Channel(ch.ID), nil
}

func (u User) Send(ctx context.Context, content string) (*discord.Message, error) {
	dm, err := u.DM(ctx)
	if err != nil {
		return nil, err
	}
	return dm.Send(ctx, content)
}

func (u User) Relationship() (discord.Relationship, bool) {
	return u.client.state.Relationship(u.ID)
}

func (u User) Friend(ctx context.Context) error {
	return u.client.api.CreateRelationship(ctx, u.ID, discord.CreateRelationshipParams{})
}

func (u User) Block(ctx context.Context) error {
	return u.client.api.CreateRelationship(ctx, u.ID, discord.CreateRelationshipParams{Type: discord.RelationshipTypeBlocked})
}

func (u User) RemoveRelationship(ctx context.Context) error {
	return u.client.api.RemoveRelationship(ctx, u.ID)
}

func (u User) Note(ctx context.Context) (string, error) {
	n, err := u.client.api.GetUserNote(ctx, u.ID)
	if err != nil {
		return "", err
	}
	return n.Note, nil
}

func (u User) SetNote(ctx context.Context, note string) error {
	return u.client.api.ModifyUserNote(ctx, u.ID, discord.ModifyUserNoteParams{Note: &note})
}
