package ztype

import (
	"io"

	zserio "github.com/woven-planet/go-zserio"
)

// ReadString reads a string from the bitstream.
func ReadString(r zserio.Reader) (string, error) {
	size, err := ReadVarsize(r)
	if err != nil {
		return "", err
	}
	buf := make([]byte, size)
	if _, err = io.ReadFull(r, buf); err != nil {
		return "", err
	}
	return string(buf), nil
}
