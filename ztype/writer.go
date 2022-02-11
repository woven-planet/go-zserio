package ztype

import (
	"io"

	"github.com/icza/bitio"
	zserio "github.com/woven-planet/go-zserio"
)

var _ zserio.Writer = (*Writer)(nil)

// Writer reimplements the counting of bits, which may be important when
// wanting to align the writer.
type Writer struct {
	writer    *bitio.Writer
	bitsCount int64
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{
		writer: bitio.NewWriter(w),
	}
}

func (w *Writer) Write(p []byte) (int, error) {
	n, err := w.writer.Write(p)
	if err == nil {
		w.bitsCount += int64(n) * 8
	}
	return n, err
}

func (w *Writer) WriteBits(r uint64, n uint8) error {
	err := w.writer.WriteBits(r, n)
	if err == nil {
		w.bitsCount += int64(n)
	}
	return err
}

func (w *Writer) WriteBitsUnsafe(r uint64, n uint8) error {
	err := w.writer.WriteBitsUnsafe(r, n)
	if err == nil {
		w.bitsCount += int64(n)
	}
	return err
}

func (w *Writer) WriteBool(b bool) error {
	err := w.writer.WriteBool(b)
	if err == nil {
		w.bitsCount += 1
	}
	return err
}

func (w *Writer) WriteByte(b byte) error {
	err := w.writer.WriteByte(b)
	if err == nil {
		w.bitsCount += 8
	}
	return err
}

// Align could replace the AlignReader function.
func (w *Writer) Align(boundary uint8) (int64, error) {
	misaligned := uint8((w.bitsCount % int64(boundary)))

	if misaligned != 0 {
		numOfBitsToDiscard := boundary - misaligned
		err := w.WriteBits(0, numOfBitsToDiscard)
		return w.bitsCount, err
	}
	return w.bitsCount, nil
}

// Close flushes the buffer, but does not close the underlying writer.
func (w *Writer) Close() error {
	return w.writer.Close()
}
