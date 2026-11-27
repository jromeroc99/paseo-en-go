package main

import "fmt"

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// show prints the values of t in order: left, the node, right.
func show(t *Tree) {
	if t == nil {
		return
	}
	show(t.Left)
	fmt.Println(t.Value)
	show(t.Right)
}

func main() {
	a := &Tree{Value: 1, Right: &Tree{Value: 2}}
	b := &Tree{Left: &Tree{Value: 1}, Value: 2}
	show(a)
	show(b)
}
