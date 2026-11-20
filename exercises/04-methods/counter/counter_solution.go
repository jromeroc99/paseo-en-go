//go:build solution

// Package counter keeps a count that can go up and start again.
package counter

// Counter counts how many times Increment was called since the last Reset.
type Counter struct {
	count int
}

// Increment adds one to the counter.
func (c *Counter) Increment() {
	c.count++
}

// Reset puts the counter back to zero.
func (c *Counter) Reset() {
	c.count = 0
}

// Value returns the current count.
func (c Counter) Value() int {
	return c.count
}
