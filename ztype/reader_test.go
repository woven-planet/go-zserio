package ztype_test

import (
	"io"

	"github.com/icza/bitio"
	"github.com/woven-planet/go-zserio/ztype"
)

func NewCountReader(r io.Reader) *ztype.CountReader {
	return ztype.NewCountReader(bitio.NewCountReader(r))
}
