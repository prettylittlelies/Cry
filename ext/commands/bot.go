package commands

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"slices"
	"sync"

	"github.com/prettylittlelies/cry"
	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
)

var ErrOwnersNotSet = errors.New("commands: owners aren't set")

type Mode int

const (
	Default Mode = iota
	SelfBot
	UserBot
)

type Prefix func(b *Bot, m *discord.MessageCreateEvent) []string

func Prefixes(p ...string) Prefix {
	return func(*Bot, *discord.MessageCreateEvent) []string { return p }
}

func WhenMentioned() Prefix {
	return func(b *Bot, _ *discord.MessageCreateEvent) []string {
		id := b.client.State().Me().ID.String()
		return []string{"<@" + id + "> ", "<@!" + id + "> "}
	}
}

func WhenMentionedOr(p ...string) Prefix {
	return func(b *Bot, m *discord.MessageCreateEvent) []string {
		return append(WhenMentioned()(b, m), p...)
	}
}

type BotOption func(*Bot)

func WithPrefix(p Prefix) BotOption       { return func(b *Bot) { b.prefix = p } }
func WithMode(m Mode) BotOption           { return func(b *Bot) { b.mode = m } }
func WithCaseInsensitive() BotOption      { return func(b *Bot) { b.fold = true } }
func WithStripAfterPrefix() BotOption     { return func(b *Bot) { b.stripAfterPrefix = true } }
func WithDescription(d string) BotOption  { return func(b *Bot) { b.Description = d } }
func WithLogger(l *slog.Logger) BotOption { return func(b *Bot) { b.logger = l } }
func WithOwnerIDs(ids ...snowflake.ID) BotOption {
	return func(b *Bot) { b.ownerIDs = ids }
}

type checkEntry struct {
	id    uint64
	check Check
}

type extensionRecord struct {
	ext      Extension
	cogs     []string
	commands []string
}

type Bot struct {
	Description string

	client           *cry.Client
	prefix           Prefix
	mode             Mode
	stripAfterPrefix bool
	ownerIDs         []snowflake.ID
	logger           *slog.Logger
	remove           []func()

	mu sync.RWMutex
	group
	checks     []checkEntry
	checkOnce  []checkEntry
	nextCheck  uint64
	before     Hook
	after      Hook
	cogs       map[string]*loadedCog
	extensions map[string]*extensionRecord
	loading    *extensionRecord
}

func NewBot(c *cry.Client, opts ...BotOption) *Bot {
	b := &Bot{
		client:     c,
		prefix:     Prefixes("!"),
		logger:     slog.Default(),
		cogs:       map[string]*loadedCog{},
		extensions: map[string]*extensionRecord{},
	}
	for _, o := range opts {
		o(b)
	}
	b.group = newGroup(b.fold)
	b.remove = append(b.remove,
		cry.On(c, func(ctx context.Context, m *discord.MessageCreateEvent) { b.ProcessCommands(ctx, m) }),
		cry.On(c, func(_ context.Context, e *CommandErrorEvent) { b.defaultError(e) }),
	)
	return b
}

func (b *Bot) Client() *cry.Client {
	return b.client
}

func (b *Bot) BeforeInvoke(h Hook) { b.before = h }

func (b *Bot) AfterInvoke(h Hook) { b.after = h }

func (b *Bot) AddCommand(c *Command) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if c.Parent != nil {
		return fmt.Errorf("commands: %s is a subcommand", c.QualifiedName())
	}
	if err := b.group.add(c); err != nil {
		return err
	}
	if b.loading != nil {
		b.loading.commands = append(b.loading.commands, c.Name)
	}
	return nil
}

func (b *Bot) RemoveCommand(name string) *Command {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.group.remove(name)
}

func (b *Bot) GetCommand(name string) *Command {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.group.GetCommand(name)
}

func (b *Bot) AddCheck(check Check, once bool) (remove func()) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.nextCheck++
	e := checkEntry{id: b.nextCheck, check: check}
	list := &b.checks
	if once {
		list = &b.checkOnce
	}
	*list = append(*list, e)
	return func() {
		b.mu.Lock()
		defer b.mu.Unlock()
		*list = slices.DeleteFunc(*list, func(x checkEntry) bool { return x.id == e.id })
	}
}

func (b *Bot) canRun(ctx *Context, once bool) (bool, error) {
	b.mu.RLock()
	list := b.checks
	if once {
		list = b.checkOnce
	}
	list = slices.Clone(list)
	b.mu.RUnlock()
	for _, e := range list {
		if ok, err := e.check(ctx); err != nil || !ok {
			return false, err
		}
	}
	return true, nil
}

