//go:build solution

// Package safecounter counts visits from many goroutines without losing any.
package safecounter

import "sync"

// SafeCounter counts visits per key; it is safe to use from many goroutines.
type SafeCounter struct {
	mu     sync.Mutex
	counts map[string]int
}

// NewSafeCounter returns an empty counter, ready to use.
func NewSafeCounter() *SafeCounter {
	return &SafeCounter{counts: map[string]int{}}
}

// Inc adds one visit to key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counts[key]++
}

// Value returns how many visits key has.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counts[key]
}
