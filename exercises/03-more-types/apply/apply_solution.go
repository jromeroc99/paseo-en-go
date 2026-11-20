//go:build solution

// Package apply transforms every element of a list with a function.
package apply

// Apply returns a new slice where each element is f applied to the element
// of nums in the same position.
func Apply(nums []int, f func(int) int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = f(n)
	}
	return result
}
