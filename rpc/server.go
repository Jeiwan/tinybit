package rpc

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/sirupsen/logrus"
)

// Server is a JSON-RPC server.
type Server struct {
	port int
	rpc  *rpc.Server
}

// NewServer returns a new Server.
func NewServer(port int, node Node) (*Server, error) {
	rpcs := rpc.NewServer()

	handlers := RPC{node: node}
	if err := rpcs.Register(handlers); err != nil {
		return nil, err
	}

	s := Server{
		port: port,
		rpc:  rpcs,
	}

	return &s, nil
}

// Run runs the Server.
func (s Server) Run() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		logrus.Errorf("failed to run JSON-RPC server: %+v", err)
		return
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			logrus.Errorf("JSON-RPC connection failed: %+v", err)
			return
		}

		go s.rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
