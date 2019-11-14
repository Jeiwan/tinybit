package main

import (
	"os"

	"github.com/Jeiwan/tinybit/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	if os.Getenv("DEBUG") != "" {
		logrus.SetLevel(logrus.DebugLevel)
	}

	cmd.Execute()
}
