package main

import (
	"fmt"
	"testing"
)

func TestTopic(t *testing.T) {
	c := Suscribe(func(s string) { fmt.Println("C", s) })
	if c.id != 1 {
		t.Log("bad counter")
		t.Fail()
	}
	d := Suscribe(func(s string) { fmt.Println("D", s) })
	if d.id != 2 {
		t.Log("bad counter")
		t.Fail()
	}
	Publish("popo")
	Publish("again")
	c.Leave()
	d.Leave()
}
