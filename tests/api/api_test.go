package api_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/prettylittlelies/cry/api"
	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/rest"
	"github.com/prettylittlelies/cry/tests/testkit"
)

type seen struct {
	method, path, rawQuery, reason string
	body                           []byte
}

func server(t *testing.T, status int, reply string) (*api.API, *seen) {
	t.Helper()
	got := &seen{}
	c, _ := testkit.RESTClient(t, func(w http.ResponseWriter, r *http.Request) {
		got.method, got.path, got.rawQuery = r.Method, r.URL.EscapedPath(), r.URL.RawQuery
		got.reason = r.Header.Get("X-Audit-Log-Reason")
		got.body, _ = io.ReadAll(r.Body)
		w.WriteHeader(status)
		io.WriteString(w, reply)
	})
	return api.New(c), got
}

func TestGetUserProfile(t *testing.T) {
	a, got := server(t, http.StatusOK, `{"user":{"id":"5","username":"cry","discriminator":"0","avatar":null},"user_profile":{"bio":"hi","pronouns":"they/them"},"badges":[],"guild_badges":[],"connected_accounts":[],"premium_type":null,"premium_since":null,"premium_guild_since":null}`)
	p, err := a.GetUserProfile(context.Background(), 5, discord.GetUserProfileQuery{WithMutualGuilds: true})
	if err != nil {
		t.Fatal(err)
	}
	if got.method != http.MethodGet || got.path != "/users/5/profile" || got.rawQuery != "with_mutual_guilds=true" {
		t.Fatalf("request %s %s ?%s", got.method, got.path, got.rawQuery)
	}
	if p.User.Username != "cry" || p.UserProfile.Pronouns != "they/them" {
		t.Fatalf("profile %+v", p)
	}
}

func TestCreateMessage(t *testing.T) {
	a, got := server(t, http.StatusOK, `{"id":"9","channel_id":"7","author":{"id":"5","username":"cry","discriminator":"0","avatar":null},"content":"hello","timestamp":"2026-11-21T04:00:00+00:00","edited_timestamp":null,"tts":false,"mention_everyone":false,"mentions":[],"mention_roles":[],"attachments":[],"embeds":[],"pinned":false,"type":0}`)
	m, err := a.CreateMessage(context.Background(), 7, discord.CreateMessageParams{Content: "hello"})
	if err != nil {
		t.Fatal(err)
	}
	var body map[string]any
	json.Unmarshal(got.body, &body)
	if got.method != http.MethodPost || got.path != "/channels/7/messages" || body["content"] != "hello" {
		t.Fatalf("request %s %s %s", got.method, got.path, got.body)
	}
	if m.ID != 9 || m.Content != "hello" {
		t.Fatalf("message %+v", m)
	}
}

func TestNoContentWithReason(t *testing.T) {
	a, got := server(t, http.StatusNoContent, "")
	if err := a.AddGuildMemberRole(context.Background(), 1, 2, 3, api.WithReason("promoted")); err != nil {
		t.Fatal(err)
	}
	if got.method != http.MethodPut || got.path != "/guilds/1/members/2/roles/3" || got.reason != "promoted" {
		t.Fatalf("request %s %s reason=%q", got.method, got.path, got.reason)
	}
}

func TestRawPathParam(t *testing.T) {
	a, got := server(t, http.StatusNoContent, "")
	if err := a.DeleteAttachment(context.Background(), "abc/cat.png"); err != nil {
		t.Fatal(err)
	}
	if got.path != "/attachments/abc/cat.png" {
		t.Fatalf("path %q", got.path)
	}
}

func TestValues(t *testing.T) {
	limit := int64(50)
	q := struct {
		Limit  *int64  `json:"limit,omitempty"`
		Before string  `json:"before,omitempty"`
		Types  []int64 `json:"types,omitempty"`
		Skip   *bool   `json:"skip,omitempty"`
		Always bool    `json:"always"`
	}{Limit: &limit, Types: []int64{1, 2}}
	if got := rest.Values(q).Encode(); got != "always=false&limit=50&types=1&types=2" {
		t.Fatalf("values %s", got)
	}
}
