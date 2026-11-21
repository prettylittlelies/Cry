package gateway_test

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/klauspost/compress/zstd"
	"github.com/prettylittlelies/cry/gateway/codec"
	"github.com/prettylittlelies/cry/internal/etf"
	"github.com/prettylittlelies/cry/transport/websocket"
)

var errReadAhead = errors.New("decoder read past the current message")

type strictSource struct {
	chunks  [][]byte
	allowed int
	served  int
}

func (s *strictSource) Read() (websocket.Opcode, []byte, error) {
	if s.served >= s.allowed || s.served >= len(s.chunks) {
		return 0, nil, errReadAhead
	}
	c := s.chunks[s.served]
	s.served++
	return websocket.Binary, c, nil
}

type flusher interface {
	io.Writer
	Flush() error
}

func stream(t *testing.T, w flusher, buf *bytes.Buffer, payloads [][]byte) [][]byte {
	t.Helper()
	var chunks [][]byte
	for _, p := range payloads {
		if _, err := w.Write(p); err != nil {
			t.Fatal(err)
		}
		if err := w.Flush(); err != nil {
			t.Fatal(err)
		}
		chunks = append(chunks, bytes.Clone(buf.Bytes()))
		buf.Reset()
	}
	return chunks
}

func events(n int) []map[string]any {
	out := make([]map[string]any, n)
	for i := range out {
		out[i] = map[string]any{"op": 0, "s": i + 1, "t": "MESSAGE_CREATE", "d": map[string]any{"content": fmt.Sprintf("event %d %s", i, bytes.Repeat([]byte("x"), 300*i))}}
	}
	return out
}

func encode(t *testing.T, enc codec.Encoding, v any) []byte {
	t.Helper()
	if enc == codec.ETF {
		b, err := etf.Marshal(v)
		if err != nil {
			t.Fatal(err)
		}
		return b
	}
	b, _ := json.Marshal(v)
	return b
}

func TestStreamsDoNotReadAhead(t *testing.T) {
	for _, comp := range []codec.Compression{codec.ZstdStream, codec.ZlibStream} {
		for _, enc := range []codec.Encoding{codec.ETF, codec.JSON} {
			t.Run(string(comp)+"/"+string(enc), func(t *testing.T) {
				evs := events(6)
				payloads := make([][]byte, len(evs))
				for i, e := range evs {
					payloads[i] = encode(t, enc, e)
				}
				var buf bytes.Buffer
				var w flusher
				if comp == codec.ZstdStream {
					zw, _ := zstd.NewWriter(&buf)
					w = zw
				} else {
					w = zlib.NewWriter(&buf)
				}
				src := &strictSource{chunks: stream(t, w, &buf, payloads), allowed: 1}
				r, err := codec.NewReader(src, enc, comp)
				if err != nil {
					t.Fatal(err)
				}
				defer r.Close()
				for i, e := range evs {
					src.allowed = i + 1
					raw, err := r.Next()
					if err != nil {
						t.Fatalf("event %d: %v", i, err)
					}
					want, _ := json.Marshal(e)
					var got, exp any
					json.Unmarshal(raw, &got)
					json.Unmarshal(want, &exp)
					if fmt.Sprint(got) != fmt.Sprint(exp) {
						t.Fatalf("event %d:\n got  %s\n want %s", i, raw, want)
					}
				}
				if _, err := r.Next(); !errors.Is(err, errReadAhead) {
					t.Fatalf("expected the reader to ask for a new message, got %v", err)
				}
			})
		}
	}
}

func TestPlainETF(t *testing.T) {
	msg := encode(t, codec.ETF, map[string]any{"op": 10, "d": map[string]any{"heartbeat_interval": 41250}})
	src := &strictSource{chunks: [][]byte{msg}, allowed: 1}
	r, err := codec.NewReader(src, codec.ETF, codec.None)
	if err != nil {
		t.Fatal(err)
	}
	raw, err := r.Next()
	if err != nil || string(raw) != `{"d":{"heartbeat_interval":41250},"op":10}` {
		t.Fatalf("got %s err=%v", raw, err)
	}
}

func TestEncodeOpcodes(t *testing.T) {
	op, b, err := codec.Encode(codec.ETF, map[string]int{"op": 1})
	if err != nil || op != websocket.Binary || b[0] != 131 {
		t.Fatalf("etf op=%v b=%x err=%v", op, b, err)
	}
	op, b, _ = codec.Encode(codec.JSON, map[string]int{"op": 1})
	if op != websocket.Text || string(b) != `{"op":1}` {
		t.Fatalf("json op=%v b=%s", op, b)
	}
}
