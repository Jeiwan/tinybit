package node

import (
	"bytes"
	"fmt"
	"io"
	"net"

	"github.com/Jeiwan/tinybit/binary"
	"github.com/Jeiwan/tinybit/protocol"
	"github.com/sirupsen/logrus"
)

// Node implements a Bitcoin node.
type Node struct {
	Network      string
	NetworkMagic protocol.Magic
	Peers        []Peer
	UserAgent    string
}

// New returns a new Node.
func New(network, userAgent string) (*Node, error) {
	networkMagic, ok := protocol.Networks[network]
	if !ok {
		return nil, fmt.Errorf("unsupported network %s", network)
	}

	return &Node{
		Network:      network,
		NetworkMagic: networkMagic,
		UserAgent:    userAgent,
	}, nil
}

// Run starts a node.
func (no Node) Run(nodeAddr string) error {
	peerAddr, err := ParseNodeAddr(nodeAddr)
	if err != nil {
		return err
	}

	version, err := protocol.NewVersionMsg(
		no.Network,
		no.UserAgent,
		peerAddr.IP,
		peerAddr.Port,
	)
	if err != nil {
		return err
	}

	msgSerialized, err := binary.Marshal(version)
	if err != nil {
		logrus.Fatalln(err)
	}

	conn, err := net.Dial("tcp", nodeAddr)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer conn.Close()

	_, err = conn.Write(msgSerialized)
	if err != nil {
		logrus.Fatalln(err)
	}

	tmp := make([]byte, protocol.MsgHeaderLength)

Loop:
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				return err
			}
			break Loop
		}

		logrus.Debugf("received header: %x", tmp[:n])
		var msgHeader protocol.MessageHeader
		if err := binary.NewDecoder(bytes.NewReader(tmp[:n])).Decode(&msgHeader); err != nil {
			logrus.Errorf("invalid header: %+v", err)
			continue
		}

		if err := msgHeader.Validate(); err != nil {
			logrus.Error(err)
			continue
		}

		logrus.Debugf("received message: %s", msgHeader.Command)

		switch msgHeader.CommandString() {
		case "version":
			if err := no.handleVersion(&msgHeader, conn); err != nil {
				logrus.Errorf("failed to handle 'version': %+v", err)
				continue
			}
		case "verack":
			if err := no.handleVerack(&msgHeader, conn); err != nil {
				logrus.Errorf("failed to handle 'verack': %+v", err)
				continue
			}
		case "ping":
			if err := no.handlePing(&msgHeader, conn); err != nil {
				logrus.Errorf("failed to handle 'ping': %+v", err)
				continue
			}
		}
	}

	return nil
}
