package ztype

import (
	"io"

	"github.com/icza/bitio"
	zserio "github.com/woven-planet/go-zserio"
)

var _ zserio.Reader = (*Reader)(nil)

// Reader reimplements the counting of bits, which may be important when
// wanting to align the reader.
type Reader struct {
	r         *bitio.Reader
	bitsCount int64
}

func NewReader(r io.Reader) *Reader {
	return &Reader{r: bitio.NewReader(r)}
}

func (r *Reader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	r.bitsCount += int64(n) * 8
	return n, err
}

func (r *Reader) ReadBits(n uint8) (uint64, error) {
	b, err := r.r.ReadBits(n)
	if err == nil {
		r.bitsCount += int64(n)
	}
	return b, err
}

func (r *Reader) ReadBool() (bool, error) {
	b, err := r.r.ReadBool()
	if err == nil {
		r.bitsCount += 1
	}
	return b, err
}

func (r *Reader) ReadByte() (byte, error) {
	b, err := r.r.ReadByte()
	if err == nil {
		r.bitsCount += 8
	}
	return b, err
}

// Align could replace the AlignReader function.
func (r *Reader) Align(boundary uint8) (int64, error) {
	misaligned := uint8((r.bitsCount % int64(boundary)))
	if misaligned == 0 {
		return r.bitsCount, nil
	}

	discard := boundary - misaligned
	_, err := r.ReadBits(discard)
	return r.bitsCount, err
}
