package ext_test

import (
	"context"
	"errors"
	"net/http"
	"slices"
	"sync"
	"testing"
	"time"

	"github.com/prettylittlelies/cry"
	"github.com/prettylittlelies/cry/build"
	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/ext/commands"
	"github.com/prettylittlelies/cry/snowflake"
	"github.com/prettylittlelies/cry/tests/testkit"
)

const me, other = snowflake.ID(1), snowflake.ID(2)

func TestStringViewMatchesDiscordPy(t *testing.T) {
	v := commands.NewStringView(`it's "two words" “three” a\"b rest of it`)
	for _, want := range []string{"it's", "two words", "three", `a"b`} {
		v.SkipWS()
		got, _, err := v.GetQuotedWord()
		if err != nil || got != want {
			t.Fatalf("got %q err=%v, want %q", got, err, want)
		}
	}
	v.SkipWS()
	if r := v.ReadRest(); r != "rest of it" {
		t.Fatalf("rest %q", r)
	}
	cases := map[string]error{
		`a"b`:   &commands.UnexpectedQuoteError{},
		`"a"b`:  &commands.InvalidEndOfQuotedStringError{},
		`"open`: &commands.ExpectedClosingQuoteError{},
	}
	for in := range cases {
		_, _, err := commands.NewStringView(in).GetQuotedWord()
		if !errors.Is(err, commands.ErrArgumentParsing) || !errors.Is(err, commands.ErrUserInput) {
			t.Errorf("%q: %v", in, err)
		}
	}
	if _, _, err := commands.NewStringView(`"open`).GetQuotedWord(); err.Error() != "Expected closing \".\"" && err.Error() != `Expected closing ".` {
		t.Errorf("message %q", err)
	}
}

type harness struct {
	bot    *commands.Bot
	client *cry.Client
	errs   chan error
}

func newBot(t *testing.T, opts ...commands.BotOption) *harness {
	t.Helper()
	srv := testkit.NewH2Server(t, http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusNoContent) }))
	v := func(s string) build.Version { out, _ := build.ParseVersion(s); return out }
	snap := &build.Snapshot{Releases: map[build.Channel]build.Release{build.Stable: {Channel: build.Stable, Number: 1, MacOS: build.MacOS{Host: v("0.0.421"), Electron: v("44.10.0"), Chrome: v("152.0.7977.169")}}}}
	c, err := cry.New(context.Background(), "token", cry.WithSnapshot(snap), cry.WithAPIBase(srv.URL), cry.WithRootCAs(srv.Pool))
	if err != nil {
		t.Fatal(err)
	}
	ready := new(discord.ReadyEvent)
	ready.User.ID = me
	c.State().Apply(ready)
	h := &harness{bot: commands.NewBot(c, opts...), client: c, errs: make(chan error, 16)}
	cry.On(c, func(_ context.Context, e *commands.CommandErrorEvent) { h.errs <- e.Err })
	t.Cleanup(func() { h.bot.Close(context.Background()) })
	return h
}

func (h *harness) run(author snowflake.ID, guild snowflake.ID, content string) {
	m := &discord.MessageCreateEvent{ID: snowflake.FromTime(time.Now()), ChannelID: 11, GuildID: guild, Content: content}
	m.Author.ID = author
	h.bot.ProcessCommands(context.Background(), m)
}

func (h *harness) err(t *testing.T) error {
	t.Helper()
	select {
	case err := <-h.errs:
		return err
	case <-time.After(2 * time.Second):
		t.Fatal("no command_error dispatched")
	}
	return nil
}

