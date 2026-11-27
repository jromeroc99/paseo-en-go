//go:build !solution

// Package sum adds up a slice of numbers, whole or decimal.
//
// Write this constraint and this function in this file:
//
//	type Number interface {
//		~int | ~float64
//	}
//
//	func Sum[T Number](nums []T) T
//
// Sum returns the sum of every number in nums, or 0 when nums is empty.
// Thanks to the ~, it also works with your own types, such as
// type Euros float64.
//
// Example: Sum([]int{1, 2, 3}) is 6, and Sum([]float64{0.5, 1.5}) is 2.
package sum
