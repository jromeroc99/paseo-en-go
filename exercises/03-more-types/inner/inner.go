//go:build !solution

// Package inner trims the ends of a slice.
//
// Write this function in this file:
//
//	func Inner(nums []int) []int
//
// Inner returns nums without its first and its last element, sharing the
// same array: it does not copy anything. When nums has fewer than two
// elements, there is nothing in between, so it returns an empty slice.
// Example: Inner([]int{1, 2, 3, 4, 5}) returns []int{2, 3, 4}.
//
// Hint: a slice expression nums[low:high] goes from low up to high, but
// high itself is left out.
package inner
