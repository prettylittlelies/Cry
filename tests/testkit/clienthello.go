package testkit

import (
	"encoding/binary"
	"errors"
	"slices"
)

const Grease uint16 = 0x0a0a

type Hello struct {
	Ciphers    []uint16
	Extensions []uint16
	Data       map[uint16][]byte
}

func IsGrease(v uint16) bool {
	return v&0x0f0f == 0x0a0a && v>>8 == v&0xff
}

func normalize(v uint16) uint16 {
	if IsGrease(v) {
		return Grease
	}
	return v
}

var errShort = errors.New("testkit: truncated client hello")

type cursor struct {
	b []byte
	p int
}

func (c *cursor) take(n int) ([]byte, error) {
	if c.p+n > len(c.b) {
		return nil, errShort
	}
	out := c.b[c.p : c.p+n]
	c.p += n
	return out, nil
}

func (c *cursor) u8() (int, error) {
	b, err := c.take(1)
	if err != nil {
		return 0, err
	}
	return int(b[0]), nil
}

func (c *cursor) u16() (int, error) {
	b, err := c.take(2)
	if err != nil {
		return 0, err
	}
	return int(binary.BigEndian.Uint16(b)), nil
}

func ParseHello(record []byte) (*Hello, error) {
	c := &cursor{b: record, p: 5 + 4 + 2 + 32}
	n, err := c.u8()
	if err != nil {
		return nil, err
	}
	if _, err := c.take(n); err != nil {
		return nil, err
	}
	h := &Hello{Data: map[uint16][]byte{}}
	if n, err = c.u16(); err != nil {
		return nil, err
	}
	ciphers, err := c.take(n)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(ciphers); i += 2 {
		h.Ciphers = append(h.Ciphers, normalize(binary.BigEndian.Uint16(ciphers[i:])))
	}
	if n, err = c.u8(); err != nil {
		return nil, err
	}
	if _, err := c.take(n); err != nil {
		return nil, err
	}
	if n, err = c.u16(); err != nil {
		return nil, err
	}
	end := c.p + n
	for c.p < end {
		t, err := c.u16()
		if err != nil {
			return nil, err
		}
		l, err := c.u16()
		if err != nil {
			return nil, err
		}
		d, err := c.take(l)
		if err != nil {
			return nil, err
		}
		id := normalize(uint16(t))
		h.Extensions = append(h.Extensions, id)
		h.Data[id] = normalizeList(id, d)
	}
	return h, nil
}

func normalizeList(id uint16, d []byte) []byte {
	out := slices.Clone(d)
	start := map[uint16]int{10: 2, 13: 2, 43: 1}
	s, ok := start[id]
	if !ok {
		return out
	}
	for i := s; i+1 < len(out); i += 2 {
		v := normalize(binary.BigEndian.Uint16(out[i:]))
		binary.BigEndian.PutUint16(out[i:], v)
	}
	return out
}

func (h *Hello) ExtensionSet() []uint16 {
	set := slices.Clone(h.Extensions)
	slices.Sort(set)
	return set
}
