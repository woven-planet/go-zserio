package ast

import (
	"errors"
	"fmt"
)

var ErrUnknownType = errors.New("unknown type")

var rootScope = &RootScope{}

var zserioTypeToGoType = map[string]string{
	// integer types
	"int8":   "int8",
	"int16":  "int16",
	"int32":  "int32",
	"int64":  "int64",
	"uint8":  "uint8",
	"uint16": "uint16",
	"uint32": "uint32",
	"uint64": "uint64",
	// varint types
	"varint":    "int64",
	"varint16":  "int16",
	"varint32":  "int32",
	"varint64":  "int64",
	"varuint":   "uint64",
	"varsize":   "uint64",
	"varuint16": "uint16",
	"varuint32": "uint32",
	"varuint64": "uint64",
	// bool types
	"bool": "bool",
	// string types
	"string": "string",
	// float types
	"float16": "float32",
	"float32": "float32",
	"float64": "float64",
}

var zserioTypeToArrayTraits = map[string]string{
	// integer types
	"int8":   "ztype.SignedBitFieldArrayTraits",
	"int16":  "ztype.SignedBitFieldArrayTraits",
	"int32":  "ztype.SignedBitFieldArrayTraits",
	"int64":  "ztype.SignedBitFieldArrayTraits",
	"uint8":  "ztype.BitFieldArrayTraits",
	"uint16": "ztype.BitFieldArrayTraits",
	"uint32": "ztype.BitFieldArrayTraits",
	"uint64": "ztype.BitFieldArrayTraits",
	"int":    "ztype.SignedBitFieldArrayTraits",
	"uint":   "ztype.BitFieldArrayTraits",
	"bit":    "ztype.BitFieldArrayTraits",
	// varint types
	"varint":    "ztype.VarIntArrayTraits",
	"varint16":  "ztype.VarInt16ArrayTraits",
	"varint32":  "ztype.VarInt32ArrayTraits",
	"varint64":  "ztype.VarInt64ArrayTraits",
	"varuint":   "ztype.VarUIntArrayTraits",
	"varsize":   "ztype.VarSizeArrayTraits",
	"varuint16": "ztype.VarUInt16ArrayTraits",
	"varuint32": "ztype.VarUInt32ArrayTraits",
	"varuint64": "ztype.VarUInt64ArrayTraits",
	// bool types
	"bool": "ztype.BooleanArrayTraits",
	// string types
	"string": "ztype.StringArrayTraits",
	// float types
	"float16": "ztype.Float16ArrayTraits",
	"float32": "ztype.Float32ArrayTraits",
	"float64": "ztype.Float64ArrayTraits",
	// other
	"extern": "ztype.ObjectArrayTraits[*ztype.ExternType]",
	"bytes":  "ztype.ObjectArrayTraits[*ztype.BytesType]",
}

// zserioDeltaPackabeTypes are all native types that can be delta-packed
var zserioDeltaPackabeTypes = map[string]bool{
	// integer types
	"int8":   true,
	"int16":  true,
	"int32":  true,
	"int64":  true,
	"uint8":  true,
	"uint16": true,
	"uint32": true,
	"uint64": true,
	"int":    true,
	"uint":   true,
	"bit":    true,
	// varint types
	"varint":    true,
	"varint16":  true,
	"varint32":  true,
	"varint64":  true,
	"varuint":   true,
	"varsize":   true,
	"varuint16": true,
	"varuint32": true,
	"varuint64": true,
}

type Scope interface {
	// GoType looks up a name in the scope, and provides the Go type name for it.
	// If the type is not found ErrUnkownType is returned.
	GoType(t *TypeReference) (string, error)

	// GoArrayTraits looks up a name in the scope, and provides the array traits
	// for it.
	GoArrayTraits(t *TypeReference) (string, error)

	// IsDeltaPackable returns if a data type can be delta-packed: when stored in an
	// array, a delta compression can be applied, which only stored the delta to
	// the previous value, not the entire value. Only works for integer types.
	IsDeltaPackable(t *TypeReference) (bool, error)

	// HasType checks if a type is defined *in this scope*. It does not check parent
	// scopes.
	HasType(t string) bool
}

type BaseScope struct {
	Parent Scope
}

// HasType checks if a type is defined *in this scope*. It does not check parent
// scopes.
func (s *BaseScope) HasType(t string) bool {
	return false
}

