package ztype

import (
	"encoding/binary"

	"github.com/icza/bitio"
	"github.com/x448/float16"
)

// ReadFloat16 reads a 16-bit floating point number.
func ReadFloat16(r *bitio.CountReader) (float32, error) {
	var v uint16
	if err := binary.Read(r, binary.BigEndian, &v); err != nil {
		return 0, err
	}
	return float16.Frombits(v).Float32(), nil
}

// ReadFloat32 reads a 32-bit floating point number.
func ReadFloat32(r *bitio.CountReader) (float32, error) {
	var v float32
	err := binary.Read(r, binary.BigEndian, &v)
	return v, err
}

// ReadFloat64 reads a 64-bit floating point number.
func ReadFloat64(r *bitio.CountReader) (float64, error) {
	var v float64
	err := binary.Read(r, binary.BigEndian, &v)
	return v, err
}
