package protocol

import (
	"errors"
	"fmt"
	"io"
)

// IPv4 ...
type IPv4 [4]byte

// VersionNetAddr ...
type VersionNetAddr struct {
	Services uint64
	IP       IPv4
	Port     uint16
}

// NewIPv4 ...
func NewIPv4(a, b, c, d uint8) IPv4 {
	return IPv4{a, b, c, d}
}

// MarshalBinary implements the binary.Marshaler interface
func (ip IPv4) MarshalBinary() ([]byte, error) {
	return append([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF}, ip[:]...), nil
}

// UnmarshalBinary implements the binary.Marshaler interface
func (ip IPv4) UnmarshalBinary(r io.Reader) error {
	data := make([]byte, 16)
	if _, err := r.Read(data); err != nil {
		return fmt.Errorf("unmarshal IPv4: %+v", err)
	}

	if len(data) != 16 {
		return errors.New("invalid IPv4: wrong length")
	}

	ipv4 := data[12:16]
	copy(ip[:], ipv4)

	return nil
}
