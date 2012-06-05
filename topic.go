package main

var clientCount = 0
var clients map[int]chan string = make(map[int]chan string)

type client struct {
	id      int
	channel chan string
	cb      func(s string)
}

func Suscribe(cb func(s string)) client {
	clientCount += 1
	ch := make(chan string)
	clients[clientCount] = ch
	cl := client{clientCount, ch, cb}
	go cl.loop()
	return cl
}

func (c *client) Leave() {
	delete(clients, c.id)
	close(c.channel)
}

func (c *client) loop() {
	var s string
	for {
		s = <-c.channel
		c.cb(s)
	}
}

func Publish(s string) {
	for i := range clients {
		clients[i] <- s
	}
}
