package rapidhash

import (
	"encoding/binary"
	"math/bits"
)

const Seed uint64 = 0xbdd89aa982704029

var secret = [3]uint64{0x2d358dccaa6c78a5, 0x8bb84b93962eacc9, 0x4b33a62ed433d4a3}

func mix(a, b uint64) uint64 {
	hi, lo := bits.Mul64(a, b)
	return hi ^ lo
}

func r32(p []byte, o int) uint64 {
	return uint64(binary.LittleEndian.Uint32(p[o:]))
}

func r64(p []byte, o int) uint64 {
	return binary.LittleEndian.Uint64(p[o:])
}

func Hash(p []byte) uint64 {
	return HashSeed(p, Seed)
}

func HashSeed(p []byte, seed uint64) uint64 {
	n := len(p)
	seed ^= mix(seed^secret[0], secret[1]) ^ uint64(n)
	var a, b uint64
	switch {
	case n >= 4 && n <= 16:
		last := n - 4
		a = r32(p, 0)<<32 | r32(p, last)
		d := (n & 24) >> (n >> 3)
		b = r32(p, d)<<32 | r32(p, last-d)
	case n > 0 && n < 4:
		a = uint64(p[0])<<56 | uint64(p[n>>1])<<32 | uint64(p[n-1])
	case n > 16:
		i, o := n, 0
		if i > 48 {
			s1, s2 := seed, seed
			for i >= 48 {
				seed = mix(r64(p, o)^secret[0], r64(p, o+8)^seed)
				s1 = mix(r64(p, o+16)^secret[1], r64(p, o+24)^s1)
				s2 = mix(r64(p, o+32)^secret[2], r64(p, o+40)^s2)
				o += 48
				i -= 48
			}
			seed ^= s1 ^ s2
		}
		if i > 16 {
			seed = mix(r64(p, o)^secret[2], r64(p, o+8)^seed^secret[1])
			if i > 32 {
				seed = mix(r64(p, o+16)^secret[2], r64(p, o+24)^seed)
			}
		}
		a, b = r64(p, o+i-16), r64(p, o+i-8)
	}
	a ^= secret[1]
	b ^= seed
	hi, lo := bits.Mul64(a, b)
	return mix(lo^secret[0]^uint64(n), hi^secret[1])
}
