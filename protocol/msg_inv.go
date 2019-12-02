package protocol

import (
	"io"

	"github.com/Jeiwan/tinybit/binary"
)

const (
	DataObjectError = iota
	DataObjectTx
	DataObjectBlock
	DataObjectFilterBlock
	DataObjectCmpctBlock
)

// MsgInv represents 'inv' message.
type MsgInv struct {
	Count     uint8 // TODO: Change to var_int
	Inventory []InvVector
}

// UnmarshalBinary implements binary.Unmarshaler interface.
func (inv *MsgInv) UnmarshalBinary(r io.Reader) error {
	d := binary.NewDecoder(r)

	if err := d.Decode(&inv.Count); err != nil {
		return err
	}

	for i := uint8(0); i < inv.Count; i++ {
		var v InvVector

		if err := d.Decode(&v); err != nil {
			return err
		}

		inv.Inventory = append(inv.Inventory, v)
	}

	return nil
}

// InvVector represents inventory vector.
type InvVector struct {
	Type uint32
	Hash [32]byte
}
