package ztype_test

import (
	"io"

	"github.com/icza/bitio"
	"github.com/woven-planet/go-zserio/ztype"
)

func NewCountWriter(w io.Writer) *ztype.CountWriter {
	return ztype.NewCountWriter(bitio.NewCountWriter(w))
}
