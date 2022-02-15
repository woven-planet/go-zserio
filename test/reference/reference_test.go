package reference

import (
	"os"
	"path"
	"testing"

	"gen/github.com/woven-planet/go-zserio/test/go/reference_modules/testobject1/testobject"
	zserio "github.com/woven-planet/go-zserio"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testWorkspace(filePath string) string {
	return path.Join(os.Getenv("TEST_SRCDIR"), os.Getenv("TEST_WORKSPACE"), filePath)
}

var (
	// ReferenceFilePath is the path to the input data.
	ReferenceFilePath string = testWorkspace(os.Getenv("TESTDATA_BIN"))
)

func TestReencoding(t *testing.T) {
	// Given
	want, err := os.ReadFile(ReferenceFilePath)
	require.NoError(t, err)

	// When
	var testObject testobject.TestObject
	require.NoError(t, zserio.Unmarshal(want, &testObject), "unmarshal")
	got, err := zserio.Marshal(&testObject)
	require.NoError(t, err, "marshal")

	// Then
	assert.Equal(t, want, got)
}

func TestEqual(t *testing.T) {
	t.Skip()
	// Given
	bytes, err := os.ReadFile(ReferenceFilePath)
	require.NoError(t, err)

	// When
	var got testobject.TestObject
	require.NoError(t, zserio.Unmarshal(bytes, &got))

	// Then
	var want testobject.TestObject
	assert.Equal(t, want, got)
}
