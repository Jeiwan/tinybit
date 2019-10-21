package main

import (
	"io"
	"math/rand"
	"net"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	protocolVerion = 70015
	userAgent      = "/Satoshi:5.64/bitcoin-qt:0.4/"
)

func main() {
	version := msgVersion{
		version:   protocolVerion,
		services:  nodeNetwork,
		timestamp: time.Now().UTC().Unix(),
		addrRecv: netAddr{
			services: nodeNetwork,
			ip:       newIP(127, 0, 0, 1),
			port:     9333,
		},
		addrFrom: netAddr{
			services: nodeNetwork,
			ip:       newIP(127, 0, 0, 1),
			port:     9334,
		},
		nonce:       nonce(),
		userAgent:   newVarStr(userAgent),
		startHeight: 0,
		relay:       false,
	}
	msg, err := newMessage("version", "simnet", version)
	if err != nil {
		logrus.Fatalln(err)
	}

	msgSerialized, err := msg.serialize()
	if err != nil {
		logrus.Fatalln(err)
	}

	conn, err := net.Dial("tcp", "127.0.0.1:9333")
	if err != nil {
		logrus.Fatalln(err)
	}
	defer conn.Close()

	_, err = conn.Write(msgSerialized)
	if err != nil {
		logrus.Fatalln(err)
	}

	var responses [][]byte

	doneCh := make(chan struct{})
	go func() {
		tmp := make([]byte, 256)

		for {
			n, err := conn.Read(tmp)
			if err != nil {
				if err != io.EOF {
					logrus.Fatalln(err)
				}
				doneCh <- struct{}{}
				return
			}
			logrus.Infof("received: %x", tmp[:n])
			responses = append(responses, tmp[:n])
		}
	}()

	<-doneCh

	logrus.Println(responses)
}

func nonce() uint64 {
	return rand.Uint64()
}
