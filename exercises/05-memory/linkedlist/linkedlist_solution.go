//go:build solution

// Package linkedlist builds a singly linked list out of nodes on the heap.
package linkedlist

// Node is one element of the list; Next is nil on the last one.
type Node struct {
	Value int
	Next  *Node
}

// Push puts a new node in front of head and returns it as the new head.
func Push(head *Node, value int) *Node {
	return &Node{Value: value, Next: head}
}

// Values returns the values of the list, from head to the end.
func Values(head *Node) []int {
	values := []int{}
	for n := head; n != nil; n = n.Next {
		values = append(values, n.Value)
	}
	return values
}
