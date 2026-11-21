package ext_test

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/ext/commands"
	"github.com/prettylittlelies/cry/ext/commands/converting"
	"github.com/prettylittlelies/cry/snowflake"
)

const (
	threadID  = snowflake.ID(600000000000000006)
	cachedID  = snowflake.ID(700000000000000007)
	remoteID  = snowflake.ID(700000000000000008)
	hiddenID  = snowflake.ID(700000000000000009)
	missingID = snowflake.ID(700000000000000010)
)

func fakeAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch {
	case strings.HasSuffix(r.URL.Path, "/messages/"+remoteID.String()):
		_, _ = w.Write([]byte(`{"id":"` + remoteID.String() + `","channel_id":"` + channelID.String() + `","content":"remote"}`))
	case strings.HasSuffix(r.URL.Path, "/messages/"+hiddenID.String()):
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"code":50001,"message":"Missing Access"}`))
	case r.URL.Path == "/invites/abc":
		if r.URL.Query().Get("with_counts") != "true" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, _ = w.Write([]byte(`{"code":"abc","type":0}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"code":10008,"message":"Unknown"}`))
	}
}

func resolverBot(t *testing.T) *harness {
	t.Helper()
	h := newBotWith(t, http.HandlerFunc(fakeAPI), commands.WithMode(commands.UserBot))
	seedGuild(h)
	h.client.State().Apply(&discord.ThreadCreateEvent{ID: threadID, GuildID: guildID, Type: discord.ChannelTypePublicThread, Name: ptr("ideas")})
	cached := &discord.MessageCreateEvent{ID: cachedID, ChannelID: channelID, GuildID: guildID, Content: "cached"}
	h.client.State().Apply(cached)
	return h
}

func TestMessageConverters(t *testing.T) {
	h := resolverBot(t)
	var got discord.Message
	h.bot.AddCommand(commands.New("quote", func(_ *commands.Context, p struct{ M discord.Message }) error { got = p.M; return nil }))
	link := "https://discord.com/channels/" + guildID.String() + "/" + channelID.String() + "/"
	for input, want := range map[string]string{
		cachedID.String(): "cached",
		channelID.String() + "-" + remoteID.String(): "remote",
		link + remoteID.String():                     "remote",
	} {
		got = discord.Message{}
		h.run(other, guildID, "!quote "+input)
		if got.Content != want {
			t.Errorf("%s: %+v", input, got)
		}
	}
	var nf *commands.NotFoundError
	var unreadable *commands.ChannelNotReadable
	for input, check := range map[string]func(error) bool{
		channelID.String() + "-" + missingID.String(): func(err error) bool { return errors.As(err, &nf) && nf.Kind == "Message" },
		channelID.String() + "-" + hiddenID.String():  func(err error) bool { return errors.As(err, &unreadable) },
		"999999999999999999-" + missingID.String():    func(err error) bool { return errors.As(err, &nf) && nf.Kind == "Channel" },
		"not-a-message": func(err error) bool { return errors.As(err, &nf) && nf.Kind == "Message" },
	} {
		h.run(other, guildID, "!quote "+input)
		if err := h.err(t); !check(err) {
			t.Errorf("%s: %v", input, err)
		}
	}

	var partial converting.PartialMessage
	h.bot.AddCommand(commands.New("peek", func(_ *commands.Context, p struct{ M converting.PartialMessage }) error { partial = p.M; return nil }))
	h.run(other, guildID, "!peek "+link+missingID.String())
	if partial.ID != missingID || partial.Channel.ID != channelID {
		t.Fatalf("partial %+v", partial)
	}
}

func TestThreadInviteConverters(t *testing.T) {
	h := resolverBot(t)
	var thread converting.Thread
	h.bot.AddCommand(commands.New("thread", func(_ *commands.Context, p struct{ T converting.Thread }) error { thread = p.T; return nil }))
	for _, input := range []string{"ideas", threadID.String(), "<#" + threadID.String() + ">"} {
		thread = converting.Thread{}
		h.run(other, guildID, "!thread "+input)
		if thread.ID != threadID {
			t.Errorf("%s: %+v", input, thread)
		}
	}
	var nf *commands.NotFoundError
	h.run(other, guildID, "!thread general")
	if err := h.err(t); !errors.As(err, &nf) || nf.Kind != "Thread" {
		t.Fatalf("text channel accepted as thread: %v", err)
	}
	var channel discord.Channel
	h.bot.AddCommand(commands.New("chan", func(_ *commands.Context, p struct{ C discord.Channel }) error { channel = p.C; return nil }))
	h.run(other, guildID, "!chan https://discord.com/channels/"+guildID.String()+"/"+channelID.String())
	if channel.ID != channelID {
		t.Fatalf("channel link %+v", channel)
	}
	h.run(other, guildID, "!chan ideas")
	if err := h.err(t); !errors.As(err, &nf) {
		t.Fatalf("thread accepted as channel: %v", err)
	}

	var invite discord.Invite
	h.bot.AddCommand(commands.New("invite", func(_ *commands.Context, p struct{ I discord.Invite }) error { invite = p.I; return nil }))
	for _, input := range []string{"abc", "discord.gg/abc", "https://discord.com/invite/abc"} {
		invite = discord.Invite{}
		h.run(other, guildID, "!invite "+input)
		if invite.Code != "abc" {
			t.Errorf("%s: %+v", input, invite)
		}
	}
	h.run(other, guildID, "!invite gone")
	var bad *commands.BadInviteArgument
	if err := h.err(t); !errors.As(err, &bad) || err.Error() != `Invite "gone" is invalid or expired.` {
		t.Fatalf("invite: %v", err)
	}
}

func TestCleanContent(t *testing.T) {
	h := resolverBot(t)
	var plain, fixed string
	h.bot.AddCommand(commands.New("say", func(_ *commands.Context, p struct {
		Text commands.CleanContent `commands:"rest"`
	}) error {
		plain = string(p.Text)
		return nil
	}))
	h.bot.AddCommand(commands.New("fix", func(_ *commands.Context, p struct {
		Text commands.CleanContent `commands:"rest" clean:"fix_channel_mentions,no_nicknames,escape_markdown"`
	}) error {
		fixed = string(p.Text)
		return nil
	}))
	input := " <@" + memberID.String() + "> <@&" + roleID.String() + "> <#" + channelID.String() + "> <@123456789012345678> @everyone *hi*"
	h.run(other, guildID, "!say"+input)
	if want := "@nicky @mods <#" + channelID.String() + "> @deleted-user @​everyone *hi*"; plain != want {
		t.Errorf("plain %q want %q", plain, want)
	}
	h.run(other, guildID, "!fix"+input)
	if want := "@alice @mods #general @deleted-user @​everyone \\*hi\\*"; fixed != want {
		t.Errorf("fixed %q want %q", fixed, want)
	}
}

func TestRangeAndAttachments(t *testing.T) {
	h := resolverBot(t)
	type rangeParams struct {
		N    int     `max:"10" min:"1"`
		F    float64 `min:"0.5"`
		Name string  `max:"3"`
	}
	var got rangeParams
	h.bot.AddCommand(commands.New("r", func(_ *commands.Context, p rangeParams) error { got = p; return nil }))
	h.run(other, guildID, "!r 5 0.5 abc")
	if got != (rangeParams{5, 0.5, "abc"}) {
		t.Fatalf("in range %+v", got)
	}
	var rangeErr *commands.RangeError
	for input, want := range map[string]string{
		"!r 11 1 a":   "value must be between 1 and 10 but received 11",
		"!r 1 0.1 a":  "value must be not less than 0.5 but received 0.1",
		"!r 1 1 abcd": "value must be no more than 3 characters but received 4 characters",
	} {
		h.run(other, guildID, input)
		if err := h.err(t); !errors.As(err, &rangeErr) || err.Error() != want || !errors.Is(err, commands.ErrBadArgument) {
			t.Errorf("%s: %v", input, err)
		}
	}
	defer func() {
		if recover() == nil {
			t.Error("min on a non-numeric type must panic")
		}
	}()
	commands.New("bad", func(*commands.Context, struct {
		B bool `min:"1"`
	}) error {
		return nil
	})
}

func TestAttachments(t *testing.T) {
	h := resolverBot(t)
	type uploadParams struct {
		File  discord.Attachment
		Extra *discord.Attachment
		Label string `default:"none"`
	}
	var got uploadParams
	cmd := commands.New("upload", func(_ *commands.Context, p uploadParams) error { got = p; return nil })
	h.bot.AddCommand(cmd)
	if sig := cmd.Signature(); sig != "<file (upload a file)> [extra (upload a file)] [label=none]" {
		t.Errorf("signature %q", sig)
	}
	send := func(content string, files ...discord.Attachment) {
		m := &discord.MessageCreateEvent{ID: 1, ChannelID: channelID, GuildID: guildID, Content: content, Attachments: files}
		m.Author.ID = other
		h.bot.ProcessCommands(t.Context(), m)
	}
	send("!upload tag", discord.Attachment{Filename: "a.png"}, discord.Attachment{Filename: "b.png"})
	if got.File.Filename != "a.png" || got.Extra == nil || got.Extra.Filename != "b.png" || got.Label != "tag" {
		t.Errorf("two files %+v", got)
	}
	send("!upload", discord.Attachment{Filename: "a.png"})
	if got.Extra != nil || got.Label != "none" {
		t.Errorf("one file %+v", got)
	}
	send("!upload")
	var missing *commands.MissingRequiredAttachment
	if err := h.err(t); !errors.As(err, &missing) {
		t.Errorf("no file: %v", err)
	}

	var all []discord.Attachment
	h.bot.AddCommand(commands.New("bulk", func(_ *commands.Context, p struct {
		Files []discord.Attachment `commands:"greedy"`
	}) error {
		all = p.Files
		return nil
	}))
	send("!bulk", discord.Attachment{Filename: "a"}, discord.Attachment{Filename: "b"}, discord.Attachment{Filename: "c"})
	if len(all) != 3 {
		t.Errorf("greedy %+v", all)
	}
}
