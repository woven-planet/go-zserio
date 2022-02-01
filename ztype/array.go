package ztype

import (
	"github.com/woven-planet/go-zserio/interface"
)

// BitSizeOfDescriptor returns the bit size of a descriptor.
func bitSizeOfDescriptor(packingNode *zserio.PackingContextNode, bitPosition int) (int, error) {
	endBitPosition := bitPosition
	if packingNode.HasContext() {
		endBitPosition += packingNode.BitSizeOfDescriptor()
	} else {
		for _, childNode := range packingNode.GetChildren() {
			delta, err := bitSizeOfDescriptor(childNode, endBitPosition)
			if err != nil {
				return 0, err
			}
			endBitPosition += delta
		}
	}
	return endBitPosition - bitPosition, nil
}

// OffsetMethod is a function used to set/check bit offsets in the buffer.
type OffsetMethod func(int, int64)

// Array allows representing arrays of any type and serialize them to the zserio format.
type Array[T any, Y IArrayTraits[T]] struct {
	// ArrayTraits are the array traits used.
	ArrayTraits Y
	// RawArray is the raw array.
	RawArray []T

	// IsAuto specifies if the array size is automatically calculated.
	IsAuto bool

	// IsPacked specifies if the array is packed.
	IsPacked bool

	// FixedSize is the size of the array, if the array is of fixed size
	FixedSize int

	// The node used by this array for packing
	PackedContext *zserio.PackingContextNode

	// SetOffsetMethod is an optional function to set the offset to the buffer.
	setOffsetMethod   OffsetMethod
	checkOffsetMethod OffsetMethod
}

// Size returns the number of elements in an array.
func (array *Array[T, Y]) Size() int {
	return len(array.RawArray)
}

// ZserioBitSize returns the total size of the unpacked array in bits.
func (array *Array[T, Y]) ZserioBitSize(bitPosition int) (int, error) {
	endBitPosition := bitPosition
	size := array.Size()
	if array.IsAuto {
		delta, err := SignedBitSize(int64(size), 4)
		if err != nil {
			return 0, err
		}
		endBitPosition += delta
	}
	if array.ArrayTraits.BitSizeOfIsConstant() && size > 0 {
		var dummy T
		elementSize := array.ArrayTraits.BitSizeOf(dummy, 0)
		if array.setOffsetMethod != nil {
			endBitPosition += size * elementSize
		} else {
			// all elements are spaced in the same way
			endBitPosition = alignTo(8, endBitPosition)
			endBitPosition += elementSize + (size-1)*alignTo(8, elementSize)
		}
	} else {
		for _, element := range array.RawArray {
			if array.setOffsetMethod != nil {
				endBitPosition = alignTo(8, endBitPosition)
			}
			endBitPosition += array.ArrayTraits.BitSizeOf(element, endBitPosition)
		}
	}
	return endBitPosition - bitPosition, nil
}

// BitSizeOfPacked returns the total size of the packed array in bits.
func (array *Array[T, Y]) ZserioBitSizePacked(bitPosition int) (int, error) {
	endBitPosition := bitPosition
	size := array.Size()
	if array.IsAuto {
		delta, err := SignedBitSize(int64(size), 4)
		if err != nil {
			return 0, err
		}
		endBitPosition += delta
	}
	if size > 0 {
		//contextNode := array.PackedContext
		/* Createzserio.PackingContextNode[T]()
		for _, element := range array.RawArray {
			array.ArrayTraits.PackedTraits().(*PackedArrayTraits[T, IArrayTraits[T]]).InitContext(contextNode, element)
		}
		*/
		delta, err := bitSizeOfDescriptor(array.PackedContext, endBitPosition)
		if err != nil {
			return 0, err
		}
		endBitPosition += delta
		for _, element := range array.RawArray {
			if array.setOffsetMethod != nil {
				endBitPosition = alignTo(8, endBitPosition)
			}
			delta, err := array.ArrayTraits.PackedTraits().BitSizeOf(array.PackedContext, endBitPosition, element)
			if err != nil {
				return 0, err
			}
			endBitPosition += delta
		}
	}
	return endBitPosition - bitPosition, nil
}

// Clone does a deep copy of the array.
func (array *Array[T, Y]) Clone() zserio.ZserioType {
	clone := Array[T, Y]{
		ArrayTraits:       array.ArrayTraits,
		RawArray:          array.RawArray,
		IsAuto:            array.IsAuto,
		IsPacked:          array.IsPacked,
		FixedSize:         array.FixedSize,
		PackedContext:     array.PackedContext,
		setOffsetMethod:   array.setOffsetMethod,
		checkOffsetMethod: array.checkOffsetMethod,
	}
	return &clone
}
