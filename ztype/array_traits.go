package ztype

import (
	"math"

	"golang.org/x/exp/constraints"

	zserio "github.com/woven-planet/go-zserio"
)

// IArrayTraits is the interface used by all array traits.
type IArrayTraits[T any] interface {
	// BitSizeOfIsConstant returns true if the bit size is constant for every
	// array element, for example int32.
	BitSizeOfIsConstant() bool

	// NeedsBitsizeOfPosition is true if the array traits need to know the bit
	// size of an array element.
	NeedsBitsizeOfPosition() bool

	// NeedsReadIndex specifies if the traits need the array index of the element.
	NeedsReadIndex() bool

	// PackedTraits will return the array trait as a packed object.
	PackedTraits() IPackedArrayTraits[T]

	// BitSizeOf returns the bit size of array element. If the sizee depends on
	// the position within the bit stream, the endBitPosition parameter is taken
	// into account.
	BitSizeOf(element T, endBitPosition int) int

	// InitializeOffsets returns the end bit position of an array element within
	// a byte stream.
	InitializeOffsets(bitPosition int, value T) int

	// Read reads an array element from a byte stream.
	Read(reader zserio.Reader, index int) (T, error)

	// Write writes an array element to a byte stream.
	Write(writer zserio.Writer, value T) error

	// AsUint64 returns the array element as a uint64. This is only needed if the
	// underlying data type supports delta encoding, such as int32, etc.
	// Other datatypes, such as string, float, can implement this empty.
	AsUint64(value T) uint64
	// FromUint64 returns an array element from a uint64. This is only needed
	// if the underlying data type supports delta encoding.
	FromUint64(value uint64) T
}

var ArrayTraits = staticArrayTraits{
	VarInt:    &VarIntArrayTraits{},
	VarInt16:  &VarInt16ArrayTraits{},
	VarInt32:  &VarInt32ArrayTraits{},
	VarInt64:  &VarInt64ArrayTraits{},
	VarUInt:   &VarUIntArrayTraits{},
	VarUInt16: &VarUInt16ArrayTraits{},
	VarUInt32: &VarUInt32ArrayTraits{},
	VarUInt64: &VarUInt64ArrayTraits{},
	VarSize:   &VarSizeArrayTraits{},
	Bool:      &BooleanArrayTraits{},
	String:    &StringArrayTraits{},
	Float16:   &Float16ArrayTraits{},
	Float32:   &Float32ArrayTraits{},
	Float64:   &Float64ArrayTraits{},
	Bytes:     &BytesArrayTraits{},
	Extern:    &BitBufferArrayTraits{},
	UInt8:     &BitFieldArrayTraits[uint8]{NumBits: 8},
	UInt16:    &BitFieldArrayTraits[uint16]{NumBits: 16},
	UInt32:    &BitFieldArrayTraits[uint32]{NumBits: 32},
	UInt64:    &BitFieldArrayTraits[uint64]{NumBits: 64},
	Int8:      &SignedBitFieldArrayTraits[int8]{NumBits: 8},
	Int16:     &SignedBitFieldArrayTraits[int16]{NumBits: 16},
	Int32:     &SignedBitFieldArrayTraits[int32]{NumBits: 32},
	Int64:     &SignedBitFieldArrayTraits[int64]{NumBits: 64},
}

type staticArrayTraits struct {
	VarInt    *VarIntArrayTraits
	VarInt16  *VarInt16ArrayTraits
	VarInt32  *VarInt32ArrayTraits
	VarInt64  *VarInt64ArrayTraits
	VarUInt   *VarUIntArrayTraits
	VarUInt16 *VarUInt16ArrayTraits
	VarUInt32 *VarUInt32ArrayTraits
	VarUInt64 *VarUInt64ArrayTraits
	VarSize   *VarSizeArrayTraits
	Bool      *BooleanArrayTraits
	String    *StringArrayTraits
	Float16   *Float16ArrayTraits
	Float32   *Float32ArrayTraits
	Float64   *Float64ArrayTraits
	Bytes     *BytesArrayTraits
	Extern    *BitBufferArrayTraits
	UInt8     *BitFieldArrayTraits[uint8]
	UInt16    *BitFieldArrayTraits[uint16]
	UInt32    *BitFieldArrayTraits[uint32]
	UInt64    *BitFieldArrayTraits[uint64]
	Int8      *SignedBitFieldArrayTraits[int8]
	Int16     *SignedBitFieldArrayTraits[int16]
	Int32     *SignedBitFieldArrayTraits[int32]
	Int64     *SignedBitFieldArrayTraits[int64]
}

