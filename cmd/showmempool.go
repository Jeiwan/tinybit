package cmd

import (
	"github.com/Jeiwan/tinybit/rpc"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	showMempoolCmd.Flags().IntVar(&jsonrpcPort, "jsonrpc-port", 9334, "JSON-RPC port to connect to.")
}

var showMempoolCmd = &cobra.Command{
	Use: "showmempool",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := rpc.NewClient(jsonrpcPort)
		if err != nil {
			return err
		}
		defer c.Close()

		var reply string
		if err := c.Call("RPC.Test", nil, &reply); err != nil {
			return err
		}

		logrus.Println(reply)

		return nil
	},
}
