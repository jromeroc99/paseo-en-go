//go:build !solution

// Package counter keeps a count that can go up and start again.
package counter

// Counter counts how many times Increment was called since the last Reset.
type Counter struct {
	count int
}

// Increment adds one to the counter.
// Example: after two calls on a new Counter, Value returns 2.
func (c *Counter) Increment() {
	// TODO: add 1 to c.count.
}

// Reset puts the counter back to zero.
// Example: after Reset, Value returns 0.
func (c *Counter) Reset() {
	// TODO: set c.count to 0.
}

// Value returns the current count.
// Example: a new Counter has Value 0.
func (c Counter) Value() int {
	// TODO: return c.count.
	return 0
}
