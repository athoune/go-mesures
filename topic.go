package main

import (
    "fmt"
)

var clientCount = 0
var clients map[int]chan string = make(map[int]chan string)

type client struct {
	id      int
	channel chan string
}

func Register() client {
	clientCount += 1
	ch := make(chan string)
    clients[clientCount] = ch
    cl := client{clientCount, ch}
    go cl.loop()
    return cl
}

func (c *client) Leave() {
	delete(clients, c.id)
}

func (c *client) loop() {
    var s string
    for {
        s = <- c.channel
        fmt.Println(c.id, s)
    }
}
func Publish(s string) {
    for i := range clients {
        clients[i] <- s
    }
}
//[TODO] PublishAsync
