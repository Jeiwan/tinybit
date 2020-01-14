package node

import (
	"io"

	"github.com/Jeiwan/tinybit/binary"
	"github.com/Jeiwan/tinybit/protocol"
	"github.com/sirupsen/logrus"
)

func (no Node) handleBlock(header *protocol.MessageHeader, conn io.ReadWriter) error {
	var block protocol.MsgBlock

	lr := io.LimitReader(conn, int64(header.Length))
	if err := binary.NewDecoder(lr).Decode(&block); err != nil {
		return err
	}

	logrus.Debugf("block: %+v", block)

	return nil
}
