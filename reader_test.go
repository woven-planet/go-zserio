package zserio

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReader_Align(t *testing.T) {
	t.Run("ok 8 bit border", func(t *testing.T) {
		r := NewReader(strings.NewReader("foo-bar-baz"))
		_, err := r.ReadBool()
		assert.NoError(t, err)

		got, err := r.Align(8)

		assert.NoError(t, err)
		assert.Equal(t, int64(8), got)
	})
	t.Run("ok", func(t *testing.T) {
		r := NewReader(strings.NewReader("foo-bar-baz"))
		_, err := r.ReadBool()
		assert.NoError(t, err)

		got, err := r.Align(16)

		assert.NoError(t, err)
		assert.Equal(t, int64(16), got)
	})
}
