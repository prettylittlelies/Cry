package state_test

import (
	"slices"
	"testing"

	"github.com/prettylittlelies/cry/snowflake"
)

func apply(t *testing.T, st interface{ Apply(any) }, events ...[2]string) {
	t.Helper()
	for _, e := range events {
		st.Apply(decode(t, e[0], e[1]))
	}
}

func TestGuildEvents(t *testing.T) {
	st := readyState(t)
	apply(t, st,
		[2]string{"GUILD_UPDATE", `{"id":"10","name":"renamed","owner_id":"2","icon":null,"banner":null,"home_header":null,"splash":null,"discovery_splash":null,"application_id":null,"description":null,"afk_channel_id":null,"afk_timeout":0}`},
		[2]string{"GUILD_ROLE_CREATE", `{"guild_id":"10","role":{"id":"30","name":"mods","permissions":"8","position":1,"color":0,"hoist":false,"managed":false,"mentionable":false}}`},
		[2]string{"GUILD_MEMBER_UPDATE", `{"guild_id":"10","user":{"id":"1","username":"me","discriminator":"0","avatar":null},"roles":["30"],"joined_at":"2021-03-01T12:00:00+00:00","nick":"boss","avatar":null,"premium_since":null,"communication_disabled_until":null,"flags":0}`},
		[2]string{"GUILD_ROLE_UPDATE", `{"guild_id":"10","role":{"id":"30","name":"admins","permissions":"8","position":1,"color":0,"hoist":false,"managed":false,"mentionable":false}}`},
	)
	g, _ := st.Guild(10)
	if g.Properties.Name != "renamed" || len(g.Roles) != 2 || g.Roles[1].Name != "admins" {
		t.Fatalf("guild %+v roles %+v", g.Properties.Name, g.Roles)
	}
	m, _ := st.Member(10, 1)
	if m.Nick == nil || *m.Nick != "boss" || !slices.Equal(m.Roles, []snowflake.ID{30}) {
		t.Fatalf("member %+v", m)
	}
	apply(t, st, [2]string{"GUILD_ROLE_DELETE", `{"guild_id":"10","role_id":"30"}`})
	g, _ = st.Guild(10)
	m, _ = st.Member(10, 1)
	if len(g.Roles) != 1 || len(m.Roles) != 0 {
		t.Fatalf("role delete left %+v / %+v", g.Roles, m.Roles)
	}

	apply(t, st,
		[2]string{"GUILD_MEMBERS_CHUNK", `{"guild_id":"10","members":[{"user":{"id":"3","username":"new","discriminator":"0","avatar":null},"roles":[],"joined_at":"2021-03-01T12:00:00+00:00","flags":0}],"chunk_index":0,"chunk_count":1}`},
		[2]string{"VOICE_STATE_UPDATE", `{"guild_id":"10","channel_id":"11","user_id":"3","session_id":"s","deaf":false,"mute":false,"self_deaf":false,"self_mute":true,"self_video":false,"suppress":false,"request_to_speak_timestamp":null}`},
	)
	if u, ok := st.User(3); !ok || u.Username != "new" {
		t.Fatal("chunked member's user not cached")
	}
	if v, ok := st.VoiceState(10, 3); !ok || !v.SelfMute || len(st.VoiceStates(10)) != 1 {
		t.Fatalf("voice %+v", v)
	}
	apply(t, st, [2]string{"VOICE_STATE_UPDATE", `{"guild_id":"10","channel_id":null,"user_id":"3","session_id":"s","deaf":false,"mute":false,"self_deaf":false,"self_mute":false,"self_video":false,"suppress":false,"request_to_speak_timestamp":null}`})
	if _, ok := st.VoiceState(10, 3); ok {
		t.Fatal("leaving voice must drop the state")
	}
}

