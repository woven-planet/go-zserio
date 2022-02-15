package ztype_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/woven-planet/go-zserio"
	"github.com/woven-planet/go-zserio/ztype"
)

func TestVarintArrayEncoding(t *testing.T) {
	tests := map[string]struct {
		arrayData []int64
		arraySize int
		isAuto    bool
		isPacked  bool
		want      []byte
	}{
		"auto-varint-nonpacked-array": {
			arrayData: []int64{15353, 3463467, 45756786, 576756, -3454, 4543, -5, 0, 3453},
			isAuto:    true,
			isPacked:  false,
			want: []byte{0x09, 0x40, 0xF7, 0x79, 0x41, 0xD3, 0xB2, 0x2B, 0x55,
				0xE8, 0xE2, 0x72, 0x63, 0x99, 0x74, 0xDA, 0x7E, 0x63, 0x3F,
				0x85, 0x00, 0x5A, 0x7D},
		},
		"auto-varint-packed-array": {
			arrayData: []int64{2353, 436547, 56774, -3523, 5, -8676879, -123},
			isAuto:    true,
			isPacked:  true,
			want: []byte{0x07, 0x29, 0x18, 0xAD, 0x69, 0x21, 0xA1, 0xDD, 0xA3,
				0x6D, 0xA1, 0x82, 0xE2, 0x48, 0xE6, 0x07, 0xE0, 0x3D, 0x80},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			arrayInstance := ztype.Array[int64, *ztype.VarIntArrayTraits]{
				ArrayTraits: &ztype.VarIntArrayTraits{},
				RawArray:    test.arrayData,
				IsAuto:      test.isAuto,
				IsPacked:    test.isPacked,
			}

			assertEqualAfterMarshaling(t, test.want, &arrayInstance)
		})
	}
}

func TestBitfieldArrayEncoding(t *testing.T) {
	tests := map[string]struct {
		arrayData []uint32
		arraySize int
		isAuto    bool
		isPacked  bool
		want      []byte
	}{
		// uncompressed (u)intxx arrays are encoded as BitFieldArrays, not as
		// varintxx arrays. That means that the binary data is stored in full
		// bit length
		"auto-uint32-nonpacked-array": {
			arrayData: []uint32{346, 345754, 312, 11, 353, 866, 0, 466},
			isAuto:    true,
			want: []byte{0x08, // size info
				0x00, 0x00, 0x01, 0x5A, 0x00, 0x05, 0x46, 0x9A, // entry 0 + 1
				0x00, 0x00, 0x01, 0x38, 0x00, 0x00, 0x00, 0x0B, // entry 2 + 3
				0x00, 0x00, 0x01, 0x61, 0x00, 0x00, 0x03, 0x62,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xD2},
		},
		"auto-uint32-packed-array": {
			arrayData: []uint32{346, 345754, 312, 11, 353, 866, 0, 466},
			isAuto:    true,
			isPacked:  true,
			want: []byte{0x08, // size info
				0xA6, 0x00, 0x00, 0x02, 0xB4, 0xA8, 0xA8, 0x15,
				0x75, 0x3D, 0xFF, 0xDA, 0x60, 0x02, 0xAC, 0x00, 0x40, 0x3F,
				0xF9, 0x3C, 0x00, 0x3A, 0x40},
		},
		// This array is using delta packing. Because all values are around a
		// similar value, delta packing will only store the first value entirely.
		// All subsequent values are just deltas from the previous value.
		"auto-uint32-delta-packed-array": {
			arrayData: []uint32{5474574, 5474553, 5474566, 5474534, 5474566, 5474511, 5474512},
			isAuto:    true,
			isPacked:  true,
			want: []byte{0x07, 0x8C, 0x00, 0xA7, 0x12, 0x1D, 0xAC, 0x6E, 0x04,
				0x12, 0x40, 0x80},
		},
		"fixed-uint32-packed-array": {
			arrayData: []uint32{346, 345754, 987, 0, 56},
			isAuto:    false,
			isPacked:  true,
			arraySize: 5,
			want: []byte{0xA6, 0x00, 0x00, 0x02, 0xB4, 0xA8, 0xA8, 0x15, 0x7A,
				0x83, 0xFF, 0x84, 0xA0, 0x00, 0x70},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			arrayTraits := &ztype.BitFieldArrayTraits[uint32]{
				NumBits: 32,
			}

			arrayInstance := ztype.Array[uint32, *ztype.BitFieldArrayTraits[uint32]]{
				ArrayTraits: arrayTraits,
				RawArray:    test.arrayData,
				IsAuto:      test.isAuto,
				IsPacked:    test.isPacked,
			}

			assertEqualAfterMarshaling(t, test.want, &arrayInstance)
		})
	}
}

