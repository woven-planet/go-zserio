package zserio

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

type errWriter struct {
	err error
}

func (w *errWriter) Write([]byte) (int, error) {
	return 0, w.err
}

// WriteByte ensures that we are not creating bufio.NewWriter in
// bitio.NewWriter in NewWriter. If we don't have this method, then the writing
// is buffered and the buffer might be flushed during a writer.Close() method
// and the errors from the underlying writer would only be propagated at that
// time.
func (w *errWriter) WriteByte(c byte) error {
	return w.err
}

func TestWriter_Write(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		w := NewWriter(io.Discard)
		n, err := w.Write([]byte("foo"))
		assert.NoError(t, w.Close())
		assert.Equal(t, 3, n)
		assert.NoError(t, err)
		wroteBits, err := w.Align(8)
		assert.Equal(t, int64(24), wroteBits)
		assert.NoError(t, err)
	})
	t.Run("err does not increase bits count", func(t *testing.T) {
		w := NewWriter(&errWriter{assert.AnError})

		_, err := w.Write([]byte("foo"))

		assert.EqualError(t, err, "assert.AnError general error for testing")

		wroteBits, err := w.Align(8)
		assert.Zero(t, wroteBits)
		assert.NoError(t, err)

		assert.NoError(t, w.Close())
	})
}

func TestWriter_Align(t *testing.T) {
	t.Run("ok 8 bit border", func(t *testing.T) {
		var buf bytes.Buffer
		w := NewWriter(&buf)
		err := w.WriteBits(42, 4)
		assert.NoError(t, err)

		got, err := w.Align(8)

		assert.NoError(t, err)
		assert.Equal(t, int64(8), got)
		assert.Equal(t, []byte{0xa0}, buf.Bytes())
	})
	t.Run("ok", func(t *testing.T) {
		var buf bytes.Buffer
		w := NewWriter(&buf)
		err := w.WriteBits(42, 4)
		assert.NoError(t, err)

		got, err := w.Align(16)

		assert.NoError(t, err)
		assert.Equal(t, int64(16), got)
		assert.Equal(t, []byte{0xa0, 0x0}, buf.Bytes())
	})
}
