package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// String makes Person a fmt.Stringer.
func (p Person) String() string {
	return fmt.Sprintf("%v (%v años)", p.Name, p.Age)
}

func main() {
	a := Person{"Ana", 42}
	z := Person{"Zoe", 9}
	fmt.Println(a, z)
}
