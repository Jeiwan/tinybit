package protocol_test

import (
	"encoding/hex"
	"testing"
	"time"

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
		UserAgent:   protocol.NewUserAgent(),
		StartHeight: -1,
		Relay:       true,
	}
	msg, err := protocol.NewMessage("version", "simnet", version)
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
		return
	}

	msgSerialized, err := msg.Serialize()
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
