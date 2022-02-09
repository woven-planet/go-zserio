package ztype

import zserio "github.com/woven-planet/go-zserio/interface"

var _ zserio.Reader = (*CountReader)(nil)

// CountReader reimplements the counting of bits, which may be important when
// wanting to align the reader.
// TODO @aignas 2022-02-09: replace usage of *bitio.CountReader with *ztype.CountReader.
type CountReader struct {
	r zserio.Reader

	bitsCount int64
}

func (r *CountReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	r.bitsCount += int64(n) * 8
	return n, err
}

func (r *CountReader) ReadBits(n uint8) (uint64, error) {
	b, err := r.r.ReadBits(n)
	r.bitsCount += int64(b)
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
func (r *CountReader) Align(boundary uint8) {
	misaligned := uint8((r.bitsCount % int64(boundary)))
	if misaligned == 0 {
		return
	}

	discard := boundary - misaligned
	r.ReadBits(discard)
}
