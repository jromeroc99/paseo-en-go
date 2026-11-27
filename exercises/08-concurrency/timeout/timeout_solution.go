//go:build solution

// Package timeout waits for a result, but not forever.
package timeout

import "time"

// WithTimeout returns the result of work and true, or 0 and false if work
// takes longer than limit.
func WithTimeout(work func() int, limit time.Duration) (int, bool) {
	// With room for one value, the goroutine can finish even after a timeout.
	result := make(chan int, 1)
	go func() {
		result <- work()
	}()
	select {
	case v := <-result:
		return v, true
	case <-time.After(limit):
		return 0, false
	}
}
