//go:build !solution

// Package bits picks the smallest unsigned integer type that holds a number.
//
// Write this function in this file:
//
//	func BitsNeeded(n uint64) int
//
// BitsNeeded returns 8, 16, 32 or 64: the size in bits of the smallest
// unsigned integer type (uint8, uint16, uint32 or uint64) that can hold n.
// Example: BitsNeeded(255) is 8, because 255 fits in a uint8, and
// BitsNeeded(256) is 16.
//
// Hint: the package math has the largest value of each type, such as
// math.MaxUint8 and math.MaxUint16.
package bits
