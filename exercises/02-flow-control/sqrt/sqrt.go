//go:build !solution

// Package sqrt computes square roots with Newton's method.
//
// Write this function in this file:
//
//	func Sqrt(x float64) float64
//
// Sqrt returns the square root of x, for x >= 0, without using math.Sqrt.
// Example: Sqrt(9) returns 3.
//
// Hint: start with z := 1.0 and repeat z -= (z*z - x) / (2*z) until z
// barely changes.
package sqrt
