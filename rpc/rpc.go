package rpc

// RPC implements RPC interface of the node.
type RPC struct {
}

// MempoolArgs are arguments of Mempool method.
type MempoolArgs interface{}

// MempoolReply is reply of Mempool method.
type MempoolReply string

// Mempool returns current mempool state information.
func (r RPC) Mempool(args *MempoolArgs, reply *MempoolReply) error {
	*reply = "Mempool"

	return nil
}
