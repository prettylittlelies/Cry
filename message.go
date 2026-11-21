package cry

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
)

type Message struct {
	client    *Client
	ChannelID snowflake.ID
	ID        snowflake.ID
}

func (c *Client) Message(channel, id snowflake.ID) Message {
	return Message{client: c, ChannelID: channel, ID: id}
}

func (m Message) Channel() Channel {
	return m.client.Channel(m.ChannelID)
}

func (m Message) Cached() (discord.MessageCreateEvent, bool) {
	return m.client.state.Message(m.ID)
}

func (m Message) Fetch(ctx context.Context) (discord.Message, error) {
	raw, err := m.client.api.GetMessage(ctx, m.ChannelID, m.ID)
	if err != nil {
		return discord.Message{}, err
	}
	var out discord.Message
	return out, json.Unmarshal(raw, &out)
}

func (m Message) ReplyMessage(ctx context.Context, params discord.CreateMessageParams) (*discord.Message, error) {
	guild := params.MessageReference.GuildID
	if cached, ok := m.Cached(); ok && guild == 0 {
		guild = cached.GuildID
	}
	params.MessageReference = discord.MessageReference{MessageID: m.ID, ChannelID: m.ChannelID, GuildID: guild}
	return m.Channel().SendMessage(ctx, params)
}

func (m Message) Reply(ctx context.Context, content string) (*discord.Message, error) {
	return m.ReplyMessage(ctx, discord.CreateMessageParams{Content: content})
}

func (m Message) EditMessage(ctx context.Context, params discord.EditMessageParams) (*discord.Message, error) {
	return m.client.api.EditMessage(ctx, m.ChannelID, m.ID, params)
}

func (m Message) Edit(ctx context.Context, content string) (*discord.Message, error) {
	return m.EditMessage(ctx, discord.EditMessageParams{Content: &content})
}

func (m Message) Delete(ctx context.Context) error {
	return m.client.api.DeleteMessage(ctx, m.ChannelID, m.ID)
}

func (m Message) Pin(ctx context.Context) error {
	return m.client.api.PinMessage(ctx, m.ChannelID, m.ID)
}

func (m Message) Unpin(ctx context.Context) error {
	return m.client.api.UnpinMessage(ctx, m.ChannelID, m.ID)
}

func ReactionEmoji(emoji string) string {
	emoji = strings.TrimSuffix(strings.TrimPrefix(emoji, "<"), ">")
	emoji = strings.TrimPrefix(emoji, "a:")
	return strings.TrimPrefix(emoji, ":")
}

func (m Message) React(ctx context.Context, emoji string) error {
	return m.client.api.CreateReaction(ctx, m.ChannelID, m.ID, ReactionEmoji(emoji), discord.CreateReactionQuery{})
}

func (m Message) SuperReact(ctx context.Context, emoji string) error {
	return m.client.api.CreateReaction(ctx, m.ChannelID, m.ID, ReactionEmoji(emoji), discord.CreateReactionQuery{Type: discord.ReactionTypeBurst})
}

func (m Message) Unreact(ctx context.Context, emoji string) error {
	return m.client.api.DeleteOwnReaction(ctx, m.ChannelID, m.ID, ReactionEmoji(emoji))
}

func (m Message) RemoveReaction(ctx context.Context, emoji string, user snowflake.ID) error {
	return m.client.api.DeleteUserReaction(ctx, m.ChannelID, m.ID, ReactionEmoji(emoji), user)
}

func (m Message) ClearReaction(ctx context.Context, emoji string) error {
	return m.client.api.DeleteReactionEmoji(ctx, m.ChannelID, m.ID, ReactionEmoji(emoji))
}

func (m Message) ClearReactions(ctx context.Context) error {
	return m.client.api.DeleteAllReactions(ctx, m.ChannelID, m.ID)
}
