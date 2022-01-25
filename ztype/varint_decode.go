package ztype

import "github.com/icza/bitio"

func ReadVarint16(r *bitio.Reader) (int16, error) {
	b, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	isNegative := (b & 0x80) != 0
	v := int16(b & 0x3f)
	if (b & 0x40) != 0 {
		if b, err = r.ReadByte(); err != nil {
			return 0, err
		}
		v = (v << 8) | int16(b)
	}
	if isNegative {
		v = -v
	}
	return v, nil
}

func ReadVarint32(r *bitio.Reader) (int32, error) {
	b, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	isNegative := (b & 0x80) != 0
	v := int32(b & 0x3f)
	hasNext := (b & 0x40) != 0

	for i := 0; i < 2; i++ {
		if !hasNext {
			break
		}
		if b, err = r.ReadByte(); err != nil {
			return 0, err
		}
		v = (v << 7) | int32(b&0x7f)
		hasNext = (b & 0x80) != 0
	}
	if hasNext {
		b, err = r.ReadByte()
		if err != nil {
			return 0, err
		}
		v = (v << 8) | int32(b)
	}
	if isNegative {
		v = -v
	}
	return v, nil
}

func ReadVarint64(r *bitio.Reader) (int64, error) {
	b, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	isNegative := (b & 0x80) != 0
	v := int64(b & 0x3f)
	hasNext := (b & 0x40) != 0

	// Loop unrolling does not have any impact on performance when benchmarking
	// with Go 1.17.
	// Another unexpected find is that using TryReadByte, and only checking
	// r.TryError at the end of loop is 15% slower than checking for errors on
	// every read.
	for i := 0; i < 6; i++ {
		if !hasNext {
			break
		}
		if b, err = r.ReadByte(); err != nil {
			return 0, err
		}
		v = (v << 7) | int64(b&0x7f)
		hasNext = (b & 0x80) != 0
	}
	if hasNext {
		b, err = r.ReadByte()
		if err != nil {
			return 0, err
		}
		v = (v << 8) | int64(b)
	}
	if isNegative {
		v = -v
	}
	return v, nil
}

func ReadVarint(r *bitio.Reader) (int64, error) {
	b, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	if b == 0x80 {
		return VARINT_MIN, nil
	}

	isNegative := (b & 0x80) != 0
	v := int64(b & 0x3f)
	hasNext := (b & 0x40) != 0

	// Loop unrolling does not have any impact on performance when benchmarking
	// with Go 1.17.
	// Another unexpected find is that using TryReadByte, and only checking
	// r.TryError at the end of loop is 15% slower than checking for errors on
	// every read.
	for i := 0; i < 7; i++ {
		if !hasNext {
			break
		}
		if b, err = r.ReadByte(); err != nil {
			return 0, err
		}
		v = (v << 7) | int64(b&0x7f)
		hasNext = (b & 0x80) != 0
	}
	if hasNext {
		b, err = r.ReadByte()
		if err != nil {
			return 0, err
		}
		v = (v << 8) | int64(b)
	}
	if isNegative {
		v = -v
	}
	return v, nil
}
