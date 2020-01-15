package protocol

import (
	"errors"
	"io"
	"math"

	"github.com/Jeiwan/tinybit/binary"
)

var errInvalidVarIntValue = errors.New("invalid varint value")

// VarInt is variable length integer.
type VarInt struct {
	value interface{}
}

// Int returns returns value as 'int'.
func (vi VarInt) Int() (int, error) {
	switch v := vi.value.(type) {
	case uint8:
		return int(v), nil

	case uint16:
		return int(v), nil

	case uint32:
		return int(v), nil

	case uint64:
		// Assume we'll never get value more than MaxInt64.
		if v > math.MaxInt64 {
			return math.MaxInt64, nil
		}

		return int(v), nil
	}

	return 0, errInvalidVarIntValue
}

// UnmarshalBinary implements binary.Unmarshaler interface.
func (vi *VarInt) UnmarshalBinary(r io.Reader) error {
	var b uint8

	lr := io.LimitReader(r, 1)
	if err := binary.NewDecoder(lr).Decode(&b); err != nil {
		return err
	}

	if b < 0xFD {
		vi.value = b
		return nil
	}

	if b == 0xFD {
		var v uint16
		lr := io.LimitReader(r, 2)
		if err := binary.NewDecoder(lr).Decode(&v); err != nil {
			return err
		}

		vi.value = v
		return nil
	}

	if b == 0xFE {
		var v uint32
		lr := io.LimitReader(r, 4)
		if err := binary.NewDecoder(lr).Decode(&v); err != nil {
			return err
		}

		vi.value = v
		return nil
	}

	if b == 0xFF {
		var v uint64
		lr := io.LimitReader(r, 8)
		if err := binary.NewDecoder(lr).Decode(&v); err != nil {
			return err
		}

		vi.value = v
		return nil
	}

	return nil
}
