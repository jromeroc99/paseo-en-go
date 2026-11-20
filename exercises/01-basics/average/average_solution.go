//go:build solution

// Package average computes the mean of two integers.
package average

// Average returns the mean of a and b, keeping the decimals.
func Average(a, b int) float64 {
	return float64(a+b) / 2
}
