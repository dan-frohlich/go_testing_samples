package hello02

import (
	"strings"
	"testing"
)

func Test_grater_uses_name(t *testing.T) {
	grtr := NewGreeter()
	who := `Gophers`
	greeting := grtr.Greet(who)
	if !strings.Contains(greeting, who) {
		t.Errorf("greeting [%s] did not contains name [%s]", greeting, who)
	}
}
