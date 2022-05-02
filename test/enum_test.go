package reference

import (
	"gen/github.com/woven-planet/go-zserio/testdata/reference_modules/core/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnumConversion(t *testing.T) {
	t.Parallel()
	for _, name := range types.SomeEnumStrings() {
		val, err := types.SomeEnumString(name)
		assert.NoError(t, err)
		assert.Equal(t, val.String(), name)
	}
}

func TestSqlRoundtrip(t *testing.T) {
	t.Parallel()
	for _, want := range types.SomeEnumValues() {
		var got types.SomeEnum
		assert.NoError(t, got.Scan(want.String()))
		assert.Equal(t, want, got)
		assert.NoError(t, got.Scan([]byte(want.String())))
		assert.Equal(t, want, got)
		assert.NoError(t, got.Scan(want))
		assert.Equal(t, want, got)
	}
}

func TestStringer(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		input types.SomeEnum
		want  string
	}{
		"known-value": {
			input: types.SomeEnumAttrA,
			want:  "ATTR_A",
		},
		"unknown-value": {
			input: types.SomeEnum(-1),
			want:  "unknown(-1)",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.input.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
