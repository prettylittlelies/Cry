package commands

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"time"
)

var (
	errInvalidContext = errors.New("commands: this context is not valid")
	errEndVarArgs     = errors.New("commands: end of variadic arguments")
)

type Hook func(ctx *Context) error

type ErrorHandler func(ctx *Context, err error)

type Check func(ctx *Context) (bool, error)

type Command struct {
	Name                 string
	Aliases              []string
	Help                 string
	Brief                string
	Usage                string
	Description          string
	Enabled              bool
	Hidden               bool
	RestIsRaw            bool
	IgnoreExtra          bool
	RequireVarPositional bool
	CooldownAfterParsing bool
	InvokeWithoutCommand bool
	Extras               map[string]any
	Checks               []Check
	Params               []*Parameter
	Parent               *Command
	Cog                  Cog
	OnError              ErrorHandler

	before      Hook
	after       Hook
	buckets     *CooldownMapping
	concurrency *MaxConcurrency
	paramsType  reflect.Type
	call        func(ctx *Context, params reflect.Value) error
	group
}

type Option func(*Command)

func Aliases(a ...string) Option   { return func(c *Command) { c.Aliases = a } }
func Help(s string) Option         { return func(c *Command) { c.Help = s } }
func Brief(s string) Option        { return func(c *Command) { c.Brief = s } }
func Usage(s string) Option        { return func(c *Command) { c.Usage = s } }
func Description(s string) Option  { return func(c *Command) { c.Description = s } }
func Hidden() Option               { return func(c *Command) { c.Hidden = true } }
func Disabled() Option             { return func(c *Command) { c.Enabled = false } }
func RestIsRaw() Option            { return func(c *Command) { c.RestIsRaw = true } }
func StrictArguments() Option      { return func(c *Command) { c.IgnoreExtra = false } }
func RequireVarPositional() Option { return func(c *Command) { c.RequireVarPositional = true } }
func CooldownAfterParsing() Option { return func(c *Command) { c.CooldownAfterParsing = true } }
func InvokeWithoutCommand() Option { return func(c *Command) { c.InvokeWithoutCommand = true } }
func CaseInsensitive() Option      { return func(c *Command) { c.setFold(true) } }
func Checks(checks ...Check) Option {
	return func(c *Command) { c.Checks = append(c.Checks, checks...) }
}
func BeforeInvoke(h Hook) Option     { return func(c *Command) { c.before = h } }
func AfterInvoke(h Hook) Option      { return func(c *Command) { c.after = h } }
func OnError(h ErrorHandler) Option  { return func(c *Command) { c.OnError = h } }
func Extras(m map[string]any) Option { return func(c *Command) { c.Extras = m } }
func Subcommands(subs ...*Command) Option {
	return func(c *Command) {
		for _, s := range subs {
			if err := c.AddCommand(s); err != nil {
				panic(err)
			}
		}
	}
}

func WithCooldown(rate int, per time.Duration, t BucketType) Option {
	return func(c *Command) { c.buckets = NewCooldownMapping(NewCooldown(rate, per), t) }
}

func WithDynamicCooldown(factory func(BucketSource) *Cooldown, t BucketType) Option {
	return func(c *Command) { c.buckets = NewDynamicCooldownMapping(factory, t) }
}

func WithMaxConcurrency(number int, per BucketType, wait bool) Option {
	return func(c *Command) { c.concurrency = NewMaxConcurrency(number, per, wait) }
}

func New[P any](name string, fn func(ctx *Context, params P) error, opts ...Option) *Command {
	t := reflect.TypeFor[P]()
	params, err := signature(t)
	if err != nil {
		panic(err)
	}
	c := &Command{
		Name:        name,
		Enabled:     true,
		IgnoreExtra: true,
		Extras:      map[string]any{},
		Params:      params,
		paramsType:  t,
		buckets:     NewCooldownMapping(nil, BucketDefault),
		call: func(ctx *Context, v reflect.Value) error {
			return fn(ctx, v.Interface().(P))
		},
		group: newGroup(false),
	}
	for _, o := range opts {
		o(c)
	}
	for _, sub := range c.all {
		sub.Parent = c
	}
	return c
}

func (c *Command) QualifiedName() string {
	if c.Parent == nil {
		return c.Name
	}
	return c.Parent.QualifiedName() + " " + c.Name
}