func TestVarInt32ArrayEncoding(t *testing.T) {
	tests := map[string]struct {
		arrayData []int32
		arraySize int
		isAuto    bool
		isPacked  bool
		want      []byte
	}{
		"auto-varint32-nonpacked-array": {
			arrayData: []int32{346, 345754, -312, 11, 353, 866, 0, -466},
			isAuto:    true,
			isPacked:  false,
			want: []byte{0x08, 0x42, 0x5A, 0x55, 0x8D, 0x1A, 0xC2, 0x38, 0x0B,
				0x42, 0x61, 0x46, 0x62, 0x00, 0xC3, 0x52},
		},
		"auto-varint32-packed-array": {
			arrayData: []int32{346, 345754, -312, 11, 353, 866, 0, -466},
			isAuto:    true,
			isPacked:  true,
			want: []byte{0x08, 0x21, 0x2D, 0x2A, 0xC6, 0x8D, 0x61, 0x1C, 0x05,
				0xA1, 0x30, 0xA3, 0x31, 0x00, 0x61, 0xA9, 0x00},
		},
		"auto-varint32-packed-delta-coded-array": {
			arrayData: []int32{5474574, 5474553, 5474566, 5474534, 5474566, 5474511, 5474512},
			isAuto:    true,
			isPacked:  true,
			want: []byte{0x07, 0x8C, 0x83, 0x4F, 0x12, 0x1D, 0xAC, 0x6E, 0x04,
				0x12, 0x40, 0x80},
		},
		"fixed-varint32-packed-array": {
			arrayData: []int32{346, -345754, 987, 0, -56, 436346},
			isAuto:    false,
			arraySize: 6,
			isPacked:  true,
			want: []byte{0x21, 0x2D, 0x6A, 0xC6, 0x8D, 0x23, 0xAD, 0x80, 0x5C,
				0x2D, 0x68, 0x3D, 0x00},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			arrayInstance := ztype.Array[int32, *ztype.VarInt32ArrayTraits]{
				ArrayTraits: &ztype.VarInt32ArrayTraits{},
				RawArray:    test.arrayData,
				IsAuto:      test.isAuto,
				IsPacked:    test.isPacked,
			}

			assertEqualAfterMarshaling(t, test.want, &arrayInstance)
		})
	}
}

func TestFloat16ArrayEncoding(t *testing.T) {
	tests := map[string]struct {
		arrayData []float32
		arraySize int
		isAuto    bool
		want      []byte
	}{
		"auto-float16-array": {
			arrayData: []float32{0.34594727, 3528, -523.5, -634, 34.21875, 0.0023002625, 0.42529297},
			isAuto:    true,
			want: []byte{0x07, 0x35, 0x89, 0x6a, 0xe4, 0xe0, 0x17, 0xe0, 0xf4,
				0x50, 0x47, 0x18, 0xb6, 0x36, 0xCE},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			arrayInstance := ztype.Array[float32, *ztype.Float16ArrayTraits]{
				ArrayTraits: &ztype.Float16ArrayTraits{},
				RawArray:    test.arrayData,
				IsAuto:      test.isAuto,
			}

			assertEqualAfterMarshaling(t, test.want, &arrayInstance)
		})
	}
}

func TestFloat32ArrayEncoding(t *testing.T) {
	tests := map[string]struct {
		arrayData []float32
		arraySize int
		isAuto    bool
		want      []byte
	}{
		"auto-float32-array": {
			arrayData: []float32{12421.43, 353.44, 0.666, -14.235363},
			isAuto:    true,
			want: []byte{0x04, 0x46, 0x42, 0x15, 0xB8, 0x43, 0xB0, 0xB8, 0x52,
				0x3F, 0x2A, 0x7E, 0xFA, 0xC1, 0x63, 0xC4, 0x0C},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			arrayInstance := ztype.Array[float32, *ztype.Float32ArrayTraits]{
				ArrayTraits: &ztype.Float32ArrayTraits{},
				RawArray:    test.arrayData,
				IsAuto:      test.isAuto,
			}

			assertEqualAfterMarshaling(t, test.want, &arrayInstance)
		})
	}
}

