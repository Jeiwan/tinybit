package binary_test

import (
	"bytes"
	"math"
	"reflect"
	"testing"

	"github.com/Jeiwan/tinybit/binary"
)

type customUnmarshaler struct {
	Value interface{}
}

func (u customUnmarshaler) UnmarshalBinary(data []byte) error {
	u.Value = data
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
			input:    []byte{0x00},
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

		{name: "command",
			input:    []byte{0xde, 0xad, 0xbe, 0xef, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			err:      nil,
			actual:   func() interface{} { var x [12]byte; return &x },
			expected: [12]byte{0xde, 0xad, 0xbe, 0xef, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},

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

			if !reflect.DeepEqual(actual, test.expected) {
				tt.Errorf(
					"expected: %v (%v), actual %v (%v)",
					test.expected,
					reflect.TypeOf(test.expected),
					actual,
					reflect.TypeOf(actual),
				)
			}
		})
	}
}
