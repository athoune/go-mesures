package main

import (
	"testing"
)

func TestTopic(t *testing.T) {
	cFlag := true
	dFlag := true
	done := make(chan bool, 4)
	c := Suscribe()
	go func() {
		for {
			<-c.channel
			done <- true
			if cFlag {
				cFlag = false
			}
		}
	}()
	if c.id != 1 {
		t.Log("bad counter")
		t.Fail()
	}
	d := Suscribe()
	go func() {
		for {
			<-d.channel
			done <- true
			if dFlag {
				dFlag = false
			}
		}
	}()
	if d.id != 2 {
		t.Log("bad counter")
		t.Fail()
	}
	Publish(msg{"popo", 42})
	Publish(msg{"again", 0})
	//waiting for async tasks
	<-done
	<-done
	<-done
	<-done
	c.Leave()
	d.Leave()
	if cFlag {
		t.Log("Trouble with closure")
		t.Fail()
	}
}
