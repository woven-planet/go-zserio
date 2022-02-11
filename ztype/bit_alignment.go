package ztype

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
