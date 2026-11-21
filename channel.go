package cry

import (
	"context"
	"encoding/json"
	"iter"
	"slices"
	"strconv"
	"time"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
)

type Channel struct {
	client *Client
	ID     snowflake.ID
}

type HistoryOptions struct {
	Limit  int
	Before snowflake.ID
	After  snowflake.ID
	Around snowflake.ID
}

const historyPage = 100

func (c *Client) Channel(id snowflake.ID) Channel {
	return Channel{client: c, ID: id}
}

func (ch Channel) Cached() (discord.Channel, bool) {
	return ch.client.state.Channel(ch.ID)
}

func (ch Channel) Fetch(ctx context.Context) (*discord.Channel, error) {
	return ch.client.api.GetChannel(ctx, ch.ID)
}

func (ch Channel) Edit(ctx context.Context, params discord.ModifyChannelParams) (*discord.Channel, error) {
	return ch.client.api.ModifyChannel(ctx, ch.ID, params)
}

func (ch Channel) Delete(ctx context.Context) error {
	_, err := ch.client.api.DeleteChannel(ctx, ch.ID, discord.DeleteChannelQuery{})
	return err
}

func (ch Channel) Message(id snowflake.ID) Message {
	return Message{client: ch.client, ChannelID: ch.ID, ID: id}
}

func nonce() json.RawMessage {
	return json.RawMessage(strconv.Quote(snowflake.FromTime(time.Now()).String()))
}

func (ch Channel) SendMessage(ctx context.Context, params discord.CreateMessageParams) (*discord.Message, error) {
	if params.Nonce == nil {
		params.Nonce = nonce()
	}
	return ch.client.api.CreateMessage(ctx, ch.ID, params)
}

func (ch Channel) Send(ctx context.Context, content string) (*discord.Message, error) {
	return ch.SendMessage(ctx, discord.CreateMessageParams{Content: content})
}

func (ch Channel) Typing(ctx context.Context) error {
	_, err := ch.client.api.TriggerTypingIndicator(ctx, ch.ID)
	return err
}

func (ch Channel) BulkDelete(ctx context.Context, ids ...snowflake.ID) error {
	return ch.client.api.BulkDeleteMessages(ctx, ch.ID, discord.BulkDeleteMessagesParams{Messages: ids})
}

func (ch Channel) Pins(ctx context.Context) ([]discord.MessagePin, error) {
	resp, err := ch.client.api.ListMessagePins(ctx, ch.ID, discord.ListMessagePinsQuery{})
	if err != nil {
		return nil, err
	}
	return resp.Items, nil
}

func (ch Channel) History(ctx context.Context, opts HistoryOptions) iter.Seq2[discord.Message, error] {
	return func(yield func(discord.Message, error) bool) {
		remaining := opts.Limit
		oldestFirst := opts.After != 0 && opts.Before == 0 && opts.Around == 0
		query := discord.ListMessagesQuery{Before: opts.Before, After: opts.After, Around: opts.Around}
		for {
			query.Limit = historyPage
			if opts.Limit > 0 {
				query.Limit = int64(min(remaining, historyPage))
			}
			page, err := ch.client.api.ListMessages(ctx, ch.ID, query)
			if err != nil {
				yield(discord.Message{}, err)
				return
			}
			if oldestFirst {
				slices.Reverse(page)
			}
			for _, m := range page {
				if opts.Limit > 0 && remaining <= 0 || !yield(m, nil) {
					return
				}
				remaining--
			}
			if len(page) < int(query.Limit) || query.Around != 0 || opts.Limit > 0 && remaining <= 0 {
				return
			}
			if oldestFirst {
				query.After = page[len(page)-1].ID
			} else {
				query.Before = page[len(page)-1].ID
			}
		}
	}
}
