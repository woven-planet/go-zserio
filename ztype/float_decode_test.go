package ztype_test

import (
	"bytes"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/icza/bitio"
	"github.com/woven-planet/go-zserio/ztype"
)

func TestReadFloat16(t *testing.T) {
	tests := map[string]struct {
		input []byte
		want  float32
	}{
		"0.0": {
			input: []byte{0x00, 0x00},
			want:  0.00,
		},
		"-0.0": {
			input: []byte{0x80, 0x00},
			want:  float32(math.Copysign(0, -1)),
		},
		"pi": {
			input: []byte{0x42, 0x48},
			want:  3.140625,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := bitio.NewCountReader(bytes.NewBuffer(test.input))
			got, err := ztype.ReadFloat16(r)
			if err != nil {
				t.Fatalf("error reading float: %v", err)
			}
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("incorrect decoding: %s", diff)
			}
			// We have to explicitly check the sign bit, since cmp.Diff
			// considers 0.0 and -0.0 to be identical.
			if math.Signbit(float64(test.want)) != math.Signbit(float64(got)) {
				t.Errorf("incorrect sign bit: -%v +%v", math.Signbit(float64(test.want)), math.Signbit(float64(got)))
			}
		})
	}
}

func TestReadFloat32(t *testing.T) {
	tests := map[string]struct {
		input []byte
		want  float32
	}{
		"0.0": {
			input: []byte{0x00, 0x00, 0x00, 0x00},
			want:  0.00,
		},
		"-0.0": {
			input: []byte{0x80, 0x00, 0x00, 0x00},
			want:  float32(math.Copysign(0, -1)),
		},
		"pi": {
			input: []byte{0x40, 0x49, 0x0f, 0xdb},
			want:  3.1415927410125732,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := bitio.NewCountReader(bytes.NewBuffer(test.input))
			got, err := ztype.ReadFloat32(r)
			if err != nil {
				t.Fatalf("error reading float: %v", err)
			}
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("incorrect decoding: %s", diff)
			}
			// We have to explicitly check the sign bit, since cmp.Diff
			// considers 0.0 and -0.0 to be identical.
			if math.Signbit(float64(test.want)) != math.Signbit(float64(got)) {
				t.Errorf("incorrect sign bit: -%v +%v", math.Signbit(float64(test.want)), math.Signbit(float64(got)))
			}
		})
	}
}

func TestReadFloat64(t *testing.T) {
	tests := map[string]struct {
		input []byte
		want  float64
	}{
		"0.0": {
			input: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			want:  0.00,
		},
		"-0.0": {
			input: []byte{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			want:  float64(math.Copysign(0, -1)),
		},
		"pi": {
			input: []byte{0x40, 0x09, 0x21, 0xfb, 0x54, 0x44, 0x2d, 0x18},
			want:  3.141592653589793,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := bitio.NewCountReader(bytes.NewBuffer(test.input))
			got, err := ztype.ReadFloat64(r)
			if err != nil {
				t.Fatalf("error reading float: %v", err)
			}
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("incorrect decoding: %s", diff)
			}
			// We have to explicitly check the sign bit, since cmp.Diff
			// considers 0.0 and -0.0 to be identical.
			if math.Signbit(float64(test.want)) != math.Signbit(float64(got)) {
				t.Errorf("incorrect sign bit: -%v +%v", math.Signbit(float64(test.want)), math.Signbit(float64(got)))
			}
		})
	}
}