var (
	mapSignedBitFieldArrayTraits []any = make([]any, 64)
	mapBitFieldArrayTraits       []any = make([]any, 64)
)

func GetSignedBitFieldArrayTraits[T constraints.Signed](numBits uint8) *SignedBitFieldArrayTraits[T] {
	traits := mapSignedBitFieldArrayTraits[numBits]
	if traits == nil {
		traits = &SignedBitFieldArrayTraits[T]{NumBits: uint8(numBits)}
		mapSignedBitFieldArrayTraits[numBits] = traits
	}

	return traits.(*SignedBitFieldArrayTraits[T])
}

func GetBitFieldArrayTraits[T constraints.Unsigned](numBits uint8) *BitFieldArrayTraits[T] {
	traits := mapBitFieldArrayTraits[numBits]
	if traits == nil {
		traits = &BitFieldArrayTraits[T]{NumBits: uint8(numBits)}
		mapBitFieldArrayTraits[numBits] = traits
	}

	return traits.(*BitFieldArrayTraits[T])
}

// Float16ArrayTraits is an array trait implementation for float16 arrays.
type Float16ArrayTraits struct{}

func (trait Float16ArrayTraits) PackedTraits() IPackedArrayTraits[float32] {
	return &PackedArrayTraits[float32, Float16ArrayTraits]{
		ArrayTraits: trait,
	}
}

func (trait Float16ArrayTraits) BitSizeOfIsConstant() bool {
	return true
}

func (trait Float16ArrayTraits) NeedsBitsizeOfPosition() bool {
	return false
}

func (trait Float16ArrayTraits) NeedsReadIndex() bool {
	return false
}

func (trait Float16ArrayTraits) BitSizeOf(element float32, endBitPosition int) int {
	// the bit size is always constant (16). The bit size of a float does not
	// depend on the position in the array, endBitPosition is therefore not used.
	return 16
}

func (trait Float16ArrayTraits) InitializeOffsets(bitPosition int, value float32) int {
	return bitPosition + trait.BitSizeOf(value, 0)
}

func (trait Float16ArrayTraits) Read(reader zserio.Reader, index int) (float32, error) {
	return ReadFloat16(reader)
}

func (trait Float16ArrayTraits) Write(writer zserio.Writer, value float32) error {
	return WriteFloat16(writer, value)
}

func (trait Float16ArrayTraits) AsUint64(value float32) uint64 {
	return 0 // delta encoding not supported
}

func (trait Float16ArrayTraits) FromUint64(value uint64) float32 {
	return 0.0 // delta encoding not supported
}

// Float32ArrayTraits is an array trait implementation for float32 arrays.
type Float32ArrayTraits struct{}

func (trait Float32ArrayTraits) PackedTraits() IPackedArrayTraits[float32] {
	return &PackedArrayTraits[float32, Float32ArrayTraits]{
		ArrayTraits: trait,
	}
}

func (trait Float32ArrayTraits) BitSizeOfIsConstant() bool {
	return true
}

func (trait Float32ArrayTraits) NeedsBitsizeOfPosition() bool {
	return false
}

func (trait Float32ArrayTraits) NeedsReadIndex() bool {
	return false
}

func (trait Float32ArrayTraits) BitSizeOf(element float32, endBitPosition int) int {
	return 32
}

func (trait Float32ArrayTraits) InitializeOffsets(bitPosition int, value float32) int {
	return bitPosition + trait.BitSizeOf(value, 0) // endBitPosition is ignored
}

func (trait Float32ArrayTraits) Read(reader zserio.Reader, index int) (float32, error) {
	return ReadFloat32(reader)
}

func (trait Float32ArrayTraits) Write(writer zserio.Writer, value float32) error {
	return WriteFloat32(writer, value)
}

