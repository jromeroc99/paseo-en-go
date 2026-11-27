//go:build !solution

// Package splitsum adds up a slice in two goroutines that report through a
// channel.
//
// Write these two functions in this file:
//
//	func Sum(nums []int, ch chan int)
//	func SplitSum(nums []int) int
//
// Sum adds up nums and sends the total through ch.
// SplitSum cuts nums in two halves, starts one goroutine with Sum for each
// half, receives both totals from the same channel and returns their sum.
//
// Example: SplitSum([]int{7, 2, 8, -9, 4, 0}) is 12: one half adds up to 17,
// the other to -5.
//
// Hint: the halves are nums[:len(nums)/2] and nums[len(nums)/2:]. The two
// totals can arrive in any order, but the sum does not care.
package splitsum
