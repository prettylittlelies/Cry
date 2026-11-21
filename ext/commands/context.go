package commands

import (
	"context"
	"github.com/prettylittlelies/cry"
	"regexp"
	"strings"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
)

type Context struct {
	context.Context
	Message           *discord.MessageCreateEvent
	Bot               *Bot
	View              *StringView
	Prefix            string
	HasPrefix         bool
	Command           *Command
	InvokedWith       string
	InvokedParents    []string
	InvokedSubcommand *Command
	SubcommandPassed  string
	CommandFailed     bool
	CurrentParameter  *Parameter
	CurrentArgument   string

	attachments int
}

type CommandEvent struct{ Context *Context }

type CommandCompletionEvent struct{ Context *Context }

type CommandErrorEvent struct {
	Context *Context
	Err     error
}

func (c *Context) Valid() bool {
	return c.HasPrefix && c.Command != nil
}

func (c *Context) CleanPrefix() string {
	if !c.HasPrefix {
		return ""
	}
	me := c.Bot.client.State().Me()
	pattern := regexp.MustCompile(`<@!?` + me.ID.String() + `>`)
	return pattern.ReplaceAllLiteralString(c.Prefix, "@"+strings.ReplaceAll(me.DisplayName(), `\`, `\\`))
}

func (c *Context) Cog() Cog {
	if c.Command == nil {
		return nil
	}
	return c.Command.Cog
}

func (c *Context) Author() discord.PartialUser {
	return c.Message.Author
}

func (c *Context) AuthorID() snowflake.ID {
	return c.Message.Author.ID
}

func (c *Context) GuildID() snowflake.ID {
	return c.Message.GuildID
}

func (c *Context) ChannelID() snowflake.ID {
	return c.Message.ChannelID
}

func (c *Context) Private() bool {
	return c.Message.GuildID == 0
}

func (c *Context) CategoryID() snowflake.ID {
	ch, ok := c.Bot.client.State().Channel(c.Message.ChannelID)
	if !ok || ch.ParentID == nil {
		return 0
	}
	return *ch.ParentID
}

func (c *Context) TopRoleID() snowflake.ID {
	st := c.Bot.client.State()
	member, ok := st.Member(c.Message.GuildID, c.Message.Author.ID)
	guild, gok := st.Guild(c.Message.GuildID)
	if !ok || !gok {
		return c.Message.GuildID
	}
	top, pos := c.Message.GuildID, int64(-1)
	for _, r := range guild.Roles {
		for _, id := range member.Roles {
			if r.ID == id && r.Position > pos {
				top, pos = r.ID, r.Position
			}
		}
	}
	return top
}

func (c *Context) Channel() cry.Channel {
	return c.Bot.client.Channel(c.Message.ChannelID)
}

func (c *Context) MessageHandle() cry.Message {
	return c.Bot.client.Message(c.Message.ChannelID, c.Message.ID)
}

func (c *Context) SendMessage(params discord.CreateMessageParams) (*discord.Message, error) {
	return c.Channel().SendMessage(c, params)
}

func (c *Context) SendAuthor(content string) (*discord.Message, error) {
	return c.Bot.client.User(c.Message.Author.ID).Send(c, content)
}

func (c *Context) Send(content string) (*discord.Message, error) {
	return c.SendMessage(discord.CreateMessageParams{Content: content})
}

func (c *Context) Reply(content string) (*discord.Message, error) {
	return c.MessageHandle().ReplyMessage(c, discord.CreateMessageParams{
		Content:          content,
		MessageReference: discord.MessageReference{GuildID: c.Message.GuildID},
	})
}

func (c *Context) Reinvoke(callHooks, restart bool) error {
	cmd := c.Command
	if cmd == nil {
		return errInvalidContext
	}
	view := c.View
	index, previous := view.Index, view.Previous
	invokedWith, invokedSub, parents := c.InvokedWith, c.InvokedSubcommand, c.InvokedParents
	subPassed := c.SubcommandPassed
	if restart {
		toCall := cmd.Root()
		view.Index = len(c.Prefix)
		view.Previous = 0
		c.InvokedParents = nil
		c.InvokedWith = view.GetWord()
		cmd = toCall
	}
	err := cmd.reinvoke(c, callHooks)
	c.Command = cmd
	view.Index, view.Previous = index, previous
	c.InvokedWith, c.InvokedSubcommand, c.InvokedParents, c.SubcommandPassed = invokedWith, invokedSub, parents, subPassed
	return err
}
