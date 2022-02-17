package ztype_test

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/woven-planet/go-zserio/zstream"
	"github.com/woven-planet/go-zserio/ztype"
)

func TestWriteVarint16(t *testing.T) {
	tests := map[string]struct {
		input []int16
		want  []byte
		err   error
	}{
		"zero": {
			input: []int16{0},
			want:  []byte{0x00},
		},
		"two-byte-value": {
			input: []int16{(1 << 6)},
			want:  []byte{0x40, 0x40},
		},
		"negative-value": {
			input: []int16{-0x1ff},
			want:  []byte{0xc1, 0xff},
		},
		"minimum": {
			input: []int16{ztype.MinVarint16},
			want:  []byte{0xff, 0xff},
		},
		"maximum": {
			input: []int16{ztype.MaxVarint16},
			want:  []byte{0x7f, 0xff},
		},
		"too-small": {
			input: []int16{ztype.MinVarint16 - 1},
			err:   ztype.ErrOutOfBounds,
		},
		"too-large": {
			input: []int16{ztype.MaxVarint16 + 1},
			err:   ztype.ErrOutOfBounds,
		},
		"multi-values": {
			input: []int16{0x40, 0x08, 0x1ff},
			want: []byte{
				0x40, 0x40,
				0x08,
				0x41, 0xff,
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			w := zstream.NewWriter(&buf)
			var err error
			for _, n := range test.input {
				err = ztype.WriteVarint16(w, n)
				if err != nil {
					break
				}
			}
			w.Close()
			if diff := cmp.Diff(test.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Fatalf("incorrect error: %s", diff)
			}
			if diff := cmp.Diff(test.want, buf.Bytes()); diff != "" {
				t.Errorf("incorrect encoding: %s", diff)
			}
		})
	}
}

func TestWriteVarint32(t *testing.T) {
	tests := map[string]struct {
		input []int32
		want  []byte
		err   error
	}{
		"zero": {
			input: []int32{0},
			want:  []byte{0x00},
		},
		"two-byte-value": {
			input: []int32{(1 << 6)},
			want:  []byte{0x40, 0x40},
		},
		"0xffff": {
			input: []int32{0xffff},
			want:  []byte{0x43, 0xff, 0x7f},
		},
		"minimum": {
			input: []int32{ztype.MinVarint32},
			want:  []byte{0xff, 0xff, 0xff, 0xff},
		},
		"maximum": {
			input: []int32{ztype.MaxVarint32},
			want:  []byte{0x7f, 0xff, 0xff, 0xff},
		},
		"too-small": {
			input: []int32{ztype.MinVarint32 - 1},
			err:   ztype.ErrOutOfBounds,
		},
		"too-large": {
			input: []int32{ztype.MaxVarint32 + 1},
			err:   ztype.ErrOutOfBounds,
		},
		"multi-values": {
			input: []int32{0x40, 0x08, 0x1ff},
			want: []byte{
				0x40, 0x40,
				0x08,
				0x43, 0x7f,
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			w := zstream.NewWriter(&buf)
			var err error
			for _, n := range test.input {
				err = ztype.WriteVarint32(w, n)
				if err != nil {
					break
				}
			}
			w.Close()
			if diff := cmp.Diff(test.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Fatalf("incorrect error: %s", diff)
			}
			if diff := cmp.Diff(test.want, buf.Bytes()); diff != "" {
				t.Errorf("incorrect encoding: %s", diff)
			}
		})
	}
}

func TestWriteVarint64(t *testing.T) {
	tests := map[string]struct {
		input []int64
		want  []byte
		err   error
	}{
		"zero": {
			input: []int64{0},
			want:  []byte{0x00},
		},
		"two-byte-value": {
			input: []int64{(1 << 6)},
			want:  []byte{0x40, 0x40},
		},
		"0xffff": {
			input: []int64{0xffff},
			want:  []byte{0x43, 0xff, 0x7f},
		},
		"minimum": {
			input: []int64{ztype.MinVarint64},
			want:  []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		},
		"maximum": {
			input: []int64{ztype.MaxVarint64},
			want:  []byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		},
		"too-small": {
			input: []int64{ztype.MinVarint64 - 1},
			err:   ztype.ErrOutOfBounds,
		},
		"too-large": {
			input: []int64{ztype.MaxVarint64 + 1},
			err:   ztype.ErrOutOfBounds,
		},
		"multi-values": {
			input: []int64{0x40, 0x08, 0x1ff},
			want: []byte{
				0x40, 0x40,
				0x08,
				0x43, 0x7f,
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			w := zstream.NewWriter(&buf)
			var err error
			for _, n := range test.input {
				err = ztype.WriteVarint64(w, n)
				if err != nil {
					break
				}
			}
			w.Close()
			if diff := cmp.Diff(test.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Fatalf("incorrect error: %s", diff)
			}
			if diff := cmp.Diff(test.want, buf.Bytes()); diff != "" {
				t.Errorf("incorrect encoding: %s", diff)
			}
		})
	}
}

