package ztype_test

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/woven-planet/go-zserio/zstream"
	"github.com/woven-planet/go-zserio/ztype"
)

func TestWriteBool(t *testing.T) {
	tests := map[string]struct {
		input []bool
		want  []byte
	}{
		"false": {
			input: []bool{false},
			want:  []byte{0x00},
		},
		"true": {
			input: []bool{true},
			want:  []byte{0x80},
		},
		"multiple": {
			input: []bool{true, false, true, false, true, true},
			want:  []byte{0b10101100},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			w := zstream.NewWriter(&buf)
			var err error
			for _, n := range test.input {
				err = ztype.WriteBool(w, n)
				if err != nil {
					break
				}
			}
			w.Close()
			if diff := cmp.Diff(test.want, buf.Bytes()); diff != "" {
				t.Errorf("incorrect encoding: %s", diff)
			}
		})
	}
}
