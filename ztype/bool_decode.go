package ztype

import zserio "github.com/woven-planet/go-zserio"

// ReadBool reads a boolean value from the bitstream
func ReadBool(r zserio.Reader) (bool, error) {
	return r.ReadBool()
}
