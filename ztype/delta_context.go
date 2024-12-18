package ztype

import (
	"math/bits"

	zserio "github.com/woven-planet/go-zserio"
)

const (
	maxBitNumberBits  = 6
	maxBitNumberLimit = 62
)

// DeltaContext is a packing context used when writing data using delta
// packing, i.e. instead of storing all values, only stores the deltas.
type DeltaContext[T any] struct {
	// previousElement is the value of the previously stored element
	previousElement *uint64

	// maxBitNumber specifies the number of bits needed per delta element
	maxBitNumber        int
	unpackedBitSize     int
	firstElementBitSize int
	numElements         int

	// specifies if delta packing is actually used (it may be skipped if normal
	// packing is more efficient)
	isPacked          bool
	processingStarted bool
}

// arrayTraitsBitsizeOf returns the bit size of an array element.
func arrayTraitsBitsizeOf[T any](arrayTraits IArrayTraits[T], bitPosition int, element T) int {
	return arrayTraits.BitSizeOf(element, bitPosition)
}

func absDiff(lhs, rhs uint64) uint64 {
	if lhs > rhs {
		return lhs - rhs
	}
	return rhs - lhs
}

// Init initializes a delta context array, and calculates the needed space per element in the final array.
func (context *DeltaContext[T]) Init(arrayTraits IArrayTraits[T], element T) {
	context.numElements++
	context.unpackedBitSize += bitsizeOfUnpacked(arrayTraits, element)

	if context.previousElement == nil {
		context.previousElement = new(uint64)
		context.firstElementBitSize = context.unpackedBitSize
	} else if context.maxBitNumber <= maxBitNumberLimit {
		context.isPacked = true
		// Calculate the delta to the previous value, and calculate how many
		// bits are needed to store the delta.
		delta := absDiff(arrayTraits.AsUint64(element), *context.previousElement)
		maxBitNumber := bits.Len64(delta)
		if maxBitNumber > context.maxBitNumber {
			context.maxBitNumber = maxBitNumber
			// cannot store using delta packing if the 64bit range is
			// exhausted
			if maxBitNumber > maxBitNumberLimit {
				context.isPacked = false
			}
		}
	}
	*context.previousElement = arrayTraits.AsUint64(element)
}

// BitSizeOf returns the size of the delta context array in bits.
func (context *DeltaContext[T]) BitSizeOf(arrayTraits IArrayTraits[T], bitPosition int, element T) (int, error) {
	if !context.processingStarted {
		context.processingStarted = true
		context.finishInit()
		return context.BitSizeOfDescriptor() + bitsizeOfUnpacked(arrayTraits, element), nil
	}
	if !context.isPacked {
		return bitsizeOfUnpacked(arrayTraits, element), nil
	}
	if context.maxBitNumber > 0 {
		return context.maxBitNumber + 1, nil
	}
	return 0, nil
}

// Read reads the next element of an array encoded using delta contexts.
func (context *DeltaContext[T]) Read(arrayTraits IArrayTraits[T], reader zserio.Reader) (T, error) {
	if !context.processingStarted {
		context.processingStarted = true
		context.ReadDescriptor(reader)
		return context.readUnpacked(arrayTraits, reader)
	}
	if !context.isPacked {
		return context.readUnpacked(arrayTraits, reader)
	}
	if context.maxBitNumber > 0 {
		delta, err := ReadSignedBits(reader, uint8(context.maxBitNumber+1))
		if err != nil {
			return arrayTraits.FromUint64(0), err
		}
		*context.previousElement = uint64(int64(*context.previousElement) + delta)
	}
	value := arrayTraits.FromUint64(*context.previousElement)
	return value, nil
}

// Write writes an element of an delta context array.
func (context *DeltaContext[T]) Write(arrayTraits IArrayTraits[T], writer zserio.Writer, element T) error {
	if !context.processingStarted {
		context.processingStarted = true
		context.finishInit()
		if err := context.WriteDescriptor(writer); err != nil {
			return err
		}
		return context.writeUnpacked(arrayTraits, writer, element)
	}
	if !context.isPacked {
		return context.writeUnpacked(arrayTraits, writer, element)
	}
	if context.maxBitNumber > 0 {
		delta := arrayTraits.AsUint64(element) - *context.previousElement
		err := writer.WriteBits(delta, uint8(context.maxBitNumber+1))
		if err != nil {
			return err
		}
		*context.previousElement = arrayTraits.AsUint64(element)
	}
	return nil
}

// finishInit decided if the array should be written packed or unpacked,
// depending on which variant is more space-efficient.
func (context *DeltaContext[T]) finishInit() {
	if context.isPacked {
		deltaBitsize := 0
		if context.maxBitNumber > 0 {
			deltaBitsize = context.maxBitNumber + 1
		}
		// decide if this array should be packed or not by comparing the array
		// bit sizes of both methods. Packed is usually more efficient if the
		// the array values are not differing too much from each other.
		packedBitsizeWithDescriptor := 1 + maxBitNumberBits +
			context.firstElementBitSize + (context.numElements-1)*deltaBitsize

		unpackedBitsizeWithDescriptor := 1 + context.unpackedBitSize

		if packedBitsizeWithDescriptor >= unpackedBitsizeWithDescriptor {
			context.isPacked = false
		}
	}
}

// BitSizeOfDescriptor returns the bit size of a delta context array descriptor.
func (context *DeltaContext[T]) BitSizeOfDescriptor() int {
	context.finishInit()
	if context.isPacked {
		return 1 + maxBitNumberBits
	}
	return 1
}

// ReadDescriptor reads the descriptor of a delta packed context.
func (context *DeltaContext[T]) ReadDescriptor(reader zserio.Reader) error {
	var err error
	if context.isPacked, err = reader.ReadBool(); err != nil {
		return err
	}
	if context.isPacked {
		// read how many bits are used for the delta encoding of each element
		var maxBitNumber uint64
		maxBitNumber, err = reader.ReadBits(maxBitNumberBits)
		context.maxBitNumber = int(maxBitNumber)
	}
	return err
}

// readUnpacked reads an unpacked array element from a delta context.
func (context *DeltaContext[T]) readUnpacked(arrayTraits IArrayTraits[T], reader zserio.Reader) (T, error) {
	element, err := arrayTraits.Read(reader, 0)
	if err != nil {
		return arrayTraits.FromUint64(0), err
	}

	if context.previousElement == nil {
		context.previousElement = new(uint64)
	}
	*context.previousElement = arrayTraits.AsUint64(element)

	return element, nil
}

// WriteDescriptor writes the descriptor of a delta packed context.
func (context *DeltaContext[T]) WriteDescriptor(writer zserio.Writer) error {
	if err := writer.WriteBool(context.isPacked); err != nil {
		return err
	}
	if context.isPacked {
		return writer.WriteBits(uint64(context.maxBitNumber), maxBitNumberBits)
	}
	return nil
}

// writeUnpacked writes an unpacked array element to a writer.
func (context *DeltaContext[T]) writeUnpacked(arrayTraits IArrayTraits[T], writer zserio.Writer, element T) error {
	if context.previousElement == nil {
		context.previousElement = new(uint64)
	}
	*context.previousElement = arrayTraits.AsUint64(element)

	return arrayTraits.Write(writer, element)
}

// bitsizeOfUnpacked returns the unpacked bit size of an array element.
func bitsizeOfUnpacked[T any](arrayTraits IArrayTraits[T], element T) int {
	return arrayTraits.BitSizeOf(element, 0)
}
