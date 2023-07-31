package ztype

import zserio "github.com/woven-planet/go-zserio"

// WriteBytes writes an zserio bytes type (variable size byte buffer) from a reader.
func WriteBytes(w zserio.Writer, b *BytesType) error {
	if err := WriteVarsize(w, b.ByteSize); err != nil {
		return err
	}

	if _, err := w.Write(b.Buffer); err != nil {
		return err
	}
	return nil
}
