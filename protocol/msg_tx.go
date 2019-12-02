package protocol

import (
	"bytes"
	"fmt"
	"io"

	"github.com/Jeiwan/tinybit/binary"
)

// MsgTx represents 'tx' message.
type MsgTx struct {
	Version    int32
	Flag       uint16
	TxInCount  uint8 // TODO: Convert to var_int
	TxIn       []TxInput
	TxOutCount uint8 // TODO: Convert to var_int
	TxOut      []TxOutput
	TxWitness  TxWitnessData
	LockTime   uint32
}

// UnmarshalBinary implements binary.Unmarshaler
func (tx *MsgTx) UnmarshalBinary(r io.Reader) error {
	d := binary.NewDecoder(r)

	if err := d.Decode(&tx.Version); err != nil {
		return err
	}

	var flagPrefix byte
	if err := d.Decode(&flagPrefix); err != nil {
		return err
	}

	if flagPrefix == 0 {
		var flag byte
		if err := d.Decode(&flag); err != nil {
			return err
		}

		r := bytes.NewBuffer([]byte{flagPrefix, flag})
		if err := binary.NewDecoder(r).Decode(&tx.Flag); err != nil {
			return err
		}

		if err := d.Decode(&tx.TxInCount); err != nil {
			return err
		}
	} else {
		tx.TxInCount = flagPrefix
	}

	for i := uint8(0); i < tx.TxInCount; i++ {
		var txin TxInput

		if err := d.Decode(&txin); err != nil {
			return err
		}

		tx.TxIn = append(tx.TxIn, txin)
	}

	if err := d.Decode(&tx.TxOutCount); err != nil {
		return err
	}

	for i := uint8(0); i < tx.TxOutCount; i++ {
		var txout TxOutput

		if err := d.Decode(&txout); err != nil {
			return err
		}

		tx.TxOut = append(tx.TxOut, txout)
	}

	if tx.Flag == 1 {
		if err := d.Decode(&tx.TxWitness); err != nil {
			return err
		}
	}

	if err := d.Decode(&tx.LockTime); err != nil {
		return err
	}

	return nil
}

// TxInput represents transaction input.
type TxInput struct {
	PreviousOutput  OutPoint
	ScriptLength    uint8 // TODO: Convert to var_int
	SignatureScript []byte
	Sequence        uint32
}

// TxOutput represents transaction output.
type TxOutput struct {
	Value          int64
	PkScriptLength uint8 // TODO: Convert to var_int
	PkScript       []byte
}

// TxWitnessData represents transaction witness data.
type TxWitnessData struct {
	Count   uint8 // TODO: Convert to var_int
	Witness []TxWitness
}

// TxWitness represents a component of transaction witness data.
type TxWitness struct {
	Length uint8 // TODO: Convert to var_int
	Data   []byte
}

// OutPoint represents previous output transaction reference.
type OutPoint struct {
	Hash  [32]byte
	Index uint32
}

// UnmarshalBinary implements binary.Unmarshaler interface.
func (txw *TxWitnessData) UnmarshalBinary(r io.Reader) error {
	d := binary.NewDecoder(r)

	if err := d.Decode(&txw.Count); err != nil {
		return err
	}

	txw.Witness = nil

	for i := uint8(0); i < txw.Count; i++ {
		var w TxWitness

		if err := d.Decode(&w); err != nil {
			return err
		}

		txw.Witness = append(txw.Witness, w)
	}

	return nil
}

// UnmarshalBinary implements binary.Unmarshaler interface.
func (txw *TxWitness) UnmarshalBinary(r io.Reader) error {
	if err := binary.NewDecoder(r).Decode(&txw.Length); err != nil {
		return err
	}

	if txw.Length == 0 {
		return nil
	}

	txw.Data = make([]byte, txw.Length)
	n, err := io.LimitReader(r, int64(txw.Length)).Read(txw.Data)
	if err != nil {
		return err
	}

	if int64(n) != int64(txw.Length) {
		return fmt.Errorf("invalid witness data was read: want %d bytes, got %d bytes", txw.Length, n)
	}

	return nil
}

// UnmarshalBinary implements binary.Unmarshaler interface.
func (txin *TxInput) UnmarshalBinary(r io.Reader) error {
	d := binary.NewDecoder(r)

	if err := d.Decode(&txin.PreviousOutput); err != nil {
		return err
	}

	if err := d.Decode(&txin.ScriptLength); err != nil {
		return err
	}

	if txin.ScriptLength != 0 {
		txin.SignatureScript = make([]byte, txin.ScriptLength)
		n, err := io.LimitReader(r, int64(txin.ScriptLength)).Read(txin.SignatureScript)
		if err != nil {
			return err
		}

		if int64(n) != int64(txin.ScriptLength) {
			return fmt.Errorf("invalid input script was read: want %d bytes, got %d bytes", txin.ScriptLength, n)
		}
	}

	if err := d.Decode(&txin.Sequence); err != nil {
		return err
	}

	return nil
}

// UnmarshalBinary implements binary.Unmarshaler interface.
func (txout *TxOutput) UnmarshalBinary(r io.Reader) error {
	d := binary.NewDecoder(r)

	if err := d.Decode(&txout.Value); err != nil {
		return err
	}

	if err := d.Decode(&txout.PkScriptLength); err != nil {
		return err
	}

	txout.PkScript = make([]byte, txout.PkScriptLength)
	n, err := io.LimitReader(r, int64(txout.PkScriptLength)).Read(txout.PkScript)
	if err != nil {
		return err
	}

	if int64(n) != int64(txout.PkScriptLength) {
		return fmt.Errorf("invalid output script was read: want %d bytes, got %d bytes", txout.PkScriptLength, n)
	}

	return nil
}
