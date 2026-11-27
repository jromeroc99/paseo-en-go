//go:build !solution

// Package factorial computes factorials with a recursive function.
//
// Write this function in this file:
//
//	func Factorial(n int) int
//
// Factorial returns n! = n × (n-1) × … × 1, and 0! is 1. It must call
// itself: Factorial(n) is n × Factorial(n-1). Each call gets its own frame
// on the stack, so Factorial(5) has six frames on top of each other at its
// deepest point.
package factorial
