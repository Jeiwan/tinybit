package node

import (
	"encoding/hex"

	"github.com/Jeiwan/tinybit/protocol"
	"github.com/sirupsen/logrus"
)

// Mempool represents mempool.
type Mempool struct {
	NewBlockCh chan protocol.MsgBlock
	NewTxCh    chan protocol.MsgTx

	txs map[string]*protocol.MsgTx
}

// NewMempool returns a new Mempool.
func NewMempool() *Mempool {
	return &Mempool{
		NewBlockCh: make(chan protocol.MsgBlock),
		NewTxCh:    make(chan protocol.MsgTx),
		txs:        make(map[string]*protocol.MsgTx),
	}
}

// Run starts mempool state handling.
func (m Mempool) Run() {
	for {
		select {
		case tx := <-m.NewTxCh:
			hash, err := tx.Hash()
			if err != nil {
				logrus.Errorf("failed to calculate tx hash: %+v", err)
				break
			}

			txid := hex.EncodeToString(hash)
			m.txs[txid] = &tx
		}
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