// GoType looks up a name in the scope, and provides the Go type name for it.
// If the type is not found ErrUnkownType is returned.
func (s *BaseScope) GoType(t *TypeReference) (string, error) {
	if s.Parent != nil {
		return s.Parent.GoType(t)
	}
	return "", fmt.Errorf("%w: %s", ErrUnknownType, t.Name)
}

func (s *BaseScope) GoArrayTraits(t *TypeReference) (string, error) {
	if s.Parent != nil {
		return s.Parent.GoArrayTraits(t)
	}
	return "", fmt.Errorf("%w: %s", ErrUnknownType, t.Name)
}

// RootScope is the root zserio scope, which contains all built-in types
type RootScope struct{}

func goBits(b int) string {
	if b == 0 {
		return ""
	}
	if b <= 8 {
		return "8"
	}
	if b <= 16 {
		return "16"
	}
	if b <= 32 {
		return "32"
	}
	return "64"
}

func (s *RootScope) GoType(t *TypeReference) (string, error) {
	if !t.IsBuiltin {
		return "", ErrUnknownType
	}

	// check for bit field types
	if t.Name == "int" {
		if t.LengthExpression != nil {
			// length is evaluated at runtime
			return "int64", nil
		}
		return "int" + goBits(t.Bits), nil
	}

	if t.Name == "bit" || t.Name == "uint" {
		if t.LengthExpression != nil {
			// length is evaluated at runtime
			return "uint64", nil
		}
		return "uint" + goBits(t.Bits), nil
	}

	// Extern and Bytes are both zserio internal types, that store
	// bit and byte buffers, respectively.
	if t.Name == "extern" {
		return "*ztype.ExternType", nil
	}

	if t.Name == "bytes" {
		return "*ztype.BytesType", nil
	}

	if goType, found := zserioTypeToGoType[t.Name]; found {
		return goType, nil
	}
	return "", fmt.Errorf("%w: %s", ErrUnknownType, t.Name)
}

// GoArrayTraits returns the array traits for a zserio basic type (int, float, ...)
func (s *RootScope) GoArrayTraits(t *TypeReference) (string, error) {
	if !t.IsBuiltin {
		return "", ErrUnknownType
	}

	if arrayTrait, found := zserioTypeToArrayTraits[t.Name]; found {
		return arrayTrait, nil
	}
	return "", fmt.Errorf("%w: %s", ErrUnknownType, t.Name)
}

// IsDeltaPackable reutrns if a zserio basic type is delta-packabe. If the type
// is not a basic type, the function will return an error.
func (s *RootScope) IsDeltaPackable(t *TypeReference) (bool, error) {
	if !t.IsBuiltin {
		return false, ErrUnknownType
	}
	return zserioDeltaPackabeTypes[t.Name], nil
}

// HasType checks if a type is defined *in this scope*. It does not check parent
// scopes.
func (s *RootScope) HasType(t string) bool {
	if t == "int" || t == "bit" || t == "uint" || t == "extern" || t == "bytes" {
		return true
	}
	_, found := zserioTypeToGoType[t]
	return found
}

type CompoundScope struct {
	Symbol any
	Scope  map[string]any
}

// SymbolScope keeps track of all symbols that are available in this scope
type SymbolScope struct {
	// TypeScope contains all types that are currently in the scope
	TypeScope map[string]any

	// OtherScope contains all symbols that are not Types (e.g. enum values)
	OtherScope map[string]any

	// CompoundScope contains all symbols that are restricted to the current
	// evaluation scope (e.g. the struct parameters when evaluating a struct)
	CompoundScopes map[string]*CompoundScope

	CurrentCompoundScope *string
}

func (v *SymbolScope) GetType(name string) (*SymbolReference, error) {
	typeSymbol, found := v.TypeScope[name]
	if found {
		return &SymbolReference{
			Name:   name,
			Symbol: typeSymbol,
		}, nil
	}
	return nil, errors.New("type not found")
}

func (v *SymbolScope) GetCompoundType(compoundType, name string) (*SymbolReference, error) {
	if compoundScope, found := v.CompoundScopes[compoundType]; found {
		if currentScopeSymbol, found := compoundScope.Scope[name]; found {
			return &SymbolReference{
				Name:         name,
				CompoundName: compoundType,
				Symbol:       currentScopeSymbol,
			}, nil
		}
	}
	return nil, errors.New("compound type not found")
}

func (v *SymbolScope) GetOtherType(name string) (*SymbolReference, error) {
	otherSymbol, found := v.OtherScope[name]
	if found {
		return &SymbolReference{
			Name:   name,
			Symbol: otherSymbol,
		}, nil
	}

	return nil, errors.New("other type not found")
}
