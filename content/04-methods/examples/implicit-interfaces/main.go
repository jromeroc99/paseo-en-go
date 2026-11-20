package main

import "fmt"

type Greeter interface {
	Greet() string
}

// Person never says it implements Greeter: having Greet is enough.
type Person struct {
	Name string
}

func (p Person) Greet() string {
	return "Hola, soy " + p.Name
}

func main() {
	var g Greeter = Person{"Ana"}
	fmt.Println(g.Greet())
}
