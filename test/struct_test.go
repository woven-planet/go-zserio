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
	// and the float types were correctly generated.
	value := &types.ValueWrapper{Value: 1, F64Value: 10.0}
	require.Equal(value.getSum(), 22.0)
}
