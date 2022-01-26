package ztype

import "github.com/icza/bitio"

func ReadVaruint16(r *bitio.CountReader) (uint16, error) {
	b, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	v := uint16(b & 0x7f)
	if (b & 0x80) != 0 {
		if b, err = r.ReadByte(); err != nil {
			return 0, err
		}
		v = (v << 8) | uint16(b)
	}
	return v, nil
}

func ReadVaruint32(r *bitio.CountReader) (uint32, error) {
	b, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	v := uint32(b & 0x7f)
	hasNext := (b & 0x80) != 0

	for i := 0; i < 2; i++ {
		if !hasNext {
			break
		}
		if b, err = r.ReadByte(); err != nil {
			return 0, err
		}
		v = (v << 7) | uint32(b&0x7f)
		hasNext = (b & 0x80) != 0
	}
	if hasNext {
		b, err = r.ReadByte()
		if err != nil {
			return 0, err
		}
		v = (v << 8) | uint32(b)
	}
	return v, nil
}

func ReadVaruint64(r *bitio.CountReader) (uint64, error) {
	b, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	v := uint64(b & 0x7f)
	hasNext := (b & 0x80) != 0

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
		v = (v << 7) | uint64(b&0x7f)
		hasNext = (b & 0x80) != 0
	}
	if hasNext {
		b, err = r.ReadByte()
		if err != nil {
			return 0, err
		}
		v = (v << 8) | uint64(b)
	}
	return v, nil
}

func ReadVaruint(r *bitio.CountReader) (uint64, error) {
	b, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	v := uint64(b & 0x7f)
	hasNext := (b & 0x80) != 0

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
		v = (v << 7) | uint64(b&0x7f)
		hasNext = (b & 0x80) != 0
	}
	if hasNext {
		b, err = r.ReadByte()
		if err != nil {
			return 0, err
		}
		v = (v << 8) | uint64(b)
	}
	return v, nil
}

func ReadVarsize(r *bitio.CountReader) (uint64, error) {
	b, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	v := uint64(b & 0x7f)
	hasNext := (b & 0x80) != 0

	for i := 0; i < 3; i++ {
		if !hasNext {
			break
		}
		if b, err = r.ReadByte(); err != nil {
			return 0, err
		}
		v = (v << 7) | uint64(b&0x7f)
		hasNext = (b & 0x80) != 0
	}
	if hasNext {
		b, err = r.ReadByte()
		if err != nil {
			return 0, err
		}
		v = (v << 8) | uint64(b)
	}
	if v > VARSIZE_MAX {
		return 0, ErrOutOfBounds
	}
	return v, nil
}
