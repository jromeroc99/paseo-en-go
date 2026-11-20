package main

import "fmt"

type Describer interface {
	Describe() string
}

type Box struct {
	Label string
}

// Describe handles a nil receiver instead of panicking.
func (b *Box) Describe() string {
	if b == nil {
		return "<caja vacía>"
	}
	return b.Label
}

func main() {
	var d Describer
	fmt.Println(d == nil)

	var b *Box
	d = b
	fmt.Println(d == nil)
	fmt.Println(d.Describe())
}
