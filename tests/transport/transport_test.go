package transport_test

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/andybalholm/brotli"
	"github.com/klauspost/compress/zstd"
	"github.com/prettylittlelies/cry/fingerprint"
	"github.com/prettylittlelies/cry/tests/testkit"
	"github.com/prettylittlelies/cry/transport"
)

func setup(t *testing.T, h http.Handler) (*testkit.H2Server, *transport.Transport) {
	t.Helper()
	srv := testkit.NewH2Server(t, h)
	tr := transport.New(fingerprint.Electron44(), transport.WithRootCAs(srv.Pool))
	t.Cleanup(tr.CloseIdleConnections)
	return srv, tr
}

func do(t *testing.T, tr *transport.Transport, req *http.Request) (*http.Response, []byte) {
	t.Helper()
	resp, err := tr.RoundTrip(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	return resp, b
}

func TestGet(t *testing.T) {
	srv, tr := setup(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Proto", r.Proto)
		w.Header().Set("X-Auth", r.Header.Get("Authorization"))
		io.WriteString(w, r.Method+" "+r.URL.RequestURI())
	}))
	req, _ := http.NewRequest(http.MethodGet, srv.URL+"/api/v9/users/@me?x=1", nil)
	req.Header.Set("Authorization", "token")
	resp, body := do(t, tr, req)
	if resp.StatusCode != http.StatusOK || string(body) != "GET /api/v9/users/@me?x=1" {
		t.Fatalf("got %d %q", resp.StatusCode, body)
	}
	if resp.Header.Get("X-Proto") != "HTTP/2.0" || resp.Header.Get("X-Auth") != "token" {
		t.Fatalf("headers %v", resp.Header)
	}
}

func TestUploadFlowControl(t *testing.T) {
	payload := bytes.Repeat([]byte("cry"), 4<<20/3)
	sum := sha256.Sum256(payload)
	srv, tr := setup(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := sha256.New()
		io.Copy(h, r.Body)
		io.WriteString(w, hex.EncodeToString(h.Sum(nil)))
	}))
	req, _ := http.NewRequest(http.MethodPost, srv.URL, bytes.NewReader(payload))
	_, body := do(t, tr, req)
	if string(body) != hex.EncodeToString(sum[:]) {
		t.Fatal("upload digest mismatch")
	}
}

func TestDownloadFlowControl(t *testing.T) {
	const size = 24 << 20
	srv, tr := setup(t, http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Length", strconv.Itoa(size))
		w.Write(make([]byte, size))
	}))
	req, _ := http.NewRequest(http.MethodGet, srv.URL, nil)
	_, body := do(t, tr, req)
	if len(body) != size {
		t.Fatalf("got %d bytes", len(body))
	}
}

func TestEncodings(t *testing.T) {
	const text = "discord userdoccers"
	encoders := map[string]func(io.Writer) io.WriteCloser{
		"br":      func(w io.Writer) io.WriteCloser { return brotli.NewWriter(w) },
		"gzip":    func(w io.Writer) io.WriteCloser { return gzip.NewWriter(w) },
		"deflate": func(w io.Writer) io.WriteCloser { fw, _ := flate.NewWriter(w, flate.DefaultCompression); return fw },
		"zstd":    func(w io.Writer) io.WriteCloser { zw, _ := zstd.NewWriter(w); return zw },
	}
	srv, tr := setup(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enc := r.URL.Query().Get("e")
		w.Header().Set("Content-Encoding", enc)
		cw := encoders[enc](w)
		io.WriteString(cw, text)
		cw.Close()
	}))
	for name := range encoders {
		t.Run(name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, srv.URL+"/?e="+name, nil)
			resp, body := do(t, tr, req)
			if string(body) != text || !resp.Uncompressed || resp.Header.Get("Content-Encoding") != "" {
				t.Fatalf("got %q uncompressed=%v", body, resp.Uncompressed)
			}
		})
	}
}

func TestMultiplexing(t *testing.T) {
	srv, tr := setup(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, r.URL.Query().Get("i"))
	}))
	var wg sync.WaitGroup
	for i := range 64 {
		wg.Go(func() {
			req, _ := http.NewRequest(http.MethodGet, srv.URL+"/?i="+strconv.Itoa(i), nil)
			if _, body := do(t, tr, req); string(body) != strconv.Itoa(i) {
				t.Errorf("stream %d got %q", i, body)
			}
		})
	}
	wg.Wait()
	if n := srv.Connections(); n != 1 {
		t.Fatalf("want 1 connection, got %d", n)
	}
}

func TestCancelKeepsConnection(t *testing.T) {
	release := make(chan struct{})
	srv, tr := setup(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/slow" {
			select {
			case <-release:
			case <-r.Context().Done():
			}
			return
		}
		io.WriteString(w, "ok")
	}))
	defer close(release)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	slow, _ := http.NewRequestWithContext(ctx, http.MethodGet, srv.URL+"/slow", nil)
	if _, err := tr.RoundTrip(slow); err == nil {
		t.Fatal("expected cancellation error")
	}
	req, _ := http.NewRequest(http.MethodGet, srv.URL+"/fast", nil)
	if _, body := do(t, tr, req); string(body) != "ok" {
		t.Fatalf("got %q", body)
	}
	if n := srv.Connections(); n != 1 {
		t.Fatalf("want 1 connection, got %d", n)
	}
}

func TestRejectsPlainHTTP(t *testing.T) {
	tr := transport.New(fingerprint.Electron44())
	req, _ := http.NewRequest(http.MethodGet, "http://discord.com", nil)
	if _, err := tr.RoundTrip(req); err != transport.ErrUnsupported {
		t.Fatalf("got %v", err)
	}
}
