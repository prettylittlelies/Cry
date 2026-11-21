package codec

import (
	"compress/zlib"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/klauspost/compress/zstd"
	"github.com/prettylittlelies/cry/internal/etf"
	"github.com/prettylittlelies/cry/transport/websocket"
)

type Encoding string

const (
	ETF  Encoding = "etf"
	JSON Encoding = "json"
)

type Compression string

const (
	None       Compression = ""
	ZstdStream Compression = "zstd-stream"
	ZlibStream Compression = "zlib-stream"
)

type Source interface {
	Read() (websocket.Opcode, []byte, error)
}

type UnsupportedError struct {
	Encoding    Encoding
	Compression Compression
}

func (e *UnsupportedError) Error() string {
	return fmt.Sprintf("codec: unsupported encoding %q with compression %q", e.Encoding, e.Compression)
}

type messages struct {
	src Source
	cur []byte
}

func (m *messages) fill() error {
	for len(m.cur) == 0 {
		_, msg, err := m.src.Read()
		if err != nil {
			return err
		}
		m.cur = msg
	}
	return nil
}

func (m *messages) Read(p []byte) (int, error) {
	if err := m.fill(); err != nil {
		return 0, err
	}
	n := copy(p, m.cur)
	m.cur = m.cur[n:]
	return n, nil
}

func (m *messages) ReadByte() (byte, error) {
	if err := m.fill(); err != nil {
		return 0, err
	}
	b := m.cur[0]
	m.cur = m.cur[1:]
	return b, nil
}

type Reader struct {
	next func() (json.RawMessage, error)
	stop func()
}

func NewReader(src Source, enc Encoding, comp Compression) (*Reader, error) {
	msgs := &messages{src: src}
	var (
		stream io.Reader
		stop   = func() {}
	)
	switch comp {
	case None:
		return plain(src, enc)
	case ZstdStream:
		dec, err := zstd.NewReader(msgs, zstd.WithDecoderConcurrency(1), zstd.WithDecoderLowmem(true), zstd.WithDecoderMaxWindow(1<<27))
		if err != nil {
			return nil, err
		}
		stream, stop = dec, dec.Close
	case ZlibStream:
		zr, err := zlib.NewReader(msgs)
		if err != nil {
			return nil, err
		}
		stream, stop = zr, func() { zr.Close() }
	default:
		return nil, &UnsupportedError{Encoding: enc, Compression: comp}
	}
	switch enc {
	case ETF:
		d := etf.NewDecoder(&byteReader{r: stream})
		return &Reader{next: func() (json.RawMessage, error) { return d.Next() }, stop: stop}, nil
	case JSON:
		d := json.NewDecoder(stream)
		return &Reader{next: func() (json.RawMessage, error) {
			var raw json.RawMessage
			err := d.Decode(&raw)
			return raw, err
		}, stop: stop}, nil
	}
	stop()
	return nil, &UnsupportedError{Encoding: enc, Compression: comp}
}

func plain(src Source, enc Encoding) (*Reader, error) {
	switch enc {
	case ETF:
		return &Reader{next: func() (json.RawMessage, error) {
			_, msg, err := src.Read()
			if err != nil {
				return nil, err
			}
			return etf.ToJSON(msg)
		}, stop: func() {}}, nil
	case JSON:
		return &Reader{next: func() (json.RawMessage, error) {
			_, msg, err := src.Read()
			return msg, err
		}, stop: func() {}}, nil
	}
	return nil, &UnsupportedError{Encoding: enc}
}

func (r *Reader) Next() (json.RawMessage, error) {
	raw, err := r.next()
	if errors.Is(err, io.ErrUnexpectedEOF) {
		return nil, io.EOF
	}
	return raw, err
}

func (r *Reader) Close() {
	r.stop()
}

func Encode(enc Encoding, v any) (websocket.Opcode, []byte, error) {
	if enc == ETF {
		b, err := etf.Marshal(v)
		return websocket.Binary, b, err
	}
	b, err := json.Marshal(v)
	return websocket.Text, b, err
}

type byteReader struct {
	r   io.Reader
	one [1]byte
}

func (b *byteReader) Read(p []byte) (int, error) {
	return b.r.Read(p)
}

func (b *byteReader) ReadByte() (byte, error) {
	if _, err := io.ReadFull(b.r, b.one[:]); err != nil {
		return 0, err
	}
	return b.one[0], nil
}
