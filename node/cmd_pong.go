package node

import (
	"io"

	"github.com/Jeiwan/tinybit/binary"
	"github.com/Jeiwan/tinybit/protocol"
	"github.com/sirupsen/logrus"
)

func (n Node) handlePong(header *protocol.MessageHeader, conn io.ReadWriter) error {
	var pong protocol.MsgPing

	logrus.Debugln("received pong")

	lr := io.LimitReader(conn, int64(header.Length))
	if err := binary.NewDecoder(lr).Decode(&pong); err != nil {
		return err
	}

	n.PongCh <- pong.Nonce

	return nil
}
