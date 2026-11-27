//go:build !solution

// Package squares computes squares in goroutines and waits for all of them.
//
// Write this function in this file:
//
//	func Squares(nums []int) []int
//
// Squares returns a new slice where position i holds nums[i] * nums[i].
// Compute each square in its own goroutine, and use a sync.WaitGroup to
// wait until every goroutine has finished before returning.
//
// Example: Squares([]int{1, 2, 3}) is [1 4 9].
//
// Hint: create the result slice first with make. Each goroutine writes only
// its own position, so no two goroutines touch the same element.
package squares
