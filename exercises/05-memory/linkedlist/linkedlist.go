//go:build !solution

// Package linkedlist builds a singly linked list out of nodes on the heap.
//
// Write this type and these two functions in this file:
//
//	type Node struct {
//		Value int
//		Next  *Node
//	}
//
//	func Push(head *Node, value int) *Node
//	func Values(head *Node) []int
//
// Push creates a node with value, links it in front of head and returns it:
// the new head of the list. head may be nil, which is the empty list.
// Values follows the Next pointers from head and returns every value in
// order; for the empty list it returns an empty slice.
//
// Example: Values(Push(Push(nil, 2), 1)) is [1 2].
package linkedlist
