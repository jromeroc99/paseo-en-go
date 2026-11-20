package main

import "fmt"

// Vertex is a point on a grid.
type Vertex struct {
	X, Y int
}

func main() {
	v1 := Vertex{1, 2} // every field, in order
	v2 := Vertex{X: 1} // Y takes its zero value
	v3 := Vertex{}     // X and Y take their zero value
	p := &Vertex{1, 2} // a pointer to a new Vertex
	fmt.Println(v1, v2, v3, p)
}