func (trait Float32ArrayTraits) AsUint64(value float32) uint64 {
	return 0 // delta encoding not supported
}

func (trait Float32ArrayTraits) FromUint64(value uint64) float32 {
	return 0.0 // delta encoding not supported
}

// Float64ArrayTraits is an array trait implementation for float64 arrays.
type Float64ArrayTraits struct{}

func (trait Float64ArrayTraits) PackedTraits() IPackedArrayTraits[float64] {
	return &PackedArrayTraits[float64, Float64ArrayTraits]{
		ArrayTraits: trait,
	}
}

func (trait Float64ArrayTraits) BitSizeOfIsConstant() bool {
	return true
}

func (trait Float64ArrayTraits) NeedsBitsizeOfPosition() bool {
	return false
}

func (trait Float64ArrayTraits) NeedsReadIndex() bool {
	return false
}

func (trait Float64ArrayTraits) BitSizeOf(element float64, endBitPosition int) int {
	return 64
}

func (trait Float64ArrayTraits) InitializeOffsets(bitPosition int, value float64) int {
	return bitPosition + trait.BitSizeOf(value, 0)
}

func (trait Float64ArrayTraits) Read(reader zserio.Reader, index int) (float64, error) {
	return ReadFloat64(reader)
}

func (trait Float64ArrayTraits) Write(writer zserio.Writer, value float64) error {
	return WriteFloat64(writer, value)
}

func (trait Float64ArrayTraits) AsUint64(value float64) uint64 {
	return 0 // delta encoding not supported
}

func (trait Float64ArrayTraits) FromUint64(value uint64) float64 {
	return 0.0 // delta encoding not supported
}

// VarIntArrayTraits is an array traits implementation for VarInt.
type VarIntArrayTraits struct{}

func (trait VarIntArrayTraits) PackedTraits() IPackedArrayTraits[int64] {
	return &PackedArrayTraits[int64, VarIntArrayTraits]{
		ArrayTraits: trait,
	}
}

func (trait VarIntArrayTraits) BitSizeOfIsConstant() bool {
	// The bit size of varint differs, array elements therefore may have a
	// different size.
	return false
}

func (trait VarIntArrayTraits) NeedsBitsizeOfPosition() bool {
	return false
}

func (trait VarIntArrayTraits) NeedsReadIndex() bool {
	return false
}

func (trait VarIntArrayTraits) BitSizeOf(element int64, endBitPosition int) int {
	// the bit size of a VarInt only depends on the value, not on the position
	// in the stream. endBitPosition is therefore ignored.
	bitSize, _ := SignedBitSize(element, 8)
	return bitSize
}

func (trait VarIntArrayTraits) InitializeOffsets(bitPosition int, value int64) int {
	return bitPosition + trait.BitSizeOf(value, 0)
}

func (trait VarIntArrayTraits) Read(reader zserio.Reader, index int) (int64, error) {
	return ReadVarint(reader)
}

func (trait VarIntArrayTraits) Write(writer zserio.Writer, value int64) error {
	return WriteVarint(writer, value)
}

func (trait VarIntArrayTraits) AsUint64(value int64) uint64 {
	return uint64(value)
}

func (trait VarIntArrayTraits) FromUint64(value uint64) int64 {
	return int64(value)
}

// VarInt16ArrayTraits is the implementation of an array traits for a VarInt16.
type VarInt16ArrayTraits struct{}

func (trait VarInt16ArrayTraits) PackedTraits() IPackedArrayTraits[int16] {
	return &PackedArrayTraits[int16, VarInt16ArrayTraits]{
		ArrayTraits: trait,
	}
}

func (trait VarInt16ArrayTraits) BitSizeOfIsConstant() bool {
	return false
}

func (trait VarInt16ArrayTraits) NeedsBitsizeOfPosition() bool {
	return false
}

func (trait VarInt16ArrayTraits) NeedsReadIndex() bool {
	return false
}

func (trait VarInt16ArrayTraits) BitSizeOf(element int16, endBitPosition int) int {
	bitSize, _ := SignedBitSize(int64(element), 2)
	return bitSize
}

