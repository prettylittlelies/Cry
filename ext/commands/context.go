package commands

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
)

type Context struct {
	context.Context
	Bot         *Bot
	Message     *discord.MessageCreateEvent
	Prefix      string
	InvokedWith string
	Command     *Command
}

func nonce() json.RawMessage {
	return json.RawMessage(strconv.Quote(snowflake.FromTime(time.Now()).String()))
}

func (c *Context) send(params discord.CreateMessageParams) (*discord.Message, error) {
	params.Nonce = nonce()
	return c.Bot.client.API().CreateMessage(c, c.Message.ChannelID, params)
}

func (c *Context) Send(content string) (*discord.Message, error) {
	return c.send(discord.CreateMessageParams{Content: content})
}

func (c *Context) Reply(content string) (*discord.Message, error) {
	return c.send(discord.CreateMessageParams{
		Content: content,
		MessageReference: discord.MessageReference{
			MessageID: c.Message.ID,
			ChannelID: c.Message.ChannelID,
			GuildID:   c.Message.GuildID,
		},
	})
}

func (c *Context) Author() discord.PartialUser {
	return c.Message.Author
}
