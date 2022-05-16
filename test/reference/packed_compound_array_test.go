package reference

import (
	"gen/github.com/woven-planet/go-zserio/testdata/reference_modules/core/array"
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
	var object array.PackedArrayOfCompounds
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
	var got array.PackedArrayOfCompounds
	require.NoError(t, zserio.Unmarshal(bytes, &got))

	// Then
	want := array.PackedArrayOfCompounds{
		List: []array.PackableStructure{
			{Value: 15, Text: "fünfzehn"},
			{Value: 14, Text: "fourteen"},
			{Value: 13, Text: "dertien"},
			{Value: 16, Text: "seksten"},
			{Value: 0, Text: "nul"},
		},
	}
	assert.Equal(t, want, got)
}
