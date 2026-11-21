package client_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/prettylittlelies/cry"
	"github.com/prettylittlelies/cry/snowflake"
	"github.com/prettylittlelies/cry/tests/testkit"
)

type call struct {
	Method, Path, Query string
	Body                map[string]any
}

type recorder struct {
	mu    sync.Mutex
	calls []call
	pages map[string][]string
}

func (r *recorder) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	raw, _ := io.ReadAll(req.Body)
	var body map[string]any
	_ = json.Unmarshal(raw, &body)
	r.mu.Lock()
	r.calls = append(r.calls, call{req.Method, req.URL.EscapedPath(), req.URL.RawQuery, body})
	r.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	switch {
	case req.Method == http.MethodGet && strings.HasSuffix(req.URL.Path, "/messages"):
		cursor := req.URL.Query().Get("before") + req.URL.Query().Get("after")
		var msgs []string
		for _, id := range r.pages[cursor] {
			msgs = append(msgs, `{"id":"`+id+`","channel_id":"11","content":"","author":{"id":"2"}}`)
		}
		_, _ = w.Write([]byte("[" + strings.Join(msgs, ",") + "]"))
	case req.URL.Path == "/users/@me/channels":
		_, _ = w.Write([]byte(`{"id":"50","type":1,"recipients":[{"id":"2","username":"pal","discriminator":"0","avatar":null}]}`))
	case req.Method == http.MethodPost:
		_, _ = w.Write([]byte(`{"id":"99","channel_id":"11","content":"ok"}`))
	default:
		w.WriteHeader(http.StatusNoContent)
	}
}

func (r *recorder) take() []call {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := r.calls
	r.calls = nil
	return out
}

func restClient(t *testing.T, r *recorder) *cry.Client {
	t.Helper()
	srv := testkit.NewH2Server(t, r)
	c, err := cry.New(context.Background(), "token", cry.WithSnapshot(snapshot(t)), cry.WithAPIBase(srv.URL), cry.WithRootCAs(srv.Pool))
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func ids(from, n int) []string {
	out := make([]string, n)
	for i := range out {
		out[i] = strconv.Itoa(from - i)
	}
	return out
}

func TestMessageHandles(t *testing.T) {
	r := &recorder{}
	c := restClient(t, r)
	ctx := context.Background()
	msg := c.Message(11, 77)
	msg.Reply(ctx, "hey")
	msg.React(ctx, "<a:party:123456789012345678>")
	msg.Unreact(ctx, "👋")
	msg.RemoveReaction(ctx, "👋", 2)
	calls := r.take()
	if len(calls) != 4 {
		t.Fatalf("calls %+v", calls)
	}
	ref, _ := calls[0].Body["message_reference"].(map[string]any)
	if calls[0].Path != "/channels/11/messages" || ref["message_id"] != "77" || calls[0].Body["nonce"] == nil {
		t.Errorf("reply %+v", calls[0])
	}
	if calls[1].Method != "PUT" || calls[1].Path != "/channels/11/messages/77/reactions/party:123456789012345678/@me" {
		t.Errorf("react %+v", calls[1])
	}
	if calls[2].Method != "DELETE" || calls[2].Path != "/channels/11/messages/77/reactions/%F0%9F%91%8B/@me" {
		t.Errorf("unreact %+v", calls[2])
	}
	if !strings.HasSuffix(calls[3].Path, "/reactions/%F0%9F%91%8B/2") {
		t.Errorf("remove reaction %+v", calls[3])
	}
}

func TestHistoryPaginates(t *testing.T) {
	r := &recorder{pages: map[string][]string{"": ids(300, 100), "201": ids(200, 100), "101": ids(100, 30)}}
	c := restClient(t, r)
	var got []snowflake.ID
	for m, err := range c.Channel(11).History(context.Background(), cry.HistoryOptions{}) {
		if err != nil {
			t.Fatal(err)
		}
		got = append(got, m.ID)
	}
	if len(got) != 230 || got[0] != 300 || got[229] != 71 {
		t.Fatalf("history %d first=%v last=%v", len(got), got[0], got[len(got)-1])
	}
	r.take()
	got = got[:0]
	for m := range c.Channel(11).History(context.Background(), cry.HistoryOptions{Limit: 150}) {
		got = append(got, m.ID)
	}
	calls := r.take()
	if len(got) != 150 || len(calls) != 2 || !strings.Contains(calls[1].Query, "limit=50") {
		t.Fatalf("limited history %d %+v", len(got), calls)
	}
}

func TestMemberAndUserHandles(t *testing.T) {
	r := &recorder{}
	c := restClient(t, r)
	ctx := context.Background()
	c.Member(10, 2).AddRoles(ctx, 30, 31)
	c.Member(10, 2).Ban(ctx, 0)
	calls := r.take()
	if len(calls) != 3 || calls[0].Path != "/guilds/10/members/2/roles/30" || calls[2].Path != "/guilds/10/bans/2" {
		t.Fatalf("member calls %+v", calls)
	}
	if _, err := c.User(2).Send(ctx, "yo"); err != nil {
		t.Fatal(err)
	}
	c.User(2).Send(ctx, "again")
	calls = r.take()
	dms := 0
	for _, cl := range calls {
		if cl.Path == "/users/@me/channels" {
			dms++
		}
	}
	if dms != 1 || calls[len(calls)-1].Path != "/channels/50/messages" {
		t.Fatalf("dm must be opened once and reused: %+v", calls)
	}
	if id, ok := c.State().DirectMessage(2); !ok || id != 50 {
		t.Fatal("dm not cached")
	}
}
