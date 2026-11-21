package etf

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"math"
	"slices"
	"strconv"
)

type encoder struct {
	buf bytes.Buffer
	dec *json.Decoder
}

func Marshal(v any) ([]byte, error) {
	raw, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return FromJSON(raw)
}

func FromJSON(raw []byte) ([]byte, error) {
	e := &encoder{dec: json.NewDecoder(bytes.NewReader(raw))}
	e.dec.UseNumber()
	e.buf.WriteByte(version)
	if err := e.value(0); err != nil {
		return nil, err
	}
	return e.buf.Bytes(), nil
}

func (e *encoder) value(depth int) error {
	if depth > maxDepth {
		return ErrDepth
	}
	tok, err := e.dec.Token()
	if err != nil {
		return err
	}
	return e.token(tok, depth)
}

func (e *encoder) token(tok json.Token, depth int) error {
	switch t := tok.(type) {
	case nil:
		e.atom("nil")
	case bool:
		e.atom(strconv.FormatBool(t))
	case string:
		e.binary(t)
	case json.Number:
		e.number(t)
	case json.Delim:
		if t == '[' {
			return e.list(depth)
		}
		return e.object(depth)
	}
	return nil
}

func (e *encoder) atom(name string) {
	e.buf.WriteByte(smallAtomExt)
	e.buf.WriteByte(byte(len(name)))
	e.buf.WriteString(name)
}

func (e *encoder) binary(s string) {
	e.buf.WriteByte(binaryExt)
	e.buf.Write(binary.BigEndian.AppendUint32(nil, uint32(len(s))))
	e.buf.WriteString(s)
}

func (e *encoder) number(n json.Number) {
	if i, err := strconv.ParseInt(string(n), 10, 64); err == nil {
		e.integer(i)
		return
	}
	f, _ := n.Float64()
	e.float(f)
}

func (e *encoder) integer(i int64) {
	switch {
	case i >= 0 && i <= 255:
		e.buf.WriteByte(smallIntegerExt)
		e.buf.WriteByte(byte(i))
	case i >= math.MinInt32 && i <= math.MaxInt32:
		e.buf.WriteByte(integerExt)
		e.buf.Write(binary.BigEndian.AppendUint32(nil, uint32(int32(i))))
	case i > math.MaxInt32 && i <= math.MaxUint32:
		e.buf.WriteByte(smallBigExt)
		digits := binary.LittleEndian.AppendUint64(nil, uint64(i))
		n := len(digits)
		for n > 0 && digits[n-1] == 0 {
			n--
		}
		e.buf.WriteByte(byte(n))
		e.buf.WriteByte(0)
		e.buf.Write(digits[:n])
	default:
		e.float(float64(i))
	}
}

func (e *encoder) float(f float64) {
	e.buf.WriteByte(newFloatExt)
	e.buf.Write(binary.BigEndian.AppendUint64(nil, math.Float64bits(f)))
}

func (e *encoder) list(depth int) error {
	start := e.buf.Len()
	e.buf.WriteByte(listExt)
	e.buf.Write([]byte{0, 0, 0, 0})
	var n uint32
	for e.dec.More() {
		if err := e.value(depth + 1); err != nil {
			return err
		}
		n++
	}
	if _, err := e.dec.Token(); err != nil {
		return err
	}
	if n == 0 {
		e.buf.Truncate(start)
		e.buf.WriteByte(nilExt)
		return nil
	}
	binary.BigEndian.PutUint32(e.buf.Bytes()[start+1:], n)
	e.buf.WriteByte(nilExt)
	return nil
}

type member struct {
	key   string
	index int64
	value []byte
}

func arrayIndex(s string) (int64, bool) {
	if s == "" || (len(s) > 1 && s[0] == '0') {
		return 0, false
	}
	n, err := strconv.ParseUint(s, 10, 32)
	if err != nil || n >= math.MaxUint32 {
		return 0, false
	}
	return int64(n), true
}

func (e *encoder) object(depth int) error {
	var indexed, named []member
	for e.dec.More() {
		tok, err := e.dec.Token()
		if err != nil {
			return err
		}
		key := tok.(string)
		sub := &encoder{dec: e.dec}
		if err := sub.value(depth + 1); err != nil {
			return err
		}
		m := member{key: key, value: sub.buf.Bytes()}
		if i, ok := arrayIndex(key); ok {
			m.index = i
			indexed = append(indexed, m)
		} else {
			named = append(named, m)
		}
	}
	if _, err := e.dec.Token(); err != nil {
		return err
	}
	slices.SortFunc(indexed, func(a, b member) int { return int(a.index - b.index) })
	e.buf.WriteByte(mapExt)
	e.buf.Write(binary.BigEndian.AppendUint32(nil, uint32(len(indexed)+len(named))))
	for _, m := range indexed {
		e.integer(m.index)
		e.buf.Write(m.value)
	}
	for _, m := range named {
		e.binary(m.key)
		e.buf.Write(m.value)
	}
	return nil
}
