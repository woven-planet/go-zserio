package ztype_test

import (
	"io"

	"github.com/woven-planet/go-zserio/ztype"
)

func NewCountReader(r io.Reader) *ztype.Reader {
	return ztype.NewReader(r)
}
