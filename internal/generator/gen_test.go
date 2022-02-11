package generator

import (
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/woven-planet/go-zserio/internal/ast"
	"github.com/woven-planet/go-zserio/internal/model"
)

func testWorkspace(filePath string) string {
	return path.Join(os.Getenv("TEST_SRCDIR"), os.Getenv("TEST_WORKSPACE"), filePath)
}

func TestStableOutputOrder(t *testing.T) {
	log.SetOutput(io.Discard) // disable logs in this test

	t.Parallel()
	for _, path := range strings.Split(os.Getenv("TEST_ZSERIO_EXAMPLES"), " ") {
		path := path
		for _, withPreamble := range []bool{true, false} {
			withPreamble := withPreamble
			t.Run(filepath.Base(path), func(t *testing.T) {
				t.Parallel()
				m, err := model.FromFiles(testWorkspace(path))
				require.NoError(t, err)
				require.Len(t, m.Packages, 1, "should have only a single package")

				var pkg *ast.Package
				for _, pkg = range m.Packages {
					break
				}

				first := newOutputs(pkg, "test", withPreamble)
				for i := range first {
					first[i].data = nil
				}
				second := newOutputs(pkg, "test", withPreamble)
				for i := range second {
					second[i].data = nil
				}
				assert.Equal(t, "pkg.go", first[0].name)
				assert.Equal(t, first, second)
			})
		}
	}
}
