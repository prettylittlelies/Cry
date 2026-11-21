package dave

import "errors"

var ErrLEB128 = errors.New("dave: malformed uleb128")

func appendULEB128(b []byte, v uint64) []byte {
	for v >= 0x80 {
		b = append(b, byte(v)|0x80)
		v >>= 7
	}
	return append(b, byte(v))
}

func readULEB128(b []byte) (uint64, int, error) {
	var v uint64
	for i, c := range b {
		if i == 10 || i == 9 && c > 1 {
			return 0, 0, ErrLEB128
		}
		v |= uint64(c&0x7f) << (7 * i)
		if c < 0x80 {
			return v, i + 1, nil
		}
	}
	return 0, 0, ErrLEB128
}
