package cmd

import (
	"os"

	"github.com/Jeiwan/tinybit/node"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	userAgent = "/Satoshi:5.64/tinybit:0.0.1/"
)

var (
	network  string
	nodeAddr string
)

func init() {
	tinybitCmd.Flags().StringVar(&nodeAddr, "node-addr", "127.0.0.1:9333", "TCP address of a Bitcoin node to connect to")
	tinybitCmd.Flags().StringVar(&network, "network", "simnet", "Bitcoin network (simnet, mainnet)")
}

var tinybitCmd = &cobra.Command{
	Use: "tinybit",
	RunE: func(cmd *cobra.Command, args []string) error {
		node, err := node.New(network, userAgent)
		if err != nil {
			return err
		}

		return node.Run(nodeAddr)
	},
}

// Execute ...
func Execute() {
	if err := tinybitCmd.Execute(); err != nil {
		logrus.Fatalln(err)
		os.Exit(1)
	}
}
