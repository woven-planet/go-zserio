package ztype

import "io"

type StringReader interface {
	io.Reader
	io.ByteReader
}

// ReadString reads a string from the bitstream.
func ReadString(r StringReader) (string, error) {
	size, err := ReadVarsize(r)
	if err != nil {
		return "", err
	}
	buf := make([]byte, size)
	if _, err = r.Read(buf); err != nil {
		return "", err
	}
	return string(buf), nil
}
