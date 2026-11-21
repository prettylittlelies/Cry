package mls

import (
	"encoding/binary"
	"errors"
	"fmt"
)

var (
	ErrTruncated     = errors.New("mls: truncated input")
	ErrTrailingData  = errors.New("mls: trailing data")
	ErrVarintTooLong = errors.New("mls: varint exceeds 30 bits")
	ErrVarintPadding = errors.New("mls: non-minimal varint")
)

type encoder interface {
	encode(w *Writer)
}

type decoder interface {
	decode(r *Reader)
}

type Writer struct {
	buf []byte
}

func (w *Writer) Bytes() []byte { return w.buf }

func (w *Writer) U8(v uint8)   { w.buf = append(w.buf, v) }
func (w *Writer) U16(v uint16) { w.buf = binary.BigEndian.AppendUint16(w.buf, v) }
func (w *Writer) U32(v uint32) { w.buf = binary.BigEndian.AppendUint32(w.buf, v) }
func (w *Writer) U64(v uint64) { w.buf = binary.BigEndian.AppendUint64(w.buf, v) }
func (w *Writer) Raw(b []byte) { w.buf = append(w.buf, b...) }

func (w *Writer) Bool(v bool) {
	if v {
		w.U8(1)
		return
	}
	w.U8(0)
}

func (w *Writer) Varint(n int) {
	switch {
	case n < 1<<6:
		w.U8(uint8(n))
	case n < 1<<14:
		w.U16(uint16(n) | 0x4000)
	case n < 1<<30:
		w.U32(uint32(n) | 0x80000000)
	default:
		panic(ErrVarintTooLong)
	}
}

func (w *Writer) Opaque(b []byte) {
	w.Varint(len(b))
	w.Raw(b)
}

func (w *Writer) Vector(body func(*Writer)) {
	var inner Writer
	body(&inner)
	w.Opaque(inner.buf)
}

func (w *Writer) Encode(v encoder) { v.encode(w) }

func (w *Writer) Optional(present bool, body func(*Writer)) {
	w.Bool(present)
	if present {
		body(w)
	}
}

func writeList[T encoder](w *Writer, items []T) {
	w.Vector(func(w *Writer) {
		for _, it := range items {
			it.encode(w)
		}
	})
}

type Reader struct {
	b   []byte
	err error
}

func NewReader(b []byte) *Reader { return &Reader{b: b} }

func (r *Reader) Err() error { return r.err }

func (r *Reader) Fail(err error) {
	if r.err == nil {
		r.err = err
	}
}

func (r *Reader) Empty() bool { return len(r.b) == 0 }

func (r *Reader) Remaining() []byte { return r.b }

func (r *Reader) take(n int) []byte {
	if r.err != nil {
		return nil
	}
	if n < 0 || n > len(r.b) {
		r.Fail(ErrTruncated)
		return nil
	}
	out := r.b[:n:n]
	r.b = r.b[n:]
	return out
}

func (r *Reader) U8() uint8 {
	if b := r.take(1); b != nil {
		return b[0]
	}
	return 0
}

func (r *Reader) U16() uint16 {
	if b := r.take(2); b != nil {
		return binary.BigEndian.Uint16(b)
	}
	return 0
}

func (r *Reader) U32() uint32 {
	if b := r.take(4); b != nil {
		return binary.BigEndian.Uint32(b)
	}
	return 0
}

func (r *Reader) U64() uint64 {
	if b := r.take(8); b != nil {
		return binary.BigEndian.Uint64(b)
	}
	return 0
}

func (r *Reader) Bool() bool {
	switch v := r.U8(); v {
	case 0:
		return false
	case 1:
		return true
	default:
		r.Fail(fmt.Errorf("mls: invalid boolean %d", v))
		return false
	}
}

func (r *Reader) Varint() int {
	first := r.U8()
	if r.err != nil {
		return 0
	}
	prefix := first >> 6
	n := int(first & 0x3f)
	switch prefix {
	case 0:
		return n
	case 1:
		n = n<<8 | int(r.U8())
		if n < 1<<6 {
			r.Fail(ErrVarintPadding)
		}
	case 2:
		for range 3 {
			n = n<<8 | int(r.U8())
		}
		if n < 1<<14 {
			r.Fail(ErrVarintPadding)
		}
	default:
		r.Fail(ErrVarintTooLong)
	}
	return n
}

func (r *Reader) Raw(n int) []byte {
	return r.take(n)
}

func (r *Reader) Opaque() []byte {
	return r.take(r.Varint())
}

func (r *Reader) Vector(body func(*Reader)) {
	inner := &Reader{b: r.Opaque(), err: r.err}
	for inner.err == nil && !inner.Empty() {
		before := len(inner.b)
		body(inner)
		if inner.err == nil && len(inner.b) == before {
			inner.Fail(fmt.Errorf("mls: vector element consumed nothing"))
		}
	}
	r.Fail(inner.err)
}

func (r *Reader) Decode(v decoder) { v.decode(r) }

func (r *Reader) Optional(body func(*Reader)) bool {
	present := r.Bool()
	if present {
		body(r)
	}
	return present
}

func readList[T any, P interface {
	*T
	decoder
}](r *Reader) []T {
	var out []T
	r.Vector(func(r *Reader) {
		var v T
		P(&v).decode(r)
		out = append(out, v)
	})
	return out
}

func Marshal(v encoder) []byte {
	var w Writer
	v.encode(&w)
	return w.buf
}

func Unmarshal(b []byte, v decoder) error {
	r := NewReader(b)
	v.decode(r)
	if r.err != nil {
		return r.err
	}
	if !r.Empty() {
		return ErrTrailingData
	}
	return nil
}
