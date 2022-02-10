package ztype

import (
	"io"

	"github.com/icza/bitio"
	zserio "github.com/woven-planet/go-zserio/interface"
)

var _ zserio.Reader = (*CountReader)(nil)

// CountReader reimplements the counting of bits, which may be important when
// wanting to align the reader.
// TODO @aignas 2022-02-09: replace usage of zserio.Reader with *ztype.CountReader.
type CountReader struct {
	r         *bitio.Reader
	bitsCount int64
}

func NewCountReader(r io.Reader) *CountReader {
	return &CountReader{r: bitio.NewReader(r)}
}

func (r *CountReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	r.bitsCount += int64(n) * 8
	return n, err
}

func (r *CountReader) ReadBits(n uint8) (uint64, error) {
	b, err := r.r.ReadBits(n)
	if err == nil {
		r.bitsCount += int64(n)
	}
	return b, err
}

func (r *CountReader) ReadBool() (bool, error) {
	b, err := r.r.ReadBool()
	if err == nil {
		r.bitsCount += 1
	}
	return b, err
}

func (r *CountReader) ReadByte() (byte, error) {
	b, err := r.r.ReadByte()
	if err == nil {
		r.bitsCount += 8
	}
	return b, err
}

// Align could replace the AlignReader function.
func (r *CountReader) Align(boundary uint8) error {
	misaligned := uint8((r.bitsCount % int64(boundary)))
	if misaligned == 0 {
		return nil
	}

	discard := boundary - misaligned
	_, err := r.ReadBits(discard)
	return err
}
