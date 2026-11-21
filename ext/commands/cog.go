package commands

import "context"

type Cog interface {
	CogName() string
	CogCommands() []*Command
}

type CogDescriber interface {
	CogDescription() string
}

type CogLoader interface {
	CogLoad(ctx context.Context, b *Bot) error
}

type CogUnloader interface {
	CogUnload(ctx context.Context, b *Bot)
}

type CogChecker interface {
	CogCheck(ctx *Context) (bool, error)
}

type BotChecker interface {
	BotCheck(ctx *Context) (bool, error)
}

type BotCheckOncer interface {
	BotCheckOnce(ctx *Context) (bool, error)
}

type CogErrorHandler interface {
	CogCommandError(ctx *Context, err error)
}

type CogBeforeInvoker interface {
	CogBeforeInvoke(ctx *Context) error
}

type CogAfterInvoker interface {
	CogAfterInvoke(ctx *Context) error
}

type CogListener interface {
	CogListeners(b *Bot) []func()
}

type loadedCog struct {
	cog       Cog
	commands  []*Command
	listeners []func()
	unchecks  []func()
}

func assignCog(c *Command, cog Cog) {
	c.Cog = cog
	for _, sub := range c.Commands() {
		assignCog(sub, cog)
	}
}

func (b *Bot) inject(ctx context.Context, cog Cog) (*loadedCog, error) {
	if l, ok := cog.(CogLoader); ok {
		if err := l.CogLoad(ctx, b); err != nil {
			return nil, err
		}
	}
	lc := &loadedCog{cog: cog}
	for _, c := range cog.CogCommands() {
		assignCog(c, cog)
		if c.Parent != nil {
			continue
		}
		if err := b.AddCommand(c); err != nil {
			for _, added := range lc.commands {
				b.RemoveCommand(added.Name)
			}
			return nil, err
		}
		lc.commands = append(lc.commands, c)
	}
	if chk, ok := cog.(BotChecker); ok {
		lc.unchecks = append(lc.unchecks, b.AddCheck(chk.BotCheck, false))
	}
	if chk, ok := cog.(BotCheckOncer); ok {
		lc.unchecks = append(lc.unchecks, b.AddCheck(chk.BotCheckOnce, true))
	}
	if l, ok := cog.(CogListener); ok {
		lc.listeners = l.CogListeners(b)
	}
	return lc, nil
}

func (b *Bot) eject(ctx context.Context, lc *loadedCog) {
	for _, c := range lc.commands {
		b.RemoveCommand(c.Name)
	}
	for _, remove := range lc.listeners {
		remove()
	}
	for _, remove := range lc.unchecks {
		remove()
	}
	if h := b.HelpCommand(); h != nil && h.helpBase().Cog == lc.cog {
		h.helpBase().SetCog(nil)
	}
	if u, ok := lc.cog.(CogUnloader); ok {
		u.CogUnload(ctx, b)
	}
}