func (trait VarInt16ArrayTraits) InitializeOffsets(bitPosition int, value int16) int {
	return bitPosition + trait.BitSizeOf(value, 0) // endBitPosition is ignored
}

func (trait VarInt16ArrayTraits) Read(reader zserio.Reader, index int) (int16, error) {
	return ReadVarint16(reader)
}

func (trait VarInt16ArrayTraits) Write(writer zserio.Writer, value int16) error {
	return WriteVarint16(writer, value)
}

func (trait VarInt16ArrayTraits) AsUint64(value int16) uint64 {
	return uint64(value)
}

func (trait VarInt16ArrayTraits) FromUint64(value uint64) int16 {
	return int16(value)
}

// VarInt32ArrayTraits is the implementation of an array traits for a VarInt32.
type VarInt32ArrayTraits struct{}

func (trait VarInt32ArrayTraits) PackedTraits() IPackedArrayTraits[int32] {
	return &PackedArrayTraits[int32, VarInt32ArrayTraits]{
		ArrayTraits: trait,
	}
}

func (trait VarInt32ArrayTraits) BitSizeOfIsConstant() bool {
	return false
}

func (trait VarInt32ArrayTraits) NeedsBitsizeOfPosition() bool {
	return false
}

func (trait VarInt32ArrayTraits) NeedsReadIndex() bool {
	return false
}

func (trait VarInt32ArrayTraits) BitSizeOf(element int32, endBitPosition int) int {
	bitSize, _ := SignedBitSize(int64(element), 4)
	return bitSize
}

func (trait VarInt32ArrayTraits) InitializeOffsets(bitPosition int, value int32) int {
	return bitPosition + trait.BitSizeOf(value, 0) // endBitPosition is ignored
}

func (trait VarInt32ArrayTraits) Read(reader zserio.Reader, index int) (int32, error) {
	return ReadVarint32(reader)
}

func (trait VarInt32ArrayTraits) Write(writer zserio.Writer, value int32) error {
	return WriteVarint32(writer, value)
}

func (trait VarInt32ArrayTraits) AsUint64(value int32) uint64 {
	return uint64(value)
}

func (trait VarInt32ArrayTraits) FromUint64(value uint64) int32 {
	return int32(value)
}

// VarInt64ArrayTraits is the implementation of an array traits for a VarInt64.
type VarInt64ArrayTraits struct{}

func (trait VarInt64ArrayTraits) PackedTraits() IPackedArrayTraits[int64] {
	return &PackedArrayTraits[int64, VarInt64ArrayTraits]{
		ArrayTraits: trait,
	}
}

func (trait VarInt64ArrayTraits) BitSizeOfIsConstant() bool {
	return false
}

func (trait VarInt64ArrayTraits) NeedsBitsizeOfPosition() bool {
	return false
}

func (trait VarInt64ArrayTraits) NeedsReadIndex() bool {
	return false
}

func (trait VarInt64ArrayTraits) BitSizeOf(element int64, endBitPosition int) int {
	bitSize, _ := SignedBitSize(int64(element), 8)
	return bitSize
}

func (trait VarInt64ArrayTraits) InitializeOffsets(bitPosition int, value int64) int {
	return bitPosition + trait.BitSizeOf(value, 0) // endBitPosition is ignored
}

func (trait VarInt64ArrayTraits) Read(reader zserio.Reader, index int) (int64, error) {
	return ReadVarint64(reader)
}

func (trait VarInt64ArrayTraits) Write(writer zserio.Writer, value int64) error {
	return WriteVarint64(writer, value)
}

func (trait VarInt64ArrayTraits) AsUint64(value int64) uint64 {
	return uint64(value)
}

func (trait VarInt64ArrayTraits) FromUint64(value uint64) int64 {
	return int64(value)
}

// VarUInt16ArrayTraits is the implementation of an array traits for a VarUint16.
type VarUInt16ArrayTraits struct{}

func (trait VarUInt16ArrayTraits) PackedTraits() IPackedArrayTraits[uint16] {
	return &PackedArrayTraits[uint16, VarUInt16ArrayTraits]{
		ArrayTraits: trait,
	}
}

func (trait VarUInt16ArrayTraits) BitSizeOfIsConstant() bool {
	return false
}

