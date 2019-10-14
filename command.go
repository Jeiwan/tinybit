package main

import "github.com/sirupsen/logrus"

const (
	cmdPing    = "ping"
	cmdPong    = "pong"
	cmdVersion = "version"
)

var commands = map[string][12]byte{
	cmdPing:    newCommand(cmdPing),
	cmdPong:    newCommand(cmdPong),
	cmdVersion: newCommand(cmdVersion),
}

func newCommand(command string) [12]byte {
	l := len(command)
	if l > 12 {
		logrus.Fatalf("command %s is too long", command)
	}

	var packed [12]byte
	buf := make([]byte, 12-l)
	copy(packed[:], append([]byte(command), buf...)[:])

	return packed
}
