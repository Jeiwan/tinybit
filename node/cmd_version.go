package node

import (
	"io"

	"github.com/Jeiwan/tinybit/binary"
	"github.com/Jeiwan/tinybit/protocol"
	"github.com/sirupsen/logrus"
)

func (n Node) handleVersion(header *protocol.MessageHeader, conn io.Reader) error {
	var version protocol.MsgVersion

	lr := io.LimitReader(conn, int64(header.Length))
	if err := binary.NewDecoder(lr).Decode(&version); err != nil {
		return err
	}

	logrus.Infof("VERSION: %+v", version.UserAgent.String)

	return nil
}