func TestWriteVarint(t *testing.T) {
	tests := map[string]struct {
		input []int64
		want  []byte
		err   error
	}{
		"zero": {
			input: []int64{0},
			want:  []byte{0x00},
		},
		"two-byte-value": {
			input: []int64{(1 << 6)},
			want:  []byte{0x40, 0x40},
		},
		"0xffff": {
			input: []int64{0xffff},
			want:  []byte{0x43, 0xff, 0x7f},
		},
		"minimum": {
			input: []int64{ztype.MinVarint},
			want:  []byte{0x80},
		},
		"non-magic-minimum": {
			input: []int64{ztype.MinVarint + 1},
			want:  []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		},
		"maximum": {
			input: []int64{ztype.MaxVarint},
			want:  []byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		},
		// Note that we do not test for (VARINT_MIN - 1) and (VARINT_MAX + 1)
		// because they do not fit in an int64, so you can never generate those
		// in Go anyway.
		"multi-values": {
			input: []int64{0x40, 0x08, 0x1ff},
			want: []byte{
				0x40, 0x40,
				0x08,
				0x43, 0x7f,
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			w := zstream.NewWriter(&buf)
			var err error
			for _, n := range test.input {
				err = ztype.WriteVarint(w, n)
				if err != nil {
					break
				}
			}
			w.Close()
			if diff := cmp.Diff(test.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Fatalf("incorrect error: %s", diff)
			}
			if diff := cmp.Diff(test.want, buf.Bytes()); diff != "" {
				t.Errorf("incorrect encoding: %s", diff)
			}
		})
	}
}

