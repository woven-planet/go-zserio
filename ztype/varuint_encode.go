package ztype

import (
	"errors"
	"math/bits"

	zserio "github.com/woven-planet/go-zserio"
)

// ErrOutOfBounds is an error that is returned when you try to write a
// value that is too large for the zserio (integer) type.
var ErrOutOfBounds = errors.New("value too large or too small for type")

// WriteVaruint16 writes a zserio varuint16 value to the bitstream.
// If you pass in a value that is outside to allowed range of
// (MinVaruint16, MaxVaruint16) ErrOutOfBounds will be returned.
func WriteVaruint16(w zserio.Writer, v uint16) error {
	return writeVarUint(w, uint64(v), 2)
}

// WriteVaruint32 writes a zserio varuint32 value to the bitstream.
// If you pass in a value that is outside to allowed range of
// (MinVaruint32, MaxVaruint32) ErrOutOfBounds will be returned.
func WriteVaruint32(w zserio.Writer, v uint32) error {
	return writeVarUint(w, uint64(v), 4)
}

// WriteVaruint64 writes a zserio varuint64 value to the bitstream.
// If you pass in a value that is outside to allowed range of
// (MinVaruint64, MaxVaruint64) ErrOutOfBounds will be returned.
func WriteVaruint64(w zserio.Writer, v uint64) error {
	return writeVarUint(w, uint64(v), 8)
}

// WriteVaruint writes a zserio varuint value to the bitstream.
// If you pass in a value that is outside to allowed range of
// (MinVaruint, MaxVaruint) ErrOutOfBounds will be returned.
func WriteVaruint(w zserio.Writer, v uint64) error {
	return writeVarUint(w, uint64(v), 9)
}

// WriteVarsize writes a zserio varsize value to the bitstream.
// If you pass in a value that is outside to allowed range of
// (MinVarsize, MaxVarsize) ErrOutOfBounds will be returned.
func WriteVarsize(w zserio.Writer, v uint64) error {
	if v > MaxVarsize {
		return ErrOutOfBounds
	}
	return writeVarUint(w, uint64(v), 5)
}

func writeVarUint(w zserio.Writer, v uint64, maxBytes int) error {
	neededBytes, err := UnsignedBitSize(v, maxBytes)
	if err != nil {
		return err
	}
	neededBytes /= 8
	needsMaxByteRange := neededBytes == maxBytes
	for i := 0; i < neededBytes; i++ {
		var b uint64 = 0x00
		remainingBits := 8
		hasNextByte := i < (neededBytes - 1)
		if hasNextByte {
			remainingBits--
			b |= 0x80
		} else if !needsMaxByteRange {
			remainingBits--
		}
		shiftBits := (neededBytes - (i + 1)) * 7
		if needsMaxByteRange && hasNextByte {
			shiftBits++
		}
		b |= (v >> uint64(shiftBits)) & (0xff >> (8 - remainingBits))
		if err == nil {
			err = w.WriteBitsUnsafe(b, 8)
		}
	}
	return err
}

// UnsignedBitSize returns the size in bits of the zserio encoding of an unsigned
// value.
func UnsignedBitSize(v uint64, maxBytes int) (int, error) {
	var max uint64 = (1 << 7) - 1
	// Unlike the Python zserio version this version is generic and does
	// not need a per-type table, but still gets identical performance.
	//
	//   goos: darwin
	//   goarch: amd64
	//   pkg: github.com/woven-planet/go-zserio
	//   cpu: Intel(R) Core(TM) i7-6567U CPU @ 3.30GHz
	//   BenchmarkBitsize-4              322097268                3.659 ns/op
	//   BenchmarkBitsizeTable-4         324647538                3.642 ns/op
	bytes := 1
	for {
		if v <= max {
			return bytes * 8, nil
		}
		bytes++
		if bytes > maxBytes {
			return 0, ErrOutOfBounds
		}
		if bytes == maxBytes {
			// The last byte can take 8 bits, since it does not need a
			// more-flag in the msb.
			max = (max << 8) | 0xff
		} else {
			max = (max << 7) | 0xff
		}
	}
}

// NumBits returns the exact number of bits needed to represent an unsigned
// integer in zserio, considering the fact that 0 is not counted, due to the
// zserio encoding.
func NumBits(v uint64) int {
	if v == 0 {
		return 0
	}
	if v == 1 {
		return 1
	}
	// note that we subtract -1 here, because we will never encode a 0
	return bits.Len64(v - 1)
}
