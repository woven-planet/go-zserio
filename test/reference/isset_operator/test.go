package reference

import (
	"gen/github.com/woven-planet/go-zserio/testdata/isset_operator/schema"
	"os"
	"testing"

	"github.com/bazelbuild/rules_go/go/tools/bazel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	zserio "github.com/woven-planet/go-zserio"
)

func testIssetOperator(t *testing.T) {
	// Create a test object with two bits set.
	obj := schema.IssetOperator{
		U8Value: schema.PermissionREADABLE | schema.PermissionWRITEABLE,
	}
	assert.False(t, obj.IsExecutable())
	assert.True(t, obj.IsReadable())
	assert.True(t, obj.IsWriteable())

	// Change the bit mask value, and ensure the functions still return
	// the correct value.
	obj.U8Value = schema.PermissionEXECUTABLE
	assert.True(t, obj.IsExecutable())
	assert.False(t, obj.IsReadable())
	assert.False(t, obj.IsWriteable())
}

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
	var object schema.IssetOperator
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
	var got schema.IssetOperator
	require.NoError(t, zserio.Unmarshal(bytes, &got))

	// Then
	want := schema.IssetOperator{
		U8Value: schema.PermissionREADABLE,
	}
	assert.Equal(t, want, got)
}
