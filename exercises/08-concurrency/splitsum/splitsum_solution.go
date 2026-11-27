//go:build solution

// Package splitsum adds up a slice in two goroutines that report through a
// channel.
package splitsum

// Sum adds up nums and sends the total through ch.
func Sum(nums []int, ch chan int) {
	total := 0
	for _, n := range nums {
		total += n
	}
	ch <- total
}

// SplitSum adds up each half of nums in its own goroutine and joins both
// totals.
func SplitSum(nums []int) int {
	ch := make(chan int)
	half := len(nums) / 2
	go Sum(nums[:half], ch)
	go Sum(nums[half:], ch)
	return <-ch + <-ch
}
