package binary_test

import (
	"bytes"
	"io"
	"math"
	"reflect"
	"testing"

	"github.com/Jeiwan/tinybit/binary"
	"github.com/google/go-cmp/cmp"
)

type customUnmarshaler struct {
	Value int
}

func (u *customUnmarshaler) UnmarshalBinary(r io.Reader) error {
	data := make([]byte, 1)
	if _, err := r.Read(data); err != nil {
		return err
	}

	u.Value = int(data[0])
	return nil
}

func TestUnmarshal(t *testing.T) {

	tests := []struct {
		name     string
		input    []byte
		err      error
		actual   func() interface{}
		expected interface{}
	}{
		{name: "bool true",
			input:    []byte{0x01},
			err:      nil,
			actual:   func() interface{} { var x bool; return &x },
			expected: true},

		{name: "bool false",
			input:    []byte{0},
			err:      nil,
			actual:   func() interface{} { var x bool; return &x },
			expected: false},

		{name: "int32",
			input:    []byte{0xFF, 0xFF, 0xFF, 0x7F},
			err:      nil,
			actual:   func() interface{} { var x int32; return &x },
			expected: int32(math.MaxInt32)},

		{name: "int64",
			input:    []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7F},
			err:      nil,
			actual:   func() interface{} { var x int64; return &x },
			expected: int64(math.MaxInt64)},

		{name: "uint8",
			input:    []byte{0xFF},
			err:      nil,
			actual:   func() interface{} { var x uint8; return &x },
			expected: uint8(math.MaxUint8)},

		{name: "uint16",
			input:    []byte{0xFF, 0xFF},
			err:      nil,
			actual:   func() interface{} { var x uint16; return &x },
			expected: uint16(math.MaxUint16)},

		{name: "uint32",
			input:    []byte{0xFF, 0xFF, 0xFF, 0xFF},
			err:      nil,
			actual:   func() interface{} { var x uint32; return &x },
			expected: uint32(math.MaxUint32)},

		{name: "uint64",
			input:    []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
			err:      nil,
			actual:   func() interface{} { var x uint64; return &x },
			expected: uint64(math.MaxUint64)},

		{name: "magic or checksum",
			input:    []byte{0xde, 0xad, 0xbe, 0xef},
			err:      nil,
			actual:   func() interface{} { var x [4]byte; return &x },
			expected: [4]byte{0xde, 0xad, 0xbe, 0xef}},

		{name: "hash",
			input:    []byte{0x31, 0x7c, 0x14, 0x4a, 0xe5, 0xb5, 0xa2, 0x24, 0x37, 0x0b, 0xd6, 0x8c, 0x92, 0x8b, 0x9f, 0x9e, 0x15, 0x2d, 0x98, 0x29, 0x23, 0x5f, 0xfb, 0xec, 0xec, 0x5e, 0xe6, 0x41, 0x13, 0x66, 0x2f, 0xc4},
			err:      nil,
			actual:   func() interface{} { var x [32]byte; return &x },
			expected: [32]byte{0x31, 0x7c, 0x14, 0x4a, 0xe5, 0xb5, 0xa2, 0x24, 0x37, 0x0b, 0xd6, 0x8c, 0x92, 0x8b, 0x9f, 0x9e, 0x15, 0x2d, 0x98, 0x29, 0x23, 0x5f, 0xfb, 0xec, 0xec, 0x5e, 0xe6, 0x41, 0x13, 0x66, 0x2f, 0xc4}},

		{name: "command",
			input:    []byte{0xde, 0xad, 0xbe, 0xef, 0, 0, 0, 0, 0, 0, 0, 0},
			err:      nil,
			actual:   func() interface{} { var x [12]byte; return &x },
			expected: [12]byte{0xde, 0xad, 0xbe, 0xef, 0, 0, 0, 0, 0, 0, 0, 0}},

		{name: "Unmarshaler",
			input:    []byte{0x01, 0x02, 0x03},
			err:      nil,
			actual:   func() interface{} { var x customUnmarshaler; return &x },
			expected: customUnmarshaler{Value: 1}},

		{name: "struct",
			input: []byte{
				11,
				0, 22,
				0xde, 0xad, 0xbe, 0xef,
			},
			err: nil,
			actual: func() interface{} {
				x := struct {
					A uint8
					B uint16
					C [4]byte
				}{}
				return &x
			},
			expected: struct {
				A uint8
				B uint16
				C [4]byte
			}{11, 22, [4]byte{0xde, 0xad, 0xbe, 0xef}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			actualPtr := test.actual()
			err := binary.NewDecoder(bytes.NewReader(test.input)).Decode(actualPtr)
			actual := reflect.ValueOf(actualPtr).Elem().Interface()

			if err == nil && test.err != nil {
				tt.Errorf("expected error: %v, actual: %v", test.err, actual)
				return
			}

			if err != nil && test.err == nil {
				tt.Errorf("didn't expect an error: %v", err)
				return
			}

			if diff := cmp.Diff(test.expected, actual); diff != "" {
				tt.Errorf("Decode() mismatch(-want +got):\n%s", diff)
				return
			}
		})
	}
}
