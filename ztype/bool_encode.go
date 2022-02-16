package ztype

import zserio "github.com/woven-planet/go-zserio"

// WriteBool reads a boolean value from the bitstream
func WriteBool(w zserio.Writer, v bool) error {
	return w.WriteBool(v)
}
