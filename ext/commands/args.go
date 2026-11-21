package commands

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
)

var (
	userMention    = regexp.MustCompile(`^<@!?(\d{15,21})>$`)
	channelMention = regexp.MustCompile(`^<#(\d{15,21})>$`)
	roleMention    = regexp.MustCompile(`^<@&(\d{15,21})>$`)
	ErrNotFound    = errors.New("commands: not found")
	ErrNotBool     = errors.New("commands: not a boolean")
)

type Args struct {
	ctx  *Context
	view *View
}

func (a *Args) next(name string) (string, error) {
	w, err := a.view.Word()
	if errors.Is(err, errEndOfInput) {
		return "", &MissingArgumentError{Name: name}
	}
	return w, err
}

func (a *Args) Empty() bool {
	a.view.SkipSpace()
	return a.view.EOF()
}

func (a *Args) Rest(name string) (string, error) {
	r := a.view.Rest()
	if r == "" {
		return "", &MissingArgumentError{Name: name}
	}
	return r, nil
}

func (a *Args) String(name string) (string, error) {
	return a.next(name)
}

func parse[T any](a *Args, name string, conv func(string) (T, error)) (T, error) {
	var zero T
	w, err := a.next(name)
	if err != nil {
		return zero, err
	}
	v, err := conv(w)
	if err != nil {
		return zero, &BadArgumentError{Name: name, Value: w, Err: err}
	}
	return v, nil
}

func (a *Args) Int(name string) (int64, error) {
	return parse(a, name, func(s string) (int64, error) { return strconv.ParseInt(s, 10, 64) })
}

func (a *Args) Float(name string) (float64, error) {
	return parse(a, name, func(s string) (float64, error) { return strconv.ParseFloat(s, 64) })
}

func (a *Args) Bool(name string) (bool, error) {
	return parse(a, name, func(s string) (bool, error) {
		switch strings.ToLower(s) {
		case "yes", "y", "true", "t", "1", "enable", "on":
			return true, nil
		case "no", "n", "false", "f", "0", "disable", "off":
			return false, nil
		}
		return false, ErrNotBool
	})
}

func idFrom(pattern *regexp.Regexp, s string) (snowflake.ID, error) {
	if m := pattern.FindStringSubmatch(s); m != nil {
		s = m[1]
	}
	return snowflake.Parse(s)
}

func (a *Args) ID(name string) (snowflake.ID, error) {
	return parse(a, name, snowflake.Parse)
}

func (a *Args) Channel(name string) (snowflake.ID, error) {
	return parse(a, name, func(s string) (snowflake.ID, error) { return idFrom(channelMention, s) })
}

func (a *Args) Role(name string) (snowflake.ID, error) {
	return parse(a, name, func(s string) (snowflake.ID, error) { return idFrom(roleMention, s) })
}

func (a *Args) User(name string) (discord.PartialUser, error) {
	return parse(a, name, func(s string) (discord.PartialUser, error) {
		st := a.ctx.Bot.client.State()
		if id, err := idFrom(userMention, s); err == nil {
			if u, ok := st.User(id); ok {
				return u, nil
			}
			return discord.PartialUser{ID: id}, nil
		}
		if u, ok := st.UserByName(s); ok {
			return u, nil
		}
		return discord.PartialUser{}, ErrNotFound
	})
}
