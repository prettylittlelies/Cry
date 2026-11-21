package rest_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"
	"sync"
	"testing"

	"github.com/prettylittlelies/cry/rest"
	"github.com/prettylittlelies/cry/snowflake"
)

type gcs struct {
	mu         sync.Mutex
	data       []byte
	failFirst  bool
	preflights []http.Header
	puts       []http.Header
	slotBody   []byte
}

func (g *gcs) handler(t *testing.T, base func() string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		g.mu.Lock()
		defer g.mu.Unlock()
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/channels/42/attachments":
			g.slotBody, _ = io.ReadAll(r.Body)
			fmt.Fprintf(w, `{"attachments":[{"id":"1","upload_url":"%s/upload/abc/cat.png?upload_id=x","upload_filename":"abc/cat.png"}]}`, base())
		case r.Method == http.MethodOptions && r.URL.Path == "/upload/abc/cat.png":
			g.preflights = append(g.preflights, r.Header.Clone())
			w.Header().Set("Access-Control-Allow-Origin", "https://discord.com")
			w.Header().Set("Access-Control-Allow-Methods", "PUT")
			w.Header().Set("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
			w.Header().Set("Access-Control-Max-Age", "3600")
			w.WriteHeader(http.StatusOK)
		case r.Method == http.MethodPut && r.URL.Path == "/upload/abc/cat.png":
			g.puts = append(g.puts, r.Header.Clone())
			body, _ := io.ReadAll(r.Body)
			if r.Header.Get("Content-Range") == "bytes */*" {
				w.Header().Set("Range", fmt.Sprintf("bytes=0-%d", len(g.data)-1))
				w.WriteHeader(http.StatusPermanentRedirect)
				return
			}
			if g.failFirst {
				g.failFirst = false
				g.data = append(g.data, body[:1024]...)
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}
			g.data = append(g.data, body...)
			w.WriteHeader(http.StatusOK)
		default:
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
	}
}

func TestCloudUploadResumes(t *testing.T) {
	if testing.Short() {
		t.Skip("waits for upload backoff")
	}
	g := &gcs{failFirst: true}
	var url string
	c, base := clientURL(t, g.handler(t, func() string { return url }))
	url = base
	payload := bytes.Repeat([]byte("discord!"), 512)
	up, err := c.UploadAttachment(context.Background(), snowflake.ID(42), rest.CloudFile{Name: "cat.png", Data: payload, Spoiler: true})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(g.data, payload) {
		t.Fatalf("stored %d bytes, want %d", len(g.data), len(payload))
	}
	if want := `{"files":[{"filename":"cat.png","file_size":4096,"id":"`; !strings.HasPrefix(string(g.slotBody), want) || !strings.HasSuffix(string(g.slotBody), `","is_clip":false,"original_content_type":"image/png"}]}`) {
		t.Fatalf("slot body %s", g.slotBody)
	}
	var requested []string
	for _, p := range g.preflights {
		requested = append(requested, p.Get("Access-Control-Request-Headers"))
	}
	if want := []string{"content-type", "content-range", "content-range,content-type"}; !slices.Equal(requested, want) {
		t.Fatalf("preflights requested %q, want %q", requested, want)
	}
	last := g.puts[len(g.puts)-1]
	if last.Get("Content-Range") != "bytes 1024-4095/4096" {
		t.Fatalf("resume range %q", last.Get("Content-Range"))
	}
	for _, h := range g.puts {
		for _, forbidden := range []string{"Authorization", "X-Super-Properties", "X-Discord-Locale", "Cookie"} {
			if h.Get(forbidden) != "" {
				t.Fatalf("cross-site PUT leaked %s", forbidden)
			}
		}
		if h.Get("Sec-Fetch-Site") != "cross-site" || h.Get("Origin") != rest.DefaultOrigin || h.Get("Referer") != rest.DefaultOrigin+"/" || h.Get("Accept-Language") != "en-US" {
			t.Fatalf("cross-site headers %v", h)
		}
	}
	att, _ := json.Marshal(up.Attachment(0))
	if string(att) != `{"id":"0","filename":"cat.png","uploaded_filename":"abc/cat.png","is_spoiler":true}` {
		t.Fatalf("attachment %s", att)
	}
}

func TestDeleteUploadKeepsSlash(t *testing.T) {
	var path string
	c := client(t, func(w http.ResponseWriter, r *http.Request) {
		path = r.URL.EscapedPath()
		w.WriteHeader(http.StatusNoContent)
	})
	if err := c.DeleteUpload(context.Background(), "abc/cat.png"); err != nil {
		t.Fatal(err)
	}
	if path != "/attachments/abc/cat.png" {
		t.Fatalf("path %q", path)
	}
}

func TestBodylessPostSendsZeroLength(t *testing.T) {
	var length string
	c := client(t, func(w http.ResponseWriter, r *http.Request) {
		length = r.Header.Get("Content-Length")
		w.WriteHeader(http.StatusNoContent)
	})
	req := &rest.Request{Route: rest.NewRoute(http.MethodPost, "/channels/{channel_id}/typing", rest.P("channel_id", "1"))}
	if _, err := c.Do(context.Background(), req); err != nil {
		t.Fatal(err)
	}
	if length != "0" {
		t.Fatalf("content-length %q", length)
	}
}
