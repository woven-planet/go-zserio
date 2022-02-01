package ztype

import "github.com/icza/bitio"

// ReadBool reads a boolean value from the bitstream
func ReadBool(r *bitio.CountReader) (bool, error) {
	return r.ReadBool()
}
