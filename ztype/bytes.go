package ztype

// BytesType is a helper struct to store byte blobs in zserio
type BytesType struct {
	// Buffer is the content of the zserio byte type
	Buffer []byte

	// BitSize is the size of the buffer in bytes
	ByteSize uint64
}
