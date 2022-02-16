package ztype_test

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/woven-planet/go-zserio/zstream"
	"github.com/woven-planet/go-zserio/ztype"
)

func TestWriteString(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  []byte
	}{
		"empty": {
			input: []string{""},
			want:  []byte{0x00},
		},
		"ascii": {
			input: []string{"Hello, world"},
			want:  append([]byte{0x0c}, []byte("Hello, world")...),
		},
		"unicode": {
			input: []string{"Hellø! 🚀"},
			want:  append([]byte{0x0c}, []byte("Hellø! 🚀")...),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			w := zstream.NewWriter(&buf)
			var err error
			for _, n := range test.input {
				err = ztype.WriteString(w, n)
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
