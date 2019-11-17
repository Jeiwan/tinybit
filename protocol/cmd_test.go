package protocol

import "testing"

func TestNewCommand(t *testing.T) {
	tests := []struct {
		input    string
		expected [commandLength]byte
		fails    bool
	}{
		{input: "test", expected: [commandLength]byte{0x74, 0x65, 0x73, 0x74, 0, 0, 0, 0, 0, 0, 0, 0}, fails: false},
	}

	for _, test := range tests {
		t.Run(test.input, func(tt *testing.T) {
			actual := newCommand(test.input)

			if actual != test.expected {
				t.Errorf("expected %x, got %x", test.expected, actual)
			}
		})
	}
}
