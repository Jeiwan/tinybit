package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

const (
	nodeNetwork = 1
)

var (
	magicMainnet = [4]byte{0xf9, 0xbe, 0xb4, 0xd9}
	magicSimnet  = [4]byte{0x16, 0x1c, 0x14, 0x12}
)

type messagePayload interface {
	serialize() ([]byte, error)
}

type message struct {
	magic    [4]byte
	command  [12]byte
	length   uint32
	checksum [4]byte
	payload  []byte
}

func newMessage(cmd, network string, payload messagePayload) (*message, error) {
	serializedPayload, err := payload.serialize()
	if err != nil {
		return nil, err
	}

	command, ok := commands[cmd]
	if !ok {
		return nil, fmt.Errorf("unsupported command %s", cmd)
	}

	msg := message{
		magic:    magicSimnet,
		command:  command,
		length:   uint32(len(serializedPayload)),
		checksum: checksum(serializedPayload),
		payload:  serializedPayload,
	}

	return &msg, nil
}

type varStr struct {
	Length uint8
	String string
}

func newVarStr(str string) varStr {
	return varStr{
		Length: uint8(len(str)), // TODO: implement var_int
		String: str,
	}
}

func (m message) serialize() ([]byte, error) {
	var buf bytes.Buffer

	if err := binary.Write(&buf, binary.LittleEndian, m.magic); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.LittleEndian, m.command); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.LittleEndian, m.length); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.LittleEndian, m.checksum); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.LittleEndian, m.payload); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (v varStr) serialize() ([]byte, error) {
	var buf bytes.Buffer

	if err := binary.Write(&buf, binary.LittleEndian, v.Length); err != nil {
		return nil, err
	}

	if _, err := buf.Write([]byte(v.String)); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func checksum(data []byte) [4]byte {
	hash := sha256.Sum256(data)
	hash = sha256.Sum256(hash[:])
	var hashArr [4]byte
	copy(hashArr[:], hash[0:4])

	return hashArr

}
