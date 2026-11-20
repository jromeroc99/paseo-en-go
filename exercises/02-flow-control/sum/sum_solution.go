//go:build solution

// Package sum adds up a run of whole numbers.
package sum

// SumTo returns 1 + 2 + ... + n, or 0 when n is less than 1.
func SumTo(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}
