package ext_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sync"
	"testing"

	"github.com/prettylittlelies/cry"
	"github.com/prettylittlelies/cry/build"
	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/ext/commands"
	"github.com/prettylittlelies/cry/snowflake"
	"github.com/prettylittlelies/cry/tests/testkit"
)

func TestViewTokenizes(t *testing.T) {
	v := commands.NewView(`one "two words" “three” 'it\'s' rest of it`)
	for _, want := range []string{"one", "two words", "three", "it's"} {
		got, err := v.Word()
		if err != nil || got != want {
			t.Fatalf("got %q err=%v, want %q", got, err, want)
		}
	}
	if r := v.Rest(); r != "rest of it" {
		t.Fatalf("rest %q", r)
	}
	var unclosed *commands.UnclosedQuoteError
	if _, err := commands.NewView(`"open`).Word(); !errors.As(err, &unclosed) {
		t.Fatalf("unclosed: %v", err)
	}
	var unexpected *commands.UnexpectedQuoteError
	if _, err := commands.NewView(`"a"b`).Word(); !errors.As(err, &unexpected) {
		t.Fatalf("unexpected: %v", err)
	}
}

type sent struct {
	mu     sync.Mutex
	bodies []map[string]any
}

func (s *sent) last() map[string]any {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.bodies) == 0 {
		return nil
	}
	return s.bodies[len(s.bodies)-1]
}

const me, other = snowflake.ID(1), snowflake.ID(2)

func bot(t *testing.T, opts ...commands.Option) (*commands.Bot, *sent, *[]error) {
	t.Helper()
	got := &sent{}
	srv := testkit.NewH2Server(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		raw, _ := io.ReadAll(r.Body)
		json.Unmarshal(raw, &body)
		got.mu.Lock()
		got.bodies = append(got.bodies, body)
		got.mu.Unlock()
		io.WriteString(w, `{"id":"500","channel_id":"11","author":{"id":"1","username":"me","discriminator":"0","avatar":null},"content":"ok","timestamp":"2026-11-21T04:00:00+00:00","edited_timestamp":null,"tts":false,"mention_everyone":false,"mentions":[],"mention_roles":[],"attachments":[],"embeds":[],"pinned":false,"type":0}`)
	}))
	v := func(s string) build.Version { out, _ := build.ParseVersion(s); return out }
	snap := &build.Snapshot{Releases: map[build.Channel]build.Release{build.Stable: {Channel: build.Stable, Number: 1, MacOS: build.MacOS{Host: v("0.0.421"), Electron: v("44.10.0"), Chrome: v("152.0.7977.169")}}}}
	c, err := cry.New(context.Background(), "token", cry.WithSnapshot(snap), cry.WithAPIBase(srv.URL), cry.WithRootCAs(srv.Pool))
	if err != nil {
		t.Fatal(err)
	}
	ready := new(discord.ReadyEvent)
	ready.User.ID = me
	c.State().Apply(ready)
	var errs []error
	var mu sync.Mutex
	opts = append([]commands.Option{commands.WithPrefix("!"), commands.WithErrorHandler(func(_ *commands.Context, err error) {
		mu.Lock()
		errs = append(errs, err)
		mu.Unlock()
	})}, opts...)
	b := commands.New(c, opts...)
	t.Cleanup(b.Close)
	return b, got, &errs
}

func msg(author snowflake.ID, guild snowflake.ID, content string) *discord.MessageCreateEvent {
	m := &discord.MessageCreateEvent{ID: 99, ChannelID: 11, GuildID: guild, Content: content}
	m.Author.ID = author
	return m
}

func TestReplyAndSelfOnly(t *testing.T) {
	b, got, _ := bot(t)
	b.Add(&commands.Command{Name: "ping", Run: func(ctx *commands.Context, _ *commands.Args) error {
		_, err := ctx.Reply("pong")
		return err
	}})
	if err := b.Process(context.Background(), msg(other, 10, "!ping")); err != nil || got.last() != nil {
		t.Fatalf("other user's message must be ignored: err=%v body=%v", err, got.last())
	}
	if err := b.Process(context.Background(), msg(me, 10, "!ping")); err != nil {
		t.Fatal(err)
	}
	body := got.last()
	ref, _ := body["message_reference"].(map[string]any)
	if body["content"] != "pong" || body["nonce"] == nil || ref["message_id"] != "99" {
		t.Fatalf("reply body %v", body)
	}
}

func TestGroupsArgsAndErrors(t *testing.T) {
	b, _, errs := bot(t, commands.WithCaseInsensitive(true))
	var name, text string
	var n int64
	b.Add(&commands.Command{
		Name:    "tag",
		Aliases: []string{"t"},
		Subcommands: []*commands.Command{{
			Name: "add",
			Run: func(_ *commands.Context, a *commands.Args) (err error) {
				if name, err = a.String("name"); err != nil {
					return err
				}
				text, err = a.Rest("text")
				return err
			},
		}},
	}, &commands.Command{
		Name:   "repeat",
		Checks: []commands.Check{commands.GuildOnly()},
		Run: func(_ *commands.Context, a *commands.Args) (err error) {
			n, err = a.Int("times")
			return err
		},
	})
	if err := b.Process(context.Background(), msg(me, 10, `!T ADD greet "hello there" friend`)); err != nil {
		t.Fatal(err)
	}
	if name != "greet" || text != `"hello there" friend` {
		t.Fatalf("name=%q text=%q", name, text)
	}
	b.Process(context.Background(), msg(me, 10, "!repeat 3"))
	if n != 3 {
		t.Fatalf("n=%d", n)
	}
	b.Process(context.Background(), msg(me, 0, "!repeat 3"))
	b.Process(context.Background(), msg(me, 10, "!repeat"))
	b.Process(context.Background(), msg(me, 10, "!repeat many"))
	b.Process(context.Background(), msg(me, 10, "!nope"))
	if len(*errs) != 4 {
		t.Fatalf("errors %v", *errs)
	}
	var missing *commands.MissingArgumentError
	var bad *commands.BadArgumentError
	if !errors.Is((*errs)[0], commands.ErrGuildOnly) || !errors.As((*errs)[1], &missing) || !errors.As((*errs)[2], &bad) || !errors.Is((*errs)[3], commands.ErrCommandNotFound) {
		t.Fatalf("errors %v", *errs)
	}
}

type greeter struct {
	loaded, unloaded bool
}

func (g *greeter) Name() string { return "greeter" }

func (g *greeter) Commands() []*commands.Command {
	return []*commands.Command{{Name: "hi", Run: func(*commands.Context, *commands.Args) error { return nil }}}
}

func (g *greeter) Load(*commands.Bot) error { g.loaded = true; return nil }

func (g *greeter) Unload(*commands.Bot) { g.unloaded = true }

func TestCogs(t *testing.T) {
	b, _, errs := bot(t)
	g := &greeter{}
	if err := b.AddCog(g); err != nil || !g.loaded || len(b.Commands()) != 1 {
		t.Fatalf("add cog: err=%v loaded=%v commands=%d", err, g.loaded, len(b.Commands()))
	}
	b.RemoveCog("greeter")
	if !g.unloaded || len(b.Commands()) != 0 {
		t.Fatalf("remove cog: unloaded=%v commands=%d", g.unloaded, len(b.Commands()))
	}
	b.Process(context.Background(), msg(me, 10, "!hi"))
	if len(*errs) != 1 || !errors.Is((*errs)[0], commands.ErrCommandNotFound) {
		t.Fatalf("errors %v", *errs)
	}
}
