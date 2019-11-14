package cmd

import (
	"bytes"
	"io"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/Jeiwan/tinybit/binary"
	"github.com/Jeiwan/tinybit/protocol"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	tinybitCmd.Flags().String("node-url", "127.0.0.1:9333", "TCP address of a Bitcoin node to connect to")
	tinybitCmd.Flags().String("network", "simnet", "Bitcoin network (simnet, mainnet)")
}

var tinybitCmd = &cobra.Command{
	Use: "tinybit",
	Run: func(cmd *cobra.Command, args []string) {
		nodeURL, err := cmd.Flags().GetString("node-url")
		if err != nil {
			logrus.Fatalln(err)
		}
		network, err := cmd.Flags().GetString("network")
		if err != nil {
			logrus.Fatalln(err)
		}

		version := protocol.MsgVersion{
			Version:   protocol.Version,
			Services:  protocol.SrvNodeNetwork,
			Timestamp: time.Now().UTC().Unix(),
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
			Nonce:       nonce(),
			UserAgent:   protocol.NewUserAgent(),
			StartHeight: -1,
			Relay:       true,
		}

		msg, err := protocol.NewMessage("version", network, version)
		if err != nil {
			logrus.Fatalln(err)
		}

		msgSerialized, err := binary.Marshal(msg)
		if err != nil {
			logrus.Fatalln(err)
		}

		conn, err := net.Dial("tcp", nodeURL)
		if err != nil {
			logrus.Fatalln(err)
		}
		defer conn.Close()

		_, err = conn.Write(msgSerialized)
		if err != nil {
			logrus.Fatalln(err)
		}

		tmp := make([]byte, protocol.MsgHeaderLength)

		for {
			n, err := conn.Read(tmp)
			if err != nil {
				if err != io.EOF {
					logrus.Fatalln(err)
				}
				return
			}

			logrus.Debugf("received: %x", tmp[:n])
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
				if err := handleVersion(&msgHeader, conn); err != nil {
					logrus.Errorf("failed to handle 'version': %+v", err)
					continue
				}
			}
		}
	},
}

func handleVersion(header *protocol.MessageHeader, conn io.Reader) error {
	var version protocol.MsgVersion

	lr := io.LimitReader(conn, int64(header.Length))
	if err := binary.NewDecoder(lr).Decode(&version); err != nil {
		return err
	}

	logrus.Infof("VERSION: %+v", version.UserAgent.String)

	return nil
}

func nonce() uint64 {
	return rand.Uint64()
}

// Execute ...
func Execute() {
	if err := tinybitCmd.Execute(); err != nil {
		logrus.Fatalln(err)
		os.Exit(1)
	}
}
