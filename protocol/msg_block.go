package protocol

import (
	"io"

	"github.com/Jeiwan/tinybit/binary"
)

// MsgBlock represents 'block' message.
type MsgBlock struct {
	Version    int32
	PrevBlock  [32]byte
	MerkleRoot [32]byte
	Timestamp  uint32
	Bits       uint32
	Nonce      uint32
	TxCount    uint8 // TODO: Convert to var_int
	Txs        []MsgTx
}

// UnmarshalBinary implements binary.Unmarshaler
func (b *MsgBlock) UnmarshalBinary(r io.Reader) error {
	d := binary.NewDecoder(r)

	if err := d.Decode(&b.Version); err != nil {
		return err
	}

	if err := d.Decode(&b.PrevBlock); err != nil {
		return err
	}

	if err := d.Decode(&b.MerkleRoot); err != nil {
		return err
	}

	if err := d.Decode(&b.Timestamp); err != nil {
		return err
	}

	if err := d.Decode(&b.Bits); err != nil {
		return err
	}

	if err := d.Decode(&b.Nonce); err != nil {
		return err
	}

	if err := d.Decode(&b.TxCount); err != nil {
		return err
	}

	for i := uint8(0); i < b.TxCount; i++ {
		var tx MsgTx

		if err := d.Decode(&tx); err != nil {
			return err
		}

		b.Txs = append(b.Txs, tx)
	}

	return nil
}