func TestMessageEvents(t *testing.T) {
	st := readyState(t)
	msg := func(id string) [2]string {
		return [2]string{"MESSAGE_CREATE", `{"id":"` + id + `","channel_id":"11","author":{"id":"2","username":"pal","discriminator":"0","avatar":null},"content":"hi","timestamp":"2021-03-01T12:00:00+00:00","edited_timestamp":null,"tts":false,"mention_everyone":false,"mentions":[],"mention_roles":[],"attachments":[],"embeds":[],"pinned":false,"type":0}`}
	}
	reaction := func(kind, user, emoji string, burst int) [2]string {
		return [2]string{kind, `{"user_id":"` + user + `","channel_id":"11","message_id":"100","message_author_id":"2","emoji":` + emoji + `,"type":` + string(rune('0'+burst)) + `}`}
	}
	wave := `{"id":null,"name":"👋"}`
	apply(t, st, msg("100"), msg("101"), msg("102"),
		[2]string{"MESSAGE_UPDATE", `{"id":"100","channel_id":"11","author":{"id":"2","username":"pal","discriminator":"0","avatar":null},"content":"edited","timestamp":"2021-03-01T12:00:00+00:00","edited_timestamp":"2021-03-01T12:01:00+00:00","tts":false,"mention_everyone":false,"mentions":[],"mention_roles":[],"attachments":[],"embeds":[],"pinned":false,"type":0}`},
		reaction("MESSAGE_REACTION_ADD", "2", wave, 0),
		reaction("MESSAGE_REACTION_ADD", "1", wave, 0),
		reaction("MESSAGE_REACTION_ADD", "2", wave, 1),
		[2]string{"MESSAGE_DELETE_BULK", `{"ids":["101","102"],"channel_id":"11"}`},
	)
	m, _ := st.Message(100)
	if m.Content != "edited" || m.EditedTimestamp == nil || len(m.Reactions) != 1 {
		t.Fatalf("message %+v", m)
	}
	r := m.Reactions[0]
	if r.Count != 3 || r.CountDetails.Normal != 2 || r.CountDetails.Burst != 1 || !r.Me || r.MeBurst {
		t.Fatalf("reaction %+v", r)
	}
	if _, ok := st.Message(101); ok {
		t.Fatal("bulk delete kept 101")
	}
	apply(t, st, reaction("MESSAGE_REACTION_REMOVE", "1", wave, 0))
	m, _ = st.Message(100)
	if m.Reactions[0].Count != 2 || m.Reactions[0].Me {
		t.Fatalf("after remove %+v", m.Reactions[0])
	}
	apply(t, st, [2]string{"MESSAGE_REACTION_REMOVE_ALL", `{"channel_id":"11","message_id":"100"}`})
	m, _ = st.Message(100)
	if m.Reactions != nil {
		t.Fatalf("remove all %+v", m.Reactions)
	}
}

func TestSocialEvents(t *testing.T) {
	st := readyState(t)
	apply(t, st,
		[2]string{"RELATIONSHIP_UPDATE", `{"id":"2","type":1,"nickname":"bestie","user_ignored":false}`},
		[2]string{"CHANNEL_RECIPIENT_ADD", `{"channel_id":"20","user":{"id":"3","username":"x","discriminator":"0","avatar":null}}`},
		[2]string{"READY_SUPPLEMENTAL", `{"guilds":[{"id":"10","voice_states":[]}],"merged_members":[],"merged_presences":{"friends":[{"user":{"id":"2"},"status":"idle","activities":[],"client_status":{}}],"guilds":[]},"lazy_private_channels":[],"disclose":[],"game_invites":[]}`},
	)
	if r, _ := st.Relationship(2); r.Nickname == nil || *r.Nickname != "bestie" {
		t.Fatalf("relationship %+v", r)
	}
	if c, _ := st.Channel(20); len(c.Recipients) != 1 {
		t.Fatalf("recipients %+v", c.Recipients)
	}
	if p, ok := st.Presence(2); !ok || p.Status != "idle" {
		t.Fatalf("presence %+v", p)
	}
	apply(t, st, [2]string{"CHANNEL_RECIPIENT_REMOVE", `{"channel_id":"20","user":{"id":"3","username":"x","discriminator":"0","avatar":null}}`})
	if c, _ := st.Channel(20); len(c.Recipients) != 0 {
		t.Fatalf("recipient remove %+v", c.Recipients)
	}
}
