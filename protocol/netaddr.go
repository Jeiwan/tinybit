package protocol

import (
	"bytes"
	"encoding/binary"
)

// IPv4 ...
type IPv4 [4]byte

// VersionNetAddr ...
type VersionNetAddr struct {
	Services uint64
	IP       *IPv4
	Port     uint16
}

// NewIPv4 ...
func NewIPv4(a, b, c, d uint8) *IPv4 {
	return &IPv4{a, b, c, d}
}

// Serialize ...
func (na NetAddr) Serialize() ([]byte, error) {
	var buf bytes.Buffer

	if na.Time != 0 {
		if err := binary.Write(&buf, binary.LittleEndian, na.Time); err != nil {
			return nil, err
		}
	}

	if err := binary.Write(&buf, binary.LittleEndian, na.Services); err != nil {
		return nil, err
	}

	if _, err := buf.Write(na.IP.ToIPv6()); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.BigEndian, na.Port); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// ToIPv6 ...
func (ip IPv4) ToIPv6() []byte {
	return append([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF}, ip[:]...)
}
