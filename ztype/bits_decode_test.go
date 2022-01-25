package ztype_test

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/icza/bitio"
	"github.com/woven-planet/go-zserio/ztype"
)

func TestReadSignedBits(t *testing.T) {
	tests := map[string]struct {
		input []byte
		bits  []uint8
		want  []int64
	}{
		"two-bit-one": {
			input: []byte{0x40},
			bits:  []uint8{2},
			want:  []int64{1},
		},
		"single-bit-negative-one": {
			input: []byte{0x80},
			bits:  []uint8{1},
			want:  []int64{-1},
		},
		"two-bit-negative-one": {
			input: []byte{0xc0},
			bits:  []uint8{2},
			want:  []int64{-1},
		},
		"multiple": {
			input: []byte{0x55},
			bits:  []uint8{1, 3, 4},
			want:  []int64{0, -3, 5},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := bitio.NewReader(bytes.NewBuffer(test.input))
			got := make([]int64, 0, len(test.input))
			for ix := range test.want {
				v, err := ztype.ReadSignedBits(r, test.bits[ix])
				if err != nil {
					t.Fatalf("error reading bits: %v", err)
				}
				got = append(got, v)
			}
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("incorrect decoding: %s", diff)
			}
		})
	}
}
