package tcp_socket

import (
	"testing"
)

func TestLanguage(t *testing.T) {
	key, value, ok := Parse("popo 42")
	if !ok || key != "popo" || value != 42 {
		t.Log("bad parsing")
		t.Fail()
	}

}
