package ztype

import (
	zserio "github.com/woven-planet/go-zserio"
)

// IPackedArrayTraits is the interface for array traits that are packed.
type IPackedArrayTraits[T any] interface {
	CreateContext() (*zserio.PackingContextNode, error)
	InitContext(contextNode *zserio.PackingContextNode, element T) error
	BitSizeOf(contextNode *zserio.PackingContextNode, bitPosition int, element T) (int, error)
	InitializeOffsets(contextNode *zserio.PackingContextNode, bitPosition int, element T) int
	Read(contextNode *zserio.PackingContextNode, reader zserio.Reader, index int) (T, error)
	Write(contextNode *zserio.PackingContextNode, writer zserio.Writer, value T)
}

// PackedArrayTraits is a wrapper around array traits used when the array is
// written in packed configruation.
type PackedArrayTraits[T any, Y IArrayTraits[T]] struct {
	ArrayTraits Y
}

func (traits *PackedArrayTraits[T, Y]) CreateContext() (*zserio.PackingContextNode, error) {
	// A packed array traits is using only one single delta context, so create
	// the delta context now and return the complete context.
	// Any compound object uses the ObjectPackedArrayTraits, where several
	// packing contexts will be used.
	return CreatePackingContextNode[T](), nil
}

// InitContext initializes a packing context for a packed array traits.
func (traits *PackedArrayTraits[T, Y]) InitContext(contextNode *zserio.PackingContextNode, element T) error {
	contextNode.Context.(*DeltaContext[T]).Init(traits.ArrayTraits, element)
	return nil
}

// BitSizeOf returns the bit size of an element of a packed array traits.
func (traits *PackedArrayTraits[T, Y]) BitSizeOf(contextNode *zserio.PackingContextNode, bitPosition int, element T) (int, error) {
	return contextNode.Context.(*DeltaContext[T]).BitSizeOf(traits.ArrayTraits, bitPosition, element)
}

// InitializeOffsets calculates the offsets of the packed array element.
func (traits *PackedArrayTraits[T, Y]) InitializeOffsets(contextNode *zserio.PackingContextNode, bitPosition int, element T) int {
	delta, _ := contextNode.Context.(*DeltaContext[T]).BitSizeOf(traits.ArrayTraits, bitPosition, element)
	return bitPosition + delta
}

// Read reads an array element of a packed array traits.
func (traits *PackedArrayTraits[T, Y]) Read(contextNode *zserio.PackingContextNode, reader zserio.Reader, index int) (T, error) {
	return contextNode.Context.(*DeltaContext[T]).Read(traits.ArrayTraits, reader)
}

// Write writes an array element of a packed array traits.
func (traits *PackedArrayTraits[T, Y]) Write(contextNode *zserio.PackingContextNode, writer zserio.Writer, value T) {
	contextNode.Context.(*DeltaContext[T]).Write(traits.ArrayTraits, writer, value)
}
