package ztype

import "github.com/icza/bitio"

func WriteBool(w *bitio.Writer, v bool) error {
	return w.WriteBool(v)
}
