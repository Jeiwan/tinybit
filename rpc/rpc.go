package rpc

import (
	"fmt"

	"github.com/Jeiwan/tinybit/protocol"
)

// Node defines the interface of interaction between JSON-RPC server and a node.
type Node interface {
	Mempool() map[string]*protocol.MsgTx
}

// RPC implements RPC interface of the node.
type RPC struct {
	node Node
}

// MempoolArgs are arguments of Mempool method.
type MempoolArgs interface{}

// MempoolReply is reply of Mempool method.
type MempoolReply string

// GetMempool returns current mempool state information.
func (r RPC) GetMempool(args *MempoolArgs, reply *MempoolReply) error {
	txs := r.node.Mempool()

	*reply = MempoolReply(formatMempoolReply(txs))

	return nil
}

func formatMempoolReply(txs map[string]*protocol.MsgTx) string {
	var result string

	for k := range txs {
		result += fmt.Sprintf("%s\n", k)
	}
	result += fmt.Sprintf("Total %d transactions", len(txs))

	return result
}
