package hello02

import "fmt"

//Greeter greets names
type Greeter interface {
	Greet(name string) string
}

type greeter struct{}

func (g *greeter) Greet(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

//CreateGreeter makes a Greeter
func NewGreeter() Greeter {
	return &greeter{}
}
