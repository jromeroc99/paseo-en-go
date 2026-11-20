//go:build !solution

// Package apply transforms every element of a list with a function.
package apply

// Apply returns a new slice where each element is f applied to the element
// of nums in the same position. It does not modify nums.
// Example: Apply([]int{1, 2, 3}, double) returns []int{2, 4, 6}.
func Apply(nums []int, f func(int) int) []int {
	// TODO: make a slice of the same length and fill it with f(n).
	return nil
}
