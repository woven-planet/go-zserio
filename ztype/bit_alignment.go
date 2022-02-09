package ztype

import (
	"github.com/icza/bitio"
)

// AlignReader aligns a bitio reader to a bit boundary.
func AlignReader(r *bitio.CountReader, boundary uint8) {
	misaligned := uint8((r.BitsCount % int64(boundary)))
	if misaligned != 0 {
		numOfBitsToDiscard := boundary - misaligned
		r.ReadBits(numOfBitsToDiscard)
	}
}

// AlignWriter aligns a bitio writer to a bit boundary.
func AlignWriter(w *bitio.CountWriter, boundary uint8) error {
	misaligned := uint8((w.BitsCount % int64(boundary)))
	if misaligned != 0 {
		numOfBitsToDiscard := boundary - misaligned
		return w.WriteBits(0, numOfBitsToDiscard)
	}
	return nil
}

// CountAlignBits calculates how many bits would need to be spend to align the
// current bitPosition to the next boundary.
func CountAlignBits(bitPosition int, boundary uint8) int {
	misaligned := (bitPosition % int(boundary))
	if misaligned != 0 {
		return int(boundary) - misaligned
	}
	return 0
}

func alignTo(alignmentValue, bitPosition int) int {
	if bitPosition <= 0 || alignmentValue == 0 {
		return bitPosition
	}
	return (((bitPosition - 1) / alignmentValue) + 1) * alignmentValue
}
