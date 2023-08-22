package ast

import "fmt"

var zserioIntegerTypeRank = map[string]int{
	// integer types
	"int8":   8,
	"int16":  16,
	"int32":  32,
	"int64":  64,
	"uint8":  8,
	"uint16": 16,
	"uint32": 32,
	"uint64": 64,
	// varint types
	"varint":    64, // a varint64 can store the same number of bits as int64
	"varint16":  15,
	"varint32":  29,
	"varint64":  57,
	"varuint":   64,
	"varsize":   31,
	"varuint16": 15,
	"varuint32": 29,
	"varuint64": 64,
}

var zserioSignedToUnsigned = map[string]string{
	// integer types
	"int8":  "uint8",
	"int16": "uint16",
	"int32": "uint32",
	"int64": "uint64",
	"int":   "bit",
	// varint types
	"varint":   "varuint",
	"varint16": "varuint16",
	"varint32": "varuint32",
	"varint64": "varuint64",
}

// DetermineArithmeticOperationResultType determines the result type of an expression
// that involves two operands in an arithmetic operation. For example, when adding
// integers with floats, the resulting type should be float, and how to handle operations
// that mix unsigned/signed types, or types of different bit length.
func DetermineArithmeticOperationResultType(operand1 *Expression, operand2 *Expression) (ExpressionType, *TypeReference, error) {
	// Short-circuit if both operands have the same type.
	if operand1.ResultType == operand2.ResultType && compareNativeTypeReferences(operand1.NativeZserioType, operand2.NativeZserioType) {
		return operand1.ResultType, operand1.NativeZserioType, nil
	}

	// check mixing floats with integers - the result type will always be a float type.
	if operand1.ResultType == ExpressionTypeFloat || operand2.ResultType == ExpressionTypeFloat {
		resultType, err := evaluateMixingFloatTypes(operand1.NativeZserioType, operand2.NativeZserioType)
		return ExpressionTypeFloat, resultType, err
	}

	// If the native types are not float types, they must be integer types.
	if operand1.ResultType != ExpressionTypeInteger || operand2.ResultType != ExpressionTypeInteger {
		return "", nil, fmt.Errorf("%w: %s, %s", ErrUnknownType, operand1.ResultType, operand2.ResultType)
	}

	resultType, err := evaluateMixingIntegerTypes(operand1.NativeZserioType, operand2.NativeZserioType)
	return ExpressionTypeInteger, resultType, err
}

// compareNativeTypeReferences compares two native type references.
// The type references can be nil, in which case the evaluation returns true if
// both type references are nil.
func compareNativeTypeReferences(type1 *TypeReference, type2 *TypeReference) bool {
	if type1 == nil && type2 == nil {
		return true
	}
	if type1 == nil || type2 == nil {
		return false
	}
	return type1.Name == type2.Name && type1.Bits == type2.Bits && type1.LengthExpression == nil && type2.LengthExpression == nil
}

// evaluateMixingFloatTypes determines the type of an expression, if operands are mixed, and one of them
// is of float type. An example is adding a float32 to an int64. The float32 has the higher priority,
// and the result type will be float32.
func evaluateMixingFloatTypes(type1 *TypeReference, type2 *TypeReference) (*TypeReference, error) {
	if compareNativeTypeReferences(type1, type2) {
		return type1, nil
	}
	// Note that the types may be nil - this is the case for hard-coded numerals,
	// e.g. "17.5". It is not clear if their type is float16, float32 or float64.
	// Their type needs to be deduced from the other operand, if set.
	if type1 == nil {
		return type2, nil
	}
	if type2 == nil {
		return type1, nil
	}

	// When mixing two float types, pick the larger one as the result.
	for _, float_type := range []string{"float64", "float32", "float16"} {
		if type1.Name == float_type || type2.Name == float_type {
			return &TypeReference{
				IsBuiltin: true,
				Name:      float_type,
			}, nil
		}
	}
	return nil, fmt.Errorf("%w: %s, %s", ErrUnknownType, type1.Name, type2.Name)
}

// getIntegerTypeRank returns the rank of an integer (the number of bits).
func getIntegerTypeRank(integerType *TypeReference) (int, error) {
	// if the bit length is dynamic, always assume 64bits
	if integerType.LengthExpression != nil {
		return 64, nil
	}
	if integerType.Bits != 0 {
		return integerType.Bits, nil
	}
	if rank, found := zserioIntegerTypeRank[integerType.Name]; found {
		return rank, nil
	}
	return 0, fmt.Errorf("%w: %s", ErrUnknownType, integerType.Name)
}

func evaluateMixingIntegerTypes(type1 *TypeReference, type2 *TypeReference) (*TypeReference, error) {
	// For now, we use the same logic as C++ when mixing integer types:
	// Mixing signed/unsigned values of the same type (in the sequence of int16, int32, int64),
	// the unsigned type will win.
	// Mixing different ranks result in the higher rank being used.
	if compareNativeTypeReferences(type1, type2) {
		return type1, nil
	}

	// Handle cases such as "u16Var + 12"
	if type1 == nil {
		return type2, nil
	}
	if type2 == nil {
		return type1, nil
	}

	// Identify if the resulting integer type should be signed or unsigned.
	_, op1IsSigned := zserioSignedToUnsigned[type1.Name]
	_, op2IsSigned := zserioSignedToUnsigned[type2.Name]
	resultIsSigned := op1IsSigned && op2IsSigned

	op1Rank, err := getIntegerTypeRank(type1)
	if err != nil {
		return nil, err
	}
	op2Rank, err := getIntegerTypeRank(type2)
	if err != nil {
		return nil, err
	}

	// Pick the higher rank.
	newType := &TypeReference{}
	if op1Rank >= op2Rank {
		if op1IsSigned == resultIsSigned {
			return type1, nil
		}
		*newType = *type1
	} else {
		if op1IsSigned == resultIsSigned {
			return type2, nil
		}
		*newType = *type2
	}

	// The result type needs to be unsigned. Convert the type
	// with the highest rank to unsigned.
	if newTypeName, ok := zserioSignedToUnsigned[newType.Name]; !ok {
		return nil, fmt.Errorf("%w: %s", ErrUnknownType, newType.Name)
	} else {
		newType.Name = newTypeName
	}
	return newType, nil
}
