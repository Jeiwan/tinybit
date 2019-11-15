package node

import (
	"io"

	"github.com/Jeiwan/tinybit/binary"
	"github.com/Jeiwan/tinybit/protocol"
	"github.com/sirupsen/logrus"
)

func (n Node) handleVersion(header *protocol.MessageHeader, conn io.ReadWriter) error {
	var version protocol.MsgVersion

	lr := io.LimitReader(conn, int64(header.Length))
	if err := binary.NewDecoder(lr).Decode(&version); err != nil {
		return err
	}

	logrus.Infof("VERSION: %+v", version.UserAgent.String)

	verack, err := protocol.NewVerackMsg(n.Network)
	if err != nil {
		return err
	}

	msg, err := binary.Marshal(verack)
	if err != nil {
		return err
	}

	if _, err := conn.Write(msg); err != nil {
		return err
	}

	return nil
}
