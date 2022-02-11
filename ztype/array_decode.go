package ztype

import (
	"fmt"

	"github.com/woven-planet/go-zserio"
)

// readDescriptor reads the descriptor of a packed array, and all of its children.
func readDescriptor(packingNode *zserio.PackingContextNode, reader zserio.Reader) error {
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
func (array *Array[T, Y]) UnmarshalZserio(reader zserio.Reader) error {
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
			if array.checkOffsetMethod != nil {
				// FIXME @aignas 2022-02-10: What is the aligning boundary for arrays?
				count, err := reader.Align(8)
				if err != nil {
					return fmt.Errorf("align: %w", err)
				}
				array.checkOffsetMethod(index, count)
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
func ArrayFromReader[T any, Y IArrayTraits[T]](reader zserio.Reader, arrayTraits Y, size int, isPacked, isAuto bool, options ...ArrayOption[T, Y]) (*Array[T, Y], error) {
	arrayInstance := Array[T, Y]{
		ArrayTraits: arrayTraits,
		RawArray:    make([]T, 0),
		IsAuto:      isAuto,
		IsPacked:    isPacked,
		FixedSize:   size,
	}
	for _, opt := range options {
		opt.apply(&arrayInstance)
	}

	err := arrayInstance.UnmarshalZserio(reader)
	if err != nil {
		return nil, err
	}
	return &arrayInstance, nil
}

// ArrayOption allows to customize array creation.
type ArrayOption[T any, Y IArrayTraits[T]] interface {
	apply(*Array[T, Y])
}

func WithSetOffsetMethod[T any, Y IArrayTraits[T]](fn OffsetMethod) ArrayOption[T, Y] {
	return setOffsetMethod[T, Y](fn)
}

type setOffsetMethod[T any, Y IArrayTraits[T]] OffsetMethod

func (fn setOffsetMethod[T, Y]) apply(array *Array[T, Y]) {
	array.setOffsetMethod = OffsetMethod(fn)
}

func WithCheckOffsetMethod[T any, Y IArrayTraits[T]](fn OffsetMethod) ArrayOption[T, Y] {
	return checkOffsetMethod[T, Y](fn)
}

type checkOffsetMethod[T any, Y IArrayTraits[T]] OffsetMethod

func (fn checkOffsetMethod[T, Y]) apply(array *Array[T, Y]) {
	array.checkOffsetMethod = OffsetMethod(fn)
}
