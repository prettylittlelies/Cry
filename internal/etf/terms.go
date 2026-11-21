package etf

import (
	"errors"
	"fmt"
)

const (
	version         = 131
	newFloatExt     = 'F'
	compressedExt   = 'P'
	smallIntegerExt = 'a'
	integerExt      = 'b'
	floatExt        = 'c'
	atomExt         = 'd'
	smallTupleExt   = 'h'
	largeTupleExt   = 'i'
	nilExt          = 'j'
	stringExt       = 'k'
	listExt         = 'l'
	binaryExt       = 'm'
	smallBigExt     = 'n'
	largeBigExt     = 'o'
	smallAtomExt    = 's'
	mapExt          = 't'
	atomUTF8Ext     = 'v'
	smallAtomUTF8   = 'w'
)

const maxDepth = 256

var (
	ErrVersion = errors.New("etf: missing format version")
	ErrDepth   = errors.New("etf: nesting limit reached")
)

type UnsupportedTagError struct {
	Tag byte
}

func (e *UnsupportedTagError) Error() string {
	return fmt.Sprintf("etf: unsupported term tag %d", e.Tag)
}
