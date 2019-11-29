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

	peer := Peer{
		Connection: conn,
		Services:   version.Services,
		IP:         version.AddrFrom.IP,
		Port:       version.AddrFrom.Port,
		UserAgent:  version.UserAgent.String,
		Version:    version.Version,
	}

	n.Peers = append(n.Peers, peer)

	logrus.Debugf("new peer %s (%s:%d)", peer.UserAgent, peer.IP.String(), peer.Port)

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
