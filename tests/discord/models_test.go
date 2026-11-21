package discord_test

import (
	"encoding/json"
	"testing"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
)

func TestReadyGuildProperties(t *testing.T) {
	raw := `{"id":81384788765712384,"joined_at":"2021-03-01T12:00:00.000000+00:00","member_count":42,"large":false,"lazy":true,"version":1700000000000,
		"properties":{"id":81384788765712384,"name":"cry","owner_id":"53908232506183680","features":["COMMUNITY"],"icon":null,"premium_tier":2},
		"channels":[{"id":"1","type":0,"name":"general","position":0}],
		"roles":[{"id":81384788765712384,"name":"@everyone","permissions":"1071698660929","position":0,"color":0,"hoist":false,"managed":false,"mentionable":false,"flags":0}],
		"emojis":[],"stickers":[]}`
	var g discord.Guild
	if err := json.Unmarshal([]byte(raw), &g); err != nil {
		t.Fatal(err)
	}
	if g.ID != 81384788765712384 || g.Name != "cry" || g.OwnerID != 53908232506183680 || g.MemberCount != 42 || !g.Lazy || g.PremiumTier != 2 {
		t.Fatalf("guild %+v", g)
	}
	if len(g.Channels) != 1 || *g.Channels[0].Name != "general" || g.JoinedAt == nil {
		t.Fatalf("channels %+v joined=%v", g.Channels, g.JoinedAt)
	}
	everyone := g.Role(g.ID)
	if everyone == nil || !everyone.Permissions.Has(discord.PermSendMessages) || everyone.Permissions.Has(discord.PermBanMembers) {
		t.Fatalf("role %+v", everyone)
	}
}

func TestPlainGuild(t *testing.T) {
	var g discord.Guild
	if err := json.Unmarshal([]byte(`{"id":"5","name":"plain","owner_id":"6","unavailable":false}`), &g); err != nil || g.ID != 5 || g.Name != "plain" || g.OwnerID != 6 {
		t.Fatalf("%+v err=%v", g, err)
	}
}

func TestMessageDecode(t *testing.T) {
	raw := `{"id":1234,"channel_id":"99","guild_id":"7","author":{"id":"5","username":"cry","discriminator":"0","global_name":"Cry","avatar":null},
		"content":"hi <@5>","timestamp":"2026-11-21T04:00:00.123000+00:00","edited_timestamp":null,"tts":false,"mention_everyone":false,
		"mentions":[],"mention_roles":[],"attachments":[],"embeds":[{"title":"t","fields":[{"name":"a","value":"b"}]}],"pinned":false,"type":0,"nonce":"1"}`
	var m discord.Message
	if err := json.Unmarshal([]byte(raw), &m); err != nil {
		t.Fatal(err)
	}
	if m.Author.DisplayName() != "Cry" || m.Embeds[0].Fields[0].Value != "b" || m.Timestamp.Year() != 2026 {
		t.Fatalf("message %+v", m)
	}
	if m.JumpURL() != "https://discord.com/channels/7/99/1234" {
		t.Fatalf("jump %s", m.JumpURL())
	}
}

func TestEmojiAndMentions(t *testing.T) {
	id := snowflake.ID(41771983429993937)
	name := "LUL"
	e := discord.Emoji{ID: &id, Name: &name, Animated: true}
	if e.String() != "<a:LUL:41771983429993937>" || e.APIName() != "LUL:41771983429993937" {
		t.Fatalf("emoji %s %s", e.String(), e.APIName())
	}
	unicode := "🔥"
	if (&discord.Emoji{Name: &unicode}).String() != "🔥" {
		t.Fatal("unicode emoji")
	}
	u := discord.User{ID: 5, Username: "cry"}
	if u.Mention() != "<@5>" || u.DisplayName() != "cry" {
		t.Fatal("user helpers")
	}
}

func TestPermissionsJSON(t *testing.T) {
	var p discord.Permissions
	if err := json.Unmarshal([]byte(`"8"`), &p); err != nil || !p.Has(discord.PermBanMembers) {
		t.Fatalf("administrator implies everything: %v %v", p, err)
	}
	out, _ := json.Marshal(discord.PermSendMessages | discord.PermViewChannel)
	if string(out) != `"3072"` {
		t.Fatalf("marshal %s", out)
	}
}

func TestMemberIdentity(t *testing.T) {
	var merged, full discord.Member
	json.Unmarshal([]byte(`{"user_id":10,"roles":["1"],"joined_at":"2021-01-01T00:00:00+00:00","nick":null,"avatar":null,"flags":0,"deaf":false,"mute":false}`), &merged)
	json.Unmarshal([]byte(`{"user":{"id":"11","username":"x","discriminator":"0"},"roles":[],"nick":"n","avatar":null,"flags":0,"deaf":false,"mute":false,"joined_at":null}`), &full)
	if merged.ID() != 10 || !merged.HasRole(1) || full.ID() != 11 {
		t.Fatalf("merged=%d full=%d", merged.ID(), full.ID())
	}
}
