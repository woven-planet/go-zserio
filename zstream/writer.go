package zstream

import (
	"fmt"
	"io"

	"github.com/icza/bitio"
)

// Writer is the interface to our BitWriter.
type Writer interface {
	io.Writer
	io.ByteWriter
}

var (
	_ io.Closer = (*BitWriter)(nil)
)

// BitWriter implements a Writer required to serialize types in the ztype package.
// It also implements io.ByteWriter, io.Writer and io.Closer interfaces. The
// user is responsible for closing the writer in order to flush the byte stream
// to the underlying writer.
type BitWriter struct {
	writer *bitio.CountWriter
}

// NewWriter creates a new instance of the writer.
func NewWriter(r Writer) *BitWriter {
	// implementation note: accepting io.ByteWriter ensures that we are not
	// creating bufio.NewWriter in bitio.NewWriter in NewWriter. If we don't
	// have this method, then the writing is buffered and the buffer might be
	// flushed during a writer.Close() method and the errors from the
	// underlying writer would only be propagated at that time.
	return &BitWriter{
		writer: bitio.NewCountWriter(r),
	}
}

// WriteBits writes out the n lowest bits of r.
func (w *BitWriter) Write(b []byte) (int, error) {
	return w.writer.Write(b)
}

// WriteBits writes out the n lowest bits of r.
func (w *BitWriter) WriteBits(r uint64, n uint8) error {
	return w.writer.WriteBits(r, n)
}

// WriteBitsUnsafe writes out the n lowest bits of r.
//
// WriteBitsUnsafe() offers slightly better performance than WriteBits() because
// the input r is not masked. Calling WriteBitsUnsafe() with an r that does
// not satisfy this is undefined behavior (might corrupt previously written bits).
func (w *BitWriter) WriteBitsUnsafe(r uint64, n uint8) error {
	return w.writer.WriteBitsUnsafe(r, n)
}

// WriteBool writes one bit: 1 if param is true, 0 otherwise.
func (w *BitWriter) WriteBool(b bool) error {
	return w.writer.WriteBool(b)
}

// WriteByte writes 8 bits.
func (w *BitWriter) WriteByte(b byte) error {
	return w.writer.WriteByte(b)
}

// Align a writer to a bit boundary.
func (w *BitWriter) Align(boundary uint8) (int64, error) {
	d := discard(w.writer.BitsCount, boundary)
	if d == 0 {
		return w.writer.BitsCount, nil
	}

	err := w.WriteBits(0, d)
	if err != nil {
		return 0, fmt.Errorf("discard: %w", err)
	}

	return w.writer.BitsCount, nil
}

// Close flushes the internal buffer, but does not close the underlying writer implementation.
func (w *BitWriter) Close() error {
	return w.writer.Close()
}

// BitPosition returns the number of bits written
func (w *BitWriter) BitPosition() int64 {
	return w.writer.BitsCount
}
