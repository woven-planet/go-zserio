package ztype

import "github.com/icza/bitio"

func WriteString(w *bitio.CountWriter, v string) error {
	if err := WriteVarsize(w, uint64(len(v))); err != nil {
		return err
	}
	_, err := w.Write([]byte(v))
	return err
}

// BitSizeOfString returns the zserio bit size of a string.
func BitSizeOfString(v string) (int, error) {
	// the length of a string is encoded as a varsize (max 5 bytes)
	headerSize, err := UnsignedBitSize(uint64(len(v)), 5)
	if err != nil {
		return 0, err
	}
	return headerSize + len(v)*8, nil
}
