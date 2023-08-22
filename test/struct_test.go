package reference

import (
	"gen/github.com/woven-planet/go-zserio/testdata/reference_modules/core/types"
	"testing"

	"github.com/stretchr/testify/require"

	zserio "github.com/woven-planet/go-zserio"
)

func TestOptionalField(t *testing.T) {
	value := &types.ValueWrapper{Parameter: 0, Value: 15, EnumValue: types.ColorRed}
	dataWithoutOptional, err := zserio.Marshal(value)
	require.NoError(t, err)

	value.Parameter = 7
	dataWithOptional, err := zserio.Marshal(value)
	require.NoError(t, err)
	require.Greater(t, len(dataWithOptional), len(dataWithoutOptional))
}

func TestZserioFunction(t *testing.T) {
	// This test case ensures that the zserio functions were correctly generated,
	// and the float types were correctly generated and casted.
	value := &types.ValueWrapper{Value: 1, F16Value: 1.5, F32Value: 1.5, F64Value: 10.0}
	require.Equal(t, 17.5, value.GetSum())
}

func TestGoIntegerTypeCasts(t *testing.T) {
	// This test case tests integer type casts, when integers of different bit length
	// and signed/unsigned are used within an expression.
	value := &types.ValueWrapper{
		Value:         1,
		Vari16Value:   16,
		Vari32Value:   32,
		Vari64Value:   64,
		VariValue:     -33,
		Varu16Value:   1,
		Varu32Value:   0,
		Varu64Value:   353,
		VaruValue:     6457,
		VarSizeValue:  4425,
		I32ValueArray: []int32{-1, -2, -3, -4, -5},
	}
	require.Equal(t, (float32)(11320.0), value.TestTypeCasts())
}
