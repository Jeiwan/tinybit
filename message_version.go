package main

import (
	"bytes"
	"encoding/binary"
)

type msgVersion struct {
	version     int32
	services    uint64
	timestamp   int64
	addrRecv    netAddr
	addrFrom    netAddr
	nonce       uint64
	userAgent   varStr
	startHeight int32
	relay       bool
}

func (v msgVersion) serialize() ([]byte, error) {
	var buf bytes.Buffer

	if err := binary.Write(&buf, binary.LittleEndian, v.version); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.LittleEndian, v.services); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.LittleEndian, v.timestamp); err != nil {
		return nil, err
	}

	serializedAddrRecv, err := v.addrRecv.serialize()
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(serializedAddrRecv); err != nil {
		return nil, err
	}

	serializedAddrFrom, err := v.addrFrom.serialize()
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(serializedAddrFrom); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.LittleEndian, v.nonce); err != nil {
		return nil, err
	}

	serializedUserAgent, err := v.userAgent.serialize()
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(serializedUserAgent); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.LittleEndian, v.startHeight); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.LittleEndian, v.relay); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
