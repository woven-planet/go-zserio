package ztype

import zserio "github.com/woven-planet/go-zserio"

// ObjectArray allows representing arrays of structs and serialize them to the zserio format.
type ObjectArray[T any, PT zserio.PackableZserioType[T]] struct {
	ArrayTraits ObjectArrayTraits[T, PT]

	// The node used by this array for packing
	PackedContext *zserio.PackingContextNode

	// SetOffsetMethod is an optional function to set the offset to the buffer.
	setOffsetMethod   OffsetMethod
	checkOffsetMethod OffsetMethod

	// SetCompoundParameterFn is an optional function to set the compound parameters
	SetCompoundParameterFn func(index int, item PT)

	// FixedSize is the size of the array, if the array is of fixed size
	FixedSize int

	// IsAuto specifies if the array size is automatically calculated.
	IsAuto bool

	// IsPacked specifies if the array is packed.
	IsPacked bool
}

// ZserioBitSize returns the total size of the unpacked object array in bits.
func (array *ObjectArray[T, PT]) ZserioBitSize(elements []T, bitPosition int) (int, error) {
	endBitPosition := bitPosition
	size := len(elements)
	if array.IsAuto {
		delta, err := SignedBitSize(int64(size), 4)
		if err != nil {
			return 0, err
		}
		endBitPosition += delta
	}
	if array.ArrayTraits.BitSizeOfIsConstant() && size > 0 {
		var dummy PT
		elementSize := array.ArrayTraits.BitSizeOf(dummy, 0)
		if array.setOffsetMethod != nil {
			endBitPosition += size * elementSize
		} else {
			// all elements are spaced in the same way
			endBitPosition = alignTo(8, endBitPosition)
			endBitPosition += elementSize + (size-1)*alignTo(8, elementSize)
		}
	} else {
		for i := 0; i < size; i++ {
			if array.setOffsetMethod != nil {
				endBitPosition = alignTo(8, endBitPosition)
			}
			endBitPosition += array.ArrayTraits.BitSizeOf(&elements[i], endBitPosition)
		}
	}
	return endBitPosition - bitPosition, nil
}

// ZserioBitSizePacked returns the total size of the packed object array in bits.
func (array *ObjectArray[T, PT]) ZserioBitSizePacked(elements []T, bitPosition int) (int, error) {
	endBitPosition := bitPosition
	size := len(elements)
	if array.IsAuto {
		delta, err := SignedBitSize(int64(size), 4)
		if err != nil {
			return 0, err
		}
		endBitPosition += delta
	}
	if size > 0 {
		packedTraits := array.ArrayTraits.PackedTraits()

		for i := 0; i < size; i++ {
			if array.setOffsetMethod != nil {
				endBitPosition = alignTo(8, endBitPosition)
			}
			delta, err := packedTraits.BitSizeOf(array.PackedContext, endBitPosition, &elements[i])
			if err != nil {
				return 0, err
			}
			endBitPosition += delta
		}
	}
	return endBitPosition - bitPosition, nil
}
