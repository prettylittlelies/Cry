package voice

import (
	"bufio"
	"bytes"
	"errors"
	"io"
)

var ErrNotOgg = errors.New("voice: not an ogg stream")

type OggOpus struct {
	r       *bufio.Reader
	pending []byte
	queue   [][]byte
	headers int
}

func NewOggOpus(r io.Reader) *OggOpus {
	return &OggOpus{r: bufio.NewReader(r)}
}

func (o *OggOpus) page() error {
	var head [27]byte
	if _, err := io.ReadFull(o.r, head[:]); err != nil {
		return err
	}
	if !bytes.Equal(head[:4], []byte("OggS")) {
		return ErrNotOgg
	}
	table := make([]byte, head[26])
	if _, err := io.ReadFull(o.r, table); err != nil {
		return err
	}
	total := 0
	for _, l := range table {
		total += int(l)
	}
	body := make([]byte, total)
	if _, err := io.ReadFull(o.r, body); err != nil {
		return err
	}
	for _, l := range table {
		o.pending = append(o.pending, body[:l]...)
		body = body[l:]
		if l < 255 {
			o.queue = append(o.queue, o.pending)
			o.pending = nil
		}
	}
	return nil
}

func (o *OggOpus) Next() ([]byte, error) {
	for {
		for len(o.queue) > 0 {
			p := o.queue[0]
			o.queue = o.queue[1:]
			if o.headers < 2 && (bytes.HasPrefix(p, []byte("OpusHead")) || bytes.HasPrefix(p, []byte("OpusTags"))) {
				o.headers++
				continue
			}
			return p, nil
		}
		if err := o.page(); err != nil {
			return nil, err
		}
	}
}
