package commands

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/prettylittlelies/cry"
	"github.com/prettylittlelies/cry/discord"
)

type PrefixFunc func(msg *discord.MessageCreateEvent) []string

type Option func(*Bot)

func WithPrefix(prefixes ...string) Option {
	return func(b *Bot) { b.prefixes = func(*discord.MessageCreateEvent) []string { return prefixes } }
}

func WithPrefixFunc(fn PrefixFunc) Option {
	return func(b *Bot) { b.prefixes = fn }
}

func WithSelfOnly(enabled bool) Option {
	return func(b *Bot) { b.selfOnly = enabled }
}

func WithCaseInsensitive(enabled bool) Option {
	return func(b *Bot) { b.fold = enabled }
}

func WithErrorHandler(fn func(ctx *Context, err error)) Option {
	return func(b *Bot) { b.onError = fn }
}

type Bot struct {
	client   *cry.Client
	prefixes PrefixFunc
	selfOnly bool
	fold     bool
	onError  func(ctx *Context, err error)
	remove   func()

	mu       sync.RWMutex
	commands []*Command
	cogs     map[string]Cog
}

func New(c *cry.Client, opts ...Option) *Bot {
	b := &Bot{
		client:   c,
		prefixes: func(*discord.MessageCreateEvent) []string { return []string{"!"} },
		selfOnly: true,
		onError:  func(*Context, error) {},
		cogs:     map[string]Cog{},
	}
	for _, o := range opts {
		o(b)
	}
	b.remove = cry.On(c, func(ctx context.Context, m *discord.MessageCreateEvent) {
		b.Process(ctx, m)
	})
	return b
}

func (b *Bot) Close() {
	b.remove()
}

func (b *Bot) Add(cmds ...*Command) {
	for _, c := range cmds {
		c.link()
	}
	b.mu.Lock()
	b.commands = append(b.commands, cmds...)
	b.mu.Unlock()
}

func (b *Bot) Remove(name string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	for i, c := range b.commands {
		if c.Name == name {
			b.commands = append(b.commands[:i:i], b.commands[i+1:]...)
			return
		}
	}
}

func (b *Bot) Commands() []*Command {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return append([]*Command(nil), b.commands...)
}

func (b *Bot) AddCog(cog Cog) error {
	if l, ok := cog.(Loader); ok {
		if err := l.Load(b); err != nil {
			return err
		}
	}
	b.mu.Lock()
	b.cogs[cog.Name()] = cog
	b.mu.Unlock()
	b.Add(cog.Commands()...)
	return nil
}

func (b *Bot) RemoveCog(name string) {
	b.mu.Lock()
	cog, ok := b.cogs[name]
	delete(b.cogs, name)
	b.mu.Unlock()
	if !ok {
		return
	}
	for _, c := range cog.Commands() {
		b.Remove(c.Name)
	}
	if u, ok := cog.(Unloader); ok {
		u.Unload(b)
	}
}

func (b *Bot) find(name string) *Command {
	b.mu.RLock()
	defer b.mu.RUnlock()
	for _, c := range b.commands {
		for _, n := range c.names() {
			if n == name || b.fold && strings.EqualFold(n, name) {
				return c
			}
		}
	}
	return nil
}

func (b *Bot) prefix(m *discord.MessageCreateEvent) (string, bool) {
	for _, p := range b.prefixes(m) {
		if p != "" && strings.HasPrefix(m.Content, p) {
			return p, true
		}
	}
	return "", false
}

func (b *Bot) Process(ctx context.Context, m *discord.MessageCreateEvent) error {
	if b.selfOnly && m.Author.ID != b.client.State().Me().ID {
		return nil
	}
	p, ok := b.prefix(m)
	if !ok {
		return nil
	}
	view := NewView(m.Content[len(p):])
	invoked, err := view.Word()
	if err != nil {
		return nil
	}
	c := &Context{Context: ctx, Bot: b, Message: m, Prefix: p, InvokedWith: invoked}
	cmd := b.find(invoked)
	if cmd == nil {
		b.onError(c, fmt.Errorf("%w: %s", ErrCommandNotFound, invoked))
		return ErrCommandNotFound
	}
	for len(cmd.Subcommands) > 0 {
		save := view.pos
		word, err := view.Word()
		sub := cmd.sub(word, b.fold)
		if err != nil || sub == nil {
			view.pos = save
			break
		}
		cmd = sub
	}
	c.Command = cmd
	if err := b.invoke(c, view); err != nil {
		b.onError(c, err)
		return err
	}
	return nil
}

func (b *Bot) invoke(c *Context, view *View) error {
	for at := c.Command; at != nil; at = at.parent {
		for _, check := range at.Checks {
			if err := check(c); err != nil {
				return err
			}
		}
	}
	if c.Command.Run == nil {
		return &CommandError{Command: c.Command.QualifiedName(), Err: errors.New("no handler")}
	}
	if err := c.Command.Run(c, &Args{ctx: c, view: view}); err != nil {
		return &CommandError{Command: c.Command.QualifiedName(), Err: err}
	}
	return nil
}
