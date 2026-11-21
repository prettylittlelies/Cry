package commands

import (
	"errors"
	"strings"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/ext/commands/converting"
	"github.com/prettylittlelies/cry/snowflake"
)

type CleanContent string

func (c *Context) Scope() converting.Scope {
	return converting.Scope{Context: c, Client: c.Bot.client, GuildID: c.GuildID()}
}

func lookup[T any](kind string, find func(converting.Scope, string) (T, bool)) func(*Context, string) (T, error) {
	return func(ctx *Context, argument string) (T, error) {
		v, ok := find(ctx.Scope(), argument)
		if !ok {
			return v, notFound(kind, argument)
		}
		return v, nil
	}
}

func guildOnly[T any](next func(*Context, string) (T, error)) func(*Context, string) (T, error) {
	return func(ctx *Context, argument string) (T, error) {
		if ctx.Private() {
			var zero T
			return zero, errNoPrivateMessage()
		}
		return next(ctx, argument)
	}
}

func parsed[T any](parse func(string) (T, bool), fail func(string) error) func(*Context, string) (T, error) {
	return func(_ *Context, argument string) (T, error) {
		v, ok := parse(argument)
		if !ok {
			return v, fail(argument)
		}
		return v, nil
	}
}

func located[T any](find func(converting.Scope, string, snowflake.ID) (T, error)) func(*Context, string) (T, error) {
	return func(ctx *Context, argument string) (T, error) {
		v, err := find(ctx.Scope(), argument, ctx.ChannelID())
		switch {
		case errors.Is(err, converting.ErrChannelNotFound):
			return v, notFound("Channel", argument)
		case errors.Is(err, converting.ErrChannelNotReadable):
			return v, errChannelNotReadable(argument)
		case errors.Is(err, converting.ErrNotFound):
			return v, notFound("Message", argument)
		}
		return v, err
	}
}

func cleanOptions(p *Parameter) converting.CleanOptions {
	opts := converting.DefaultCleanOptions()
	if p == nil {
		return opts
	}
	for opt := range strings.SplitSeq(p.Tag.Get("clean"), ",") {
		switch strings.TrimSpace(opt) {
		case "fix_channel_mentions":
			opts.FixChannelMentions = true
		case "no_nicknames":
			opts.UseNicknames = false
		case "escape_markdown":
			opts.EscapeMarkdown = true
		case "remove_markdown":
			opts.RemoveMarkdown = true
		}
	}
	return opts
}

func convertClean(ctx *Context, argument string) (CleanContent, error) {
	return CleanContent(ctx.Scope().Clean(cleanOptions(ctx.CurrentParameter), ctx.Message.Mentions, argument)), nil
}

func init() {
	RegisterConverter(lookup("User", converting.Scope.User))
	RegisterConverter(lookup("Member", converting.Scope.Member))
	RegisterConverter(guildOnly(lookup("Role", converting.Scope.Role)))
	RegisterConverter(lookup("Guild", converting.Scope.Guild))
	RegisterConverter(lookup("Channel", converting.Scope.Channel))
	RegisterConverter(lookup("Thread", converting.Scope.Thread))
	RegisterConverter(lookup("Emoji", converting.Scope.Emoji))
	RegisterConverter(located(converting.Scope.Message))
	RegisterConverter(located(converting.Scope.PartialMessage))
	RegisterConverter(func(ctx *Context, argument string) (discord.Invite, error) {
		invite, ok := ctx.Scope().Invite(argument)
		if !ok {
			return invite, errBadInvite(argument)
		}
		return invite, nil
	})
	RegisterConverter(convertClean)
	RegisterConverter(lookup("Sticker", converting.Scope.Sticker))
	RegisterConverter(lookup("Scheduled event", converting.Scope.ScheduledEvent))
	RegisterConverter(parsed(converting.ParseTimestamp, func(a string) error { return errBadTimestamp(a) }))
	RegisterConverter(parsed(converting.ParsePartialEmoji, func(a string) error { return errPartialEmoji(a) }))
	RegisterConverter(parsed(converting.ParseColour, func(a string) error { return errBadColour(a) }))
	RegisterConverter(parsed(converting.Object, func(a string) error { return errObjectNotFound(a) }))
}
