package protocol

// IPv4 ...
type IPv4 [4]byte

// VersionNetAddr ...
type VersionNetAddr struct {
	Services uint64
	IP       *IPv4
	Port     uint16
}

// NewIPv4 ...
func NewIPv4(a, b, c, d uint8) *IPv4 {
	return &IPv4{a, b, c, d}
}

// MarshalBinary implements the binary.Marshaler interface
func (ip IPv4) MarshalBinary() ([]byte, error) {
	return append([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF}, ip[:]...), nil
}
