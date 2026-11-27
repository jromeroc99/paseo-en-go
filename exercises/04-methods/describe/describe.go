//go:build !solution

// Package describe tells what kind of value it receives.
//
// Write this function in this file:
//
//	func Describe(v any) string
//
// Describe returns the type and the value of v, such as "int: 42",
// "string: hi" or "bool: true". It returns "nil" for nil and "unknown"
// for any other type.
// Example: Describe(42) returns "int: 42".
//
// Hint: use a type switch with cases for int, string, bool and nil.
package describe
