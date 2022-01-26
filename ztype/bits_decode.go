package ztype

import "github.com/icza/bitio"

// ReadUnsignedBits reads an unsigned number with a given number of bits
func ReadUnsignedBits(r *bitio.CountReader, bits uint8) (uint64, error) {
	return r.ReadBits(bits)
}

// ReadSignedBits reads an signed number with a given number of bits
func ReadSignedBits(r *bitio.CountReader, bits uint8) (int64, error) {
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
