package main

import "fmt"

type Describer interface {
	Describe() string
}

type Temp float64

func (t Temp) Describe() string {
	return fmt.Sprintf("%.1f grados", float64(t))
}

type Word string

func (w Word) Describe() string {
	return "la palabra " + string(w)
}

func main() {
	var d Describer
	d = Temp(21.5)
	fmt.Printf("(%v, %T) %s\n", d, d, d.Describe())
	d = Word("gopher")
	fmt.Printf("(%v, %T) %s\n", d, d, d.Describe())
}
