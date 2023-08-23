package reference

import (
	"gen/github.com/woven-planet/go-zserio/testdata/subtype_arithmetics/schema"
	"os"
	"testing"

	"github.com/bazelbuild/rules_go/go/tools/bazel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	zserio "github.com/woven-planet/go-zserio"
)

func testWorkspace(t require.TestingT, filePath string) string {
	actualPath, err := bazel.Runfile(filePath)
	require.NoError(t, err)
	return actualPath
}

// ReferenceFilePath is the path to the input data.
var ReferenceFilePath string = os.Getenv("TESTDATA_BIN")

func TestRoundTrip(t *testing.T) {
	// Given
	want, err := os.ReadFile(testWorkspace(t, ReferenceFilePath))
	require.NoError(t, err)

	// When
	var object schema.SubTypeArithmetics
	require.NoError(t, zserio.Unmarshal(want, &object), "unmarshal")
	got, err := zserio.Marshal(&object)
	require.NoError(t, err, "marshal")

	// Then
	assert.Equal(t, want, got)
}

func TestEqual(t *testing.T) {
	// Given
	bytes, err := os.ReadFile(testWorkspace(t, ReferenceFilePath))
	require.NoError(t, err)

	// When
	var got schema.SubTypeArithmetics
	require.NoError(t, zserio.Unmarshal(bytes, &got))

	// Then
	want := schema.SubTypeArithmetics{
		Value1:      -55543,
		Value2:      43634,
		Value3:      22.0,
		Value4:      64565,
		Value5:      -24.0,
		ColourValue: schema.ColourPurple,
	}
	assert.Equal(t, want, got)
}
