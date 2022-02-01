package ztype

import "github.com/icza/bitio"

// WriteBool reads a boolean value from the bitstream
func WriteBool(w *bitio.CountWriter, v bool) error {
	return w.WriteBool(v)
}
