package node

import (
	"io"

	"github.com/Jeiwan/tinybit/protocol"
)

// Peer describes a peer node in a network.
type Peer struct {
	Connection io.ReadWriter
	IP         protocol.IPv4
	Port       uint16
	Services   uint64
	UserAgent  string
	Version    int32
}
