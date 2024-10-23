package ztype

// ExternType is a helper struct to store zserio extern types
type ExternType struct {
	// Buffer is the content of the zserio extern type
	Buffer []byte

	// BitSize is the size of the buffer in bits
	BitSize uint64
}
