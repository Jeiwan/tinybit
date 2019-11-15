package protocol

import "fmt"

// NewVerackMsg returns a new 'verack' message.
func NewVerackMsg(network string) (*Message, error) {
	magic, ok := Networks[network]
	if !ok {
		return nil, fmt.Errorf("unsupported network '%s'", network)
	}

	head := MessageHeader{
		Magic:    magic,
		Command:  newCommand("verack"),
		Length:   0,
		Checksum: checksum([]byte{}),
	}

	msg := Message{
		MessageHeader: head,
		Payload:       []byte{},
	}

	return &msg, nil
}
