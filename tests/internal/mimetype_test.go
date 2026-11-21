package internal_test

import (
	"testing"

	"github.com/prettylittlelies/cry/internal/mimetype"
)

func TestChromiumMimePrecedence(t *testing.T) {
	cases := map[string]string{
		"cat.png":   "image/png",
		"song.mp3":  "audio/mpeg",
		"clip.webm": "video/webm",
		"app.js":    "text/javascript",
		"icon.ICO":  "image/x-icon",
		"notes.txt": "text/plain",
		"data.json": "application/json",
		"pic.jfif":  "image/jpeg",
		"raw.xyz":   "",
		"noext":     "",
	}
	for name, want := range cases {
		if got := mimetype.ForFilename(name); got != want {
			t.Errorf("%s: got %q, want %q", name, got, want)
		}
	}
}
