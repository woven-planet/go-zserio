package ztype

import "errors"

type BitWriter interface {
	WriteBits(r uint64, n uint8) error
}

// ErrValueOutOfBitRange is returned when you try to write a bit field type
// with a value that is larger than the number of bits allows for.
var ErrValueOutOfBitRange = errors.New("value is out of range for the given number of bits")

// WriteUnsignedBits writes an unsigned number with a given number of bits
func WriteUnsignedBits(w BitWriter, value uint64, bits uint8) error {
	if value > ((1 << bits) - 1) {
		return ErrValueOutOfBitRange
	}
	return w.WriteBits(value, bits)
}

// WriteSignedBits writes an signed number with a given number of bits
func WriteSignedBits(w BitWriter, value int64, bits uint8) error {
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

// WriteUint8 writes an unsigned 8-bit integer.
func WriteUint8(w BitWriter, value uint8) error {
	return WriteUnsignedBits(w, uint64(value), 8)
}

// WriteUint16 writes an unsigned 16-bit integer.
func WriteUint16(w BitWriter, value uint16) error {
	return WriteUnsignedBits(w, uint64(value), 16)
}

// WriteUint32 writes an unsigned 32-bit integer.
func WriteUint32(w BitWriter, value uint32) error {
	return WriteUnsignedBits(w, uint64(value), 32)
}

// WriteUint64 writes an unsigned 64-bit integer.
func WriteUint64(w BitWriter, value uint64) error {
	return WriteUnsignedBits(w, uint64(value), 64)
}

// WriteInt8 writes a signed 8-bit integer.
func WriteInt8(w BitWriter, value int8) error {
	return WriteSignedBits(w, int64(value), 8)
}

// WriteInt16 writes a signed 16-bit integer.
func WriteInt16(w BitWriter, value int16) error {
	return WriteSignedBits(w, int64(value), 16)
}

// WriteInt32 writes a signed 32-bit integer.
func WriteInt32(w BitWriter, value int32) error {
	return WriteSignedBits(w, int64(value), 32)
}

// WriteInt64 writes a signed 64-bit integer.
func WriteInt64(w BitWriter, value int64) error {
	return WriteSignedBits(w, int64(value), 64)
}
