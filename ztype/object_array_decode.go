package ztype

import zserio "github.com/woven-planet/go-zserio"

// UnmarshalZserio reads an array of structs from a bit reader, in either packed or unpacked configuration.
func (array *ObjectArray[T, PT]) UnmarshalZserio(reader zserio.Reader) ([]T, error) {
	arraySize := array.FixedSize
	if array.IsAuto {
		// The array size is passed as a varsize inside the byte stream
		u64ReadSize, err := ReadVarsize(reader)
		if err != nil {
			return nil, err
		}
		arraySize = int(u64ReadSize)
	}

	result := make([]T, arraySize)
	arrayTraits := ObjectArrayTraits[T, PT]{}

	if arraySize > 0 {
		var err error
		var element PT
		var packedTraits IPackedObjectArrayTraits[T, PT]

		if array.IsPacked {
			packedTraits = arrayTraits.PackedTraits()
			// A descriptor is only written for packed arrays.
			array.PackedContext, err = packedTraits.CreateContext()
			if err != nil {
				return nil, err
			}
		}

		for index := 0; index < arraySize; index++ {
			element = &result[index]
			if array.SetCompoundParameterFn != nil {
				array.SetCompoundParameterFn(index, element)
			}

			if array.checkOffsetMethod != nil {
				count, err := reader.Align(8)
				if err != nil {
					return nil, err
				}
				array.checkOffsetMethod(index, count)
			}

			if array.IsPacked {
				err = packedTraits.Read(array.PackedContext, reader, element)
			} else {
				err = arrayTraits.Read(reader, element)
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return result, nil
}
