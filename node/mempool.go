package node

import (
	"encoding/hex"

	"github.com/Jeiwan/tinybit/protocol"
)

// Mempool represents mempool.
type Mempool struct {
	txs map[string]*protocol.MsgTx
}

// NewMempool returns a new Mempool.
func NewMempool() *Mempool {
	return &Mempool{
		txs: make(map[string]*protocol.MsgTx),
	}
}

// AddTx adds transaction to the mempool.
func (m Mempool) AddTx(tx *protocol.MsgTx) error {
	hash, err := tx.Hash()
	if err != nil {
		return err
	}

	k := hex.EncodeToString(hash)
	m.txs[k] = tx

	return nil
}

// Mempool ...
func (n Node) Mempool() map[string]*protocol.MsgTx {
	m := make(map[string]*protocol.MsgTx)

	for k, v := range n.mempool.txs {
		m[string(k)] = v
	}

	return m
}
