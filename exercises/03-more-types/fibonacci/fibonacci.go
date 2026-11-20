//go:build !solution

// Package fibonacci generates the Fibonacci sequence with a closure.
package fibonacci

// Fibonacci returns a function that, on each call, returns the next
// number of the Fibonacci sequence, starting at 0. Each function returned
// by Fibonacci keeps its own count.
// Example: f := Fibonacci(); then f(), f(), f(), f() return 0, 1, 1, 2.
func Fibonacci() func() int {
	// TODO: declare the two last numbers and return a closure that uses them.
	return nil
}
