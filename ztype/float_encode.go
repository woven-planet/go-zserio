package ztype

import (
	"encoding/binary"

	"github.com/icza/bitio"
	"github.com/x448/float16"
)

func WriteFloat16(w *bitio.CountWriter, v float32) error {
	f := float16.Fromfloat32(v).Bits()
	return binary.Write(w, binary.BigEndian, f)
}

func WriteFloat32(w *bitio.CountWriter, v float32) error {
	return binary.Write(w, binary.BigEndian, v)
}

func WriteFloat64(w *bitio.CountWriter, v float64) error {
	return binary.Write(w, binary.BigEndian, v)
}
