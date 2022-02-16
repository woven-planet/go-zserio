package zserio

import (
	"fmt"
	"io"

	"github.com/icza/bitio"
)

var (
	_ io.ByteReader = (*Reader)(nil)
	_ io.Reader     = (*Reader)(nil)
)

// Reader implements a reader required to deserialize types in the ztype
// package. It also implements io.ByteReader and io.Reader interfaces.
type Reader struct {
	reader *bitio.CountReader
}

// NewReader creates a new instance of the reader.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		reader: bitio.NewCountReader(r),
	}
}

// Read reads up to len(p) bytes (8 * len(p) bits) from the underlying reader,
// and counts the number of bits read.
func (r *Reader) Read(p []byte) (int, error) {
	return r.reader.Read(p)
}

// ReadByte reads the next 8 bits and returns them as a byte.
func (r *Reader) ReadByte() (byte, error) {
	return r.reader.ReadByte()
}

// ReadBits reads n bits and returns them as the lowest n bits of u.
func (r *Reader) ReadBits(n uint8) (uint64, error) {
	return r.reader.ReadBits(n)
}

// ReadBool reads the next bit, and returns true if it is 1.
func (r *Reader) ReadBool() (bool, error) {
	return r.reader.ReadBool()
}

// Align the reader to a bit boundary
func (r *Reader) Align(boundary uint8) (int64, error) {
	d := discard(r.reader.BitsCount, boundary)
	if d == 0 {
		return r.reader.BitsCount, nil
	}

	_, err := r.ReadBits(d)
	if err != nil {
		return 0, fmt.Errorf("discard: %w", err)
	}

	return r.reader.BitsCount, err
}

func discard(count int64, boundary uint8) uint8 {
	if misaligned := uint8(count % int64(boundary)); misaligned != 0 {
		return boundary - misaligned
	}

	return 0
}
