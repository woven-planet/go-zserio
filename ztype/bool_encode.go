package ztype

type BoolWriter interface {
	WriteBool(bool) error
}

// WriteBool reads a boolean value from the bitstream
func WriteBool(w BoolWriter, v bool) error {
	return w.WriteBool(v)
}
