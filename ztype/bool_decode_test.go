package ztype_test

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/woven-planet/go-zserio/zstream"
	"github.com/woven-planet/go-zserio/ztype"
)

func TestReadBool(t *testing.T) {
	tests := map[string]struct {
		input []byte
		want  []bool
	}{
		"false": {
			input: []byte{0x00},
			want:  []bool{false},
		},
		"true": {
			input: []byte{0x80},
			want:  []bool{true},
		},
		"multiple": {
			input: []byte{0b10101100},
			want:  []bool{true, false, true, false, true, true},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := zstream.NewReader(bytes.NewBuffer(test.input))
			got := make([]bool, 0, len(test.input))
			for range test.want {
				b, err := ztype.ReadBool(r)
				if err != nil {
					t.Fatalf("error reading bool: %v", err)
				}
				got = append(got, b)
			}
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("incorrect decoding: %s", diff)
			}
		})
	}
}
