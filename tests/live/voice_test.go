//go:build live

package live_test

import (
	"context"
	"encoding/hex"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/prettylittlelies/cry"
	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
	"github.com/prettylittlelies/cry/voice"
)

func envID(t *testing.T, key string, fallback snowflake.ID) snowflake.ID {
	t.Helper()
	raw := os.Getenv(key)
	if raw == "" {
		return fallback
	}
	n, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		t.Fatalf("%s: %v", key, err)
	}
	return snowflake.ID(n)
}

type silence struct{ left int }

func (s *silence) Next() ([]byte, error) {
	if s.left == 0 {
		return nil, io.EOF
	}
	s.left--
	return []byte{0xF8, 0xFF, 0xFE}, nil
}

func TestVoiceJoin(t *testing.T) {
	tok := token(t)
	guild := envID(t, "CRY_VOICE_GUILD", 1105071842064076830)
	channel := envID(t, "CRY_VOICE_CHANNEL", 1254193760045236354)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
	c, err := cry.New(ctx, tok, cry.WithLogger(logger))
	if err != nil {
		t.Fatal(err)
	}
	ready := make(chan struct{}, 1)
	cry.On(c, func(_ context.Context, _ *discord.ReadyEvent) {
		select {
		case ready <- struct{}{}:
		default:
		}
	})
	cry.On(c, func(_ context.Context, e *cry.RawEvent) {
		if e.Type == "READY" {
			t.Errorf("READY arrived but did not decode into discord.ReadyEvent (see the warning above)")
		}
	})
	runCtx, stop := context.WithCancel(ctx)
	defer stop()
	go c.Run(runCtx)
	select {
	case <-ready:
	case <-time.After(45 * time.Second):
		t.Fatal("no decoded READY within 45s")
	}
	t.Logf("logged in as %d", c.State().Me().ID)

	joinCtx, joinCancel := context.WithTimeout(ctx, 30*time.Second)
	defer joinCancel()
	rec := newRecorder(t)
	opts := voice.Options{OnOpus: rec.write}
	if os.Getenv("CRY_VOICE_CAMERA") != "" {
		opts.Video = voice.Camera
	}
	conn, err := voice.Join(joinCtx, c, guild, channel, opts)
	if err != nil {
		t.Fatalf("join: %v", err)
	}
	defer func() {
		closeCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := conn.Close(closeCtx); err != nil {
			t.Logf("close: %v", err)
		}
	}()
	t.Log("voice handshake complete")

	wait := 20 * time.Second
	if raw := os.Getenv("CRY_VOICE_WAIT"); raw != "" {
		if wait, err = time.ParseDuration(raw); err != nil {
			t.Fatalf("CRY_VOICE_WAIT: %v", err)
		}
	}
	t.Logf("waiting up to %s for dave; join the channel from another account now", wait)
	deadline := time.Now().Add(wait)
	for !conn.Encrypted() && time.Now().Before(deadline) {
		if err := conn.Err(); err != nil {
			t.Fatalf("connection failed: %v", err)
		}
		time.Sleep(200 * time.Millisecond)
	}
	if conn.Encrypted() {
		t.Logf("dave active, epoch authenticator %s", hex.EncodeToString(conn.EpochAuthenticator()))
	} else {
		t.Logf("dave not active after %s (expected if nobody else joined)", wait)
	}

	if path := os.Getenv("CRY_VOICE_OGG"); path != "" {
		f, err := os.Open(path)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		if err := conn.Play(ctx, voice.NewOggOpus(f)); err != nil {
			t.Fatalf("play: %v", err)
		}
		t.Log("played", path)
	} else if err := conn.Play(ctx, &silence{left: 250}); err != nil {
		t.Fatalf("play silence: %v", err)
	}
	if path := os.Getenv("CRY_VOICE_CAMERA"); path != "" {
		playIVF(t, ctx, conn, path, "camera")
	}
	if path := os.Getenv("CRY_VOICE_STREAM"); path != "" {
		streamCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		stream, err := conn.StartStream(streamCtx, voice.Options{})
		cancel()
		if err != nil {
			t.Fatalf("start stream: %v", err)
		}
		t.Log("go live started")
		playIVF(t, ctx, stream, path, "stream")
		closeCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		if err := stream.Close(closeCtx); err != nil {
			t.Logf("stop stream: %v", err)
		}
		cancel()
	}
	if raw := os.Getenv("CRY_VOICE_RECORD"); raw != "" {
		d, err := time.ParseDuration(raw)
		if err != nil {
			t.Fatalf("CRY_VOICE_RECORD: %v", err)
		}
		rec.start()
		t.Logf("recording for %s, talk from another account now", d)
		time.Sleep(d)
		rec.stop()
	}
	if err := conn.Err(); err != nil {
		t.Fatalf("connection failed: %v", err)
	}
}

func playIVF(t *testing.T, ctx context.Context, conn *voice.Conn, path, what string) {
	t.Helper()
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	t.Logf("%s: negotiated video codec %q", what, conn.VideoCodec())
	start := time.Now()
	if err := conn.PlayIVF(ctx, f, voice.VideoOptions{}); err != nil {
		t.Fatalf("%s: %v", what, err)
	}
	t.Logf("%s: played %s in %s", what, path, time.Since(start).Round(time.Millisecond))
}

type recorder struct {
	t       *testing.T
	mu      sync.Mutex
	on      bool
	files   map[snowflake.ID]*os.File
	writers map[snowflake.ID]*voice.OggWriter
	frames  map[snowflake.ID]int
}

func newRecorder(t *testing.T) *recorder {
	return &recorder{t: t, files: map[snowflake.ID]*os.File{}, writers: map[snowflake.ID]*voice.OggWriter{}, frames: map[snowflake.ID]int{}}
}

func (r *recorder) start() {
	r.mu.Lock()
	r.on = true
	r.mu.Unlock()
}

func (r *recorder) write(user snowflake.ID, frame []byte) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if !r.on {
		return
	}
	w, ok := r.writers[user]
	if !ok {
		if err := os.MkdirAll("recordings", 0o755); err != nil {
			r.t.Error(err)
			return
		}
		f, err := os.Create(filepath.Join("recordings", user.String()+".ogg"))
		if err != nil {
			r.t.Error(err)
			return
		}
		r.files[user], w = f, voice.NewOggWriter(f)
		r.writers[user] = w
	}
	if err := w.WritePacket(frame); err == nil {
		r.frames[user]++
	}
}

func (r *recorder) stop() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.on = false
	if len(r.writers) == 0 {
		r.t.Log("heard nobody")
	}
	for user, w := range r.writers {
		w.Close()
		r.files[user].Close()
		abs, _ := filepath.Abs(r.files[user].Name())
		r.t.Logf("user %d: %d frames (%.1fs) -> %s", user, r.frames[user], float64(r.frames[user])*0.02, abs)
	}
}
