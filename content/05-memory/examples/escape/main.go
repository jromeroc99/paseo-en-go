package main

import "fmt"

type Point struct {
	X, Y int
}

func newPoint(x, y int) *Point {
	p := Point{x, y}
	return &p
}

func main() {
	q := newPoint(1, 2)
	fmt.Println(*q)
}
