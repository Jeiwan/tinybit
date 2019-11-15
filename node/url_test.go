package node_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Jeiwan/tinybit/node"
	"github.com/Jeiwan/tinybit/protocol"
)

func TestParseNodeAddr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		err      error
		expected *node.Addr
	}{
		{name: "ok",
			input:    "127.0.0.1:8333",
			err:      nil,
			expected: &node.Addr{IP: protocol.NewIPv4(127, 0, 0, 1), Port: 8333}},
		{name: "empty input",
			input:    "",
			err:      errors.New("malformed node address"),
			expected: nil},
		{name: "missing port",
			input:    "127.0.0.1",
			err:      errors.New("malformed node address"),
			expected: nil},
		{name: "missing ip",
			input:    ":1234",
			err:      errors.New("malformed node address"),
			expected: nil},
		{name: "invalid ip",
			input:    "300.300.300.300:1234",
			err:      nil,
			expected: &node.Addr{IP: protocol.NewIPv4(0, 0, 0, 0), Port: 1234}},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			actual, err := node.ParseNodeAddr(test.input)
			if err != nil && test.err == nil {
				t.Errorf("unexpected error: %+v", err)
			}

			if err == nil && test.err != nil {
				t.Errorf("expected error: %+v, got: %+v", err, actual)
			}

			if err != nil && test.err != nil && err.Error() != test.err.Error() {
				t.Errorf("expected error: %+v, got: %+v", err, test.err)
			}

			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected: %+v, actual: %+v", test.expected, actual)
			}
		})
	}
}
