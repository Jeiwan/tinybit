package protocol

import (
	"math/rand"
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
	payload := MsgPing{
		Nonce: rand.Uint64(),
	}

	msg, err := NewMessage("ping", network, payload)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

// NewPongMsg returns a new MsgPong.
func NewPongMsg(network string, nonce uint64) (*Message, error) {
	payload := MsgPong{
		Nonce: nonce,
	}

	msg, err := NewMessage("pong", network, payload)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
