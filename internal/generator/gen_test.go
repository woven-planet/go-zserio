package generator

import (
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/bazelbuild/rules_go/go/tools/bazel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/woven-planet/go-zserio/internal/ast"
	"github.com/woven-planet/go-zserio/internal/model"
)

func testWorkspace(t require.TestingT, filePath string) string {
	actualPath, err := bazel.Runfile(filePath)
	require.NoError(t, err)
	return actualPath
}

func TestStableOutputOrder(t *testing.T) {
	log.SetOutput(io.Discard) // disable logs in this test

	t.Parallel()
	for _, path := range strings.Split(os.Getenv("TEST_ZSERIO_EXAMPLES"), " ") {
		path := path
		for _, withPreamble := range []bool{true, false} {
			opts := &options{
				outputToStdout: !withPreamble,
			}
			t.Run(filepath.Base(path), func(t *testing.T) {
				t.Parallel()
				m, err := model.FromFiles(testWorkspace(t, path))
				require.NoError(t, err)
				require.Len(t, m.Packages, 1, "should have only a single package")

				var pkg *ast.Package
				for _, pkg = range m.Packages {
					break
				}

				first := newOutputs(pkg, "test", opts)
				for i := range first {
					first[i].data = nil
				}
				second := newOutputs(pkg, "test", opts)
				for i := range second {
					second[i].data = nil
				}
				assert.Equal(t, "pkg.go", first[0].name)
				assert.Equal(t, first, second)
			})
		}
	}
}

func TestAssembleUniqueFilePath(t *testing.T) {
	longRootDirectory := "/Users/my_user/my_projects/my_project_working_with_zserio/my_project_resources/my_zserio_schema_data"
	longFileName := "my_very_sophisticated_zserio_template_based_data_type"
	otherLongFileName := "my_other_very_sophisticated_zserio_template_based_data_type"

	// Do not change current behaviour if limit is disabled
	assert.Equal(t, path.Join(longRootDirectory, longFileName)+FileSuffix, assembleUniqueFilePath(longRootDirectory, longFileName, false, DefaultMaxPathLength))

	// Do not change file path if length is below maximum
	assert.Equal(t, path.Join(longRootDirectory, otherLongFileName)+FileSuffix, assembleUniqueFilePath(longRootDirectory, otherLongFileName, true, DefaultMaxPathLength))

	// Shorten file name to stay below maximum
	assert.Equal(t, path.Join(longRootDirectory, "my_very_sophisticated")+FileSuffix, assembleUniqueFilePath(longRootDirectory, longFileName, true, 130))

	// Append indexed suffix if same file appears again
	assert.Equal(t, path.Join(longRootDirectory, "my_very_sophisticated")+"_1"+FileSuffix, assembleUniqueFilePath(longRootDirectory, longFileName, true, 130))
	assert.Equal(t, path.Join(longRootDirectory, "my_very_sophisticated")+"_2"+FileSuffix, assembleUniqueFilePath(longRootDirectory, longFileName, true, 130))

	assert.Equal(t, path.Join(longRootDirectory, "my_other_very")+FileSuffix, assembleUniqueFilePath(longRootDirectory, otherLongFileName, true, 130))
	assert.Equal(t, path.Join(longRootDirectory, "my_other_very")+"_1"+FileSuffix, assembleUniqueFilePath(longRootDirectory, otherLongFileName, true, 130))
	assert.Equal(t, path.Join(longRootDirectory, "my_other_very")+"_2"+FileSuffix, assembleUniqueFilePath(longRootDirectory, otherLongFileName, true, 130))
	assert.Equal(t, path.Join(longRootDirectory, "my_other_very")+"_3"+FileSuffix, assembleUniqueFilePath(longRootDirectory, otherLongFileName, true, 130))
}
