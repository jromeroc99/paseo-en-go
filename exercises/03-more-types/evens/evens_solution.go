//go:build solution

// Package evens picks the even numbers out of a list.
package evens

// Evens returns a new slice with the even numbers of nums, in the same order.
func Evens(nums []int) []int {
	var result []int
	for _, n := range nums {
		if n%2 == 0 {
			result = append(result, n)
		}
	}
	return result
}
