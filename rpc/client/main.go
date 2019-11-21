package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	rpcdemo "practice/rpc"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)
	var result float64
	err = client.Call("DemoService.Div", rpcdemo.Args{25, 100}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	err = client.Call("DemoService.Div", rpcdemo.Args{25, 0}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