func (trait VarUInt16ArrayTraits) NeedsBitsizeOfPosition() bool {
	return false
}

func (trait VarUInt16ArrayTraits) NeedsReadIndex() bool {
	return false
}

func (trait VarUInt16ArrayTraits) BitSizeOf(element uint16, endBitPosition int) int {
	bitSize, _ := UnsignedBitSize(uint64(element), 4)
	return bitSize
}

func (trait VarUInt16ArrayTraits) InitializeOffsets(bitPosition int, value uint16) int {
	return bitPosition + trait.BitSizeOf(value, 0) // endBitPosition is ignored
}

func (trait VarUInt16ArrayTraits) Read(reader zserio.Reader, index int) (uint16, error) {
	return ReadVaruint16(reader)
}

func (trait VarUInt16ArrayTraits) Write(writer zserio.Writer, value uint16) error {
	return WriteVaruint16(writer, value)
}

func (trait VarUInt16ArrayTraits) AsUint64(value uint16) uint64 {
	return uint64(value)
}

func (trait VarUInt16ArrayTraits) FromUint64(value uint64) uint16 {
	return uint16(value)
}

type VarUInt32ArrayTraits struct{}

func (trait VarUInt32ArrayTraits) PackedTraits() IPackedArrayTraits[uint32] {
	return &PackedArrayTraits[uint32, VarUInt32ArrayTraits]{
		ArrayTraits: trait,
	}
}

func (trait VarUInt32ArrayTraits) BitSizeOfIsConstant() bool {
	return false
}

func (trait VarUInt32ArrayTraits) NeedsBitsizeOfPosition() bool {
	return false
}

func (trait VarUInt32ArrayTraits) NeedsReadIndex() bool {
	return false
}

func (trait VarUInt32ArrayTraits) BitSizeOf(element uint32, endBitPosition int) int {
	bitSize, _ := UnsignedBitSize(uint64(element), 4)
	return bitSize
}

func (trait VarUInt32ArrayTraits) InitializeOffsets(bitPosition int, value uint32) int {
	return bitPosition + trait.BitSizeOf(value, 0) // endBitPosition is ignored
}

func (trait VarUInt32ArrayTraits) Read(reader zserio.Reader, index int) (uint32, error) {
	return ReadVaruint32(reader)
}

func (trait VarUInt32ArrayTraits) Write(writer zserio.Writer, value uint32) error {
	return WriteVaruint32(writer, value)
}

func (trait VarUInt32ArrayTraits) AsUint64(value uint32) uint64 {
	return uint64(value)
}

func (trait VarUInt32ArrayTraits) FromUint64(value uint64) uint32 {
	return uint32(value)
}

type VarUInt64ArrayTraits struct{}

func (trait VarUInt64ArrayTraits) PackedTraits() IPackedArrayTraits[uint64] {
	return &PackedArrayTraits[uint64, VarUInt64ArrayTraits]{
		ArrayTraits: trait,
	}
}

func (trait VarUInt64ArrayTraits) BitSizeOfIsConstant() bool {
	return false
}

func (trait VarUInt64ArrayTraits) NeedsBitsizeOfPosition() bool {
	return false
}

func (trait VarUInt64ArrayTraits) NeedsReadIndex() bool {
	return false
}

func (trait VarUInt64ArrayTraits) BitSizeOf(element uint64, endBitPosition int) int {
	bitSize, _ := UnsignedBitSize(uint64(element), 8)
	return bitSize
}

func (trait VarUInt64ArrayTraits) InitializeOffsets(bitPosition int, value uint64) int {
	return bitPosition + trait.BitSizeOf(value, 0) // endBitPosition is ignored
}

func (trait VarUInt64ArrayTraits) Read(reader zserio.Reader, index int) (uint64, error) {
	return ReadVaruint64(reader)
}

func (trait VarUInt64ArrayTraits) Write(writer zserio.Writer, value uint64) error {
	return WriteVaruint64(writer, value)
}

func (trait VarUInt64ArrayTraits) AsUint64(value uint64) uint64 {
	return value
}

func (trait VarUInt64ArrayTraits) FromUint64(value uint64) uint64 {
	return value
}

type VarUIntArrayTraits struct{}