func (b *Bot) IsOwner(user snowflake.ID) (bool, error) {
	if len(b.ownerIDs) == 0 {
		if b.mode == SelfBot {
			return user == b.client.State().Me().ID, nil
		}
		return false, ErrOwnersNotSet
	}
	return slices.Contains(b.ownerIDs, user), nil
}

func (b *Bot) AddCog(ctx context.Context, cog Cog, override bool) error {
	b.mu.RLock()
	_, exists := b.cogs[cog.CogName()]
	b.mu.RUnlock()
	if exists {
		if !override {
			return fmt.Errorf("commands: cog named %q already loaded", cog.CogName())
		}
		b.RemoveCog(ctx, cog.CogName())
	}
	lc, err := b.inject(ctx, cog)
	if err != nil {
		return err
	}
	b.mu.Lock()
	b.cogs[cog.CogName()] = lc
	if b.loading != nil {
		b.loading.cogs = append(b.loading.cogs, cog.CogName())
	}
	b.mu.Unlock()
	return nil
}

func (b *Bot) RemoveCog(ctx context.Context, name string) Cog {
	b.mu.Lock()
	lc, ok := b.cogs[name]
	delete(b.cogs, name)
	b.mu.Unlock()
	if !ok {
		return nil
	}
	b.eject(ctx, lc)
	return lc.cog
}

func (b *Bot) GetCog(name string) Cog {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if lc, ok := b.cogs[name]; ok {
		return lc.cog
	}
	return nil
}

func (b *Bot) Cogs() map[string]Cog {
	b.mu.RLock()
	defer b.mu.RUnlock()
	out := make(map[string]Cog, len(b.cogs))
	for k, v := range b.cogs {
		out[k] = v.cog
	}
	return out
}

func (b *Bot) GetPrefix(m *discord.MessageCreateEvent) []string {
	return b.prefix(b, m)
}

func (b *Bot) skip(author snowflake.ID) bool {
	me := b.client.State().Me().ID
	switch b.mode {
	case SelfBot:
		return author != me
	case UserBot:
		return false
	}
	return author == me
}

func (b *Bot) GetContext(ctx context.Context, m *discord.MessageCreateEvent) *Context {
	view := NewStringView(m.Content)
	c := &Context{Context: ctx, Message: m, Bot: b, View: view}
	if b.skip(m.Author.ID) {
		return c
	}
	matched := false
	for _, p := range b.GetPrefix(m) {
		if view.SkipString(p) {
			c.Prefix, c.HasPrefix, matched = p, true, true
			break
		}
	}
	if !matched {
		return c
	}
	if b.stripAfterPrefix {
		view.SkipWS()
	}
	c.InvokedWith = view.GetWord()
	b.mu.RLock()
	c.Command = b.group.Get(c.InvokedWith)
	b.mu.RUnlock()
	return c
}

func (b *Bot) Invoke(c *Context) {
	if c.Command == nil {
		if c.InvokedWith != "" {
			b.client.Dispatch(c, &CommandErrorEvent{Context: c, Err: errCommandNotFound(c.InvokedWith)})
		}
		return
	}
	b.client.Dispatch(c, &CommandEvent{Context: c})
	cmd := c.Command
	err := func() error {
		ok, err := b.canRun(c, true)
		if err != nil {
			return err
		}
		if !ok {
			return NewCheckFailure("The global check once functions failed.")
		}
		return cmd.invoke(c)
	}()
	if err != nil {
		if !errors.Is(err, ErrCommand) {
			err = errInvoke(err)
		}
		cmd.dispatchError(c, err)
		return
	}
	b.client.Dispatch(c, &CommandCompletionEvent{Context: c})
}

func (b *Bot) ProcessCommands(ctx context.Context, m *discord.MessageCreateEvent) {
	if m.Author.Bot {
		return
	}
	b.Invoke(b.GetContext(ctx, m))
}

func (b *Bot) defaultError(e *CommandErrorEvent) {
	if cmd := e.Context.Command; cmd != nil {
		if cmd.OnError != nil {
			return
		}
		if _, ok := cmd.Cog.(CogErrorHandler); ok && cmd.Cog != nil {
			return
		}
	}
	name := "<none>"
	if e.Context.Command != nil {
		name = e.Context.Command.QualifiedName()
	}
	b.logger.Error("ignoring exception in command", "command", name, "err", e.Err)
}

func (b *Bot) Close(ctx context.Context) {
	b.mu.RLock()
	exts := make([]string, 0, len(b.extensions))
	for name := range b.extensions {
		exts = append(exts, name)
	}
	cogs := make([]string, 0, len(b.cogs))
	for name := range b.cogs {
		cogs = append(cogs, name)
	}
	b.mu.RUnlock()
	for _, name := range exts {
		b.UnloadExtension(ctx, name)
	}
	for _, name := range cogs {
		b.RemoveCog(ctx, name)
	}
	for _, remove := range b.remove {
		remove()
	}
}
