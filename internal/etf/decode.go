package etf

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"encoding/json"
	"io"
	"math"
	"math/big"
	"strconv"
)

type source interface {
	io.Reader
	io.ByteReader
}

type Decoder struct {
	r   source
	buf bytes.Buffer
}

func NewDecoder(r io.Reader) *Decoder {
	if s, ok := r.(source); ok {
		return &Decoder{r: s}
	}
	return &Decoder{r: bufio.NewReader(r)}
}

func ToJSON(b []byte) ([]byte, error) {
	return NewDecoder(bytes.NewReader(b)).Next()
}

func Unmarshal(b []byte, v any) error {
	raw, err := ToJSON(b)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, v)
}

func (d *Decoder) Next() ([]byte, error) {
	v, err := d.r.ReadByte()
	if err != nil {
		return nil, err
	}
	if v != version {
		return nil, ErrVersion
	}
	d.buf.Reset()
	if err := d.term(d.r, 0); err != nil {
		return nil, err
	}
	return bytes.Clone(d.buf.Bytes()), nil
}

func (d *Decoder) Decode(v any) error {
	raw, err := d.Next()
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, v)
}

func read(r source, n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := io.ReadFull(r, b)
	return b, err
}

func u16(r source) (int, error) {
	b, err := read(r, 2)
	if err != nil {
		return 0, err
	}
	return int(binary.BigEndian.Uint16(b)), nil
}

func u32(r source) (int, error) {
	b, err := read(r, 4)
	if err != nil {
		return 0, err
	}
	return int(binary.BigEndian.Uint32(b)), nil
}

func (d *Decoder) str(s []byte) {
	raw, _ := json.Marshal(string(s))
	d.buf.Write(raw)
}

func (d *Decoder) bytesList(b []byte) {
	d.buf.WriteByte('[')
	for i, c := range b {
		if i > 0 {
			d.buf.WriteByte(',')
		}
		d.buf.WriteString(strconv.Itoa(int(c)))
	}
	d.buf.WriteByte(']')
}

func (d *Decoder) atom(name []byte) {
	switch string(name) {
	case "nil", "null":
		d.buf.WriteString("null")
	case "true", "false":
		d.buf.Write(name)
	default:
		d.str(name)
	}
}

func (d *Decoder) term(r source, depth int) error {
	if depth > maxDepth {
		return ErrDepth
	}
	tag, err := r.ReadByte()
	if err != nil {
		return err
	}
	switch tag {
	case smallIntegerExt:
		b, err := r.ReadByte()
		if err != nil {
			return err
		}
		d.buf.WriteString(strconv.Itoa(int(b)))
	case integerExt:
		b, err := read(r, 4)
		if err != nil {
			return err
		}
		d.buf.WriteString(strconv.Itoa(int(int32(binary.BigEndian.Uint32(b)))))
	case newFloatExt:
		b, err := read(r, 8)
		if err != nil {
			return err
		}
		d.buf.WriteString(strconv.FormatFloat(math.Float64frombits(binary.BigEndian.Uint64(b)), 'g', -1, 64))
	case floatExt:
		b, err := read(r, 31)
		if err != nil {
			return err
		}
		f, err := strconv.ParseFloat(string(bytes.TrimRight(b, "\x00")), 64)
		if err != nil {
			return err
		}
		d.buf.WriteString(strconv.FormatFloat(f, 'g', -1, 64))
	case smallAtomExt, smallAtomUTF8:
		n, err := r.ReadByte()
		if err != nil {
			return err
		}
		name, err := read(r, int(n))
		if err != nil {
			return err
		}
		d.atom(name)
	case atomExt, atomUTF8Ext:
		n, err := u16(r)
		if err != nil {
			return err
		}
		name, err := read(r, n)
		if err != nil {
			return err
		}
		d.atom(name)
	case binaryExt:
		n, err := u32(r)
		if err != nil {
			return err
		}
		b, err := read(r, n)
		if err != nil {
			return err
		}
		d.str(b)
	case stringExt:
		n, err := u16(r)
		if err != nil {
			return err
		}
		b, err := read(r, n)
		if err != nil {
			return err
		}
		d.bytesList(b)
	case nilExt:
		d.buf.WriteString("[]")
	case listExt:
		n, err := u32(r)
		if err != nil {
			return err
		}
		if err := d.sequence(r, n, depth); err != nil {
			return err
		}
		return d.tail(r, depth)
	case smallTupleExt:
		n, err := r.ReadByte()
		if err != nil {
			return err
		}
		return d.sequence(r, int(n), depth)
	case largeTupleExt:
		n, err := u32(r)
		if err != nil {
			return err
		}
		return d.sequence(r, n, depth)
	case mapExt:
		n, err := u32(r)
		if err != nil {
			return err
		}
		return d.object(r, n, depth)
	case smallBigExt:
		n, err := r.ReadByte()
		if err != nil {
			return err
		}
		return d.bignum(r, int(n))
	case largeBigExt:
		n, err := u32(r)
		if err != nil {
			return err
		}
		return d.bignum(r, n)
	case compressedExt:
		return d.compressed(r, depth)
	default:
		return &UnsupportedTagError{Tag: tag}
	}
	return nil
}