func (trait VarUIntArrayTraits) PackedTraits() IPackedArrayTraits[uint64] {
	return &PackedArrayTraits[uint64, VarUIntArrayTraits]{
		ArrayTraits: trait,
	}
}

func (trait VarUIntArrayTraits) BitSizeOfIsConstant() bool {
	return false
}

func (trait VarUIntArrayTraits) NeedsBitsizeOfPosition() bool {
	return false
}

func (trait VarUIntArrayTraits) NeedsReadIndex() bool {
	return false
}

func (trait VarUIntArrayTraits) BitSizeOf(element uint64, endBitPosition int) int {
	bitSize, _ := UnsignedBitSize(element, 8)
	return bitSize
}

func (trait VarUIntArrayTraits) InitializeOffsets(bitPosition int, value uint64) int {
	return bitPosition + trait.BitSizeOf(value, 0) // endBitPosition is ignored
}

func (trait VarUIntArrayTraits) Read(reader zserio.Reader, index int) (uint64, error) {
	return ReadVaruint(reader)
}

func (trait VarUIntArrayTraits) Write(writer zserio.Writer, value uint64) error {
	return WriteVaruint(writer, value)
}

func (trait VarUIntArrayTraits) AsUint64(value uint64) uint64 {
	return value
}

func (trait VarUIntArrayTraits) FromUint64(value uint64) uint64 {
	return value
}

type VarSizeArrayTraits struct{}

func (trait VarSizeArrayTraits) PackedTraits() IPackedArrayTraits[uint64] {
	return &PackedArrayTraits[uint64, VarSizeArrayTraits]{
		ArrayTraits: trait,
	}
}

func (trait VarSizeArrayTraits) BitSizeOfIsConstant() bool {
	return false
}

func (trait VarSizeArrayTraits) NeedsBitsizeOfPosition() bool {
	return false
}

func (trait VarSizeArrayTraits) NeedsReadIndex() bool {
	return false
}

func (trait VarSizeArrayTraits) BitSizeOf(element uint64, endBitPosition int) int {
	bitSize, _ := UnsignedBitSize(element, 5)
	return bitSize
}

func (trait VarSizeArrayTraits) InitializeOffsets(bitPosition int, value uint64) int {
	return bitPosition + trait.BitSizeOf(value, 0) // endBitPosition is ignored
}

func (trait VarSizeArrayTraits) Read(reader zserio.Reader, index int) (uint64, error) {
	return ReadVarsize(reader)
}

func (trait VarSizeArrayTraits) Write(writer zserio.Writer, value uint64) error {
	return WriteVarsize(writer, value)
}

func (trait VarSizeArrayTraits) AsUint64(value uint64) uint64 {
	return value
}

func (trait VarSizeArrayTraits) FromUint64(value uint64) uint64 {
	return value
}

type BitFieldArrayTraits[T constraints.Integer] struct {
	NumBits uint8
}

func (trait BitFieldArrayTraits[T]) PackedTraits() IPackedArrayTraits[T] {
	return &PackedArrayTraits[T, BitFieldArrayTraits[T]]{
		ArrayTraits: trait,
	}
}

func (trait BitFieldArrayTraits[T]) BitSizeOfIsConstant() bool {
	return true
}

func (trait BitFieldArrayTraits[T]) NeedsBitsizeOfPosition() bool {
	return false
}

func (trait BitFieldArrayTraits[T]) NeedsReadIndex() bool {
	return false
}

func (trait BitFieldArrayTraits[T]) BitSizeOf(element T, endBitPosition int) int {
	return int(trait.NumBits)
}

func (trait BitFieldArrayTraits[T]) InitializeOffsets(bitPosition int, value T) int {
	return bitPosition + trait.BitSizeOf(value, 0) // endBitPosition is ignored
}

func (trait BitFieldArrayTraits[T]) Read(reader zserio.Reader, index int) (T, error) {
	value, err := reader.ReadBits(uint8(trait.NumBits))
	return T(value), err
}

func (trait BitFieldArrayTraits[T]) Write(writer zserio.Writer, value T) error {
	return writer.WriteBits(uint64(value), trait.NumBits)
}

