package commands

import (
	"cmp"
	"reflect"
	"regexp"
	"slices"
	"strings"
)

type VerifyChecks int

const (
	VerifyAlways VerifyChecks = iota
	VerifyNever
	VerifyInGuild
)

type DMHelp int

const (
	DMNever DMHelp = iota
	DMAlways
	DMOverThreshold
)

type HelpCommand interface {
	helpBase() *HelpBase
	PrepareHelpCommand(ctx *Context, command string) error
	SendBotHelp(mapping map[Cog][]*Command) error
	SendCogHelp(cog Cog) error
	SendGroupHelp(group *Command) error
	SendCommandHelp(command *Command) error
	SendErrorMessage(message string) error
	CommandNotFound(name string) string
	SubcommandNotFound(command *Command, name string) string
}

type HelpErrorHandler interface {
	OnHelpCommandError(ctx *Context, err error)
}

type HelpBase struct {
	ShowHidden     bool
	VerifyChecks   VerifyChecks
	CommandName    string
	CommandOptions []Option
	Cog            Cog
	Context        *Context

	impl *Command
}

var (
	mentionTransforms = map[string]string{"@everyone": "@​everyone", "@here": "@​here"}
	mentionPattern    = regexp.MustCompile(`@everyone|@here|<@!?[0-9]{17,22}>|<@&[0-9]{17,22}>`)
)

func (h *HelpBase) helpBase() *HelpBase { return h }

func (h *HelpBase) PrepareHelpCommand(*Context, string) error { return nil }

func (h *HelpBase) SendBotHelp(map[Cog][]*Command) error { return nil }

func (h *HelpBase) SendCogHelp(Cog) error { return nil }

func (h *HelpBase) SendGroupHelp(*Command) error { return nil }

func (h *HelpBase) SendCommandHelp(*Command) error { return nil }

func (h *HelpBase) SendErrorMessage(message string) error {
	_, err := h.Context.Send(message)
	return err
}

func (h *HelpBase) CommandNotFound(name string) string {
	return `No command called "` + name + `" found.`
}

func (h *HelpBase) SubcommandNotFound(command *Command, name string) string {
	if command.IsGroup() {
		return `Command "` + command.QualifiedName() + `" has no subcommand named ` + name
	}
	return `Command "` + command.QualifiedName() + `" has no subcommands.`
}

func (h *HelpBase) Command() *Command { return h.impl }

func (h *HelpBase) SetCog(cog Cog) {
	h.Cog = cog
	if h.impl != nil {
		h.impl.Cog = cog
	}
}

func (h *HelpBase) InvokedWith() string {
	name := h.impl.Name
	ctx := h.Context
	if ctx == nil || ctx.Command == nil || ctx.Command.QualifiedName() != name {
		return name
	}
	return ctx.InvokedWith
}

func (h *HelpBase) BotMapping() map[Cog][]*Command {
	bot := h.Context.Bot
	mapping := map[Cog][]*Command{}
	for _, cog := range bot.Cogs() {
		mapping[cog] = h.CogCommands(cog)
	}
	mapping[nil] = slices.DeleteFunc(bot.Commands(), func(c *Command) bool { return c.Cog != nil })
	return mapping
}

func (h *HelpBase) CogCommands(cog Cog) []*Command {
	out := slices.DeleteFunc(slices.Clone(cog.CogCommands()), func(c *Command) bool { return c.Parent != nil })
	if h.impl != nil && h.impl.Cog == cog && !slices.Contains(out, h.impl) {
		out = append(out, h.impl)
	}
	return out
}

func (h *HelpBase) CommandSignature(command *Command) string {
	var entries []string
	for parent := command.Parent; parent != nil; parent = parent.Parent {
		if sig := parent.Signature(); sig == "" || parent.InvokeWithoutCommand {
			entries = append(entries, parent.Name)
		} else {
			entries = append(entries, parent.Name+" "+sig)
		}
	}
	slices.Reverse(entries)
	alias := command.Name
	if len(command.Aliases) > 0 {
		alias = "[" + command.Name + "|" + strings.Join(command.Aliases, "|") + "]"
	}
	if parent := strings.Join(entries, " "); parent != "" {
		alias = parent + " " + alias
	}
	return h.Context.CleanPrefix() + alias + " " + command.Signature()
}

