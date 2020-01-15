package node

import (
	"fmt"
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

	hash, err := tx.Hash()
	if err != nil {
		return fmt.Errorf("tx.Hash: %+v", err)
	}

	logrus.Debugf("transaction: %x", hash)

	if err := tx.Verify(); err != nil {
		return fmt.Errorf("rejected invalid transaction %x", hash)
	}

	no.mempool.NewTxCh <- tx

	return nil
}
