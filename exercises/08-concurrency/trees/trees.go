//go:build !solution

// Package trees decides whether two binary trees hold the same values.
//
// Write this type and these two functions in this file:
//
//	type Tree struct {
//		Left  *Tree
//		Value int
//		Right *Tree
//	}
//
//	func Walk(t *Tree, ch chan int)
//	func Same(t1, t2 *Tree) bool
//
// Walk visits the tree in order: first the whole Left subtree, then the
// node's own Value, then the whole Right subtree. It sends every value
// through ch in that order and closes ch at the end. A nil tree has no
// values.
//
// Same walks both trees at the same time, each one in its own goroutine,
// and reports whether they send exactly the same values in the same order.
// Two trees with different shapes can still be the same.
//
// Example: a tree with 1 at the root and 2 on its right, and another one
// with 2 at the root and 1 on its left, both send 1, 2: they are the same.
//
// Hint: Walk can only close ch once, at the very end, so let it call a
// recursive helper that does the walking and then close ch.
package trees
