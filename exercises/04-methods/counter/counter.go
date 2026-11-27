//go:build !solution

// Package counter keeps a count that can go up and start again.
//
// Write this type and these methods in this file:
//
//	type Counter struct {
//		count int
//	}
//
//	func (c *Counter) Increment()
//	func (c *Counter) Reset()
//	func (c Counter) Value() int
//
// Counter counts how many times Increment was called since the last Reset.
//
// Increment adds one to the counter.
// Example: after two calls on a new Counter, Value returns 2.
//
// Reset puts the counter back to zero.
// Example: after Reset, Value returns 0.
//
// Value returns the current count.
// Example: a new Counter has Value 0.
//
// Hint: Increment and Reset change c.count, so they need a pointer
// receiver. Value only reads it.
package counter
