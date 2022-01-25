package ztype

import "github.com/icza/bitio"

func ReadString(r *bitio.Reader) (string, error) {
	size, err := ReadVarsize(r)
	if err != nil {
		return "", err
	}
	buf := make([]byte, size)
	if _, err = r.Read(buf); err != nil {
		return "", err
	}
	return string(buf), nil
}
