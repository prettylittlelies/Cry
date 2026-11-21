package voice_test

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/prettylittlelies/cry/voice"
)

func oggPage(packets ...[]byte) []byte {
	var table, body []byte
	for _, p := range packets {
		n := len(p)
		for n >= 255 {
			table = append(table, 255)
			n -= 255
		}
		table = append(table, byte(n))
		body = append(body, p...)
	}
	head := append([]byte("OggS"), make([]byte, 22)...)
	head = append(head, byte(len(table)))
	return append(append(head, table...), body...)
}

func TestOggOpusDemux(t *testing.T) {
	long := bytes.Repeat([]byte{0xAB}, 300)
	short := []byte{0xF8, 1, 2}
	stream := append(oggPage([]byte("OpusHead........"), []byte("OpusTags....")), oggPage(long, short)...)
	r := voice.NewOggOpus(bytes.NewReader(stream))
	for _, want := range [][]byte{long, short} {
		got, err := r.Next()
		if err != nil || !bytes.Equal(got, want) {
			t.Fatalf("packet %d bytes err=%v", len(got), err)
		}
	}
	if _, err := r.Next(); !errors.Is(err, io.EOF) {
		t.Fatalf("end: %v", err)
	}
	if _, err := voice.NewOggOpus(bytes.NewReader([]byte("RIFF....................................."))).Next(); !errors.Is(err, voice.ErrNotOgg) {
		t.Fatalf("not ogg: %v", err)
	}
}

func TestOggWriterRoundTrip(t *testing.T) {
	var buf bytes.Buffer
	w := voice.NewOggWriter(&buf)
	packets := [][]byte{bytes.Repeat([]byte{1}, 300), {0xF8, 0xFF, 0xFE}}
	for _, p := range packets {
		if err := w.WritePacket(p); err != nil {
			t.Fatal(err)
		}
	}
	w.Close()
	r := voice.NewOggOpus(&buf)
	for _, want := range packets {
		got, err := r.Next()
		if err != nil || !bytes.Equal(got, want) {
			t.Fatalf("got %d bytes, %v", len(got), err)
		}
	}
}
