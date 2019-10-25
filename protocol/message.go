package protocol

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

const (
	checksumLength = 4
	nodeNetwork    = 1
	magicLength    = 4
)

var (
	magicMainnet = [magicLength]byte{0xf9, 0xbe, 0xb4, 0xd9}
	magicSimnet  = [magicLength]byte{0x16, 0x1c, 0x14, 0x12}
	networks     = map[string][magicLength]byte{
		"mainnet": magicMainnet,
		"simnet":  magicSimnet,
	}
)

// MessagePayload ...
type MessagePayload interface {
	Serialize() ([]byte, error)
}

// Message ...
type Message struct {
	Magic    [magicLength]byte
	Command  [commandLength]byte
	Length   uint32
	Checksum [checksumLength]byte
	Payload  []byte
}

// NewMessage ...
func NewMessage(cmd, network string, payload MessagePayload) (*Message, error) {
	serializedPayload, err := payload.Serialize()
	if err != nil {
		return nil, err
	}

	command, ok := commands[cmd]
	if !ok {
		return nil, fmt.Errorf("unsupported command %s", cmd)
	}

	magic, ok := networks[network]
	if !ok {
		return nil, fmt.Errorf("unsupported network %s", network)
	}

	msg := Message{
		Magic:    magic,
		Command:  command,
		Length:   uint32(len(serializedPayload)),
		Checksum: checksum(serializedPayload),
		Payload:  serializedPayload,
	}

	return &msg, nil
}

// VarStr ...
type VarStr struct {
	Length uint8
	String string
}

func newVarStr(str string) VarStr {
	return VarStr{
		Length: uint8(len(str)), // TODO: implement var_int
		String: str,
	}
}

// Serialize ...
func (m Message) Serialize() ([]byte, error) {
	var buf bytes.Buffer

	if _, err := buf.Write(m.Magic[:]); err != nil {
		return nil, err
	}

	if _, err := buf.Write(m.Command[:]); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.LittleEndian, m.Length); err != nil {
		return nil, err
	}

	if _, err := buf.Write(m.Checksum[:]); err != nil {
		return nil, err
	}

	if _, err := buf.Write(m.Payload); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Serialize ...
func (v VarStr) Serialize() ([]byte, error) {
	var buf bytes.Buffer

	if err := binary.Write(&buf, binary.LittleEndian, v.Length); err != nil {
		return nil, err
	}

	if _, err := buf.Write([]byte(v.String)); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func checksum(data []byte) [checksumLength]byte {
	hash := sha256.Sum256(data)
	hash = sha256.Sum256(hash[:])
	var hashArr [checksumLength]byte
	copy(hashArr[:], hash[0:checksumLength])

	return hashArr
}
