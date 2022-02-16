package zserio

import (
	"fmt"
	"io"

	"github.com/icza/bitio"
)

var (
	_ io.ByteWriter = (*Writer)(nil)
	_ io.Closer     = (*Writer)(nil)
	_ io.Writer     = (*Writer)(nil)
)

// Writer implements a writer required to serialize types in the ztype package.
// It also implements io.ByteWriter, io.Writer and io.Closer interfaces. The
// user is responsible for closing the writer in order to flush the byte stream
// to the underlying writer.
type Writer struct {
	writer *bitio.CountWriter
}

// NewWriter creates a new instance of the writer.
func NewWriter(r io.Writer) *Writer {
	return &Writer{
		writer: bitio.NewCountWriter(r),
	}
}

// WriteBits writes out the n lowest bits of r.
func (w *Writer) Write(b []byte) (int, error) {
	return w.writer.Write(b)
}

// WriteBits writes out the n lowest bits of r.
func (w *Writer) WriteBits(r uint64, n uint8) error {
	return w.writer.WriteBits(r, n)
}

// WriteBitsUnsafe writes out the n lowest bits of r.
//
// WriteBitsUnsafe() offers slightly better performance than WriteBits() because
// the input r is not masked. Calling WriteBitsUnsafe() with an r that does
// not satisfy this is undefined behavior (might corrupt previously written bits).
func (w *Writer) WriteBitsUnsafe(r uint64, n uint8) error {
	return w.writer.WriteBitsUnsafe(r, n)
}

// WriteBool writes one bit: 1 if param is true, 0 otherwise.
func (w *Writer) WriteBool(b bool) error {
	return w.writer.WriteBool(b)
}

// WriteByte writes 8 bits.
func (w *Writer) WriteByte(b byte) error {
	return w.writer.WriteByte(b)
}

// Align a writer to a bit boundary.
func (w *Writer) Align(boundary uint8) (int64, error) {
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
func (w *Writer) Close() error {
	return w.writer.Close()
}
