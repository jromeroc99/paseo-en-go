//go:build solution

// Package safesqrt computes square roots and reports an error for negative numbers.
package safesqrt

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt is the error for a square root of a negative number.
type ErrNegativeSqrt float64

// Error describes the problem. It converts e to float64 first: passing e
// itself to fmt would call Error again, forever.
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

// Sqrt returns the square root of x, or ErrNegativeSqrt(x) when x is negative.
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}
