//go:build !solution

// Package safesqrt computes square roots and reports an error for negative numbers.
package safesqrt

// ErrNegativeSqrt is the error for a square root of a negative number.
type ErrNegativeSqrt float64

// Error describes the problem.
// Example: ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".
func (e ErrNegativeSqrt) Error() string {
	// TODO: use fmt.Sprint with float64(e), not with e itself.
	return ""
}

// Sqrt returns the square root of x, or ErrNegativeSqrt(x) when x is negative.
// Example: Sqrt(4) returns (2, nil) and Sqrt(-2) returns (0, ErrNegativeSqrt(-2)).
func Sqrt(x float64) (float64, error) {
	// TODO: return the error when x < 0; otherwise use math.Sqrt.
	return 0, nil
}
