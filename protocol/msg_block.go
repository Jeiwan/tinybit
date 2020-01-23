package protocol

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"io"
	"math/big"
	"sort"

	"github.com/Jeiwan/tinybit/binary"
)

var errInvalidBlockHash = errors.New("invalid block hash")

// MsgBlock represents 'block' message.
type MsgBlock struct {
	Version    int32
	PrevBlock  [32]byte
	MerkleRoot [32]byte
	Timestamp  uint32
	Bits       [4]byte
	Nonce      uint32
	TxCount    uint8 // TODO: Convert to var_int
	Txs        []MsgTx
}

// Hash calculates and returns block hash.
func (blck MsgBlock) Hash() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	b, err := binary.Marshal(blck.Version)
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(b); err != nil {
		return nil, err
	}

	b, err = binary.Marshal(blck.PrevBlock)
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(b); err != nil {
		return nil, err
	}

	b, err = binary.Marshal(blck.MerkleRoot)
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(b); err != nil {
		return nil, err
	}

	b, err = binary.Marshal(blck.Timestamp)
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(b); err != nil {
		return nil, err
	}

	b, err = binary.Marshal(blck.Bits)
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(b); err != nil {
		return nil, err
	}

	b, err = binary.Marshal(blck.Nonce)
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(b); err != nil {
		return nil, err
	}

	hash := sha256.Sum256(buf.Bytes())
	hash = sha256.Sum256(hash[:])
	blockHash := hash[:]

	sort.SliceStable(blockHash, func(i, j int) bool { return true })

	return blockHash, nil
}

// UnmarshalBinary implements binary.Unmarshaler
func (blck *MsgBlock) UnmarshalBinary(r io.Reader) error {
	d := binary.NewDecoder(r)

	if err := d.Decode(&blck.Version); err != nil {
		return err
	}

	if err := d.Decode(&blck.PrevBlock); err != nil {
		return err
	}

	if err := d.Decode(&blck.MerkleRoot); err != nil {
		return err
	}

	if err := d.Decode(&blck.Timestamp); err != nil {
		return err
	}

	if err := d.Decode(&blck.Bits); err != nil {
		return err
	}

	if err := d.Decode(&blck.Nonce); err != nil {
		return err
	}

	if err := d.Decode(&blck.TxCount); err != nil {
		return err
	}

	for i := uint8(0); i < blck.TxCount; i++ {
		var tx MsgTx

		if err := d.Decode(&tx); err != nil {
			return err
		}

		blck.Txs = append(blck.Txs, tx)
	}

	return nil
}

func (blck MsgBlock) unpackBits() []byte {
	bits := make([]byte, len(blck.Bits))
	copy(bits, blck.Bits[:])
	sort.SliceStable(bits, func(i, j int) bool { return true })

	target := make([]byte, 32)
	i := 32 - bits[0]
	target[i] = bits[1]
	target[i+1] = bits[2]
	target[i+2] = bits[3]

	return target
}

// Verify checks if the block message is valid.
func (blck MsgBlock) Verify() error {
	target := blck.unpackBits()

	hash, err := blck.Hash()
	if err != nil {
		return err
	}

	targetNum := big.NewInt(0).SetBytes(target)
	hashNum := big.NewInt(0).SetBytes(hash)

	// Block hash must be <= target threshold
	if hashNum.Cmp(targetNum) > 0 {
		return errInvalidBlockHash
	}

	return nil
}
