package fingerprint

import "math/rand/v2"

func grease() uint16 {
	n := uint16(rand.N(16))<<4 | 0x0a
	return n<<8 | n
}