func (h *HelpBase) RemoveMentions(s string) string {
	return mentionPattern.ReplaceAllStringFunc(s, func(m string) string {
		if t, ok := mentionTransforms[m]; ok {
			return t
		}
		if strings.HasPrefix(m, "<@&") {
			return "@deleted-role"
		}
		return "@deleted-user"
	})
}

func (h *HelpBase) FilterCommands(commands []*Command, sort bool, key func(*Command) string) []*Command {
	out := slices.Clone(commands)
	if !h.ShowHidden {
		out = slices.DeleteFunc(out, func(c *Command) bool { return c.Hidden })
	}
	if h.VerifyChecks == VerifyAlways || h.VerifyChecks == VerifyInGuild && !h.Context.Private() {
		out = slices.DeleteFunc(out, func(c *Command) bool {
			ok, err := c.CanRun(h.Context)
			return err != nil || !ok
		})
	}
	if sort {
		if key == nil {
			key = func(c *Command) string { return c.Name }
		}
		slices.SortStableFunc(out, func(a, b *Command) int { return cmp.Compare(key(a), key(b)) })
	}
	return out
}

func MaxWidth(names []string) int {
	n := 0
	for _, name := range names {
		n = max(n, StringWidth(name))
	}
	return n
}

func commandNames(commands []*Command) []string {
	names := make([]string, len(commands))
	for i, c := range commands {
		names[i] = c.Name
	}
	return names
}

type helpParams struct {
	Command *string `commands:"rest"`
}

func cloneHelp(proto HelpCommand, ctx *Context) HelpCommand {
	v := reflect.ValueOf(proto)
	dup := reflect.New(v.Elem().Type())
	dup.Elem().Set(v.Elem())
	h := dup.Interface().(HelpCommand)
	h.helpBase().Context = ctx
	return h
}

func newHelpImpl(proto HelpCommand) *Command {
	base := proto.helpBase()
	name := cmp.Or(base.CommandName, "help")
	opts := append([]Option{Help("Shows this message")}, base.CommandOptions...)
	if _, ok := proto.(HelpErrorHandler); ok {
		opts = append(opts, OnError(func(ctx *Context, err error) {
			cloneHelp(proto, ctx).(HelpErrorHandler).OnHelpCommandError(ctx, err)
		}))
	}
	cmd := New(name, func(ctx *Context, p helpParams) error { return runHelp(cloneHelp(proto, ctx), ctx, p.Command) }, opts...)
	cmd.Cog = base.Cog
	base.impl = cmd
	return cmd
}

func runHelp(h HelpCommand, ctx *Context, command *string) error {
	arg := ""
	if command != nil {
		arg = *command
	}
	if err := h.PrepareHelpCommand(ctx, arg); err != nil {
		return err
	}
	base := h.helpBase()
	if command == nil {
		return h.SendBotHelp(base.BotMapping())
	}
	if cog := ctx.Bot.GetCog(arg); cog != nil {
		return h.SendCogHelp(cog)
	}
	keys := strings.Split(arg, " ")
	cmd := ctx.Bot.Get(keys[0])
	if cmd == nil {
		return h.SendErrorMessage(h.CommandNotFound(base.RemoveMentions(keys[0])))
	}
	for _, key := range keys[1:] {
		found := cmd.Get(key)
		if found == nil {
			return h.SendErrorMessage(h.SubcommandNotFound(cmd, base.RemoveMentions(key)))
		}
		cmd = found
	}
	if cmd.IsGroup() {
		return h.SendGroupHelp(cmd)
	}
	return h.SendCommandHelp(cmd)
}
