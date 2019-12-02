package protocol

import "github.com/sirupsen/logrus"

const (
	cmdGetData    = "getdata"
	cmdInv        = "inv"
	cmdPing       = "ping"
	cmdPong       = "pong"
	cmdTx         = "tx"
	cmdVerack     = "verack"
	cmdVersion    = "version"
	commandLength = 12
)

var commands = map[string][commandLength]byte{
	cmdGetData: newCommand(cmdGetData),
	cmdInv:     newCommand(cmdInv),
	cmdPing:    newCommand(cmdPing),
	cmdPong:    newCommand(cmdPong),
	cmdTx:      newCommand(cmdTx),
	cmdVerack:  newCommand(cmdVerack),
	cmdVersion: newCommand(cmdVersion),
}

func newCommand(command string) [commandLength]byte {
	l := len(command)
	if l > commandLength {
		logrus.Fatalf("command %s is too long", command)
	}

	var packed [commandLength]byte
	buf := make([]byte, commandLength-l)
	copy(packed[:], append([]byte(command), buf...)[:])

	return packed
}
