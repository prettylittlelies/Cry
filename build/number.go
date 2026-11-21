package build

import "strconv"

type Number uint32

func (n Number) String() string {
	return strconv.FormatUint(uint64(n), 10)
}

type NativeNumber uint32

func (n NativeNumber) String() string {
	return strconv.FormatUint(uint64(n), 10)
}