func TestSignedBitSize(t *testing.T) {
	// Test values taken from https://github.com/ndsev/zserio/blob/9aac67d3300f1b18db86180a843ed8eaf4c9c296/compiler/extensions/python/runtime/tests/test_bitsizeof.py
	type testitem struct {
		expected int
		value    int64
		err      error
	}
	tests := map[string]struct {
		maxBytes int
		items    []testitem
	}{
		"int16": {
			maxBytes: 2,
			items: []testitem{
				{8, 0, nil},

				{8, (1 << 0), nil},
				{8, -(1 << 0), nil},
				{8, (1 << 6) - 1, nil},
				{8, -((1 << 6) - 1), nil},

				{16, (1 << 6), nil},
				{16, -(1 << 6), nil},
				{16, (1 << (6 + 8)) - 1, nil},
				{16, -((1 << (6 + 8)) - 1), nil},

				{0, -(1 << (6 + 8)), ztype.ErrOutOfBounds},
				{0, (1 << (6 + 8)), ztype.ErrOutOfBounds},
			},
		},
		"int32": {
			maxBytes: 4,
			items: []testitem{
				{8, 0, nil},

				{8, (1 << 0), nil},
				{8, -(1 << 0), nil},
				{8, (1 << 6) - 1, nil},
				{8, -((1 << 6) - 1), nil},

				{16, (1 << 6), nil},
				{16, -(1 << 6), nil},
				{16, (1 << (6 + 7)) - 1, nil},
				{16, -((1 << (6 + 7)) - 1), nil},

				{24, (1 << (6 + 7)), nil},
				{24, -(1 << (6 + 7)), nil},
				{24, (1 << (6 + 7 + 7)) - 1, nil},
				{24, -((1 << (6 + 7 + 7)) - 1), nil},

				{32, (1 << (6 + 7 + 7)), nil},
				{32, -(1 << (6 + 7 + 7)), nil},
				{32, (1 << (6 + 7 + 7 + 8)) - 1, nil},
				{32, -((1 << (6 + 7 + 7 + 8)) - 1), nil},

				{32, ztype.MinVarint32, nil},
				{32, ztype.MaxVarint32, nil},

				{0, (1 << (6 + 7 + 7 + 8)), ztype.ErrOutOfBounds},
				{0, -(1 << (6 + 7 + 7 + 8)), ztype.ErrOutOfBounds},
			},
		},
		"int64": {
			maxBytes: 8,
			items: []testitem{
				{8, 0, nil},

				{8, (1 << 0), nil},
				{8, -(1 << 0), nil},
				{8, (1 << 6) - 1, nil},
				{8, -((1 << 6) - 1), nil},

				{16, (1 << 6), nil},
				{16, -(1 << 6), nil},
				{16, (1 << (6 + 7)) - 1, nil},
				{16, -((1 << (6 + 7)) - 1), nil},

				{24, (1 << (6 + 7)), nil},
				{24, -(1 << (6 + 7)), nil},
				{24, (1 << (6 + 7 + 7)) - 1, nil},
				{24, -((1 << (6 + 7 + 7)) - 1), nil},

				{32, (1 << (6 + 7 + 7)), nil},
				{32, -(1 << (6 + 7 + 7)), nil},
				{32, (1 << (6 + 7 + 7 + 7)) - 1, nil},
				{32, -((1 << (6 + 7 + 7 + 7)) - 1), nil},

				{40, (1 << (6 + 7 + 7 + 7)), nil},
				{40, -(1 << (6 + 7 + 7 + 7)), nil},
				{40, (1 << (6 + 7 + 7 + 7 + 7)) - 1, nil},
				{40, -((1 << (6 + 7 + 7 + 7 + 7)) - 1), nil},

				{48, (1 << (6 + 7 + 7 + 7 + 7)), nil},
				{48, -(1 << (6 + 7 + 7 + 7 + 7)), nil},
				{48, (1 << (6 + 7 + 7 + 7 + 7 + 7)) - 1, nil},
				{48, -((1 << (6 + 7 + 7 + 7 + 7 + 7)) - 1), nil},

				{56, (1 << (6 + 7 + 7 + 7 + 7 + 7)), nil},
				{56, -(1 << (6 + 7 + 7 + 7 + 7 + 7)), nil},
				{56, (1 << (6 + 7 + 7 + 7 + 7 + 7 + 7)) - 1, nil},
				{56, -((1 << (6 + 7 + 7 + 7 + 7 + 7 + 7)) - 1), nil},

				{64, (1 << (6 + 7 + 7 + 7 + 7 + 7 + 7)), nil},
				{64, -(1 << (6 + 7 + 7 + 7 + 7 + 7 + 7)), nil},
				{64, (1 << (6 + 7 + 7 + 7 + 7 + 7 + 7 + 8)) - 1, nil},
				{64, -((1 << (6 + 7 + 7 + 7 + 7 + 7 + 7 + 8)) - 1), nil},

				{0, -(1 << (6 + 7 + 7 + 7 + 7 + 7 + 7 + 8)), ztype.ErrOutOfBounds},
				{0, (1 << (6 + 7 + 7 + 7 + 7 + 7 + 7 + 8)), ztype.ErrOutOfBounds},
			},
		},
		"int": {
			// The different notation used in block is a to keep the test in
			// sync with the ndsev implementation.
			maxBytes: 9,
			items: []testitem{
				{8, 0, nil},
				{8, -(1 << 6) + 1, nil},
				{8, (1 << 6) - 1, nil},
				{16, -(1 << 6), nil},
				{16, (1 << 6), nil},
				{16, -(1 << 13) + 1, nil},
				{16, (1 << 13) - 1, nil},
				{24, -(1 << 13), nil},
				{24, 1 << 13, nil},
				{24, -(1 << 20) + 1, nil},
				{24, (1 << 20) - 1, nil},
				{32, -(1 << 20), nil},
				{32, 1 << 20, nil},
				{32, -(1 << 27) + 1, nil},
				{32, (1 << 27) - 1, nil},
				{40, -(1 << 27), nil},
				{40, 1 << 27, nil},
				{40, -(1 << 34) + 1, nil},
				{40, (1 << 34) - 1, nil},
				{48, -(1 << 34), nil},
				{48, 1 << 34, nil},
				{48, -(1 << 41) + 1, nil},
				{48, (1 << 41) - 1, nil},
				{56, -(1 << 41), nil},
				{56, 1 << 41, nil},
				{56, -(1 << 48) + 1, nil},
				{56, (1 << 48) - 1, nil},
				{64, -(1 << 48), nil},
				{64, 1 << 48, nil},
				{64, -(1 << 55) + 1, nil},
				{64, (1 << 55) - 1, nil},
				{72, -(1 << 55), nil},
				{72, 1 << 55, nil},
				{72, -(1 << 63) + 1, nil},
				{72, (1 << 63) - 1, nil},
				{8, -(1 << 63), nil},
				// Go can not support 128-bit uints, so we do not include the error test for 1<<64
				// {8, -(1 << 63) - 1, nztype.ErrSizeToLargeil},
				// {0, 1 << 63, ztype.ErrSizeToLarge},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			for _, tt := range test.items {
				got, err := ztype.SignedBitSize(tt.value, test.maxBytes)
				if diff := cmp.Diff(tt.err, err, cmpopts.EquateErrors()); diff != "" {
					t.Errorf("incorrect error for %d: %s", tt.value, diff)
				} else if got != tt.expected {
					t.Errorf("incorrect bitsize for %d: got %d, want %d", tt.value, got, tt.expected)
				}
			}
		})
	}
}
