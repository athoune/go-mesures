package main

import (
	"fmt"
	"net"
)

func startSocket() {
	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}
		fmt.Println("Accepting a new connection")
		go doServeStuff(conn)
	}

}

func doServeStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error reading", err.Error())
			return
		}
		kv, ok := parse_cmd(string(buf))
		if ok {
			fmt.Println("Received ", kv.key, kv.value)
			conn.Write([]byte("ok\n"))
			Memento <- kv
			Publish(kv)
		} else {
			conn.Write([]byte("error bad parsing\n"))
		}
	}
}
