package client_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/prettylittlelies/cry"
	"github.com/prettylittlelies/cry/build"
	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/gateway"
	"github.com/prettylittlelies/cry/tests/testkit"
)

func snapshot(t *testing.T) *build.Snapshot {
	t.Helper()
	v := func(s string) build.Version { out, _ := build.ParseVersion(s); return out }
	return &build.Snapshot{Releases: map[build.Channel]build.Release{
		build.Stable: {Channel: build.Stable, Number: 678818, MacOS: build.MacOS{Host: v("0.0.421"), Electron: v("44.10.0"), Chrome: v("152.0.7977.169")}},
	}}
}

func client(t *testing.T) (*cry.Client, *testkit.WSServer) {
	t.Helper()
	srv := testkit.NewWSServer(t)
	c, err := cry.New(context.Background(), "token", cry.WithSnapshot(snapshot(t)), cry.WithGatewayEndpoint(srv.URL))
	if err != nil {
		t.Fatal(err)
	}
	c.Gateway().Dialer().RootCAs = srv.Pool
	return c, srv
}

func run(t *testing.T, c *cry.Client) {
	t.Helper()
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	go c.Run(ctx)
}

func readyPayload() map[string]any {
	return map[string]any{
		"session_id": "abc",
		"user":       map[string]any{"id": "1", "username": "me", "discriminator": "0", "avatar": nil},
		"users":      []any{},
		"guilds": []any{map[string]any{
			"id": "10", "joined_at": "2021-03-01T12:00:00+00:00", "large": false, "member_count": 1,
			"properties": map[string]any{"id": "10", "name": "cry", "owner_id": "1"},
			"channels":   []any{map[string]any{"id": "11", "type": 0, "name": "general"}},
			"roles":      []any{},
		}},
		"private_channels": []any{},
		"relationships":    []any{},
	}
}

func message(id, content string) map[string]any {
	return map[string]any{
		"id": id, "channel_id": "11", "guild_id": "10", "content": content,
		"author":    map[string]any{"id": "2", "username": "pal", "discriminator": "0", "avatar": nil},
		"timestamp": "2026-11-21T04:00:00+00:00", "edited_timestamp": nil, "tts": false, "mention_everyone": false,
		"mentions": []any{}, "mention_roles": []any{}, "attachments": []any{}, "embeds": []any{}, "pinned": false, "type": 0,
	}
}

func TestEventsFlowThroughStateAndHandlers(t *testing.T) {
	c, srv := client(t)
	readies := make(chan *discord.ReadyEvent, 1)
	cry.On(c, func(_ context.Context, e *discord.ReadyEvent) { readies <- e })
	messages := make(chan *discord.MessageCreateEvent, 4)
	cry.On(c, func(_ context.Context, e *discord.MessageCreateEvent) { messages <- e })
	run(t, c)
	p := testkit.AcceptGatewayPeer(t, srv)
	p.Hello(10000)
	p.Expect(gateway.OpIdentify)
	p.Dispatch("READY", 1, readyPayload())
	select {
	case e := <-readies:
		if e.User.ID != 1 {
			t.Fatalf("ready user %d", e.User.ID)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("no READY handler call")
	}
	if g, ok := c.State().Guild(10); !ok || g.Properties.Name != "cry" {
		t.Fatalf("state guild %+v", g)
	}
	waited := make(chan *discord.MessageCreateEvent, 1)
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		m, err := cry.WaitFor(ctx, c, func(e *discord.MessageCreateEvent) bool { return e.Content == "second" })
		if err == nil {
			waited <- m
		}
	}()
	time.Sleep(50 * time.Millisecond)
	p.Dispatch("MESSAGE_CREATE", 2, message("100", "first"))
	p.Dispatch("MESSAGE_CREATE", 3, message("101", "second"))
	select {
	case m := <-waited:
		if m.ID != 101 || m.GuildID != 10 {
			t.Fatalf("waited %+v", m)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("WaitFor never matched")
	}
	got := map[string]bool{}
	for range 2 {
		select {
		case m := <-messages:
			got[m.Content] = true
		case <-time.After(5 * time.Second):
			t.Fatal("missing message handler call")
		}
	}
	if !got["first"] || !got["second"] {
		t.Fatalf("handlers saw %v", got)
	}
	if m, ok := c.State().Message(100); !ok || m.Content != "first" {
		t.Fatalf("cached %+v", m)
	}
}

func TestUnknownEventsAreRaw(t *testing.T) {
	c, srv := client(t)
	raws := make(chan *cry.RawEvent, 1)
	cry.On(c, func(_ context.Context, e *cry.RawEvent) { raws <- e })
	run(t, c)
	p := testkit.AcceptGatewayPeer(t, srv)
	p.Hello(10000)
	p.Expect(gateway.OpIdentify)
	p.Dispatch("SOMETHING_NEW", 1, map[string]any{"x": 1})
	select {
	case e := <-raws:
		if e.Type != "SOMETHING_NEW" || string(e.Data) != `{"x":1}` {
			t.Fatalf("raw %+v %s", e, e.Data)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("no raw event")
	}
}

type custom struct{ n int }

func TestRemoveAndWaitForCancel(t *testing.T) {
	c, _ := client(t)
	calls := make(chan int, 4)
	remove := cry.On(c, func(_ context.Context, e *custom) { calls <- e.n })
	c.Dispatch(context.Background(), &custom{n: 1})
	if n := <-calls; n != 1 {
		t.Fatalf("got %d", n)
	}
	remove()
	remove()
	c.Dispatch(context.Background(), &custom{n: 2})
	select {
	case n := <-calls:
		t.Fatalf("removed listener still called with %d", n)
	case <-time.After(100 * time.Millisecond):
	}
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	if _, err := cry.WaitFor[custom](ctx, c, nil); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("WaitFor err %v", err)
	}
}
