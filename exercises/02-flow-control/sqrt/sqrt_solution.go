//go:build solution

// Package sqrt computes square roots with Newton's method.
package sqrt

import "math"

// Sqrt returns the square root of x, for x >= 0, without using math.Sqrt.
func Sqrt(x float64) float64 {
	z := 1.0
	for {
		next := z - (z*z-x)/(2*z)
		// Stop when a step barely changes z: comparing floats with == could loop forever.
		if math.Abs(next-z) < 1e-12 {
			return next
		}
		z = next
	}
}