func (trait BitFieldArrayTraits[T]) AsUint64(value T) uint64 {
	return uint64(value)
}

func (trait BitFieldArrayTraits[T]) FromUint64(value uint64) T {
	return T(value)
}

type SignedBitFieldArrayTraits[T constraints.Signed] struct {
	NumBits uint8
}

func fillUpperBits(value uint64, numBits uint8) uint64 {
	// check if the value is negative. If yes,
	// it will be later casted to the correct final type,
	// and the uint64 needs to be made negative.
	isNegative := (value & (1 << uint64(numBits-1))) != 0
	if isNegative {
		// if yes, mask all bits from NumBits to 64 with a 1, to make
		// the uint64 negative when casting to an int type
		upperMask := uint64(math.MaxUint64) - (1<<uint64(numBits) - 1)
		value = value | upperMask
	}
	return value
}

func (trait SignedBitFieldArrayTraits[T]) PackedTraits() IPackedArrayTraits[T] {
	return &PackedArrayTraits[T, SignedBitFieldArrayTraits[T]]{
		ArrayTraits: trait,
	}
}

func (trait SignedBitFieldArrayTraits[T]) BitSizeOfIsConstant() bool {
	return true
}

func (trait SignedBitFieldArrayTraits[T]) NeedsBitsizeOfPosition() bool {
	return false
}

func (trait SignedBitFieldArrayTraits[T]) NeedsReadIndex() bool {
	return false
}

func (trait SignedBitFieldArrayTraits[T]) BitSizeOf(element T, endBitPosition int) int {
	return int(trait.NumBits)
}

func (trait SignedBitFieldArrayTraits[T]) InitializeOffsets(bitPosition int, value T) int {
	return bitPosition + trait.BitSizeOf(value, 0) // endBitPosition is ignored
}

func (trait SignedBitFieldArrayTraits[T]) Read(reader zserio.Reader, index int) (T, error) {
	value, err := reader.ReadBits(uint8(trait.NumBits))
	return T(fillUpperBits(value, trait.NumBits)), err
}

func (trait SignedBitFieldArrayTraits[T]) Write(writer zserio.Writer, value T) error {
	return writer.WriteBits(uint64(value), trait.NumBits)
}

func (trait SignedBitFieldArrayTraits[T]) AsUint64(value T) uint64 {
	return uint64(value)
}

func (trait SignedBitFieldArrayTraits[T]) FromUint64(value uint64) T {
	return T(fillUpperBits(value, trait.NumBits))
}

type BooleanArrayTraits struct{}

func (trait BooleanArrayTraits) PackedTraits() IPackedArrayTraits[bool] {
	return &PackedArrayTraits[bool, BooleanArrayTraits]{
		ArrayTraits: trait,
	}
}

func (trait BooleanArrayTraits) BitSizeOfIsConstant() bool {
	return true
}

func (trait BooleanArrayTraits) NeedsBitsizeOfPosition() bool {
	return false
}

func (trait BooleanArrayTraits) NeedsReadIndex() bool {
	return false
}

func (trait BooleanArrayTraits) BitSizeOf(element bool, endBitPosition int) int {
	// The bit size is always constant (1). The bit size of a boolean does not
	// depend on the position in the array, endBitPosition is therefore not used.
	return 1
}

func (trait BooleanArrayTraits) InitializeOffsets(bitPosition int, value bool) int {
	return bitPosition + trait.BitSizeOf(value, 0) // endBitPosition is ignored
}

func (trait BooleanArrayTraits) Read(reader zserio.Reader, index int) (bool, error) {
	return ReadBool(reader)
}

func (trait BooleanArrayTraits) Write(writer zserio.Writer, value bool) error {
	return WriteBool(writer, value)
}

func (trait BooleanArrayTraits) AsUint64(value bool) uint64 {
	if value {
		return 1
	}

	return 0
}

func (trait BooleanArrayTraits) FromUint64(value uint64) bool {
	return value != 0
}

type StringArrayTraits struct{}

func (trait StringArrayTraits) PackedTraits() IPackedArrayTraits[string] {
	return &PackedArrayTraits[string, StringArrayTraits]{
		ArrayTraits: trait,
	}
}

