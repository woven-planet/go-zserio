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
				assert.Equal(t, "pkg", first[0].name)
				assert.Equal(t, first, second)
			})
		}
	}
}

func TestAssembleUniqueFilePath(t *testing.T) {
	longRootDirectory := "/Users/my_user/my_projects/my_project_working_with_zserio/my_project_resources/my_zserio_schema_data"

	longFileName := "my_very_sophisticated_zserio_template_based_data_type"
	otherLongFileName := "my_other_very_sophisticated_zserio_template_based_data_type"

	tables := []struct {
		fileName    string
		maxLength   int
		expected    string
		expectedErr string
	}{
		{ // Do not change file path if length is below maximum
			longFileName,
			DefaultMaxPathLength,
			path.Join(longRootDirectory, longFileName) + FileSuffix,
			"",
		},
		{ // Shorten file name to stay below maximum
			otherLongFileName,
			130,
			path.Join(longRootDirectory, "my_other_very") + FileSuffix,
			"",
		},
		{
			otherLongFileName,
			130,
			path.Join(longRootDirectory, "my_other_very_1") + FileSuffix,
			"",
		},
		{
			otherLongFileName,
			130,
			path.Join(longRootDirectory, "my_other_very_2") + FileSuffix,
			"",
		},
		{ // Fail if maximum length is too short to fit the path
			longFileName,
			90,
			"",
			"maximum path length 90 is too short, at least 107 required to fit output directory",
		},
		{ // Fail if maximum length is too short to fit the path
			longFileName,
			107,
			path.Join(longRootDirectory, "m") + FileSuffix,
			"",
		},
		{ // Fail if maximum length is too short to fit the path
			longFileName,
			107,
			path.Join(longRootDirectory, "m_1") + FileSuffix,
			"",
		},
		{ // Fail if maximum length is too short to fit the path
			longFileName,
			107,
			path.Join(longRootDirectory, "m_2") + FileSuffix,
			"",
		},
	}

	for _, table := range tables {
		result, err := assembleUniqueFilePath(longRootDirectory, table.fileName, table.maxLength)
		assert.Equal(t, table.expected, result)
		assert.LessOrEqual(t, len(result), table.maxLength)
		if table.expectedErr != "" {
			assert.EqualError(t, err, table.expectedErr)
		}
	}
}
