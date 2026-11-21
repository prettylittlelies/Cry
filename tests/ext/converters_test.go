package ext_test

import (
	"errors"
	"testing"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/ext/commands"
	"github.com/prettylittlelies/cry/ext/commands/converting"
	"github.com/prettylittlelies/cry/snowflake"
)

const (
	guildID   = snowflake.ID(100000000000000001)
	memberID  = snowflake.ID(200000000000000002)
	roleID    = snowflake.ID(300000000000000003)
	channelID = snowflake.ID(400000000000000004)
	emojiID   = snowflake.ID(500000000000000005)
)

func ptr[T any](v T) *T { return &v }

func seedGuild(h *harness) {
	g := discord.GatewayGuild{ID: guildID}
	g.Properties.Name = "lab"
	g.Roles = []discord.Role{{ID: roleID, Name: "mods"}}
	g.Emojis = []discord.Emoji{{ID: ptr(emojiID), Name: "wave"}}
	g.Channels = []discord.Channel{{ID: channelID, Name: ptr("general")}}
	m := discord.GuildMember{Nick: ptr("nicky")}
	m.User.ID, m.User.Username, m.User.Discriminator, m.User.GlobalName = memberID, "alice", "0", ptr("Alice")
	g.Members = []discord.GuildMember{m}
	ready := &discord.ReadyEvent{Guilds: []discord.GatewayGuild{g}, Users: []discord.PartialUser{m.User}}
	ready.User.ID = me
	h.client.State().Apply(ready)
}

type objectParams struct {
	Member  discord.GuildMember
	User    discord.PartialUser
	Role    discord.Role
	Channel discord.Channel
	Guild   discord.GatewayGuild
	Emoji   discord.Emoji
	Partial converting.PartialEmoji
	Colour  converting.Colour
	Object  snowflake.ID
}

func TestDiscordConverters(t *testing.T) {
	h := newBot(t, commands.WithMode(commands.UserBot))
	seedGuild(h)
	var got objectParams
	h.bot.AddCommand(commands.New("probe", func(_ *commands.Context, p objectParams) error { got = p; return nil }))
	for _, input := range []string{
		`!probe <@!200000000000000002> alice <@&300000000000000003> <#400000000000000004> lab <:wave:500000000000000005> <a:x:500000000000000005> #ff8000 <#400000000000000004>`,
		`!probe nicky @Alice mods general 100000000000000001 wave <a:x:500000000000000005> "rgb(255, 128, 0)" 400000000000000004`,
	} {
		got = objectParams{}
		h.run(other, guildID, input)
		switch {
		case got.Member.User.ID != memberID, got.User.ID != memberID, got.Role.ID != roleID,
			got.Channel.ID != channelID, got.Guild.ID != guildID, got.Emoji.ID == nil || *got.Emoji.ID != emojiID,
			got.Partial != converting.PartialEmoji{ID: emojiID, Name: "x", Animated: true},
			got.Colour != converting.RGB(255, 128, 0), got.Object != channelID:
			t.Fatalf("%s: %+v", input, got)
		}
	}
}

func TestConverterErrors(t *testing.T) {
	h := newBot(t, commands.WithMode(commands.UserBot))
	seedGuild(h)
	type roleParams struct{ Role discord.Role }
	h.bot.AddCommand(commands.New("role", func(*commands.Context, roleParams) error { return nil }))
	h.run(other, guildID, "!role nobody")
	var nf *commands.NotFoundError
	if err := h.err(t); !errors.As(err, &nf) || nf.Kind != "Role" || !errors.Is(err, commands.ErrBadArgument) {
		t.Fatalf("role: %v", err)
	}
	h.run(other, 0, "!role mods")
	var dm *commands.NoPrivateMessage
	if err := h.err(t); !errors.As(err, &dm) {
		t.Fatalf("dm: %v", err)
	}
	type colourParams struct{ Colour converting.Colour }
	h.bot.AddCommand(commands.New("colour", func(*commands.Context, colourParams) error { return nil }))
	h.run(other, guildID, "!colour #zzzzzz")
	var bad *commands.BadColourArgument
	if err := h.err(t); !errors.As(err, &bad) || bad.Argument != "#zzzzzz" {
		t.Fatalf("colour: %v", err)
	}
	for _, in := range []string{"rgb(1,2)", "0x1000000", "nope"} {
		if _, ok := converting.ParseColour(in); ok {
			t.Errorf("%q parsed", in)
		}
	}
	if c, ok := converting.ParseColour("Dark Teal"); !ok || c.String() != "#11806a" {
		t.Errorf("named: %v", c)
	}
}
