package ztype

import (
	"errors"

	"github.com/icza/bitio"
)

// ReadExtern reads an zserio extern type (variable size bitbuffer) from a reader.
func ReadExtern(r *bitio.CountReader) (*ExternType, error) {
	var err error
	e := &ExternType{}
	e.BitSize, err = ReadVarsize(r)
	if err != nil {
		return nil, err
	}
	numOfFullBytes := int(e.BitSize / 8)
	remainingBits := uint8(e.BitSize % 8)

	e.Buffer = make([]byte, numOfFullBytes)
	n, err := r.Read(e.Buffer)
	if err != nil {
		return nil, err
	}
	if n != numOfFullBytes {
		return nil, errors.New("read number of bytes didn't match")
	}
	if remainingBits != 0 {
		lastByte, err := r.ReadBits(remainingBits)
		if err != nil {
			return nil, err
		}
		e.Buffer = append(e.Buffer, byte(lastByte))
	}
	return e, nil
}
