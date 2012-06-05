package main

import (
    "testing"
)

func TestTopic(t *testing.T) {
    c := Register()
    if c.id != 1 {
        t.Log("bad counter")
        t.Fail()
    }
    d := Register()
    if d.id != 2 {
        t.Log("bad counter")
        t.Fail()
    }
    Publish("popo")
    Publish("again")
    c.Leave()
    d.Leave()
}

