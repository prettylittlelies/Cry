package fingerprint

import (
	"strings"
	"unicode/utf16"

	"github.com/prettylittlelies/cry/internal/rapidhash"
)

const (
	blinkMinimumTableSize = 8
	blinkMaxLoad          = 2
	blinkFlagBits         = 8
)

func blinkHash(name string) uint32 {
	units := utf16.Encode([]rune(strings.ToLower(name)))
	buf := make([]byte, 2*len(units))
	for i, u := range units {
		buf[2*i], buf[2*i+1] = byte(u), byte(u>>8)
	}
	h := uint32(rapidhash.Hash(buf) & (1<<(32-blinkFlagBits) - 1))
	if h == 0 {
		h = 0x80000000 >> blinkFlagBits
	}
	return h
}

type blinkTable struct {
	slots []string
	keys  int
}

func (t *blinkTable) place(name string) {
	mask := uint32(len(t.slots) - 1)
	i := blinkHash(name) & mask
	for probe := uint32(1); t.slots[i] != ""; probe++ {
		i = (i + probe) & mask
	}
	t.slots[i] = name
}

func (t *blinkTable) insert(name string) {
	if t.slots == nil {
		t.slots = make([]string, blinkMinimumTableSize)
	}
	t.place(name)
	t.keys++
	if t.keys*blinkMaxLoad < len(t.slots) {
		return
	}
	old := t.slots
	t.slots = make([]string, 2*len(old))
	for _, n := range old {
		if n != "" {
			t.place(n)
		}
	}
}

func (t *blinkTable) order() []string {
	out := make([]string, 0, t.keys)
	for _, n := range t.slots {
		if n != "" {
			out = append(out, n)
		}
	}
	return out
}

func BlinkOrder(names ...[]string) []string {
	var t blinkTable
	for _, group := range names {
		for _, n := range group {
			t.insert(n)
		}
	}
	return t.order()
}
