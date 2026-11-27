//go:build solution

// Package inner trims the ends of a slice.
package inner

// Inner returns nums without its first and its last element, sharing the
// same array.
func Inner(nums []int) []int {
	if len(nums) < 2 {
		return []int{}
	}
	return nums[1 : len(nums)-1]
}
