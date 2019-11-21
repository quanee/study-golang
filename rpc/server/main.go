package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	rpcdemo "practice/rpc"
)

func main() {
	err := rpc.Register(rpcdemo.DemoService{})
	if err != nil {
		log.Print(err)
	}
	listener, err := net.Listen("tcp", ":8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
