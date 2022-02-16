package ztype

import zserio "github.com/woven-planet/go-zserio"

// ReadUnsignedBits reads an unsigned number with a given number of bits
func ReadUnsignedBits(r *zserio.Reader, bits uint8) (uint64, error) {
	return r.ReadBits(bits)
}

// ReadSignedBits reads an signed number with a given number of bits
func ReadSignedBits(r *zserio.Reader, bits uint8) (int64, error) {
	if bits > 64 {
		panic("can not read more than 64 bits")
	}

	u, err := r.ReadBits(bits)
	if err != nil {
		return 0, err
	}

	if (u & (1 << (bits - 1))) != 0 {
		return int64(u) - (1 << bits), nil
	}
	return int64(u), nil
}

// ReadUint8 reads an unsigned 8-bit integer.
func ReadUint8(r *zserio.Reader) (uint8, error) {
	value, err := ReadUnsignedBits(r, 8)
	return uint8(value), err
}

// ReadUint16 reads an unsigned 16-bit integer.
func ReadUint16(r *zserio.Reader) (uint16, error) {
	value, err := ReadUnsignedBits(r, 16)
	return uint16(value), err
}

// ReadUint32 reads an unsigned 32-bit integer.
func ReadUint32(r *zserio.Reader) (uint32, error) {
	value, err := ReadUnsignedBits(r, 32)
	return uint32(value), err
}

// ReadUint64 reads an unsigned 64-bit integer.
func ReadUint64(r *zserio.Reader) (uint64, error) {
	value, err := ReadUnsignedBits(r, 64)
	return uint64(value), err
}

// ReadInt8 reads a signed 8-bit integer.
func ReadInt8(r *zserio.Reader) (int8, error) {
	value, err := ReadSignedBits(r, 8)
	return int8(value), err
}

// ReadInt16 reads a signed 16-bit integer.
func ReadInt16(r *zserio.Reader) (int16, error) {
	value, err := ReadSignedBits(r, 16)
	return int16(value), err
}

// ReadInt32 reads a signed 32-bit integer.
func ReadInt32(r *zserio.Reader) (int32, error) {
	value, err := ReadSignedBits(r, 32)
	return int32(value), err
}

// ReadInt64 reads a signed 64-bit integer.
func ReadInt64(r *zserio.Reader) (int64, error) {
	value, err := ReadSignedBits(r, 64)
	return int64(value), err
}
