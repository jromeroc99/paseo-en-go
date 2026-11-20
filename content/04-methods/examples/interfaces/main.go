package main

import (
	"fmt"
	"math"
)

// Abser is anything that knows its absolute value.
type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	return math.Abs(float64(f))
}

type Vertex struct{ X, Y float64 }

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser = MyFloat(-1.5)
	fmt.Println(a.Abs())
	a = Vertex{3, 4}
	fmt.Println(a.Abs())
}
