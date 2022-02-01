package ztype

import "github.com/icza/bitio"

func WriteVarint16(w *bitio.CountWriter, v int16) error {
	return writeVarInt(w, int64(v), 2)
}

func WriteVarint32(w *bitio.CountWriter, v int32) error {
	return writeVarInt(w, int64(v), 4)
}

func WriteVarint64(w *bitio.CountWriter, v int64) error {
	return writeVarInt(w, int64(v), 8)
}

func WriteVarint(w *bitio.CountWriter, v int64) error {
	if v == MinInt64 {
		return w.WriteByte(0x80)
	}
	return writeVarInt(w, int64(v), 9)
}

func writeVarInt(w *bitio.CountWriter, v int64, maxBytes int) error {
	var absValue uint64
	if v < 0 {
		absValue = uint64(-v)
	} else {
		absValue = uint64(v)
	}
	neededBytes, err := SignedBitSize(v, maxBytes)
	if err != nil {
		return err
	}
	neededBytes /= 8
	needsMaxByteRange := neededBytes == maxBytes
	for i := 0; i < neededBytes; i++ {
		var b uint64 = 0x00
		remainingBits := 8
		hasNextByte := i < (neededBytes - 1)
		hasSignBit := i == 0
		if hasSignBit {
			if v < 0 {
				b |= 0x80
			}
			remainingBits--
		}
		if hasNextByte {
			remainingBits--
			b |= (1 << remainingBits)
		} else if !needsMaxByteRange {
			remainingBits--
		}
		shiftBits := (neededBytes - (i + 1)) * 7
		if needsMaxByteRange && hasNextByte {
			shiftBits++
		}
		b |= (absValue >> int64(shiftBits)) & (0xff >> (8 - remainingBits))
		w.TryWriteBitsUnsafe(b, 8)
	}
	return w.TryError
}

// SignedBitSize returns the size in bits of the zserio encoding of a signed
// value. Unlike the Python zserio version this version is generic and does
// not need a per-type table, but still gets identical performance.
//
func SignedBitSize(v int64, maxBytes int) (int, error) {
	if v < 0 {
		v = -v
	}
	var max int64 = (1 << 6) - 1
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
