//go:build solution

// Package bits picks the smallest unsigned integer type that holds a number.
package bits

import "math"

// BitsNeeded returns the size in bits of the smallest unsigned integer
// type that can hold n.
func BitsNeeded(n uint64) int {
	switch {
	case n <= math.MaxUint8:
		return 8
	case n <= math.MaxUint16:
		return 16
	case n <= math.MaxUint32:
		return 32
	default:
		return 64
	}
}
