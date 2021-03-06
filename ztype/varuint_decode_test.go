package ztype_test

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/woven-planet/go-zserio/zstream"
	"github.com/woven-planet/go-zserio/ztype"
)

func TestReadVaruint16(t *testing.T) {
	tests := map[string]struct {
		input []byte
		want  uint16
	}{
		"minimum": {
			input: []byte{0x00},
			want:  ztype.MinVaruint16,
		},
		"maximum": {
			input: []byte{0xff, 0xff},
			want:  ztype.MaxVaruint16,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := zstream.NewReader(bytes.NewBuffer(test.input))
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
			want:  ztype.MinVaruint32,
		},
		"maximum": {
			input: []byte{0xff, 0xff, 0xff, 0xff},
			want:  ztype.MaxVaruint32,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := zstream.NewReader(bytes.NewBuffer(test.input))
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
			want:  ztype.MinVaruint64,
		},
		"maximum": {
			input: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			want:  ztype.MaxVaruint64,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := zstream.NewReader(bytes.NewBuffer(test.input))
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
			want:  ztype.MinVaruint,
		},
		"maximum": {
			input: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			want:  ztype.MaxVaruint,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := zstream.NewReader(bytes.NewBuffer(test.input))
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
			want:  ztype.MinVarsize,
		},
		"maximum": {
			input: []byte{0x83, 0xff, 0xff, 0xff, 0xff},
			want:  ztype.MaxVarsize,
		},
		"tool-large": {
			input: []byte{0xff, 0xff, 0xff, 0xff, 0xff},
			err:   ztype.ErrOutOfBounds,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := zstream.NewReader(bytes.NewBuffer(test.input))
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
