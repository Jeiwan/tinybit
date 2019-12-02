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
	Peers        map[string]*Peer
	PingCh       chan peerPing
	PongCh       chan uint64
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
		Peers:        make(map[string]*Peer),
		PingCh:       make(chan peerPing),
		PongCh:       make(chan uint64),
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
		return err
	}

	conn, err := net.Dial("tcp", nodeAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(msgSerialized)
	if err != nil {
		return err
	}

	go no.monitorPeers()

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
		case "pong":
			if err := no.handlePong(&msgHeader, conn); err != nil {
				logrus.Errorf("failed to handle 'pong': %+v", err)
				continue
			}
		case "inv":
			if err := no.handleInv(&msgHeader, conn); err != nil {
				logrus.Errorf("failed to handle 'inv': %+v", err)
				continue
			}
		case "tx":
			if err := no.handleTx(&msgHeader, conn); err != nil {
				logrus.Errorf("failed to handle 'tx': %+v", err)
				continue
			}
		}
	}

	return nil
}

func (no Node) disconnectPeer(peerID string) {
	logrus.Debugf("disconnecting peer %s", peerID)

	peer := no.Peers[peerID]
	if peer == nil {
		return
	}

	peer.Connection.Close()
}
