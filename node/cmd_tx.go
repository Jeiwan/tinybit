package node

import (
	"io"

	"github.com/Jeiwan/tinybit/binary"
	"github.com/Jeiwan/tinybit/protocol"
	"github.com/sirupsen/logrus"
)

func (no Node) handleTx(header *protocol.MessageHeader, conn io.ReadWriter) error {
	var tx protocol.MsgTx

	lr := io.LimitReader(conn, int64(header.Length))
	if err := binary.NewDecoder(lr).Decode(&tx); err != nil {
		return err
	}

	logrus.Debugf("transaction: %+v", tx)

	if err := no.mempool.AddTx(&tx); err != nil {
		logrus.Errorf("failed to add tx to mempool: %+v", err)
	}

	return nil
}
