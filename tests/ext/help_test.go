package ext_test

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/prettylittlelies/cry/ext/commands"
)

type outbox struct {
	mu   sync.Mutex
	sent []string
	dms  int
}

func (o *outbox) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.URL.Path == "/users/@me/channels" {
		o.mu.Lock()
		o.dms++
		o.mu.Unlock()
		_, _ = w.Write([]byte(`{"id":"900000000000000009","type":1}`))
		return
	}
	var body struct{ Content string }
	_ = json.NewDecoder(r.Body).Decode(&body)
	o.mu.Lock()
	o.sent = append(o.sent, body.Content)
	o.mu.Unlock()
	_, _ = w.Write([]byte(`{"id":"1","channel_id":"11","content":""}`))
}

func (o *outbox) take() []string {
	o.mu.Lock()
	defer o.mu.Unlock()
	out := o.sent
	o.sent = nil
	return out
}

type musicCog struct{ cmds []*commands.Command }

func (c *musicCog) CogName() string                  { return "Music" }
func (c *musicCog) CogDescription() string           { return "Plays music." }
func (c *musicCog) CogCommands() []*commands.Command { return c.cmds }

type playParams struct {
	Song   string `description:"what to play"`
	Volume int    `default:"50"`
}

func helpBot(t *testing.T, opts ...commands.BotOption) (*harness, *outbox) {
	t.Helper()
	box := &outbox{}
	h := newBotWith(t, box, append([]commands.BotOption{commands.WithMode(commands.UserBot), commands.WithDescription("A bot.")}, opts...)...)
	h.bot.AddCommand(commands.New("ping", func(*commands.Context, struct{}) error { return nil }, commands.Help("Pong.\nMore detail.")))
	h.bot.AddCommand(commands.New("secret", func(*commands.Context, struct{}) error { return nil }, commands.Hidden()))
	h.bot.AddCommand(commands.New("tag", func(*commands.Context, struct{}) error { return nil },
		commands.Brief("Tags."), commands.Aliases("t"),
		commands.Subcommands(commands.New("create", func(*commands.Context, struct{ Name string }) error { return nil }, commands.Help("Create a tag.")))))
	cog := &musicCog{cmds: []*commands.Command{commands.New("play", func(*commands.Context, playParams) error { return nil }, commands.Help("Play a song."))}}
	if err := h.bot.AddCog(t.Context(), cog, false); err != nil {
		t.Fatal(err)
	}
	return h, box
}

func waitSent(t *testing.T, box *outbox) string {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		if sent := box.take(); len(sent) > 0 {
			return strings.Join(sent, "\n<page>\n")
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatal("nothing sent")
	return ""
}

func TestDefaultHelp(t *testing.T) {
	h, box := helpBot(t)
	for _, tc := range []struct{ input, want string }{
		{"!help", "```\nA bot.\n\nMusic:\n  play Play a song.\n\u200bNo Category:\n  help Shows this message\n  ping Pong.\n  tag  Tags.\n\nType !help command for more info on a command.\nYou can also type !help category for more info on a category.\n```"},
		{"!help play", "```\n!play\n\nPlay a song.\n\nArguments:\n  song   what to play\n  volume No description given (default: 50)\n```"},
		{"!help tag", "```\n![tag|t]\n\nCommands:\n  create Create a tag.\n\nType !help command for more info on a command.\nYou can also type !help category for more info on a category.\n```"},
		{"!help Music", "```\nPlays music.\n\nCommands:\n  play Play a song.\n\nType !help command for more info on a command.\nYou can also type !help category for more info on a category.\n```"},
		{"!help nope", `No command called "nope" found.`},
		{"!help ping x", `Command "ping" has no subcommands.`},
		{"!help tag x", `Command "tag" has no subcommand named x`},
		{"!help @everyone", "No command called \"@\u200beveryone\" found."},
	} {
		h.run(other, 10, tc.input)
		if got := waitSent(t, box); got != tc.want {
			t.Errorf("%s:\n%q\nwant\n%q", tc.input, got, tc.want)
		}
	}
}

func TestMinimalHelp(t *testing.T) {
	h, box := helpBot(t, commands.WithHelpCommand(commands.NewMinimalHelp()))
	h.run(other, 10, "!help")
	want := "A bot.\n\nUse `!help [command]` for more info on a command.\nYou can also use `!help [category]` for more info on a category.\n\n" +
		"__**Music**__\nplay\n__**\u200bNo Category**__\nhelp\u2002ping\u2002tag"
	if got := waitSent(t, box); got != want {
		t.Errorf("bot help:\n%q\nwant\n%q", got, want)
	}
	h.run(other, 10, "!help tag create")
	if got := waitSent(t, box); got != "!tag create <name>\n\nCreate a tag.\n" {
		t.Errorf("command help %q", got)
	}
	h.run(other, 10, "!help tag")
	if got := waitSent(t, box); got != "!tag \n**Aliases:** t\n\nUse `!help [command]` for more info on a command.\nYou can also use `!help [category]` for more info on a category.\n\n**Commands**\n!tag create \u2013 Create a tag." {
		t.Errorf("group help %q", got)
	}
}

func TestHelpConfiguration(t *testing.T) {
	dm := commands.NewDefaultHelp()
	dm.DMHelp = commands.DMAlways
	dm.CommandName = "commands"
	h, box := helpBot(t, commands.WithHelpCommand(dm))
	if h.bot.Get("help") != nil || h.bot.Get("commands") == nil {
		t.Fatal("custom name not registered")
	}
	h.run(other, 10, "!commands ping")
	waitSent(t, box)
	if box.dms == 0 {
		t.Fatal("DMAlways must open a DM")
	}
	if err := h.bot.SetHelpCommand(nil); err != nil || h.bot.Get("commands") != nil {
		t.Fatalf("removal: %v", err)
	}

	p := commands.NewPaginator("```", "```")
	p.MaxSize = 20
	for range 3 {
		if err := p.AddLine("0123456789"); err != nil {
			t.Fatal(err)
		}
	}
	if pages := p.Pages(); len(pages) != 3 || pages[0] != "```\n0123456789\n```" {
		t.Fatalf("pages %q", pages)
	}
	if err := p.AddLine(strings.Repeat("x", 13)); err == nil {
		t.Fatal("oversized line must fail")
	}
}
