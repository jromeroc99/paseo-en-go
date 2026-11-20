//go:build solution

// Package halvings counts integer divisions by two.
package halvings

// Halvings returns how many times n can be divided by two, with integer
// division, until it reaches zero.
func Halvings(n int) int {
	count := 0
	for n > 0 {
		n /= 2
		count++
	}
	return count
}
