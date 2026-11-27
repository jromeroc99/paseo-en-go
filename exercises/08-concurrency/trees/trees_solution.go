//go:build solution

// Package trees decides whether two binary trees hold the same values.
package trees

// Tree is a node of a binary tree; nil is the empty tree.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk sends the values of t through ch in order and then closes ch.
func Walk(t *Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

// walk sends the values of t in order: left subtree, node, right subtree.
func walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

// Same reports whether t1 and t2 hold the same values in the same order.
func Same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 {
			return true // both channels are closed
		}
	}
}
