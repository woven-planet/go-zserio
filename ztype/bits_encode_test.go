package ztype_test

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/woven-planet/go-zserio/zstream"
	"github.com/woven-planet/go-zserio/ztype"
)

func TestWriteSignedBits(t *testing.T) {
	tests := map[string]struct {
		input []int64
		bits  []uint8
		want  []byte
	}{
		"two-bit-one": {
			input: []int64{1},
			bits:  []uint8{2},
			want:  []byte{0x40},
		},
		"single-bit-negative-one": {
			input: []int64{-1},
			bits:  []uint8{1},
			want:  []byte{0x80},
		},
		"two-bit-negative-one": {
			input: []int64{-1},
			bits:  []uint8{2},
			want:  []byte{0xc0},
		},
		"multiple": {
			input: []int64{0, -3, 5},
			bits:  []uint8{1, 3, 4},
			want:  []byte{0x55},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Run(name, func(t *testing.T) {
				var buf bytes.Buffer
				w := zstream.NewWriter(&buf)
				var err error
				for ix := range test.input {
					err = ztype.WriteSignedBits(w, test.input[ix], test.bits[ix])
					if err != nil {
						break
					}
				}
				w.Close()
				if diff := cmp.Diff(test.want, buf.Bytes()); diff != "" {
					t.Errorf("incorrect encoding: %s", diff)
				}
			})
		})
	}
}
