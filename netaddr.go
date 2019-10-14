package main

import (
	"bytes"
	"encoding/binary"
)

type ip struct {
	v4 []byte
}

type netAddr struct {
	time     uint32
	services uint64
	ip       *ip
	port     uint16
}

func newIP(a, b, c, d uint8) *ip {
	return &ip{
		v4: []byte{a, b, c, d},
	}
}

func (na netAddr) serialize() ([]byte, error) {
	var buf bytes.Buffer

	if err := binary.Write(&buf, binary.LittleEndian, na.services); err != nil {
		return nil, err
	}

	if _, err := buf.Write(na.ip.toIPv6()); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.BigEndian, na.port); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (ip ip) toIPv6() []byte {
	return append([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF}, ip.v4...)
}
