package ztype

import "github.com/icza/bitio"

func ReadBool(r *bitio.Reader) (bool, error) {
	return r.ReadBool()
}
