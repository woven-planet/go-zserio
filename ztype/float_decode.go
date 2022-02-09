package ztype

import (
	"io"
	"encoding/binary"

	"github.com/x448/float16"
)

// ReadFloat16 reads a 16-bit floating point number.
func ReadFloat16(r io.Reader) (float32, error) {
	var v uint16
	if err := binary.Read(r, binary.BigEndian, &v); err != nil {
		return 0, err
	}
	return float16.Frombits(v).Float32(), nil
}

// ReadFloat32 reads a 32-bit floating point number.
func ReadFloat32(r io.Reader) (float32, error) {
	var v float32
	err := binary.Read(r, binary.BigEndian, &v)
	return v, err
}

// ReadFloat64 reads a 64-bit floating point number.
func ReadFloat64(r io.Reader) (float64, error) {
	var v float64
	err := binary.Read(r, binary.BigEndian, &v)
	return v, err
}
