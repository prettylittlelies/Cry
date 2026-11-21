package gateway_test

import (
	"context"
	"encoding/json"
	"slices"
	"testing"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/gateway"
	"github.com/prettylittlelies/cry/snowflake"
	"github.com/prettylittlelies/cry/tests/testkit"
)

func ready(t *testing.T, h *harness) *testkit.GatewayPeer {
	t.Helper()
	p := testkit.AcceptGatewayPeer(t, h.srv)
	p.Hello(10000)
	p.Expect(gateway.OpIdentify)
	p.Send(map[string]any{"op": 0, "s": 1, "t": "READY", "d": map[string]any{"session_id": "abc", "resume_gateway_url": h.srv.URL}})
	<-h.events
	return p
}

func TestPresenceIsSentAndReidentified(t *testing.T) {
	h := start(t, false)
	p := ready(t, h)
	state := "listening"
	err := h.gw.SetPresence(context.Background(), gateway.Presence{
		Status:     discord.StatusTypeDnd,
		Activities: []discord.Activity{{Name: "Custom Status", Type: discord.ActivityTypeCustom, State: &state}},
	})
	if err != nil {
		t.Fatal(err)
	}
	sent := p.Expect(gateway.OpPresenceUpdate)
	if got := testkit.KeysOf(t, sent.D); !slices.Equal(got, []string{"status", "since", "activities", "afk"}) {
		t.Fatalf("presence keys %v", got)
	}
	var body struct {
		Status     string             `json:"status"`
		Activities []discord.Activity `json:"activities"`
	}
	json.Unmarshal(sent.D, &body)
	if body.Status != "dnd" || len(body.Activities) != 1 || *body.Activities[0].State != "listening" {
		t.Fatalf("presence %s", sent.D)
	}

	p.Send(map[string]any{"op": 9, "d": false})
	id := p.Expect(gateway.OpIdentify)
	var identify struct {
		Presence struct {
			Status string `json:"status"`
		} `json:"presence"`
	}
	json.Unmarshal(id.D, &identify)
	if identify.Presence.Status != "dnd" {
		t.Fatalf("re-identify lost presence: %s", id.D)
	}
}

func TestMemberAndSubscriptionCommands(t *testing.T) {
	h := start(t, false)
	p := ready(t, h)
	ctx := context.Background()
	guild := snowflake.ID(100000000000000001)
	query, limit := "", 0
	h.gw.RequestMembers(ctx, gateway.MemberRequest{GuildIDs: []snowflake.ID{guild}, Query: &query, Limit: &limit, Presences: true, Nonce: "n"})
	members := p.Expect(gateway.OpRequestGuildMembers)
	if got := testkit.KeysOf(t, members.D); !slices.Equal(got, []string{"guild_id", "query", "limit", "presences", "user_ids", "nonce"}) {
		t.Fatalf("member request keys %v: %s", got, members.D)
	}

	yes := true
	h.gw.SubscribeGuilds(ctx, map[snowflake.ID]gateway.GuildSubscription{
		guild: {Typing: &yes, Activities: &yes, Channels: map[snowflake.ID][][2]int{400000000000000004: {{0, 99}}}},
	})
	sub := p.Expect(gateway.OpGuildSubscriptionsBulk)
	var bulk struct {
		Subscriptions map[string]struct {
			Typing   bool               `json:"typing"`
			Channels map[string][][]int `json:"channels"`
			Threads  *bool              `json:"threads"`
		} `json:"subscriptions"`
	}
	json.Unmarshal(sub.D, &bulk)
	s, ok := bulk.Subscriptions[guild.String()]
	if !ok || !s.Typing || s.Threads != nil || !slices.Equal(s.Channels["400000000000000004"][0], []int{0, 99}) {
		t.Fatalf("subscriptions %s", sub.D)
	}

	channel := snowflake.ID(400000000000000004)
	h.gw.UpdateVoiceState(ctx, discord.UpdateVoiceState{GuildID: &guild, ChannelID: &channel, SelfMute: true, PreferredRegions: []string{"rotterdam", "frankfurt"}})
	voice := p.Expect(gateway.OpVoiceStateUpdate)
	var v struct {
		ChannelID       string   `json:"channel_id"`
		SelfMute        bool     `json:"self_mute"`
		PreferredRegion string   `json:"preferred_region"`
		Regions         []string `json:"preferred_regions"`
	}
	json.Unmarshal(voice.D, &v)
	if v.ChannelID != channel.String() || !v.SelfMute || v.PreferredRegion != "rotterdam" || len(v.Regions) != 2 {
		t.Fatalf("voice %s", voice.D)
	}
	h.gw.UpdateVoiceState(ctx, discord.UpdateVoiceState{GuildID: &guild})
	leave := p.Expect(gateway.OpVoiceStateUpdate)
	var l map[string]any
	json.Unmarshal(leave.D, &l)
	if l["channel_id"] != nil || l["guild_id"] != guild.String() {
		t.Fatalf("leave %s", leave.D)
	}

	h.gw.StreamSetPaused(ctx, "guild:1:2:3", true)
	paused := p.Expect(gateway.OpStreamSetPaused)
	if got := testkit.KeysOf(t, paused.D); !slices.Equal(got, []string{"stream_key", "paused"}) {
		t.Fatalf("stream keys %v", got)
	}
}