func (d *Decoder) sequence(r source, n, depth int) error {
	d.buf.WriteByte('[')
	for i := range n {
		if i > 0 {
			d.buf.WriteByte(',')
		}
		if err := d.term(r, depth+1); err != nil {
			return err
		}
	}
	d.buf.WriteByte(']')
	return nil
}

func (d *Decoder) tail(r source, depth int) error {
	b, err := r.ReadByte()
	if err != nil {
		return err
	}
	if b == nilExt {
		return nil
	}
	var discard Decoder
	if err := discard.term(&prefixed{first: b, r: r}, depth+1); err != nil {
		return err
	}
	return nil
}

func (d *Decoder) object(r source, n, depth int) error {
	d.buf.WriteByte('{')
	for i := range n {
		if i > 0 {
			d.buf.WriteByte(',')
		}
		var key Decoder
		if err := key.term(r, depth+1); err != nil {
			return err
		}
		k := key.buf.Bytes()
		if len(k) == 0 || k[0] != '"' {
			d.str(k)
		} else {
			d.buf.Write(k)
		}
		d.buf.WriteByte(':')
		if err := d.term(r, depth+1); err != nil {
			return err
		}
	}
	d.buf.WriteByte('}')
	return nil
}

var maxSafeInteger = big.NewInt(1<<53 - 1)

func (d *Decoder) bignum(r source, n int) error {
	sign, err := r.ReadByte()
	if err != nil {
		return err
	}
	digits, err := read(r, n)
	if err != nil {
		return err
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	v := new(big.Int).SetBytes(digits)
	if sign != 0 {
		v.Neg(v)
	}
	if v.CmpAbs(maxSafeInteger) <= 0 {
		d.buf.WriteString(v.String())
		return nil
	}
	d.str([]byte(v.String()))
	return nil
}

func (d *Decoder) compressed(r source, depth int) error {
	size, err := u32(r)
	if err != nil {
		return err
	}
	zr, err := zlib.NewReader(r)
	if err != nil {
		return err
	}
	defer zr.Close()
	inner, err := read(bufio.NewReader(io.LimitReader(zr, int64(size))), size)
	if err != nil {
		return err
	}
	return d.term(bytes.NewReader(inner), depth+1)
}

type prefixed struct {
	first byte
	used  bool
	r     source
}

func (p *prefixed) ReadByte() (byte, error) {
	if !p.used {
		p.used = true
		return p.first, nil
	}
	return p.r.ReadByte()
}

func (p *prefixed) Read(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	if !p.used {
		p.used = true
		b[0] = p.first
		return 1, nil
	}
	return p.r.Read(b)
}
