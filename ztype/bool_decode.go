package ztype

import "github.com/icza/bitio"

func ReadBool(r *bitio.CountReader) (bool, error) {
	return r.ReadBool()
}
