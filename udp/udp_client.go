package main

import (
	"fmt"
	"net"
	"os"
)

//var connMap map[string]*net.Conn

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8003")
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()
	go send(conn)
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("服务器写来：", string(buf[:n]))
	}
}

func send(conn net.Conn) {
	str := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(str) //从键盘读取内容， 放在str
		if err != nil {
			fmt.Println("os.Stdin. err1 = ", err)
			return
		}
		_, err = conn.Write(str[:n])       // 给服务器发送
	}
}
