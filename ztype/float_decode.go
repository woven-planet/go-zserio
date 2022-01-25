package ztype

import (
	"encoding/binary"

	"github.com/icza/bitio"
	"github.com/x448/float16"
)

func ReadFloat16(r *bitio.Reader) (float32, error) {
	var v uint16
	if err := binary.Read(r, binary.BigEndian, &v); err != nil {
		return 0, err
	}
	return float16.Frombits(v).Float32(), nil
}

func ReadFloat32(r *bitio.Reader) (float32, error) {
	var v float32
	err := binary.Read(r, binary.BigEndian, &v)
	return v, err
}

func ReadFloat64(r *bitio.Reader) (float64, error) {
	var v float64
	err := binary.Read(r, binary.BigEndian, &v)
	return v, err
}
