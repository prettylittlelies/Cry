//go:build live

package live_test

import (
	"context"
	"encoding/hex"
	"io"
	"log/slog"
	"os"
	"strconv"
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
	guild := envID(t, "CRY_VOICE_GUILD", 1573803464159662120)
	channel := envID(t, "CRY_VOICE_CHANNEL", 1573803465031950469)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
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
	conn, err := voice.Join(joinCtx, c, guild, channel, voice.Options{SelfMute: false})
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

	deadline := time.Now().Add(20 * time.Second)
	for !conn.Encrypted() && time.Now().Before(deadline) {
		if err := conn.Err(); err != nil {
			t.Fatalf("connection failed: %v", err)
		}
		time.Sleep(200 * time.Millisecond)
	}
	if conn.Encrypted() {
		t.Logf("dave active, epoch authenticator %s", hex.EncodeToString(conn.EpochAuthenticator()))
	} else {
		t.Log("dave not active after 20s (fine if you are alone in the channel and the server keeps transport-only encryption)")
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
	if err := conn.Err(); err != nil {
		t.Fatalf("connection failed: %v", err)
	}
}
