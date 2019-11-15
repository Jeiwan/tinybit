package protocol_test

import (
	"encoding/hex"
	"testing"
	"time"

	"github.com/Jeiwan/tinybit/binary"
	"github.com/Jeiwan/tinybit/protocol"
)

func TestMessageSerialization(t *testing.T) {
	version := protocol.MsgVersion{
		Version:   protocol.Version,
		Services:  protocol.SrvNodeNetwork,
		Timestamp: time.Date(2019, 11, 11, 0, 0, 0, 0, time.UTC).Unix(),
		AddrRecv: protocol.VersionNetAddr{
			Services: protocol.SrvNodeNetwork,
			IP:       protocol.NewIPv4(127, 0, 0, 1),
			Port:     9333,
		},
		AddrFrom: protocol.VersionNetAddr{
			Services: protocol.SrvNodeNetwork,
			IP:       protocol.NewIPv4(127, 0, 0, 1),
			Port:     9334,
		},
		Nonce:       31337,
		UserAgent:   protocol.NewUserAgent("/Satoshi:5.64/tinybit:0.0.1/"),
		StartHeight: -1,
		Relay:       true,
	}
	msg, err := protocol.NewMessage("version", "simnet", version)
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
		return
	}

	msgSerialized, err := binary.Marshal(msg)
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
		return
	}

	actual := hex.EncodeToString(msgSerialized)
	expected := "161c141276657273696f6e000000000072000000463d41ca7f110100010000000000000080a4c85d00000000010000000000000000000000000000000000ffff7f0000012475010000000000000000000000000000000000ffff7f0000012476697a0000000000001c2f5361746f7368693a352e36342f74696e796269743a302e302e312fffffffff01"
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}

}

func TestHasValidCommand(t *testing.T) {
	tests := []struct {
		name     string
		expected bool
	}{
		{"version", true},
		{"invalid", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			var packed [12]byte
			buf := make([]byte, 12-len(test.name))
			copy(packed[:], append([]byte(test.name), buf...)[:])

			mh := protocol.MessageHeader{
				Command: packed,
			}
			actual := mh.HasValidCommand()

			if actual != test.expected {
				t.Errorf("expected: %v, actual: %v", test.expected, actual)
			}
		})
	}
}

func TestHasValidMagic(t *testing.T) {
	magicMainnet := [4]byte{0xf9, 0xbe, 0xb4, 0xd9}
	magicSimnet := [4]byte{0x16, 0x1c, 0x14, 0x12}

	tests := []struct {
		name     string
		magic    [4]byte
		expected bool
	}{
		{"mainner", magicMainnet, true},
		{"simner", magicSimnet, true},
		{"invalid", [4]byte{0xde, 0xad, 0xbe, 0xef}, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			mh := protocol.MessageHeader{
				Magic: test.magic,
			}
			actual := mh.HasValidMagic()

			if actual != test.expected {
				t.Errorf("expected: %v, actual: %v", test.expected, actual)
			}
		})
	}
}
