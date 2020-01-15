package node

import (
	"fmt"
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

	hash, err := block.Hash()
	if err != nil {
		return fmt.Errorf("block.Hash: %+v", err)
	}

	logrus.Debugf("block: %+v", hash)

	if err := block.Verify(); err != nil {
		return fmt.Errorf("rejected invalid block %x", hash)
	}

	return nil
}
