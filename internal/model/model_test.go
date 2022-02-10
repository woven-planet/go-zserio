package model

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/woven-planet/go-zserio/internal/ast"
)

func testWorkspace(filePath string) string {
	return path.Join(os.Getenv("TEST_SRCDIR"), os.Getenv("TEST_WORKSPACE"), filePath)
}

func TestCanLoadEachExample(t *testing.T) {
	t.Parallel()
	for _, path := range strings.Split(os.Getenv("TEST_ZSERIO_EXAMPLES"), " ") {
		path := path
		t.Run(filepath.Base(path), func(t *testing.T) {
			t.Parallel()
			_, err := FromFiles(testWorkspace(path))
			assert.NoError(t, err)
		})
	}
}

func TestEnumResultTypes(t *testing.T) {
	m, err := FromFiles(testWorkspace("testdata/enum_expressions.zs"))
	require.NoError(t, err)

	for _, pkg := range m.Packages {
		for _, enum := range pkg.Enums {
			for i, item := range enum.Items {
				if item.ResultType == ast.ExpressionTypeInteger {
					continue
				}

				if item.ResultType == ast.ExpressionTypeString {
					continue
				}

				// TODO @aignas 2022-02-10: add a test where we ensure that recursive
				// evaluation of enums is being done.
				t.Errorf("item[%d] has an unexpected type %q (evaluation state: %q): %q", i, item.ResultType, item.EvaluationState, item)
			}
		}
	}
}
