//go:build !solution

// Package sumints adds up the integers hidden in a list of any values.
//
// Write this function in this file:
//
//	func SumInts(values []any) int
//
// SumInts returns the sum of the values that hold an int and skips all
// the others: strings, floats, bools, nil...
// Example: SumInts([]any{1, "2", 3.0, 4}) returns 5.
//
// Hint: the type assertion n, ok := v.(int) gives ok == false, instead of
// a panic, when v does not hold an int.
package sumints