func TestModes(t *testing.T) {
	for _, tc := range []struct {
		mode     commands.Mode
		own, oth bool
	}{
		{commands.Default, false, true},
		{commands.SelfBot, true, false},
		{commands.UserBot, true, true},
	} {
		h := newBot(t, commands.WithMode(tc.mode), commands.WithPrefix(commands.Prefixes("?", "!")))
		var calls []snowflake.ID
		h.bot.AddCommand(commands.New("ping", func(ctx *commands.Context, _ struct{}) error {
			calls = append(calls, ctx.AuthorID())
			return nil
		}))
		h.run(me, 10, "!ping")
		h.run(other, 10, "?ping")
		bot := &discord.MessageCreateEvent{Content: "!ping"}
		bot.Author.ID, bot.Author.Bot = 3, true
		h.bot.ProcessCommands(context.Background(), bot)
		if slices.Contains(calls, me) != tc.own || slices.Contains(calls, other) != tc.oth || slices.Contains(calls, 3) {
			t.Errorf("mode %d: calls %v", tc.mode, calls)
		}
	}
}

type banParams struct {
	Target string
	Days   *int   `default:"1"`
	Reason string `commands:"rest" default:"no reason"`
}

func TestArgumentParsing(t *testing.T) {
	h := newBot(t, commands.WithMode(commands.UserBot))
	var got banParams
	h.bot.AddCommand(commands.New("ban", func(_ *commands.Context, p banParams) error { got = p; return nil }))
	h.run(other, 10, `!ban someone 7 spamming   a lot`)
	if got.Target != "someone" || got.Days == nil || *got.Days != 7 || got.Reason != "spamming   a lot" {
		t.Fatalf("got %+v", got)
	}
	h.run(other, 10, `!ban "two words" not-a-number here`)
	if got.Target != "two words" || got.Days != nil || got.Reason != "not-a-number here" {
		t.Fatalf("optional must undo on failure: %+v days=%v", got, got.Days)
	}
	h.run(other, 10, `!ban`)
	err := h.err(t)
	var missing *commands.MissingRequiredArgument
	if !errors.As(err, &missing) || err.Error() != "target is a required argument that is missing." || !errors.Is(err, commands.ErrUserInput) {
		t.Fatalf("missing: %v", err)
	}

	type sumParams struct {
		Numbers []int  `commands:"greedy"`
		Label   string `default:"total"`
	}
	var sum sumParams
	h.bot.AddCommand(commands.New("sum", func(_ *commands.Context, p sumParams) error { sum = p; return nil }))
	h.run(other, 10, `!sum 1 2 3 done`)
	if !slices.Equal(sum.Numbers, []int{1, 2, 3}) || sum.Label != "done" {
		t.Fatalf("greedy %+v", sum)
	}

	type strictParams struct{ N int }
	h.bot.AddCommand(commands.New("strict", func(*commands.Context, strictParams) error { return nil }, commands.StrictArguments()))
	h.run(other, 10, `!strict 1 2`)
	var tooMany *commands.TooManyArguments
	if err := h.err(t); !errors.As(err, &tooMany) {
		t.Fatalf("strict: %v", err)
	}
	h.run(other, 10, `!strict x`)
	if err := h.err(t); !errors.Is(err, commands.ErrBadArgument) || err.Error() != `Converting to "int" failed for parameter "n".` {
		t.Fatalf("bad int: %v", err)
	}
}

func TestGroups(t *testing.T) {
	h := newBot(t, commands.WithMode(commands.UserBot))
	var order []string
	var parents []string
	tag := commands.New("tag", func(*commands.Context, struct{}) error { order = append(order, "tag"); return nil },
		commands.Subcommands(commands.New("add", func(ctx *commands.Context, p struct{ Name string }) error {
			order, parents = append(order, "add:"+p.Name), ctx.InvokedParents
			return nil
		})))
	h.bot.AddCommand(tag)
	h.run(other, 10, `!tag add hello`)
	if !slices.Equal(order, []string{"tag", "add:hello"}) || !slices.Equal(parents, []string{"tag"}) {
		t.Fatalf("order=%v parents=%v", order, parents)
	}
	order = nil
	quiet := commands.New("q", func(*commands.Context, struct{}) error { order = append(order, "q"); return nil }, commands.InvokeWithoutCommand(),
		commands.Subcommands(commands.New("x", func(*commands.Context, struct{}) error { order = append(order, "x"); return nil })))
	h.bot.AddCommand(quiet)
	h.run(other, 10, `!q x`)
	h.run(other, 10, `!q`)
	if !slices.Equal(order, []string{"x", "q"}) {
		t.Fatalf("invoke_without_command order %v", order)
	}
	if h.bot.GetCommand("tag add") == nil || h.bot.GetCommand("tag add").QualifiedName() != "tag add" {
		t.Fatal("qualified lookup")
	}
}

