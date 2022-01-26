package ztype

import (
	"github.com/icza/bitio"
	"github.com/woven-planet/go-zserio/interface"
)

// readDescriptor reads the descriptor of a packed array, and all of its children.
func readDescriptor(packingNode *zserio.PackingContextNode, reader *bitio.CountReader) error {
	if packingNode.HasContext() {
		return packingNode.ReadDescriptor(reader)
	}
	for _, childNode := range packingNode.GetChildren() {
		if err := readDescriptor(childNode, reader); err != nil {
			return err
		}
	}
	return nil
}

// UnmarshalZserio reads an array from a bit reader, in either packed or unpacked configuration.
func (array *Array[T, Y]) UnmarshalZserio(reader *bitio.CountReader) error {
	arraySize := array.FixedSize
	if array.IsAuto {
		// The array size is passed as a varsize inside the byte stream
		u64ReadSize, err := ReadVarsize(reader)
		if err != nil {
			return err
		}
		arraySize = int(u64ReadSize)
	}

	array.RawArray = make([]T, arraySize)
	if arraySize > 0 {
		if array.IsPacked {
			var err error
			// A descriptor is only written for packed arrays.
			array.PackedContext, err = array.ArrayTraits.PackedTraits().CreateContext()
			if err != nil {
				return err
			}
			err = readDescriptor(array.PackedContext, reader)
			if err != nil {
				return err
			}
		}
		for index := 0; index < arraySize; index++ {
			if array.CheckOffsetMethod != nil {
				reader.Align()
				(*array.CheckOffsetMethod)(index, reader.BitsCount)
			}
			var err error
			var element T
			if array.IsPacked {
				element, err = array.ArrayTraits.PackedTraits().Read(array.PackedContext, reader, index)
			} else {
				element, err = array.ArrayTraits.Read(reader, index)
			}
			if err != nil {
				return err
			}
			array.RawArray[index] = element
		}
	}
	return nil
}

// ArrayFromReader is a helper function to read an array as a one-liner.
func ArrayFromReader[T any, Y IArrayTraits[T]](
	reader *bitio.CountReader, arrayTraits Y, size int, isPacked, isAuto bool, setOffsetMethod, checkOffsetMethod *OffsetMethod) (*Array[T, Y], error) {
	arrayInstance := Array[T, Y]{
		ArrayTraits:       arrayTraits,
		RawArray:          make([]T, 0),
		IsAuto:            isAuto,
		IsPacked:          isPacked,
		SetOffsetMethod:   setOffsetMethod,
		CheckOffsetMethod: checkOffsetMethod,
		FixedSize:         size,
	}
	err := arrayInstance.UnmarshalZserio(reader)
	if err != nil {
		return nil, err
	}
	return &arrayInstance, nil
}
