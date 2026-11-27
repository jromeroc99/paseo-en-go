//go:build !solution

// Package safesqrt computes square roots and reports an error for negative numbers.
//
// Write this type, its method and this function in this file:
//
//	type ErrNegativeSqrt float64
//
//	func (e ErrNegativeSqrt) Error() string
//	func Sqrt(x float64) (float64, error)
//
// ErrNegativeSqrt is the error for a square root of a negative number.
//
// Error describes the problem.
// Example: ErrNegativeSqrt(-2).Error() returns
// "cannot Sqrt negative number: -2".
//
// Sqrt returns the square root of x, or ErrNegativeSqrt(x) when x is
// negative.
// Example: Sqrt(4) returns (2, nil) and Sqrt(-2) returns
// (0, ErrNegativeSqrt(-2)).
//
// Hint: in Error, use fmt.Sprint with float64(e), not with e itself. In
// Sqrt, return the error when x < 0; otherwise use math.Sqrt.
package safesqrt
