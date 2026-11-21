package rest_test

import (
	"context"
	"io"
	"mime"
	"net/http"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/prettylittlelies/cry/rest"
)

type part struct {
	name, contentType, body string
	rawDisposition          string
}

func TestMessageFormUpload(t *testing.T) {
	var (
		ctype string
		parts []part
	)
	c := client(t, func(w http.ResponseWriter, r *http.Request) {
		ctype = r.Header.Get("Content-Type")
		mr, err := r.MultipartReader()
		if err != nil {
			t.Error(err)
			return
		}
		for {
			p, err := mr.NextRawPart()
			if err != nil {
				break
			}
			b, _ := io.ReadAll(p)
			parts = append(parts, part{
				name:           p.FormName(),
				contentType:    p.Header.Get("Content-Type"),
				body:           string(b),
				rawDisposition: p.Header.Get("Content-Disposition"),
			})
		}
		io.WriteString(w, `{}`)
	})
	form, err := rest.MessageForm(map[string]string{"content": "hi"},
		rest.File{Name: `cat "1".png`, ContentType: "image/png", Data: []byte{0x89, 'P', 'N', 'G'}},
		rest.File{Name: "notes.bin", Data: []byte("raw")},
	)
	if err != nil {
		t.Fatal(err)
	}
	req := &rest.Request{Route: rest.NewRoute(http.MethodPost, "/channels/{channel_id}/messages", rest.P("channel_id", "1")), Form: form}
	if _, err := c.Do(context.Background(), req); err != nil {
		t.Fatal(err)
	}
	mt, params, err := mime.ParseMediaType(ctype)
	if err != nil || mt != "multipart/form-data" || !regexp.MustCompile(`^----WebKitFormBoundary[A-Za-z0-9]{16}$`).MatchString(params["boundary"]) {
		t.Fatalf("content-type %q", ctype)
	}
	if len(parts) != 3 {
		t.Fatalf("got %d parts", len(parts))
	}
	if parts[0].name != "files[0]" || parts[0].contentType != "image/png" || !strings.Contains(parts[0].rawDisposition, `filename="cat %221%22.png"`) {
		t.Errorf("part 0 %+v", parts[0])
	}
	if parts[1].name != "files[1]" || parts[1].contentType != "application/octet-stream" || parts[1].body != "raw" {
		t.Errorf("part 1 %+v", parts[1])
	}
	if parts[2].name != "payload_json" || parts[2].contentType != "" || parts[2].body != `{"content":"hi"}` {
		t.Errorf("part 2 %+v", parts[2])
	}
}

func TestIdleBucketsPruned(t *testing.T) {
	now := time.Unix(1_700_000_000, 0)
	clock := func() time.Time { return now }
	c := client(t, func(w http.ResponseWriter, _ *http.Request) { io.WriteString(w, `{}`) }, rest.WithClock(clock))
	route := func(id string) *rest.Request {
		return &rest.Request{Route: rest.NewRoute(http.MethodGet, "/channels/{channel_id}", rest.P("channel_id", id))}
	}
	for _, id := range []string{"1", "2", "3"} {
		if _, err := c.Do(context.Background(), route(id)); err != nil {
			t.Fatal(err)
		}
	}
	if n := c.Stats().Buckets; n != 3 {
		t.Fatalf("want 3 buckets, got %d", n)
	}
	now = now.Add(4 * time.Minute)
	if _, err := c.Do(context.Background(), route("1")); err != nil {
		t.Fatal(err)
	}
	now = now.Add(2 * time.Minute)
	if _, err := c.Do(context.Background(), route("4")); err != nil {
		t.Fatal(err)
	}
	if n := c.Stats().Buckets; n != 2 {
		t.Fatalf("want buckets 1 and 4 kept, got %d", n)
	}
}
