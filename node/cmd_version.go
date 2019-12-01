package node

import (
	"io"
	"net"

	"github.com/Jeiwan/tinybit/binary"
	"github.com/Jeiwan/tinybit/protocol"
	"github.com/sirupsen/logrus"
)

func (n Node) handleVersion(header *protocol.MessageHeader, conn net.Conn) error {
	var version protocol.MsgVersion

	lr := io.LimitReader(conn, int64(header.Length))
	if err := binary.NewDecoder(lr).Decode(&version); err != nil {
		return err
	}

	peer := Peer{
		Address:    conn.RemoteAddr(),
		Connection: conn,
		PongCh:     make(chan uint64),
		Services:   version.Services,
		UserAgent:  version.UserAgent.String,
		Version:    version.Version,
	}

	n.Peers[peer.ID()] = &peer
	go n.monitorPeer(&peer)

	logrus.Debugf("new peer %s", peer)

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
