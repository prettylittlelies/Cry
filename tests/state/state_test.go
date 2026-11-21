package state_test

import (
	"encoding/json"
	"testing"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
	"github.com/prettylittlelies/cry/state"
)

func decode(t *testing.T, wire, raw string) any {
	t.Helper()
	newEvent, ok := discord.Events[wire]
	if !ok {
		t.Fatalf("no event type for %s", wire)
	}
	v := newEvent()
	if err := json.Unmarshal([]byte(raw), v); err != nil {
		t.Fatalf("%s: %v", wire, err)
	}
	return v
}

const ready = `{
	"session_id": "abc",
	"user": {"id": "1", "username": "me", "discriminator": "0", "avatar": null},
	"users": [{"id": "2", "username": "pal", "discriminator": "0", "avatar": null}],
	"guilds": [{
		"id": "10", "joined_at": "2021-03-01T12:00:00+00:00", "large": false, "member_count": 3,
		"properties": {"id": "10", "name": "cry", "owner_id": "1"},
		"channels": [{"id": "11", "type": 0, "name": "general"}],
		"roles": [{"id": "10", "name": "@everyone", "permissions": "2048", "position": 0, "color": 0, "hoist": false, "managed": false, "mentionable": false, "flags": 0}]
	}],
	"merged_members": [[{"user_id": "1", "roles": [], "joined_at": "2021-03-01T12:00:00+00:00", "nick": "me", "avatar": null, "flags": 0, "deaf": false, "mute": false}]],
	"private_channels": [{"id": "20", "type": 1, "recipient_ids": ["2"]}],
	"relationships": [{"id": "2", "type": 1, "nickname": null}]
}`

func readyState(t *testing.T) *state.State {
	t.Helper()
	s := state.New(10)
	s.Apply(decode(t, "READY", ready))
	return s
}

func TestReadyPopulatesCache(t *testing.T) {
	s := readyState(t)
	if s.Me().ID != 1 || s.SessionID() != "abc" {
		t.Fatalf("me=%d session=%q", s.Me().ID, s.SessionID())
	}
	if u, ok := s.User(2); !ok || u.Username != "pal" {
		t.Fatalf("user %+v", u)
	}
	g, ok := s.Guild(10)
	if !ok || g.Properties.Name != "cry" || g.MemberCount != 3 {
		t.Fatalf("guild %+v", g)
	}
	c, ok := s.Channel(11)
	if !ok || c.GuildID != 10 || c.Type != discord.ChannelTypeGuildText {
		t.Fatalf("channel %+v", c)
	}
	if dm, ok := s.Channel(20); !ok || dm.Type != discord.ChannelTypeDM {
		t.Fatalf("dm %+v", dm)
	}
	if m, ok := s.Member(10, 1); !ok || m.Nick == nil || *m.Nick != "me" {
		t.Fatalf("merged member not keyed by user_id: %+v ok=%v", m, ok)
	}
	if r, ok := s.Relationship(2); !ok || r.Type != discord.RelationshipTypeFriend {
		t.Fatalf("relationship %+v", r)
	}
}

func TestDispatchUpdates(t *testing.T) {
	s := readyState(t)
	s.Apply(decode(t, "CHANNEL_CREATE", `{"id": "12", "type": 2, "guild_id": "10", "name": "voice"}`))
	if c, ok := s.Channel(12); !ok || c.Type != discord.ChannelTypeGuildVoice {
		t.Fatalf("channel create %+v", c)
	}
	s.Apply(decode(t, "CHANNEL_DELETE", `{"id": "12", "type": 2, "guild_id": "10"}`))
	if _, ok := s.Channel(12); ok {
		t.Fatal("channel delete")
	}
	s.Apply(decode(t, "GUILD_MEMBER_ADD", `{"guild_id": "10", "user": {"id": "3", "username": "new", "discriminator": "0", "avatar": null}, "roles": [], "joined_at": "2026-01-01T00:00:00+00:00", "deaf": false, "mute": false, "flags": 0}`))
	if _, ok := s.Member(10, 3); !ok {
		t.Fatal("member add")
	}
	s.Apply(decode(t, "GUILD_MEMBER_REMOVE", `{"guild_id": "10", "user": {"id": "3", "username": "new", "discriminator": "0", "avatar": null}}`))
	if _, ok := s.Member(10, 3); ok {
		t.Fatal("member remove")
	}
	s.Apply(decode(t, "MESSAGE_CREATE", `{"id": "100", "channel_id": "11", "guild_id": "10", "author": {"id": "2", "username": "pal", "discriminator": "0", "avatar": null}, "content": "hi", "timestamp": "2026-11-21T04:00:00+00:00", "edited_timestamp": null, "tts": false, "mention_everyone": false, "mentions": [], "mention_roles": [], "attachments": [], "embeds": [], "pinned": false, "type": 0}`))
	if m, ok := s.Message(100); !ok || m.Content != "hi" || m.GuildID != 10 {
		t.Fatalf("message %+v", m)
	}
	s.Apply(decode(t, "MESSAGE_DELETE", `{"id": "100", "channel_id": "11"}`))
	if _, ok := s.Message(100); ok {
		t.Fatal("message delete")
	}
	s.Apply(decode(t, "GUILD_DELETE", `{"id": "10", "unavailable": true}`))
	if g, ok := s.Guild(10); !ok || !g.Unavailable {
		t.Fatal("outage keeps guild marked unavailable")
	}
	s.Apply(decode(t, "GUILD_DELETE", `{"id": "10"}`))
	if _, ok := s.Guild(10); ok {
		t.Fatal("leave removes guild")
	}
	if _, ok := s.Channel(11); ok {
		t.Fatal("leave removes guild channels")
	}
}

func TestMessageCacheBounded(t *testing.T) {
	s := state.New(2)
	for i := range 3 {
		id := snowflake.ID(100 + i)
		raw, _ := json.Marshal(map[string]any{"id": id.String(), "channel_id": "1", "author": map[string]any{"id": "2", "username": "x", "discriminator": "0", "avatar": nil}, "content": "m", "timestamp": "2026-11-21T04:00:00+00:00", "edited_timestamp": nil, "tts": false, "mention_everyone": false, "mentions": []any{}, "mention_roles": []any{}, "attachments": []any{}, "embeds": []any{}, "pinned": false, "type": 0})
		s.Apply(decode(t, "MESSAGE_CREATE", string(raw)))
	}
	if _, ok := s.Message(100); ok {
		t.Fatal("oldest message should be evicted")
	}
	if _, ok := s.Message(102); !ok {
		t.Fatal("newest message missing")
	}
}
