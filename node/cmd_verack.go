package node

import (
	"io"

	"github.com/Jeiwan/tinybit/protocol"
)

func (n Node) handleVerack(header *protocol.MessageHeader, conn io.ReadWriter) error {
	return nil
}
