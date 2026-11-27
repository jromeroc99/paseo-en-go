//go:build solution

// Package factorial computes factorials with a recursive function.
package factorial

// Factorial returns n!, with 0! being 1.
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