func (c *Command) Root() *Command {
	for c.Parent != nil {
		c = c.Parent
	}
	return c
}

func (c *Command) Parents() []*Command {
	var out []*Command
	for p := c.Parent; p != nil; p = p.Parent {
		out = append(out, p)
	}
	return out
}

func (c *Command) IsGroup() bool {
	return len(c.all) > 0
}

func (c *Command) AddCommand(sub *Command) error {
	if err := c.group.add(sub); err != nil {
		return err
	}
	sub.Parent = c
	return nil
}

func (c *Command) CanRun(ctx *Context) (bool, error) {
	if !c.Enabled {
		return false, errDisabled(c.Name)
	}
	original := ctx.Command
	ctx.Command = c
	defer func() { ctx.Command = original }()
	if ok, err := ctx.Bot.canRun(ctx, false); err != nil || !ok {
		if err != nil {
			return false, err
		}
		return false, NewCheckFailure("The global check functions for command " + c.QualifiedName() + " failed.")
	}
	if checker, ok := c.Cog.(CogChecker); ok && c.Cog != nil {
		if ok, err := checker.CogCheck(ctx); err != nil || !ok {
			return false, err
		}
	}
	for _, check := range c.Checks {
		if ok, err := check(ctx); err != nil || !ok {
			return false, err
		}
	}
	return true, nil
}

func messageTime(ctx *Context) time.Time {
	if ctx.Message.EditedTimestamp != nil {
		return *ctx.Message.EditedTimestamp
	}
	return ctx.Message.ID.Time()
}

func (c *Command) prepareCooldowns(ctx *Context) error {
	if !c.buckets.Valid() {
		return nil
	}
	bucket, retry := c.buckets.UpdateRateLimit(ctx, messageTime(ctx), 1)
	if bucket != nil && retry > 0 {
		return &CommandOnCooldown{newKind(commandErr, "You are on cooldown. Try again in %.2fs", retry.Seconds()), bucket, retry, c.buckets.Type}
	}
	return nil
}

func (c *Command) IsOnCooldown(ctx *Context) bool {
	if !c.buckets.Valid() {
		return false
	}
	b := c.buckets.Bucket(ctx, messageTime(ctx))
	return b != nil && b.Tokens(messageTime(ctx)) == 0
}

func (c *Command) ResetCooldown(ctx *Context) {
	if c.buckets.Valid() {
		if b := c.buckets.Bucket(ctx, messageTime(ctx)); b != nil {
			b.Reset()
		}
	}
}

func (c *Command) CooldownRetryAfter(ctx *Context) time.Duration {
	if !c.buckets.Valid() {
		return 0
	}
	if b := c.buckets.Bucket(ctx, messageTime(ctx)); b != nil {
		return b.RetryAfter(messageTime(ctx))
	}
	return 0
}

func (c *Command) transform(ctx *Context, p *Parameter) (reflect.Value, error) {
	view := ctx.View
	view.SkipWS()
	if p.Greedy {
		if p.Kind == VarPositional {
			return c.greedyVarPos(ctx, p)
		}
		return c.greedyPos(ctx, p)
	}
	if view.EOF() {
		if p.Kind == VarPositional {
			return reflect.Value{}, errEndVarArgs
		}
		if p.Required() {
			if p.Optional() {
				return reflect.Zero(p.Type), nil
			}
			return reflect.Value{}, errMissingArgument(p)
		}
		return c.defaultValue(ctx, p)
	}
	previous := view.Index
	var argument string
	if p.Kind == KeywordOnly && !c.RestIsRaw {
		argument = strings.TrimSpace(view.ReadRest())
	} else {
		word, _, err := view.GetQuotedWord()
		if err != nil {
			if p.Optional() {
				view.Index = previous
				return reflect.Zero(p.Type), nil
			}
			return reflect.Value{}, err
		}
		argument = word
	}
	ctx.CurrentArgument = argument
	view.Previous = previous
	return runConverters(ctx, p, argument)
}

func (c *Command) greedyPos(ctx *Context, p *Parameter) (reflect.Value, error) {
	view := ctx.View
	result := reflect.MakeSlice(p.Type, 0, 0)
	for !view.EOF() {
		previous := view.Index
		view.SkipWS()
		word, _, err := view.GetQuotedWord()
		if err == nil {
			ctx.CurrentArgument = word
			var v reflect.Value
			if v, err = runConverters(ctx, p, word); err == nil {
				result = reflect.Append(result, v)
				continue
			}
		}
		view.Index = previous
		break
	}
	if result.Len() == 0 && !p.Required() {
		return c.defaultValue(ctx, p)
	}
	return result, nil
}

