package main

import (
	"fmt"
	"testing"
)

func TestTopic(t *testing.T) {
	cFlag := true
	dFlag := true
	c := Suscribe(func(s string) {
		fmt.Println("C", s)
		if cFlag {
			cFlag = false
		}
	})
	if c.id != 1 {
		t.Log("bad counter")
		t.Fail()
	}
	d := Suscribe(func(s string) {
		fmt.Println("D", s)
		if dFlag {
			dFlag = false
		}
	})
	if d.id != 2 {
		t.Log("bad counter")
		t.Fail()
	}
	Publish("popo")
	c.Leave()
	Publish("again")
	d.Leave()
    //[FIXME] a new channel must be used to wait all message
	if cFlag {
		t.Log("Trouble with closure")
		t.Fail()
	}
}
