package protocol

import (
	"fmt"
	"math/rand"

	"github.com/Jeiwan/tinybit/binary"
)

// MsgPing describes 'ping' message.
type MsgPing struct {
	Nonce uint64
}

// MsgPong describes 'pong' message.
type MsgPong struct {
	Nonce uint64
}

// NewPingMsg returns a new MsgPing.
func NewPingMsg(network string) (*Message, error) {
	magic, ok := Networks[network]
	if !ok {
		return nil, fmt.Errorf("unsupported network '%s'", network)
	}

	body := MsgPing{
		Nonce: rand.Uint64(),
	}

	serialized, err := binary.Marshal(body)
	if err != nil {
		return nil, err
	}

	head := MessageHeader{
		Magic:    magic,
		Command:  newCommand("ping"),
		Length:   uint32(len(serialized)),
		Checksum: checksum(serialized),
	}

	msg := Message{
		MessageHeader: head,
		Payload:       serialized,
	}

	return &msg, nil
}

// NewPongMsg returns a new MsgPong.
func NewPongMsg(network string, nonce uint64) (*Message, error) {
	magic, ok := Networks[network]
	if !ok {
		return nil, fmt.Errorf("unsupported network '%s'", network)
	}

	body := MsgPong{
		Nonce: nonce,
	}

	serialized, err := binary.Marshal(body)
	if err != nil {
		return nil, err
	}

	head := MessageHeader{
		Magic:    magic,
		Command:  newCommand("pong"),
		Length:   uint32(len(serialized)),
		Checksum: checksum(serialized),
	}

	msg := Message{
		MessageHeader: head,
		Payload:       serialized,
	}

	return &msg, nil
}
