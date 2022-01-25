package ztype_test

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/icza/bitio"
	"github.com/woven-planet/go-zserio/ztype"
)

func TestReadVarint16(t *testing.T) {
	tests := map[string]struct {
		input []byte
		want  int16
	}{
		"zero": {
			input: []byte{0x00},
			want:  0,
		},
		"minimum": {
			input: []byte{0xff, 0xff},
			want:  ztype.VARINT16_MIN,
		},
		"maximum": {
			input: []byte{0x7f, 0xff},
			want:  ztype.VARINT16_MAX,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := bitio.NewReader(bytes.NewBuffer(test.input))
			v, err := ztype.ReadVarint16(r)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if diff := cmp.Diff(test.want, v); diff != "" {
				t.Errorf("incorrect encoding: %s", diff)
			}
		})
	}
}

func TestReadVarint32(t *testing.T) {
	tests := map[string]struct {
		input []byte
		want  int32
	}{
		"zero": {
			input: []byte{0x00},
			want:  0,
		},
		"minimum": {
			input: []byte{0xff, 0xff, 0xff, 0xff},
			want:  ztype.VARINT32_MIN,
		},
		"maximum": {
			input: []byte{0x7f, 0xff, 0xff, 0xff},
			want:  ztype.VARINT32_MAX,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := bitio.NewReader(bytes.NewBuffer(test.input))
			v, err := ztype.ReadVarint32(r)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if diff := cmp.Diff(test.want, v); diff != "" {
				t.Errorf("incorrect encoding: %s", diff)
			}
		})
	}
}

func TestReadVarint64(t *testing.T) {
	tests := map[string]struct {
		input []byte
		want  int64
	}{
		"zero": {
			input: []byte{0x00},
			want:  0,
		},
		"minimum": {
			input: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			want:  ztype.VARINT64_MIN,
		},
		"maximum": {
			input: []byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			want:  ztype.VARINT64_MAX,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := bitio.NewReader(bytes.NewBuffer(test.input))
			v, err := ztype.ReadVarint64(r)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if diff := cmp.Diff(test.want, v); diff != "" {
				t.Errorf("incorrect encoding: %s", diff)
			}
		})
	}
}

func TestReadVarint(t *testing.T) {
	tests := map[string]struct {
		input []byte
		want  int64
	}{
		"zero": {
			input: []byte{0x00},
			want:  0,
		},
		"minimum": {
			input: []byte{0x80},
			want:  ztype.VARINT_MIN,
		},
		"non-magic-minimum": {
			input: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			want:  ztype.VARINT_MIN + 1,
		},
		"maximum": {
			input: []byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			want:  ztype.VARINT_MAX,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := bitio.NewReader(bytes.NewBuffer(test.input))
			v, err := ztype.ReadVarint(r)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if diff := cmp.Diff(test.want, v); diff != "" {
				t.Errorf("incorrect encoding: %s", diff)
			}
		})
	}
}

func BenchmarkReadVarint64(b *testing.B) {
	input := []byte{
		0xff, 0xff, 0x00,
		0xff, 0xff, 0xff, 0x7,
		0xff, 0xff, 0xff, 0x7,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	}
	for n := 0; n < b.N; n++ {
		r := bitio.NewReader(bytes.NewBuffer(input))
		_, _ = ztype.ReadVarint64(r)
		_, _ = ztype.ReadVarint64(r)
		_, _ = ztype.ReadVarint64(r)
		_, _ = ztype.ReadVarint64(r)
	}
}
