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
	var object array.PackedNestedArray
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
	var got array.PackedNestedArray
	require.NoError(t, zserio.Unmarshal(bytes, &got))

	// Then
	want := array.PackedNestedArray{
		List: []array.PackableNestedStructure{
			{Value32: 15, Text: "fünfzehn", InnerStructure: array.InnerStructure{Value64: 15 * 15, Value16: 15 * 2}},
			{Value32: 14, Text: "fourteen", InnerStructure: array.InnerStructure{Value64: 14 * 14, Value16: 14 * 2}},
			{Value32: 13, Text: "dertien", InnerStructure: array.InnerStructure{Value64: 13 * 13, Value16: 13 * 2}},
			{Value32: 16, Text: "seksten", InnerStructure: array.InnerStructure{Value64: 16 * 16, Value16: 16 * 2}},
			{Value32: 0, Text: "nul", InnerStructure: array.InnerStructure{Value64: 0, Value16: 0 * 2}},
		},
	}
	assert.Equal(t, want, got)
}
