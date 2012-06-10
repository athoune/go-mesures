package main

import (
	"fmt"
	"net"
)

var Mesures map[string]int64 = make(map[string]int64)

func main() {
	go startSocket()
	startHttp()
}

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
			Mesures[kv.key] = kv.value
			Publish(kv)
		} else {
			conn.Write([]byte("error bad parsing\n"))
		}
	}
}
