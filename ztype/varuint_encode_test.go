package ztype_test

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/woven-planet/go-zserio/zstream"
	"github.com/woven-planet/go-zserio/ztype"
)

func TestWriteVaruint16(t *testing.T) {
	tests := map[string]struct {
		input []uint16
		want  []byte
		err   error
	}{
		"zero": {
			input: []uint16{0},
			want:  []byte{0x00},
		},
		"two-byte-value": {
			input: []uint16{0x80},
			want:  []byte{0x80, 0x80},
		},
		"minimum": {
			input: []uint16{ztype.MinVaruint16},
			want:  []byte{0x00},
		},
		"maximum": {
			input: []uint16{ztype.MaxVaruint16},
			want:  []byte{0xff, 0xff},
		},
		"too-large": {
			input: []uint16{ztype.MaxVaruint16 + 1},
			err:   ztype.ErrOutOfBounds,
		},
		"multi-values": {
			input: []uint16{0x40, 0x08, 0x1ff},
			want: []byte{
				0x40,
				0x08,
				0x81, 0xff,
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			w := zstream.NewWriter(&buf)
			var err error
			for _, n := range test.input {
				err = ztype.WriteVaruint16(w, n)
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

func TestWriteVaruint32(t *testing.T) {
	tests := map[string]struct {
		input []uint32
		want  []byte
		err   error
	}{
		"zero": {
			input: []uint32{0},
			want:  []byte{0x00},
		},
		"two-byte-value": {
			input: []uint32{0x80},
			want:  []byte{0x81, 0x00},
		},
		"minimum": {
			input: []uint32{ztype.MinVaruint32},
			want:  []byte{0x00},
		},
		"maximum": {
			input: []uint32{ztype.MaxVaruint32},
			want:  []byte{0xff, 0xff, 0xff, 0xff},
		},
		"too-large": {
			input: []uint32{ztype.MaxVaruint32 + 1},
			err:   ztype.ErrOutOfBounds,
		},
		"multi-values": {
			input: []uint32{0x40, 0x08, 0x1ff},
			want: []byte{
				0x40,
				0x08,
				0x83, 0x7f,
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			w := zstream.NewWriter(&buf)
			var err error
			for _, n := range test.input {
				err = ztype.WriteVaruint32(w, n)
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

func TestWriteVaruint64(t *testing.T) {
	tests := map[string]struct {
		input []uint64
		want  []byte
		err   error
	}{
		"zero": {
			input: []uint64{0},
			want:  []byte{0x00},
		},
		"two-byte-value": {
			input: []uint64{0x80},
			want:  []byte{0x81, 0x00},
		},
		"minimum": {
			input: []uint64{ztype.MinVaruint64},
			want:  []byte{0x00},
		},
		"maximum": {
			input: []uint64{ztype.MaxVaruint64},
			want:  []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		},
		"too-large": {
			input: []uint64{ztype.MaxVaruint64 + 1},
			err:   ztype.ErrOutOfBounds,
		},
		"multi-values": {
			input: []uint64{0x40, 0x08, 0x1ff},
			want: []byte{
				0x40,
				0x08,
				0x83, 0x7f,
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			w := zstream.NewWriter(&buf)
			var err error
			for _, n := range test.input {
				err = ztype.WriteVaruint64(w, n)
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

func TestWriteVaruint(t *testing.T) {
	tests := map[string]struct {
		input []uint64
		want  []byte
		err   error
	}{
		"zero": {
			input: []uint64{0},
			want:  []byte{0x00},
		},
		"two-byte-value": {
			input: []uint64{0x80},
			want:  []byte{0x81, 0x00},
		},
		"minimum": {
			input: []uint64{ztype.MinVaruint},
			want:  []byte{0x00},
		},
		"maximum": {
			input: []uint64{ztype.MaxVaruint},
			want:  []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		},
		// Note that we do not test for VARUINT_MAX + 1 becausea that does
		// not fit in an uint64, so you can never generate those in Go anyway.
		"multi-values": {
			input: []uint64{0x40, 0x08, 0x1ff},
			want: []byte{
				0x40,
				0x08,
				0x83, 0x7f,
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			w := zstream.NewWriter(&buf)
			var err error
			for _, n := range test.input {
				err = ztype.WriteVaruint(w, n)
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

func TestWriteVarsize(t *testing.T) {
	tests := map[string]struct {
		input []uint64
		want  []byte
		err   error
	}{
		"zero": {
			input: []uint64{0},
			want:  []byte{0x00},
		},
		"two-byte-value": {
			input: []uint64{0x80},
			want:  []byte{0x81, 0x00},
		},
		"minimum": {
			input: []uint64{ztype.MinVarsize},
			want:  []byte{0x00},
		},
		"maximum": {
			input: []uint64{ztype.MaxVarsize},
			want:  []byte{0x83, 0xff, 0xff, 0xff, 0xff},
		},
		"too-large": {
			input: []uint64{ztype.MaxVarsize + 1},
			err:   ztype.ErrOutOfBounds,
		},
		"multi-values": {
			input: []uint64{0x40, 0x08, 0x1ff},
			want: []byte{
				0x40,
				0x08,
				0x83, 0x7f,
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			w := zstream.NewWriter(&buf)
			var err error
			for _, n := range test.input {
				err = ztype.WriteVarsize(w, n)
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

func TestUnsignedBitSize(t *testing.T) {
	// Test values taken from https://github.com/ndsev/zserio/blob/9aac67d3300f1b18db86180a843ed8eaf4c9c296/compiler/extensions/python/runtime/tests/test_bitsizeof.py
	type testitem struct {
		expected int
		value    uint64
		err      error
	}
	tests := map[string]struct {
		maxBytes int
		items    []testitem
	}{
		"uint16": {
			maxBytes: 2,
			items: []testitem{
				{8, 0, nil},
				{8, (1 << 0), nil},
				{8, (1 << 7) - 1, nil},
				{16, (1 << 7), nil},
				{16, (1 << (7 + 8)) - 1, nil},
				{0, (1 << (7 + 8)), ztype.ErrOutOfBounds},
			},
		},
		"uint32": {
			maxBytes: 4,
			items: []testitem{
				{8, 0, nil},
				{8, (1 << 0), nil},
				{8, (1 << 7) - 1, nil},
				{16, (1 << 7), nil},
				{16, (1 << (7 + 7)) - 1, nil},
				{24, (1 << (7 + 7)), nil},
				{24, (1 << (7 + 7 + 7)) - 1, nil},
				{32, (1 << (7 + 7 + 7)), nil},
				{32, (1 << (7 + 7 + 7 + 8)) - 1, nil},
				{0, (1 << (7 + 7 + 7 + 8)), ztype.ErrOutOfBounds},
			},
		},
		"uint64": {
			maxBytes: 8,
			items: []testitem{
				{8, 0, nil},
				{8, (1 << 0), nil},
				{8, (1 << 7) - 1, nil},
				{16, (1 << 7), nil},
				{16, (1 << (7 + 7)) - 1, nil},
				{24, (1 << (7 + 7)), nil},
				{24, (1 << (7 + 7 + 7)) - 1, nil},
				{32, (1 << (7 + 7 + 7)), nil},
				{32, (1 << (7 + 7 + 7 + 7)) - 1, nil},
				{40, (1 << (7 + 7 + 7 + 7)), nil},
				{40, (1 << (7 + 7 + 7 + 7 + 7)) - 1, nil},
				{48, (1 << (7 + 7 + 7 + 7 + 7)), nil},
				{48, (1 << (7 + 7 + 7 + 7 + 7 + 7)) - 1, nil},
				{56, (1 << (7 + 7 + 7 + 7 + 7 + 7)), nil},
				{56, (1 << (7 + 7 + 7 + 7 + 7 + 7 + 7)) - 1, nil},
				{64, (1 << (7 + 7 + 7 + 7 + 7 + 7 + 7)), nil},
				{64, (1 << (7 + 7 + 7 + 7 + 7 + 7 + 7 + 8)) - 1, nil},
				{0, (1 << (7 + 7 + 7 + 7 + 7 + 7 + 7 + 8)), ztype.ErrOutOfBounds},
			},
		},
		"uint": {
			// The different notation used in block is a to keep the test in
			// sync with the ndsev implementation.
			maxBytes: 9,
			items: []testitem{
				{8, 0, nil},
				{8, (1 << 0), nil},
				{8, (1 << 7) - 1, nil},
				{16, (1 << 7), nil},
				{16, (1 << 14) - 1, nil},
				{24, (1 << 14), nil},
				{24, (1 << 21) - 1, nil},
				{32, (1 << 21), nil},
				{32, (1 << 28) - 1, nil},
				{40, (1 << 28), nil},
				{40, (1 << 35) - 1, nil},
				{48, (1 << 35), nil},
				{48, (1 << 42) - 1, nil},
				{56, (1 << 42), nil},
				{56, (1 << 49) - 1, nil},
				{64, (1 << 49), nil},
				{64, (1 << 56) - 1, nil},
				{72, (1 << 56), nil},
				{72, (1 << 64) - 1, nil},
				// Go can not support 128-bit uints, so we do not include the error test for 1<<64
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			for _, tt := range test.items {
				got, err := ztype.UnsignedBitSize(tt.value, test.maxBytes)
				if diff := cmp.Diff(tt.err, err, cmpopts.EquateErrors()); diff != "" {
					t.Errorf("incorrect error for %d: %s", tt.value, diff)
				} else if got != tt.expected {
					t.Errorf("incorrect bitsize for %d: got %d, want %d", tt.value, got, tt.expected)
				}
			}
		})
	}
}

func BenchmarkBitsize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = ztype.UnsignedBitSize((1<<28)-1, 8)
	}
}

func TestNumBits(t *testing.T) {

	tests := map[string]struct {
		expected int
		value    uint64
	}{
		"empty":  {0, 0},
		"one":    {1, 1},
		"random": {9, 264},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := ztype.NumBits(test.value)
			if got != test.expected {
				t.Errorf("incorrect number of bits for %d: got %d, want %d", test.value, got, test.expected)
			}
		})
	}
}
