package ztype

import zserio "github.com/woven-planet/go-zserio"

// MarshalZserio writes an array to a bit writer.
func (array *ObjectArray[T, PT]) MarshalZserio(elements []T, writer zserio.Writer) error {
	size := len(elements)
	if array.IsAuto {
		err := WriteVarsize(writer, uint64(size))
		if err != nil {
			return err
		}
	}

	if size > 0 {
		var packedTraits IPackedObjectArrayTraits[T, PT]
		var element PT

		if array.IsPacked {
			// descriptors are only needed if the array is packed
			var err error
			packedTraits = array.ArrayTraits.PackedTraits()
			array.PackedContext, err = packedTraits.CreateContext()
			if err != nil {
				return err
			}

			for index := 0; index < size; index++ {
				packedTraits.InitContext(array.PackedContext, &elements[index])
			}
		}

		for index := 0; index < size; index++ {
			element = &elements[index]

			if array.checkOffsetMethod != nil {
				count, err := writer.Align(8)
				if err != nil {
					return err
				}
				array.checkOffsetMethod(index, count)
			}
			if array.IsPacked {
				packedTraits.Write(array.PackedContext, writer, element)
			} else {
				array.ArrayTraits.Write(writer, element)
			}
		}
	}
	return nil
}
