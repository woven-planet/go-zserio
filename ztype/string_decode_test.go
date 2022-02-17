package ztype_test

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/woven-planet/go-zserio/zstream"
	"github.com/woven-planet/go-zserio/ztype"
)

func TestReadString(t *testing.T) {
	tests := map[string]struct {
		input []byte
		want  string
	}{
		"empty": {
			input: []byte{0x00},
			want:  "",
		},
		"ascii": {
			input: append([]byte{0x0c}, []byte("Hello, world")...),
			want:  "Hello, world",
		},
		"unicode": {
			input: append([]byte{0x0c}, []byte("Hellø! 🚀")...),
			want:  "Hellø! 🚀",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := zstream.NewReader(bytes.NewBuffer(test.input))
			got, err := ztype.ReadString(r)
			if err != nil {
				t.Fatalf("error reading string: %v", err)
			}
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("incorrect decoding: %s", diff)
			}
		})
	}
}
