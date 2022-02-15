package ztype

import (
	"github.com/icza/bitio"
	"github.com/woven-planet/go-zserio"
)

// writeDescriptor writes the descriptor of a packing context.
func writeDescriptor(packingNode *zserio.PackingContextNode, writer *bitio.CountWriter) error {
	if packingNode.HasContext() {
		return packingNode.WriteDescriptor(writer)
	}
	for _, childNode := range packingNode.GetChildren() {
		if err := writeDescriptor(childNode, writer); err != nil {
			return err
		}
	}
	return nil
}

// MarshalZserio writes an array to a bit writer.
func (array *Array[T, Y]) MarshalZserio(writer *bitio.CountWriter) error {
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
			writeDescriptor(array.PackedContext, writer)
		}
		for index, element := range array.RawArray {
			if array.checkOffsetMethod != nil {
				writer.Align()
				array.checkOffsetMethod(index, writer.BitsCount)
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
