//go:build !solution

// Package fibonacci generates the Fibonacci sequence with a closure.
//
// Write this function in this file:
//
//	func Fibonacci() func() int
//
// Fibonacci returns a function that, on each call, returns the next
// number of the Fibonacci sequence, starting at 0. Each function returned
// by Fibonacci keeps its own count.
// Example: f := Fibonacci(); then f(), f(), f(), f() return 0, 1, 1, 2.
//
// Hint: declare the two last numbers and return a closure that uses them.
package fibonacci
