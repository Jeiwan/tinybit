package protocol

import (
	"bytes"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestVarintUnmarshalBinary(t *testing.T) {
	tests := []struct {
		name     string
		raw      []byte
		expected interface{}
		err      error
	}{
		{name: "1 byte min",
			raw:      []byte{0x00},
			expected: uint8(0),
			err:      nil},
		{name: "1 byte max",
			raw:      []byte{0xFC},
			expected: uint8(0xfc),
			err:      nil},
		{name: "2 bytes min",
			raw:      []byte{0xFD, 0x00, 0x00},
			expected: uint16(0x0000),
			err:      nil},
		{name: "2 bytes max",
			raw:      []byte{0xFD, 0xFF, 0xFF},
			expected: uint16(0xFFFF),
			err:      nil},
		{name: "4 bytes min",
			raw:      []byte{0xFE, 0x00, 0x00, 0x00, 0x00},
			expected: uint32(0x00000000),
			err:      nil},
		{name: "4 bytes max",
			raw:      []byte{0xFE, 0xFF, 0xFF, 0xFF, 0xFF},
			expected: uint32(0xFFFFFFFF),
			err:      nil},
		{name: "8 bytes min",
			raw:      []byte{0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			expected: uint64(0x0000000000000000),
			err:      nil},
		{name: "8 bytes max",
			raw:      []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
			expected: uint64(0xFFFFFFFFFFFFFFFF),
			err:      nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			r := bytes.NewBuffer(test.raw)
			varint := VarInt{}
			err := varint.UnmarshalBinary(r)

			if err == nil && test.err != nil {
				tt.Errorf("expected error: %+v, got: %+v", test.err, err)
				return
			}

			if err != nil && test.err == nil {
				tt.Errorf("unexpected error: %+v", err)
				return
			}

			if err != nil && test.err != nil && err != test.err {
				tt.Errorf("expected error: %+v, got: %+v", test.err, err)
				return
			}

			got := varint.value
			if diff := cmp.Diff(test.expected, got); diff != "" {
				tt.Errorf("varint.UnmarshalBinary() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}

func TestVarintInt(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		expected int
		err      error
	}{
		{name: "1 byte min",
			value:    uint8(0),
			expected: 0,
			err:      nil},
		{name: "1 byte max",
			value:    uint8(0xFF),
			expected: 0xFF,
			err:      nil},
		{name: "2 bytes min",
			value:    uint16(0x0000),
			expected: 0,
			err:      nil},
		{name: "2 bytes max",
			value:    uint16(0xFFFF),
			expected: 0xffff,
			err:      nil},
		{name: "4 bytes min",
			value:    uint32(0x00000000),
			expected: 0,
			err:      nil},
		{name: "4 bytes max",
			value:    uint32(0xFFFFFFFF),
			expected: 0xffffffff,
			err:      nil},
		{name: "8 bytes min",
			value:    uint64(0x0000000000000000),
			expected: 0,
			err:      nil},
		{name: "8 bytes max",
			value:    uint64(0xFFFFFFFFFFFFFFFF),
			expected: math.MaxInt64,
			err:      nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			varint := VarInt{value: test.value}
			got, err := varint.Int()

			if err == nil && test.err != nil {
				tt.Errorf("expected error: %+v, got: %+v", test.err, err)
				return
			}

			if err != nil && test.err == nil {
				tt.Errorf("unexpected error: %+v", err)
				return
			}

			if err != nil && test.err != nil && err != test.err {
				tt.Errorf("expected error: %+v, got: %+v", test.err, err)
				return
			}

			if got != test.expected {
				tt.Errorf("expected: %v, got: %v", test.expected, got)
				return
			}
		})
	}
}
