package ztype

import (
	zserio "github.com/woven-planet/go-zserio"
)

// MarshalZserio writes an array to a bit writer.
func (array *Array[T, Y]) MarshalZserio(writer zserio.Writer) error {
	size := array.Size()
	if array.IsAuto {
		err := WriteVarsize(writer, uint64(size))
		if err != nil {
			return err
		}
	}

	if size > 0 {
		var packedTraits IPackedArrayTraits[T]
		if array.IsPacked {
			// descriptors are only needed if the array is packed
			var err error
			packedTraits = array.ArrayTraits.PackedTraits()
			array.PackedContext, err = packedTraits.CreateContext()
			if err != nil {
				return err
			}
			for _, element := range array.RawArray {
				packedTraits.InitContext(array.PackedContext, element)
			}
		}
		for index, element := range array.RawArray {
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
