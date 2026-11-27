//go:build solution

// Package squares computes squares in goroutines and waits for all of them.
package squares

import "sync"

// Squares returns the square of every number, each one computed in its own
// goroutine.
func Squares(nums []int) []int {
	result := make([]int, len(nums))
	var wg sync.WaitGroup
	for i, n := range nums {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result[i] = n * n
		}()
	}
	wg.Wait()
	return result
}
