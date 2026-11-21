package commands

import (
	"slices"

	"github.com/prettylittlelies/cry/snowflake"
)

type Check func(ctx *Context) error

func GuildOnly() Check {
	return func(ctx *Context) error {
		if ctx.Message.GuildID == 0 {
			return ErrGuildOnly
		}
		return nil
	}
}

func DMOnly() Check {
	return func(ctx *Context) error {
		if ctx.Message.GuildID != 0 {
			return ErrDMOnly
		}
		return nil
	}
}

func AllowUsers(ids ...snowflake.ID) Check {
	return func(ctx *Context) error {
		if !slices.Contains(ids, ctx.Message.Author.ID) {
			return ErrNotAllowed
		}
		return nil
	}
}