func (trait StringArrayTraits) BitSizeOfIsConstant() bool {
	return false
}

func (trait StringArrayTraits) NeedsBitsizeOfPosition() bool {
	return false
}

func (trait StringArrayTraits) NeedsReadIndex() bool {
	return false
}

func (trait StringArrayTraits) BitSizeOf(element string, endBitPosition int) int {
	byteLength := len(element)
	bitSizeOfLength, err := SignedBitSize(int64(byteLength), 4)
	if err != nil {
		return 0
	}

	// the total size of a string is 8 bits for each (UTF-8) character, and a
	// varsize type for the number of characters
	return bitSizeOfLength + byteLength*8
}

func (trait StringArrayTraits) InitializeOffsets(bitPosition int, value string) int {
	return bitPosition + trait.BitSizeOf(value, 0) // endBitPosition is ignored
}

func (trait StringArrayTraits) Read(reader zserio.Reader, index int) (string, error) {
	return ReadString(reader)
}

func (trait StringArrayTraits) Write(writer zserio.Writer, value string) error {
	return WriteString(writer, value)
}

func (trait StringArrayTraits) AsUint64(value string) uint64 {
	return 0 // not supported for string objects
}

func (trait StringArrayTraits) FromUint64(value uint64) string {
	return "" // not supported for string objects
}

type BitBufferArrayTraits struct{}

func (trait BitBufferArrayTraits) PackedTraits() IPackedArrayTraits[*ExternType] {
	return &PackedArrayTraits[*ExternType, BitBufferArrayTraits]{
		ArrayTraits: trait,
	}
}

func (trait BitBufferArrayTraits) BitSizeOfIsConstant() bool {
	return false
}

func (trait BitBufferArrayTraits) NeedsBitsizeOfPosition() bool {
	return false
}

func (trait BitBufferArrayTraits) NeedsReadIndex() bool {
	return false
}

func (trait BitBufferArrayTraits) BitSizeOf(element *ExternType, endBitPosition int) int {
	return int(element.BitSize)
}

func (trait BitBufferArrayTraits) InitializeOffsets(bitPosition int, value *ExternType) int {
	return bitPosition + trait.BitSizeOf(value, 0) // endBitPosition is ignored
}

func (trait BitBufferArrayTraits) Read(reader zserio.Reader, endBitPosition int) (*ExternType, error) {
	return ReadExtern(reader)
}

func (trait BitBufferArrayTraits) Write(writer zserio.Writer, value *ExternType) error {
	return WriteExtern(writer, value)
}

func (trait BitBufferArrayTraits) AsUint64(value *ExternType) uint64 {
	return 0 // not supported for Extern type objects
}

func (trait BitBufferArrayTraits) FromUint64(value uint64) *ExternType {
	return nil // not supported for Extern type objects
}

type BytesArrayTraits struct{}

func (trait BytesArrayTraits) PackedTraits() IPackedArrayTraits[*BytesType] {
	return &PackedArrayTraits[*BytesType, BytesArrayTraits]{
		ArrayTraits: trait,
	}
}

func (trait BytesArrayTraits) BitSizeOfIsConstant() bool {
	return false
}

func (trait BytesArrayTraits) NeedsBitsizeOfPosition() bool {
	return false
}

func (trait BytesArrayTraits) NeedsReadIndex() bool {
	return false
}

func (trait BytesArrayTraits) BitSizeOf(element *BytesType, endBitPosition int) int {
	return int(element.ByteSize) * 8
}

func (trait BytesArrayTraits) InitializeOffsets(bitPosition int, value *BytesType) int {
	return bitPosition + trait.BitSizeOf(value, 0) // endBitPosition is ignored
}

func (trait BytesArrayTraits) Read(reader zserio.Reader, endBitPosition int) (*BytesType, error) {
	return ReadBytes(reader)
}

func (trait BytesArrayTraits) Write(writer zserio.Writer, value *BytesType) error {
	return WriteBytes(writer, value)
}

func (trait BytesArrayTraits) AsUint64(value *BytesType) uint64 {
	return 0 // not supported for Extern type objects
}

func (trait BytesArrayTraits) FromUint64(value uint64) *BytesType {
	return nil // not supported for Extern type objects
}
