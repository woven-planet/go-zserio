package ztype_test

import (
	"bytes"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/icza/bitio"
	"github.com/woven-planet/go-zserio/ztype"
)

func TestWriteFloat16(t *testing.T) {
	tests := map[string]struct {
		input float32
		want  []byte
	}{
		"0.0": {
			input: 0.0,
			want:  []byte{0x00, 0x00},
		},
		"-0.0": {
			input: float32(math.Copysign(0, -1)),
			want:  []byte{0x80, 0x00},
		},
		"pi": {
			input: 3.14159265358979323846,
			want:  []byte{0x42, 0x48},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			w := bitio.NewWriter(&buf)
			if err := ztype.WriteFloat16(w, test.input); err != nil {
				t.Fatalf("error writing float32: %v", err)
			}
			w.Close()
			if diff := cmp.Diff(test.want, buf.Bytes()); diff != "" {
				t.Errorf("incorrect encoding: %s", diff)
			}
		})
	}
}

func TestWriteFloat32(t *testing.T) {
	tests := map[string]struct {
		input float32
		want  []byte
	}{
		"0.0": {
			input: 0.0,
			want:  []byte{0x00, 0x00, 0x00, 0x00},
		},
		"-0.0": {
			input: float32(math.Copysign(0, -1)),
			want:  []byte{0x80, 0x00, 0x00, 0x00},
		},
		"pi": {
			input: 3.14159265358979323846,
			want:  []byte{0x40, 0x49, 0x0f, 0xdb},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			w := bitio.NewWriter(&buf)
			if err := ztype.WriteFloat32(w, test.input); err != nil {
				t.Fatalf("error writing float32: %v", err)
			}
			w.Close()
			if diff := cmp.Diff(test.want, buf.Bytes()); diff != "" {
				t.Errorf("incorrect encoding: %s", diff)
			}
		})
	}
}

func TestWriteFloat64(t *testing.T) {
	tests := map[string]struct {
		input float64
		want  []byte
	}{
		"0.0": {
			input: 0.0,
			want:  []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		"-0.0": {
			input: math.Copysign(0, -1),
			want:  []byte{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		"pi": {
			input: 3.14159265358979323846,
			want:  []byte{0x40, 0x09, 0x21, 0xfb, 0x54, 0x44, 0x2d, 0x18},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			w := bitio.NewWriter(&buf)
			if err := ztype.WriteFloat64(w, test.input); err != nil {
				t.Fatalf("error writing float32: %v", err)
			}
			w.Close()
			if diff := cmp.Diff(test.want, buf.Bytes()); diff != "" {
				t.Errorf("incorrect encoding: %s", diff)
			}
		})
	}
}
