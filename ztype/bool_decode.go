package ztype

type BoolReader interface {
	ReadBool() (bool, error)
}

// ReadBool reads a boolean value from the bitstream
func ReadBool(r BoolReader) (bool, error) {
	return r.ReadBool()
}
