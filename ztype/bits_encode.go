package ztype

import (
	"errors"

	"github.com/icza/bitio"
)

var ErrValueOutOfBitRange = errors.New("value is out of range for the given number of bits")

// WriteUnsignedBits writes an unsigned number with a given number of bits
func WriteUnsignedBits(w *bitio.Writer, value uint64, bits uint8) error {
	if value > ((1 << bits) - 1) {
		return ErrValueOutOfBitRange
	}
	return w.WriteBits(value, bits)
}

// WriteSignedBits writes an signed number with a given number of bits
func WriteSignedBits(w *bitio.Writer, value int64, bits uint8) error {
	if bits > 64 {
		panic("can not read more than 64 bits")
	}

	var min int64 = -(1 << (bits - 1))
	var max int64 = (1 << (bits - 1)) - 1
	if value < min || value > max {
		return ErrValueOutOfBitRange
	}

	if value >= 0 {
		return w.WriteBits(uint64(value), bits)
	}
	var v uint64 = (1 << bits) | uint64(value+(1<<bits))
	return w.WriteBits(v, bits)
}
