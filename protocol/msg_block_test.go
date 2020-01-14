package protocol

import (
	"testing"
)

func TestMsgBlockVerify(t *testing.T) {
	tests := []struct {
		name     string
		block    MsgBlock
		expected error
	}{
		{name: "ok",
			block: MsgBlock{Version: 536870912,
				PrevBlock:  [32]byte{172, 4, 9, 193, 214, 108, 234, 44, 154, 29, 10, 102, 174, 80, 88, 194, 234, 43, 37, 146, 145, 71, 180, 34, 119, 42, 87, 208, 209, 150, 233, 97},
				MerkleRoot: [32]byte{177, 101, 232, 30, 115, 112, 196, 254, 111, 9, 84, 131, 131, 183, 20, 35, 189, 41, 123, 171, 233, 60, 151, 170, 229, 8, 240, 37, 236, 173, 138, 127},
				Timestamp:  1579005423,
				Bits:       [4]byte{0xff, 0xff, 0x7f, 0x20},
				Nonce:      0,
				TxCount:    1,
				Txs: []MsgTx{
					{Version: 1,
						Flag:      0,
						TxInCount: 1,
						TxIn: []TxInput{
							{PreviousOutput: OutPoint{
								Hash:  [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
								Index: 4294967295},
								ScriptLength:    24,
								SignatureScript: []byte{2, 146, 2, 8, 34, 125, 127, 126, 105, 249, 169, 94, 11, 47, 80, 50, 83, 72, 47, 98, 116, 99, 100, 47},
								Sequence:        4294967295}},
						TxOutCount: 1,
						TxOut: []TxOutput{
							{Value: 5000000000,
								PkScriptLength: 25,
								PkScript:       []byte{118, 169, 20, 146, 142, 130, 144, 21, 113, 202, 194, 219, 252, 254, 137, 153, 27, 24, 23, 160, 102, 44, 148, 136, 172}}},
						TxWitness: TxWitnessData{
							Count:   0,
							Witness: []TxWitness{}},
						LockTime: 0}}},
			expected: nil},
		{name: "invalid hash",
			block: MsgBlock{Version: 536870912,
				PrevBlock:  [32]byte{172, 4, 9, 193, 214, 108, 234, 44, 154, 29, 10, 102, 174, 80, 88, 194, 234, 43, 37, 146, 145, 71, 180, 34, 119, 42, 87, 208, 209, 150, 233, 97},
				MerkleRoot: [32]byte{177, 101, 232, 30, 115, 112, 196, 254, 111, 9, 84, 131, 131, 183, 20, 35, 189, 41, 123, 171, 233, 60, 151, 170, 229, 8, 240, 37, 236, 173, 138, 127},
				Timestamp:  1579005423,
				Bits:       [4]byte{0xff, 0xff, 0x7f, 0x10},
				Nonce:      0,
				TxCount:    1,
				Txs: []MsgTx{
					{Version: 1,
						Flag:      0,
						TxInCount: 1,
						TxIn: []TxInput{
							{PreviousOutput: OutPoint{
								Hash:  [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
								Index: 4294967295},
								ScriptLength:    24,
								SignatureScript: []byte{2, 146, 2, 8, 34, 125, 127, 126, 105, 249, 169, 94, 11, 47, 80, 50, 83, 72, 47, 98, 116, 99, 100, 47},
								Sequence:        4294967295}},
						TxOutCount: 1,
						TxOut: []TxOutput{
							{Value: 5000000000,
								PkScriptLength: 25,
								PkScript:       []byte{118, 169, 20, 146, 142, 130, 144, 21, 113, 202, 194, 219, 252, 254, 137, 153, 27, 24, 23, 160, 102, 44, 148, 136, 172}}},
						TxWitness: TxWitnessData{
							Count:   0,
							Witness: []TxWitness{}},
						LockTime: 0}}},
			expected: errInvalidBlockHash},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			got := test.block.Verify()

			if got != test.expected {
				tt.Errorf("expected: %+v, got: %+v", test.expected, got)
			}
		})
	}
}
