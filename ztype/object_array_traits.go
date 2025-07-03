package ztype

import zserio "github.com/woven-planet/go-zserio"

type IPackedObjectArrayTraits[T any, PT zserio.PackableZserioType[T]] interface {
	CreateContext() (*zserio.PackingContextNode, error)
	InitContext(contextNode *zserio.PackingContextNode, element PT) error
	BitSizeOf(contextNode *zserio.PackingContextNode, bitPosition int, element PT) (int, error)
	InitializeOffsets(contextNode *zserio.PackingContextNode, bitPosition int, element PT) int
	Read(contextNode *zserio.PackingContextNode, reader zserio.Reader, element PT) error
	Write(contextNode *zserio.PackingContextNode, writer zserio.Writer, value PT)
}

// ObjectArrayTraits is an array traits for zserio structs, choice, union or enum types
type ObjectArrayTraits[T any, PT zserio.PackableZserioType[T]] struct{}

func (trait ObjectArrayTraits[T, PT]) PackedTraits() IPackedObjectArrayTraits[T, PT] {
	return &ObjectArrayPackedTraits[T, PT]{}
}

func (trait ObjectArrayTraits[T, PT]) BitSizeOfIsConstant() bool {
	return false
}

func (trait ObjectArrayTraits[T, PT]) NeedsBitsizeOfPosition() bool {
	return true
}

func (trait ObjectArrayTraits[T, PT]) NeedsReadIndex() bool {
	return true
}

func (trait ObjectArrayTraits[T, PT]) BitSizeOf(element PT, endBitPosition int) int {
	bitSize, _ := element.ZserioBitSize(endBitPosition)
	return bitSize
}

func (trait ObjectArrayTraits[T, PT]) InitializeOffsets(bitPosition int, value PT) int {
	offset := 0
	return offset
}

func (trait ObjectArrayTraits[T, PT]) Read(reader zserio.Reader, element PT) error {
	return element.UnmarshalZserio(reader)
}

func (trait ObjectArrayTraits[T, PT]) Write(writer zserio.Writer, value PT) error {
	return value.MarshalZserio(writer)
}

func (trait ObjectArrayTraits[T, PT]) AsUint64(value T) uint64 {
	return 0 // not supported
}

func (trait ObjectArrayTraits[T, PT]) FromUint64(value uint64) T {
	var dummy T
	return dummy // not supported
}

// ObjectPackedArrayTraits is a wrapper around array traits for zserio objects.
type ObjectArrayPackedTraits[T any, PT zserio.PackableZserioType[T]] struct{}

func (traits *ObjectArrayPackedTraits[T, PT]) CreateContext() (*zserio.PackingContextNode, error) {
	// Create a packing context without a context set (it will be set by the
	// compound object fields).
	node := &zserio.PackingContextNode{}

	// this is deliberately a null pointer (simulating static methods), as the
	// objects are not created yet.
	var dummyObjectPtr PT

	// Iterate over all fields in the object, an generate one packing context
	// for each field.
	err := dummyObjectPtr.ZserioCreatePackingContext(node)
	return node, err
}

// InitContext initializes a packing context for a packed array traits.
func (traits *ObjectArrayPackedTraits[T, PT]) InitContext(contextNode *zserio.PackingContextNode, element PT) error {
	return element.ZserioInitPackingContext(contextNode)
}

// BitSizeOf returns the bit size of an element of a packed array traits.
func (traits *ObjectArrayPackedTraits[T, PT]) BitSizeOf(contextNode *zserio.PackingContextNode, bitPosition int, element PT) (int, error) {
	return element.ZserioBitSizePacked(contextNode, bitPosition)
}

// InitializeOffsets calculates the offsets of the packed array element.
func (traits *ObjectArrayPackedTraits[T, PT]) InitializeOffsets(contextNode *zserio.PackingContextNode, bitPosition int, element PT) int {
	return element.ZserioInitializeOffsetsPacked(contextNode, bitPosition)
}

// Read reads an array element of a packed array traits.
func (traits *ObjectArrayPackedTraits[T, PT]) Read(contextNode *zserio.PackingContextNode, reader zserio.Reader, element PT) error {
	return element.UnmarshalZserioPacked(contextNode, reader)
}

// Write writes an array element of a packed array traits.
func (traits *ObjectArrayPackedTraits[T, PT]) Write(contextNode *zserio.PackingContextNode, writer zserio.Writer, element PT) {
	element.MarshalZserioPacked(contextNode, writer)
}