func (c *Command) greedyVarPos(ctx *Context, p *Parameter) (reflect.Value, error) {
	view := ctx.View
	previous := view.Index
	word, _, err := view.GetQuotedWord()
	if err == nil {
		ctx.CurrentArgument = word
		var v reflect.Value
		if v, err = runConverters(ctx, p, word); err == nil {
			return v, nil
		}
	}
	view.Index = previous
	return reflect.Value{}, errEndVarArgs
}

func (c *Command) defaultValue(ctx *Context, p *Parameter) (reflect.Value, error) {
	if p.Greedy || p.Kind == VarPositional {
		return reflect.Zero(p.Type), nil
	}
	return runConverters(ctx, p, p.Default)
}

func (c *Command) parseArguments(ctx *Context) (reflect.Value, error) {
	values := reflect.New(c.paramsType).Elem()
	view := ctx.View
	for _, p := range c.Params {
		ctx.CurrentParameter = p
		field := values.FieldByIndex(p.field)
		switch p.Kind {
		case Positional:
			v, err := c.transform(ctx, p)
			if err != nil {
				return values, err
			}
			field.Set(v)
		case KeywordOnly:
			if c.RestIsRaw {
				ctx.CurrentArgument = view.ReadRest()
				v, err := runConverters(ctx, p, ctx.CurrentArgument)
				if err != nil {
					return values, err
				}
				field.Set(v)
			} else {
				v, err := c.transform(ctx, p)
				if err != nil {
					return values, err
				}
				field.Set(v)
			}
		case VarPositional:
			if view.EOF() && c.RequireVarPositional {
				return values, errMissingArgument(p)
			}
			for !view.EOF() {
				v, err := c.transform(ctx, p)
				if errors.Is(err, errEndVarArgs) {
					break
				}
				if err != nil {
					return values, err
				}
				field.Set(reflect.Append(field, v))
			}
		}
		if p.Kind == KeywordOnly {
			break
		}
	}
	if !c.IgnoreExtra && !view.EOF() {
		return values, errTooManyArguments(c.QualifiedName())
	}
	return values, nil
}

func (c *Command) callBeforeHooks(ctx *Context) error {
	if c.before != nil {
		if err := c.before(ctx); err != nil {
			return err
		}
	}
	if h, ok := c.Cog.(CogBeforeInvoker); ok && c.Cog != nil {
		if err := h.CogBeforeInvoke(ctx); err != nil {
			return err
		}
	}
	if ctx.Bot.before != nil {
		return ctx.Bot.before(ctx)
	}
	return nil
}

func (c *Command) callAfterHooks(ctx *Context) error {
	if c.after != nil {
		if err := c.after(ctx); err != nil {
			return err
		}
	}
	if h, ok := c.Cog.(CogAfterInvoker); ok && c.Cog != nil {
		if err := h.CogAfterInvoke(ctx); err != nil {
			return err
		}
	}
	if ctx.Bot.after != nil {
		return ctx.Bot.after(ctx)
	}
	return nil
}

func (c *Command) prepare(ctx *Context) (reflect.Value, error) {
	ctx.Command = c
	ok, err := c.CanRun(ctx)
	if err != nil {
		return reflect.Value{}, err
	}
	if !ok {
		return reflect.Value{}, NewCheckFailure("The check functions for command " + c.QualifiedName() + " failed.")
	}
	if c.concurrency != nil {
		if err := c.concurrency.Acquire(ctx, ctx); err != nil {
			return reflect.Value{}, err
		}
	}
	params, err := c.prepareArguments(ctx)
	if err == nil {
		err = c.callBeforeHooks(ctx)
	}
	if err != nil && c.concurrency != nil {
		c.concurrency.Release(ctx)
	}
	return params, err
}

func (c *Command) prepareArguments(ctx *Context) (reflect.Value, error) {
	if c.CooldownAfterParsing {
		params, err := c.parseArguments(ctx)
		if err != nil {
			return params, err
		}
		return params, c.prepareCooldowns(ctx)
	}
	if err := c.prepareCooldowns(ctx); err != nil {
		return reflect.Value{}, err
	}
	return c.parseArguments(ctx)
}

