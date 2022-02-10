package ztype

import (
	"io"

	zserio "github.com/woven-planet/go-zserio/interface"
)

var _ zserio.Writer = (*CountWriter)(nil)

// CountWriter reimplements the counting of bits, which may be important when
// wanting to align the writer.
// TODO @aignas 2022-02-09: replace usage of zserio.Writer with *ztype.CountWriter.
type CountWriter struct {
	w Writer

	bitsCount int64
}

type Writer interface {
	io.Writer
	io.ByteWriter

	WriteBits(r uint64, n uint8) error
	WriteBool(bool) error
	WriteBitsUnsafe(r uint64, n uint8) error
}

func NewCountWriter(w Writer) *CountWriter {
	return &CountWriter{
		w: w,
	}
}

func (w *CountWriter) Write(p []byte) (int, error) {
	n, err := w.w.Write(p)
	if err == nil {
		w.bitsCount += int64(n) * 8
	}
	return n, err
}

func (w *CountWriter) WriteBits(r uint64, n uint8) error {
	err := w.w.WriteBits(r, n)
	if err == nil {
		w.bitsCount += int64(n)
	}
	return err
}

func (w *CountWriter) WriteBitsUnsafe(r uint64, n uint8) error {
	err := w.w.WriteBitsUnsafe(r, n)
	if err == nil {
		w.bitsCount += int64(n)
	}
	return err
}

func (w *CountWriter) WriteBool(b bool) error {
	err := w.w.WriteBool(b)
	if err == nil {
		w.bitsCount += 1
	}
	return err
}

func (w *CountWriter) WriteByte(b byte) error {
	err := w.w.WriteByte(b)
	if err == nil {
		w.bitsCount += 8
	}
	return err
}

// Align could replace the AlignReader function.
func (w *CountWriter) Align(boundary uint8) (int64, error) {
	misaligned := uint8((w.bitsCount % int64(boundary)))

	if misaligned != 0 {
		numOfBitsToDiscard := boundary - misaligned
		err := w.WriteBits(0, numOfBitsToDiscard)
		return w.bitsCount, err
	}
	return w.bitsCount, nil
}
