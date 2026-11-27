package main

import (
	"fmt"
	"runtime"
)

type Node struct {
	Value int
	Next  *Node
}

func main() {
	a := &Node{Value: 1}
	a.Next = &Node{Value: 2}
	a.Next = &Node{Value: 3}
	runtime.GC()
	fmt.Println(a.Value, a.Next.Value)
}
