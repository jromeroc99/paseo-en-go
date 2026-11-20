package main

import "fmt"

// Vertex is a point on a grid.
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)

	v.X = 4
	fmt.Println(v.X, v.Y)
}
