package main

import "github.com/sirupsen/logrus"

const (
	cmdPing       = "ping"
	cmdPong       = "pong"
	cmdVersion    = "version"
	commandLength = 12
)

var commands = map[string][commandLength]byte{
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
