package ztype

import (
	zserio "github.com/woven-planet/go-zserio"
)

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

	array.RawArray = make([]T, 0, arraySize)

	if arraySize > 0 {
		var err error
		var element T
		var packedTraits IPackedArrayTraits[T] = nil

		if array.IsPacked {
			packedTraits = array.ArrayTraits.PackedTraits()
			// A descriptor is only written for packed arrays.
			array.PackedContext, err = packedTraits.CreateContext()
			if err != nil {
				return err
			}
		}

		for index := 0; index < arraySize; index++ {
			if array.checkOffsetMethod != nil {
				count, err := reader.Align(8)
				if err != nil {
					return err
				}
				array.checkOffsetMethod(index, count)
			}

			if array.IsPacked {
				element, err = packedTraits.Read(array.PackedContext, reader, index)
			} else {
				element, err = array.ArrayTraits.Read(reader, index)
			}
			if err != nil {
				return err
			}
			array.RawArray = append(array.RawArray, element)
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
