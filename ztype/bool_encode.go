package ztype

import "github.com/icza/bitio"

func WriteBool(w *bitio.CountWriter, v bool) error {
	return w.WriteBool(v)
}
