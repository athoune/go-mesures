package main

import (
	"testing"
)

func TestLanguage(t *testing.T) {
	kv, ok := parse_cmd("popo 42")
	if !ok || kv.key != "popo" || kv.value != 42 {
		t.Log("bad parsing")
		t.Fail()
	}
	kv, ok = parse_cmd("popo42")
	if ok {
		t.Log("parse bad text")
		t.Fail()
	}
}
