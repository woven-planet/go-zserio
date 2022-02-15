package ztype_test

import (
	"github.com/woven-planet/go-zserio"
	"github.com/woven-planet/go-zserio/ztype"
)

// This is a compile-time check to see that we are implementing the interface.
var _ zserio.ZserioType = (*ztype.Array[int64, *ztype.VarIntArrayTraits])(nil)
