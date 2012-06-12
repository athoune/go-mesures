package main

import "fmt"

var Mesures map[string]int64 = make(map[string]int64)
var Memento chan msg = make(chan msg)

func startMemento() {
	for {
		m := <-Memento
		fmt.Println(m)
		Mesures[m.key] = m.value
	}
}