func TestChecksCooldownsConcurrency(t *testing.T) {
	h := newBot(t, commands.WithMode(commands.UserBot), commands.WithOwnerIDs(me))
	h.bot.AddCommand(commands.New("guild", func(*commands.Context, struct{}) error { return nil }, commands.Checks(commands.GuildOnly())))
	h.run(other, 0, "!guild")
	if err := h.err(t); err.Error() != "This command cannot be used in private messages." || !errors.Is(err, commands.ErrCheckFailure) {
		t.Fatalf("guild only: %v", err)
	}
	h.bot.AddCommand(commands.New("owner", func(*commands.Context, struct{}) error { return nil }, commands.Checks(commands.IsOwner())))
	h.run(other, 10, "!owner")
	var notOwner *commands.NotOwner
	if err := h.err(t); !errors.As(err, &notOwner) {
		t.Fatalf("not owner: %v", err)
	}
	h.bot.AddCommand(commands.New("slow", func(*commands.Context, struct{}) error { return nil }, commands.WithCooldown(1, time.Minute, commands.BucketUser)))
	h.run(other, 10, "!slow")
	h.run(other, 10, "!slow")
	var cooldown *commands.CommandOnCooldown
	if err := h.err(t); !errors.As(err, &cooldown) || cooldown.RetryAfter <= 0 {
		t.Fatalf("cooldown: %v", err)
	}
	release := make(chan struct{})
	started := make(chan struct{})
	h.bot.AddCommand(commands.New("busy", func(*commands.Context, struct{}) error { close(started); <-release; return nil }, commands.WithMaxConcurrency(1, commands.BucketDefault, false)))
	go h.run(other, 10, "!busy")
	<-started
	h.run(other, 10, "!busy")
	var reached *commands.MaxConcurrencyReached
	if err := h.err(t); !errors.As(err, &reached) || err.Error() != "Too many people are using this command. It can only be used 1 time globally concurrently." {
		t.Fatalf("concurrency: %v", err)
	}
	close(release)
	remove := h.bot.AddCheck(func(*commands.Context) (bool, error) { return false, nil }, false)
	h.bot.AddCommand(commands.New("any", func(*commands.Context, struct{}) error { return nil }))
	h.run(other, 10, "!any")
	if err := h.err(t); err.Error() != "The global check functions for command any failed." {
		t.Fatalf("global check: %v", err)
	}
	remove()
}

type hookCog struct {
	mu    sync.Mutex
	trace []string
	cmd   *commands.Command
}

func (c *hookCog) log(s string) {
	c.mu.Lock()
	c.trace = append(c.trace, s)
	c.mu.Unlock()
}

func (c *hookCog) CogName() string { return "hooks" }

func (c *hookCog) CogCommands() []*commands.Command { return []*commands.Command{c.cmd} }

func (c *hookCog) CogBeforeInvoke(*commands.Context) error { c.log("cog-before"); return nil }

func (c *hookCog) CogAfterInvoke(*commands.Context) error { c.log("cog-after"); return nil }

func (c *hookCog) CogCommandError(_ *commands.Context, err error) { c.log("cog-error") }

