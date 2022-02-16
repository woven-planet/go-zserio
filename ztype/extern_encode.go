package ztype

import zserio "github.com/woven-planet/go-zserio"

// WriteExtern writes an zserio extern type (variable size bitbuffer) from a reader.
func WriteExtern(w zserio.Writer, e *ExternType) error {
	if err := WriteVarsize(w, e.BitSize); err != nil {
		return err
	}

	// write full bytes
	numOfFullBytes := e.BitSize / 8
	remainingBits := uint8(e.BitSize % 8)
	if _, err := w.Write(e.Buffer[0:numOfFullBytes]); err != nil {
		return err
	}

	// and the remaining 1-7 bits
	if remainingBits != 0 {
		lastByte := uint64(e.Buffer[len(e.Buffer)-1])
		if err := w.WriteBits(lastByte, remainingBits); err != nil {
			return err
		}
	}
	return nil
}
