package reference

import (
	"gen/github.com/woven-planet/go-zserio/testdata/extern_bytes_array/schema"
	"os"
	"testing"

	"github.com/bazelbuild/rules_go/go/tools/bazel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	zserio "github.com/woven-planet/go-zserio"
	"github.com/woven-planet/go-zserio/ztype"
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
	var object schema.ExternBytesArray
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
	var got schema.ExternBytesArray
	require.NoError(t, zserio.Unmarshal(bytes, &got))

	// Then
	want := schema.ExternBytesArray{
		Extern_array: []*ztype.ExternType{
			{
				BitSize: 19,
				Buffer:  []byte{0xFF, 0x00, 0xE0},
			},
			{
				BitSize: 19,
				Buffer:  []byte{0xFF, 0x01, 0xE0},
			},
			{
				BitSize: 19,
				Buffer:  []byte{0xFF, 0x02, 0xE0},
			},
		},
		Bytes_array: []*ztype.BytesType{
			{
				ByteSize: 2,
				Buffer:   []byte{0xFF, 0x00},
			},
			{
				ByteSize: 2,
				Buffer:   []byte{0xFF, 0x01},
			},
			{
				ByteSize: 2,
				Buffer:   []byte{0xFF, 0x02},
			},
		},
	}
	assert.Equal(t, want, got)
}
