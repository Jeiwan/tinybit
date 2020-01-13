package node

import (
	"encoding/hex"

	"github.com/Jeiwan/tinybit/protocol"
)

// TransactionID is transaction hash.
type TransactionID string

// Mempool represents mempool.
type Mempool map[TransactionID]*protocol.MsgTx

// NewMempool returns a new Mempool.
func NewMempool() *Mempool { return &Mempool{} }

// AddTx adds transaction to the mempool.
func (m Mempool) AddTx(tx *protocol.MsgTx) error {
	hash, err := tx.Hash()
	if err != nil {
		return err
	}

	k := hex.EncodeToString(hash)
	m[TransactionID(k)] = tx

	return nil
}
