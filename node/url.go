package node

import (
	"errors"
	"math"
	"net"
	"strconv"
	"strings"

	"github.com/Jeiwan/tinybit/protocol"
)

// Addr ...
type Addr struct {
	IP   protocol.IPv4
	Port uint16
}

// ParseNodeAddr ...
func ParseNodeAddr(nodeAddr string) (*Addr, error) {
	parts := strings.Split(nodeAddr, ":")
	if len(parts) != 2 {
		return nil, errors.New("malformed node address")
	}

	hostnamePart := parts[0]
	portPart := parts[1]
	if hostnamePart == "" || portPart == "" {
		return nil, errors.New("malformed node address")
	}

	port, err := strconv.Atoi(portPart)
	if err != nil {
		return nil, errors.New("malformed node address")
	}

	if port < 0 || port > math.MaxUint16 {
		return nil, errors.New("malformed node address")
	}

	var addr Addr
	ip := net.ParseIP(hostnamePart)
	copy(addr.IP[:], []byte(ip.To4()))

	addr.Port = uint16(port)

	return &addr, nil
}
