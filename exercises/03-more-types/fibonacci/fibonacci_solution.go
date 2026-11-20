//go:build solution

// Package fibonacci generates the Fibonacci sequence with a closure.
package fibonacci

// Fibonacci returns a function that, on each call, returns the next
// number of the Fibonacci sequence, starting at 0.
func Fibonacci() func() int {
	current, next := 0, 1
	return func() int {
		result := current
		current, next = next, current+next
		return result
	}
}
