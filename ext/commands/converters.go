package commands

import (
	"github.com/prettylittlelies/cry/ext/commands/converting"
)

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

func init() {
	RegisterConverter(lookup("User", converting.Scope.User))
	RegisterConverter(lookup("Member", converting.Scope.Member))
	RegisterConverter(guildOnly(lookup("Role", converting.Scope.Role)))
	RegisterConverter(lookup("Guild", converting.Scope.Guild))
	RegisterConverter(lookup("Channel", converting.Scope.Channel))
	RegisterConverter(lookup("Emoji", converting.Scope.Emoji))
	RegisterConverter(parsed(converting.ParsePartialEmoji, func(a string) error { return errPartialEmoji(a) }))
	RegisterConverter(parsed(converting.ParseColour, func(a string) error { return errBadColour(a) }))
	RegisterConverter(parsed(converting.Object, func(a string) error { return errObjectNotFound(a) }))
}
