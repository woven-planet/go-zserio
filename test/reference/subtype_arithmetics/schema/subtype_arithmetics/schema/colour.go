// Code generated by go-zserio. DO NOT EDIT.

package schema

import (
	"errors"
	"fmt"
	zserio "github.com/woven-planet/go-zserio"
	"github.com/woven-planet/go-zserio/ztype"
)

type Colour uint32

const (
	ColourRed    Colour = 0
	ColourPurple Colour = 1
	ColourBlue   Colour = 2
	ColourOrange Colour = 3
	ColourCyan   Colour = 4
)

var _ColourNames = []string{
	"RED",
	"PURPLE",
	"BLUE",
	"ORANGE",
	"CYAN",
}

var _ColourValues = []Colour{
	ColourRed,
	ColourPurple,
	ColourBlue,
	ColourOrange,
	ColourCyan,
}

var _ColourName = map[Colour]string{
	ColourRed:    "RED",
	ColourPurple: "PURPLE",
	ColourBlue:   "BLUE",
	ColourOrange: "ORANGE",
	ColourCyan:   "CYAN",
}

var _ColourNameToValue = map[string]Colour{
	"RED":    ColourRed,
	"PURPLE": ColourPurple,
	"BLUE":   ColourBlue,
	"ORANGE": ColourOrange,
	"CYAN":   ColourCyan,
}

// ColourValues returns all values of the enum
func ColourValues() []Colour {
	return _ColourValues
}

// ColourStrings returns a slice of all string values of the enum
func ColourStrings() []string {
	strs := make([]string, len(_ColourNames))
	copy(strs, _ColourNames)
	return strs
}

// ColourString converts a string to an enum value. If the string
// does not match a known value an error is returned.
func ColourString(s string) (Colour, error) {
	if value, found := _ColourNameToValue[s]; found {
		return value, nil
	}

	var zero Colour
	return zero, fmt.Errorf("%s does not belong to Colour values", s)
}

// String implements fmt.Stringer
func (v Colour) String() string {
	if name, found := _ColourName[v]; found {
		return name
	}
	return fmt.Sprintf("unknown(%v)", uint32(v))
}

func (v *Colour) Clone() zserio.ZserioType {
	clone := *v
	return &clone
}

func (v *Colour) LoadDefaultValues() error {
	return nil
}

// MarshalZserio implements the zserio.Marshaler interface.
func (v *Colour) MarshalZserio(w zserio.Writer) error {
	return ztype.WriteUint32(w, uint32(*v))
}

// UnmarshalZserio implements the zserio.Unmarshaler interface.
func (v *Colour) UnmarshalZserio(r zserio.Reader) error {
	if value, err := ztype.ReadUint32(r); err != nil {
		return err
	} else {
		*v = Colour(value)
	}
	return nil
}

// ZserioBitSize implements the zserio.Marshaler interface.
func (v *Colour) ZserioBitSize(bitPosition int) (int, error) {
	endBitPosition := bitPosition
	endBitPosition += 32

	return endBitPosition - bitPosition, nil
}

func (v *Colour) ZserioCreatePackingContext(contextNode *zserio.PackingContextNode) error {
	contextNode.Context = &ztype.DeltaContext[uint32]{}
	return nil
}

func (v *Colour) ZserioInitPackingContext(contextNode *zserio.PackingContextNode) error {
	if !contextNode.HasContext() {
		return errors.New("context node has no packing")
	}
	context, ok := contextNode.Context.(*ztype.DeltaContext[uint32])
	if !ok {
		return errors.New("unsupported packing context type")
	}

	traits := ztype.BitFieldArrayTraits[uint32]{NumBits: uint8(32)}
	context.Init(&traits, uint32(*v))
	return nil
}

func (v *Colour) UnmarshalZserioPacked(contextNode *zserio.PackingContextNode, r zserio.Reader) error {
	context, ok := contextNode.Context.(*ztype.DeltaContext[uint32])
	if !ok {
		return errors.New("unsupported packing context type")
	}

	traits := ztype.BitFieldArrayTraits[uint32]{NumBits: uint8(32)}
	if tempValue, err := context.Read(&traits, r); err != nil {
		return err
	} else {
		(*v) = Colour(tempValue)
		return nil
	}
}

func (v *Colour) MarshalZserioPacked(contextNode *zserio.PackingContextNode, w zserio.Writer) error {
	context, ok := contextNode.Context.(*ztype.DeltaContext[uint32])
	if !ok {
		return errors.New("unsupported packing context type")
	}

	traits := ztype.BitFieldArrayTraits[uint32]{NumBits: uint8(32)}
	return context.Write(&traits, w, uint32(*v))
}

func (v *Colour) ZserioInitializeOffsetsPacked(contextNode *zserio.PackingContextNode, bitPosition int) int {
	return 0
}

func (v *Colour) ZserioBitSizePacked(contextNode *zserio.PackingContextNode, bitPosition int) (int, error) {
	context, ok := contextNode.Context.(*ztype.DeltaContext[uint32])
	if !ok {
		return 0, errors.New("invalid packing context")
	}
	delta, err := context.BitSizeOf(&ztype.BitFieldArrayTraits[uint32]{NumBits: uint8(32)}, bitPosition, uint32(*v))
	if err != nil {
		return 0, err
	}
	return delta, nil
}