func TestFloat64ArrayEncoding(t *testing.T) {
	tests := map[string]struct {
		arrayData []float64
		arraySize int
		isAuto    bool
		want      []byte
	}{
		"auto-float64-array": {
			arrayData: []float64{46.54754, 54785685689.659569568, 45.742131, 124124.1213915815, 1.644, 45747.457845854, 547547.4565},
			isAuto:    true,
			want: []byte{0x07, 0x40, 0x47, 0x46, 0x15, 0xCA, 0x6C, 0xA0, 0x3C,
				0x42, 0x29, 0x82, 0xF6, 0x71, 0x73, 0x51, 0xB3, 0x40, 0x46,
				0xDE, 0xFE, 0x26, 0x0B, 0x2C, 0x84, 0x40, 0xFE, 0x4D, 0xC1,
				0xF1, 0x38, 0x4C, 0x89, 0x3F, 0xFA, 0x4D, 0xD2, 0xF1, 0xA9,
				0xFB, 0xE7, 0x40, 0xE6, 0x56, 0x6E, 0xA6, 0xAC, 0x59, 0x31,
				0x41, 0x20, 0xB5, 0xB6, 0xE9, 0xBA, 0x5E, 0x35},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			arrayInstance := ztype.Array[float64, *ztype.Float64ArrayTraits]{
				ArrayTraits: &ztype.Float64ArrayTraits{},
				RawArray:    test.arrayData,
				IsAuto:      test.isAuto,
			}

			assertEqualAfterMarshaling(t, test.want, &arrayInstance)
		})
	}
}
func TestStringArrayEncoding(t *testing.T) {
	tests := map[string]struct {
		arrayData []string
		arraySize int
		isAuto    bool
		want      []byte
	}{
		"fixed-size-string-array": {
			arrayData: []string{"try", "eating", "delicious", "kaya", "toast"},
			isAuto:    false,
			want: []byte{0x03, 0x74, 0x72, 0x79, 0x06, 0x65, 0x61, 0x74, 0x69,
				0x6E, 0x67, 0x09, 0x64, 0x65, 0x6C, 0x69, 0x63, 0x69, 0x6F, 0x75,
				0x73, 0x04, 0x6B, 0x61, 0x79, 0x61, 0x05, 0x74, 0x6F, 0x61, 0x73, 0x74},
		},
		"auto-sized-string-array": {
			arrayData: []string{"cantonese", "char", "siu", "is", "also", "very", "delicious"},
			isAuto:    true,
			want: []byte{0x07, 0x09, 0x63, 0x61, 0x6E, 0x74, 0x6F, 0x6E, 0x65, 0x73,
				0x65, 0x04, 0x63, 0x68, 0x61, 0x72, 0x03, 0x73, 0x69, 0x75, 0x02,
				0x69, 0x73, 0x04, 0x61, 0x6C, 0x73, 0x6F, 0x04, 0x76, 0x65, 0x72,
				0x79, 0x09, 0x64, 0x65, 0x6C, 0x69, 0x63, 0x69, 0x6F, 0x75, 0x73},
		},
		"auto-sized-string-with-japanese-characters": {
			arrayData: []string{"日本の", "ラメン", "はとても", "美味しい", "です"},
			isAuto:    true,
			want: []byte{0x05, 0x09, 0xE6, 0x97, 0xA5, 0xE6, 0x9C, 0xAC, 0xE3,
				0x81, 0xAE, 0x09, 0xE3, 0x83, 0xA9, 0xE3, 0x83, 0xA1, 0xE3,
				0x83, 0xB3, 0x0C, 0xE3, 0x81, 0xAF, 0xE3, 0x81, 0xA8, 0xE3,
				0x81, 0xA6, 0xE3, 0x82, 0x82, 0x0C, 0xE7, 0xBE, 0x8E, 0xE5,
				0x91, 0xB3, 0xE3, 0x81, 0x97, 0xE3, 0x81, 0x84, 0x06, 0xE3,
				0x81, 0xA7, 0xE3, 0x81, 0x99},
		},
		"auto-sized-string-with-chinese-characters": {
			arrayData: []string{"福建", "虾面", "非常", "好吃"},
			isAuto:    true,
			want: []byte{0x04, 0x06, 0xE7, 0xA6, 0x8F, 0xE5, 0xBB, 0xBA, 0x06,
				0xE8, 0x99, 0xBE, 0xE9, 0x9D, 0xA2, 0x06, 0xE9, 0x9D, 0x9E,
				0xE5, 0xB8, 0xB8, 0x06, 0xE5, 0xA5, 0xBD, 0xE5, 0x90, 0x83},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			arrayInstance := ztype.Array[string, *ztype.StringArrayTraits]{
				ArrayTraits: &ztype.StringArrayTraits{},
				RawArray:    test.arrayData,
				IsAuto:      test.isAuto,
			}

			assertEqualAfterMarshaling(t, test.want, &arrayInstance)
		})
	}
}

func assertEqualAfterMarshaling(t *testing.T, want []byte, input zserio.Marshaler) {
	got, err := zserio.Marshal(input)
	if err != nil {
		t.Fatalf("error writing array: %v", err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("incorrect encoding: %s", diff)
	}
}
