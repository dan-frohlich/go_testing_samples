package hello01

import (
	"strings"
	"testing"
)

func Test_greeting_contains_name(t *testing.T) {
	name := "Gophers"
	greeting := greet(name)
	if !strings.Contains(greeting, name) {
		t.Errorf("expected a greeting witrh [%s] but got [%s].", name, greeting)
	}
}
