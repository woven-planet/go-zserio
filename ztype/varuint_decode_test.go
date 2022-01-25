package ztype_test

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/icza/bitio"
	"github.com/woven-planet/go-zserio/ztype"
)

func TestReadVaruint16(t *testing.T) {
	tests := map[string]struct {
		input []byte
		want  uint16
	}{
		"minimum": {
			input: []byte{0x00},
			want:  ztype.VARUINT16_MIN,
		},
		"maximum": {
			input: []byte{0xff, 0xff},
			want:  ztype.VARUINT16_MAX,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := bitio.NewReader(bytes.NewBuffer(test.input))
			v, err := ztype.ReadVaruint16(r)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if diff := cmp.Diff(test.want, v); diff != "" {
				t.Errorf("incorrect encoding: %s", diff)
			}
		})
	}
}

func TestReadVaruint32(t *testing.T) {
	tests := map[string]struct {
		input []byte
		want  uint32
	}{
		"minimum": {
			input: []byte{0x00},
			want:  ztype.VARUINT32_MIN,
		},
		"maximum": {
			input: []byte{0xff, 0xff, 0xff, 0xff},
			want:  ztype.VARUINT32_MAX,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := bitio.NewReader(bytes.NewBuffer(test.input))
			v, err := ztype.ReadVaruint32(r)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if diff := cmp.Diff(test.want, v); diff != "" {
				t.Errorf("incorrect encoding: %s", diff)
			}
		})
	}
}

func TestReadVaruint64(t *testing.T) {
	tests := map[string]struct {
		input []byte
		want  uint64
	}{
		"minimum": {
			input: []byte{0x00},
			want:  ztype.VARUINT64_MIN,
		},
		"maximum": {
			input: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			want:  ztype.VARUINT64_MAX,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := bitio.NewReader(bytes.NewBuffer(test.input))
			v, err := ztype.ReadVaruint64(r)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if diff := cmp.Diff(test.want, v); diff != "" {
				t.Errorf("incorrect encoding: %s", diff)
			}
		})
	}
}

func TestReadVaruint(t *testing.T) {
	tests := map[string]struct {
		input []byte
		want  uint64
	}{
		"minimum": {
			input: []byte{0x00},
			want:  ztype.VARUINT_MIN,
		},
		"maximum": {
			input: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			want:  ztype.VARUINT_MAX,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := bitio.NewReader(bytes.NewBuffer(test.input))
			v, err := ztype.ReadVaruint(r)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if diff := cmp.Diff(test.want, v); diff != "" {
				t.Errorf("incorrect encoding: %s", diff)
			}
		})
	}
}

func TestReadVarsize(t *testing.T) {
	tests := map[string]struct {
		input []byte
		want  uint64
		err   error
	}{
		"minimum": {
			input: []byte{0x00},
			want:  ztype.VARSIZE_MIN,
		},
		"maximum": {
			input: []byte{0x83, 0xff, 0xff, 0xff, 0xff},
			want:  ztype.VARSIZE_MAX,
		},
		"tool-large": {
			input: []byte{0xff, 0xff, 0xff, 0xff, 0xff},
			err:   ztype.ErrOutOfBounds,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := bitio.NewReader(bytes.NewBuffer(test.input))
			v, err := ztype.ReadVarsize(r)
			if diff := cmp.Diff(test.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Fatalf("incorrect error: %s", diff)
			}
			if diff := cmp.Diff(test.want, v); diff != "" {
				t.Errorf("incorrect encoding: %s", diff)
			}
		})
	}
}
