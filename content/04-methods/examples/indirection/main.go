package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Sum() float64 {
	return v.X + v.Y
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2) // Go reads it as (&v).Scale(2)
	fmt.Println(v)

	p := &Vertex{1, 2}
	fmt.Println(p.Sum()) // Go reads it as (*p).Sum()
}
