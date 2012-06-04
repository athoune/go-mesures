package main

import (
	"fmt"
	"net"
)

func main() {
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
		key, value, ok := parse_cmd(string(buf))
		if ok {
			fmt.Println("Received ", key, value)
			conn.Write([]byte("ok\n"))
		} else {
			conn.Write([]byte("error bad parsing\n"))
		}
	}
}
