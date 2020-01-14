package rpc

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Client is a JSON-RPC client.
type Client struct {
	conn    net.Conn
	jsonrpc *rpc.Client
}

// NewClient returns a new Client.
func NewClient(port int) (*Client, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		return nil, err
	}

	c := jsonrpc.NewClient(conn)

	client := &Client{
		conn:    conn,
		jsonrpc: c,
	}

	return client, nil
}

// Call calls a remote RPC method.
func (c Client) Call(serviceMethod string, args interface{}, reply interface{}) error {
	return c.jsonrpc.Call(serviceMethod, args, reply)
}

// Close closes a connection.
func (c Client) Close() {
	c.conn.Close()
}
