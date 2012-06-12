package main

var clientCount = 0
var clients map[int]chan msg = make(map[int]chan msg)

type client struct {
	id      int
	channel chan msg
}

func Suscribe() client {
	clientCount += 1
	ch := make(chan msg)
	clients[clientCount] = ch
	cl := client{clientCount, ch}
	return cl
}

func (c *client) Leave() {
	delete(clients, c.id)
	close(c.channel)
}

func Publish(m msg) {
	for i := range clients {
		clients[i] <- m
	}
}
