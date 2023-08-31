package reference

import (
	"gen/github.com/woven-planet/go-zserio/testdata/index_operator/schema"
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
	var object schema.IndexOperator
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
	var got schema.IndexOperator
	require.NoError(t, zserio.Unmarshal(bytes, &got))

	// Then
	want_headers := []schema.BlockHeader{
		{
			NumItems: 3,
		},
		{
			NumItems: 0,
		},
		{
			NumItems: 1,
		},
	}
	want := schema.IndexOperator{
		NumBlocks: 3,
		Headers:   want_headers,
		Blocks: []schema.Block{
			{
				Header:        want_headers[0],
				Items:         []int64{1, 2, 3},
				ConditionItem: 100,
			},
			{
				Header: want_headers[1],
				// len(Items) is 0, since BlockHeader.NumItems is 0.
				// For the same reason, ConditionItem won't be serialized.
				Items: []int64{},
			},
			{
				Header:        want_headers[2],
				Items:         []int64{10},
				ConditionItem: 200,
			},
		},
	}
	assert.Equal(t, want, got)
}