func (c *Command) hooked(ctx *Context, params reflect.Value) (err error) {
	defer func() {
		if c.concurrency != nil {
			c.concurrency.Release(ctx)
		}
		if hookErr := c.callAfterHooks(ctx); err == nil {
			err = hookErr
		}
	}()
	return wrap(ctx, c.call(ctx, params))
}

func wrap(ctx *Context, err error) error {
	switch {
	case err == nil:
		return nil
	case errors.Is(err, context.Canceled):
		ctx.CommandFailed = true
		return nil
	case errors.Is(err, ErrCommand):
		ctx.CommandFailed = true
		return err
	}
	ctx.CommandFailed = true
	return errInvoke(err)
}

func (c *Command) invoke(ctx *Context) error {
	if !c.IsGroup() {
		params, err := c.prepare(ctx)
		if err != nil {
			return err
		}
		ctx.InvokedSubcommand, ctx.SubcommandPassed = nil, ""
		return c.hooked(ctx, params)
	}
	ctx.InvokedSubcommand, ctx.SubcommandPassed = nil, ""
	early := !c.InvokeWithoutCommand
	var params reflect.Value
	if early {
		var err error
		if params, err = c.prepare(ctx); err != nil {
			return err
		}
	}
	view := ctx.View
	previous := view.Index
	view.SkipWS()
	trigger := view.GetWord()
	if trigger != "" {
		ctx.SubcommandPassed = trigger
		ctx.InvokedSubcommand = c.Get(trigger)
	}
	if early {
		if err := c.hooked(ctx, params); err != nil {
			return err
		}
	}
	ctx.InvokedParents = append(ctx.InvokedParents, ctx.InvokedWith)
	if trigger != "" && ctx.InvokedSubcommand != nil {
		ctx.InvokedWith = trigger
		return ctx.InvokedSubcommand.invoke(ctx)
	}
	if !early {
		view.Index, view.Previous = previous, previous
		params, err := c.prepare(ctx)
		if err != nil {
			return err
		}
		return c.hooked(ctx, params)
	}
	return nil
}

func (c *Command) reinvoke(ctx *Context, callHooks bool) error {
	ctx.Command = c
	run := func() error {
		params, err := c.parseArguments(ctx)
		if err != nil {
			return err
		}
		if callHooks {
			if err := c.callBeforeHooks(ctx); err != nil {
				return err
			}
			defer c.callAfterHooks(ctx)
		}
		if err := c.call(ctx, params); err != nil {
			ctx.CommandFailed = true
			return err
		}
		return nil
	}
	if !c.IsGroup() {
		ctx.InvokedSubcommand = nil
		return run()
	}
	ctx.InvokedSubcommand = nil
	early := !c.InvokeWithoutCommand
	view := ctx.View
	var params reflect.Value
	if early {
		var err error
		if params, err = c.parseArguments(ctx); err != nil {
			return err
		}
		if callHooks {
			if err := c.callBeforeHooks(ctx); err != nil {
				return err
			}
		}
	}
	previous := view.Index
	view.SkipWS()
	trigger := view.GetWord()
	if trigger != "" {
		ctx.SubcommandPassed = trigger
		ctx.InvokedSubcommand = c.Get(trigger)
	}
	if early {
		err := c.call(ctx, params)
		if callHooks {
			c.callAfterHooks(ctx)
		}
		if err != nil {
			ctx.CommandFailed = true
			return err
		}
	}
	ctx.InvokedParents = append(ctx.InvokedParents, ctx.InvokedWith)
	if trigger != "" && ctx.InvokedSubcommand != nil {
		ctx.InvokedWith = trigger
		return ctx.InvokedSubcommand.reinvoke(ctx, callHooks)
	}
	if !early {
		view.Index, view.Previous = previous, previous
		return run()
	}
	return nil
}

func (c *Command) dispatchError(ctx *Context, err error) {
	ctx.CommandFailed = true
	if c.OnError != nil {
		c.OnError(ctx, err)
	}
	if h, ok := c.Cog.(CogErrorHandler); ok && c.Cog != nil {
		h.CogCommandError(ctx, err)
	}
	ctx.Bot.client.Dispatch(ctx, &CommandErrorEvent{Context: ctx, Err: err})
}
