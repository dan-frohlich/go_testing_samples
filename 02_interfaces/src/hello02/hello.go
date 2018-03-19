package hello02

//Greeter foo
type Greeter interface {
	Greet(string) string
}

//NewGreeter  new greeter
func NewGreeter() Greeter {
	return &greeter{}
}

type greeter struct{}

func (g *greeter) Greet(name string) string {
	return "Hello, " + name
}
