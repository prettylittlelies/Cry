package commands

import (
	"errors"
	"slices"
	"strings"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
)

type RoleRef struct {
	ID   snowflake.ID
	Name string
}

func (r RoleRef) String() string {
	if r.ID != 0 {
		return r.ID.String()
	}
	return r.Name
}

func RoleID(id snowflake.ID) RoleRef { return RoleRef{ID: id} }

func RoleName(name string) RoleRef { return RoleRef{Name: name} }

func CheckAny(checks ...Check) Check {
	return func(ctx *Context) (bool, error) {
		var errs []error
		for _, c := range checks {
			ok, err := c(ctx)
			if err != nil {
				if !errors.Is(err, ErrCheckFailure) {
					return false, err
				}
				errs = append(errs, err)
				continue
			}
			if ok {
				return true, nil
			}
		}
		return false, errCheckAny(errs)
	}
}

func hasRole(ctx *Context, user snowflake.ID, ref RoleRef) bool {
	st := ctx.Bot.client.State()
	member, ok := st.Member(ctx.GuildID(), user)
	if !ok {
		return false
	}
	if ref.ID != 0 {
		return slices.Contains(member.Roles, ref.ID)
	}
	guild, ok := st.Guild(ctx.GuildID())
	if !ok {
		return false
	}
	for _, r := range guild.Roles {
		if r.Name == ref.Name && slices.Contains(member.Roles, r.ID) {
			return true
		}
	}
	return false
}

func roleCheck(bot bool, refs []RoleRef, any bool) Check {
	return func(ctx *Context) (bool, error) {
		if ctx.Private() {
			return false, errNoPrivateMessage()
		}
		user := ctx.AuthorID()
		if bot {
			user = ctx.Bot.client.State().Me().ID
		}
		names := make([]string, len(refs))
		for i, r := range refs {
			names[i] = r.String()
			if hasRole(ctx, user, r) && any {
				return true, nil
			}
		}
		if any {
			return false, errMissingAnyRole(names, bot)
		}
		if !hasRole(ctx, user, refs[0]) {
			return false, errMissingRole(names[0], bot)
		}
		return true, nil
	}
}

func HasRole(r RoleRef) Check          { return roleCheck(false, []RoleRef{r}, false) }
func HasAnyRole(r ...RoleRef) Check    { return roleCheck(false, r, true) }
func BotHasRole(r RoleRef) Check       { return roleCheck(true, []RoleRef{r}, false) }
func BotHasAnyRole(r ...RoleRef) Check { return roleCheck(true, r, true) }

func missing(have discord.Permissions, want []discord.Permission) []string {
	var out []string
	for _, p := range want {
		if !have.Has(p) {
			out = append(out, strings.ToLower(p.String()))
		}
	}
	return out
}

func permissionCheck(bot, guildLevel bool, want []discord.Permission) Check {
	return func(ctx *Context) (bool, error) {
		st := ctx.Bot.client.State()
		user := ctx.AuthorID()
		if bot {
			user = st.Me().ID
		}
		var have discord.Permissions
		switch {
		case guildLevel && ctx.Private():
			return false, errNoPrivateMessage()
		case guildLevel:
			have = st.GuildPermissions(ctx.GuildID(), user)
		default:
			have = st.ChannelPermissions(ctx.ChannelID(), user)
		}
		if m := missing(have, want); len(m) > 0 {
			return false, errMissingPermissions(m, bot)
		}
		return true, nil
	}
}

func HasPermissions(p ...discord.Permission) Check         { return permissionCheck(false, false, p) }
func BotHasPermissions(p ...discord.Permission) Check      { return permissionCheck(true, false, p) }
func HasGuildPermissions(p ...discord.Permission) Check    { return permissionCheck(false, true, p) }
func BotHasGuildPermissions(p ...discord.Permission) Check { return permissionCheck(true, true, p) }

func DMOnly() Check {
	return func(ctx *Context) (bool, error) {
		if !ctx.Private() {
			return false, errPrivateMessageOnly()
		}
		return true, nil
	}
}

func GuildOnly() Check {
	return func(ctx *Context) (bool, error) {
		if ctx.Private() {
			return false, errNoPrivateMessage()
		}
		return true, nil
	}
}

func IsOwner() Check {
	return func(ctx *Context) (bool, error) {
		ok, err := ctx.Bot.IsOwner(ctx.AuthorID())
		if err != nil {
			return false, err
		}
		if !ok {
			return false, errNotOwner()
		}
		return true, nil
	}
}

func IsNSFW() Check {
	return func(ctx *Context) (bool, error) {
		if ctx.Private() {
			return true, nil
		}
		ch, ok := ctx.Bot.client.State().Channel(ctx.ChannelID())
		if ok && ch.NSFW {
			return true, nil
		}
		name := ctx.ChannelID().String()
		if ok && ch.Name != nil {
			name = *ch.Name
		}
		return false, errNSFW(name)
	}
}
