package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// ScaleCopy changes a copy of v, so the caller sees nothing.
func (v Vertex) ScaleCopy(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Scale changes the original v through the pointer.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.ScaleCopy(10)
	fmt.Println(v)
	v.Scale(10)
	fmt.Println(v)
}
