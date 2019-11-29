package node

import (
	"io"

	"github.com/Jeiwan/tinybit/protocol"
	"github.com/sirupsen/logrus"
)

func (n Node) handleVerack(header *protocol.MessageHeader, conn io.ReadWriter) error {
	logrus.Debugln("received 'verack'")

	return nil
}