func TestHookAndErrorOrder(t *testing.T) {
	h := newBot(t, commands.WithMode(commands.UserBot))
	cog := &hookCog{}
	cog.cmd = commands.New("boom", func(*commands.Context, struct{}) error { cog.log("run"); return errors.New("kaboom") },
		commands.BeforeInvoke(func(*commands.Context) error { cog.log("cmd-before"); return nil }),
		commands.AfterInvoke(func(*commands.Context) error { cog.log("cmd-after"); return nil }),
		commands.OnError(func(*commands.Context, error) { cog.log("cmd-error") }))
	h.bot.BeforeInvoke(func(*commands.Context) error { cog.log("bot-before"); return nil })
	h.bot.AfterInvoke(func(*commands.Context) error { cog.log("bot-after"); return nil })
	if err := h.bot.AddCog(context.Background(), cog, false); err != nil {
		t.Fatal(err)
	}
	h.run(other, 10, "!boom")
	err := h.err(t)
	var invoke *commands.CommandInvokeError
	if !errors.As(err, &invoke) || invoke.Original.Error() != "kaboom" {
		t.Fatalf("invoke error: %v", err)
	}
	want := []string{"cmd-before", "cog-before", "bot-before", "run", "cmd-after", "cog-after", "bot-after", "cmd-error", "cog-error"}
	if !slices.Equal(cog.trace, want) {
		t.Fatalf("trace %v\nwant  %v", cog.trace, want)
	}
	if err := h.bot.AddCog(context.Background(), cog, false); err == nil {
		t.Fatal("duplicate cog must fail without override")
	}
	if err := h.bot.AddCog(context.Background(), cog, true); err != nil {
		t.Fatalf("override: %v", err)
	}
	h.bot.RemoveCog(context.Background(), "hooks")
	if h.bot.GetCommand("boom") != nil {
		t.Fatal("cog commands must be removed")
	}
	h.run(other, 10, "!boom")
	var notFound *commands.CommandNotFound
	if err := h.err(t); !errors.As(err, &notFound) || err.Error() != `Command "boom" is not found` {
		t.Fatalf("not found: %v", err)
	}
}

func TestExtensions(t *testing.T) {
	h := newBot(t)
	torn := false
	commands.RegisterExtension("cry.test.ext", commands.Extension{
		Setup: func(_ context.Context, b *commands.Bot) error {
			return b.AddCommand(commands.New("extcmd", func(*commands.Context, struct{}) error { return nil }))
		},
		Teardown: func(context.Context, *commands.Bot) { torn = true },
	})
	commands.RegisterExtension("cry.test.broken", commands.Extension{})
	ctx := context.Background()
	var notFound *commands.ExtensionNotFound
	if err := h.bot.LoadExtension(ctx, "missing"); !errors.As(err, &notFound) {
		t.Fatalf("not found: %v", err)
	}
	var noEntry *commands.NoEntryPointError
	if err := h.bot.LoadExtension(ctx, "cry.test.broken"); !errors.As(err, &noEntry) {
		t.Fatalf("no entry: %v", err)
	}
	if err := h.bot.LoadExtension(ctx, "cry.test.ext"); err != nil || h.bot.GetCommand("extcmd") == nil {
		t.Fatalf("load: %v", err)
	}
	var already *commands.ExtensionAlreadyLoaded
	if err := h.bot.LoadExtension(ctx, "cry.test.ext"); !errors.As(err, &already) {
		t.Fatalf("already loaded: %v", err)
	}
	if err := h.bot.ReloadExtension(ctx, "cry.test.ext"); err != nil || !torn || h.bot.GetCommand("extcmd") == nil {
		t.Fatalf("reload: %v torn=%v", err, torn)
	}
	if err := h.bot.UnloadExtension(ctx, "cry.test.ext"); err != nil || h.bot.GetCommand("extcmd") != nil {
		t.Fatalf("unload: %v", err)
	}
	var notLoaded *commands.ExtensionNotLoaded
	if err := h.bot.UnloadExtension(ctx, "cry.test.ext"); !errors.As(err, &notLoaded) {
		t.Fatalf("not loaded: %v", err)
	}
}
