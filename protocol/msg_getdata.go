package protocol

import (
	"bytes"

	"github.com/Jeiwan/tinybit/binary"
)

// MsgGetData represents 'getdata' message.
type MsgGetData struct {
	Count     uint8 // TODO: Change to var_int
	Inventory []InvVector
}

// MarshalBinary implements binary.Marshaler interface.
func (gd MsgGetData) MarshalBinary() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	b, err := binary.Marshal(gd.Count)
	if err != nil {
		return nil, err
	}

	if _, err := buf.Write(b); err != nil {
		return nil, err
	}

	for _, i := range gd.Inventory {
		b, err := binary.Marshal(i)
		if err != nil {
			return nil, err
		}

		if _, err := buf.Write(b); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}
