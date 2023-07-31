package ztype

import (
	"errors"

	zserio "github.com/woven-planet/go-zserio"
)

// ReadBytes reads an zserio bytes type (variable size byte buffer) from a reader.
func ReadBytes(r zserio.Reader) (*BytesType, error) {
	var err error
	b := &BytesType{}
	b.ByteSize, err = ReadVarsize(r)
	if err != nil {
		return nil, err
	}

	b.Buffer = make([]byte, b.ByteSize)
	n, err := r.Read(b.Buffer)
	if err != nil {
		return nil, err
	}
	if n != int(b.ByteSize) {
		return nil, errors.New("read number of bytes didn't match")
	}
	return b, nil
}
