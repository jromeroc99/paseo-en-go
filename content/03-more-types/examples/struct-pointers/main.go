package main

import "fmt"

// Vertex is a point on a grid.
type Vertex struct {
	X int
	Y int
}

// moveCopy receives a copy, so the caller's vertex never changes.
func moveCopy(v Vertex) {
	v.X++
}

// move receives a pointer, so it changes the caller's vertex.
func move(v *Vertex) {
	v.X++
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 10 // the same as (*p).X = 10
	fmt.Println(v)

	moveCopy(v)
	fmt.Println(v)

	move(&v)
	fmt.Println(v)
}
