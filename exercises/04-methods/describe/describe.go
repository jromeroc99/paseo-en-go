//go:build !solution

// Package describe tells what kind of value it receives.
package describe

// Describe returns the type and the value of v, such as "int: 42",
// "string: hi" or "bool: true". It returns "nil" for nil and "unknown"
// for any other type.
// Example: Describe(42) returns "int: 42".
func Describe(v any) string {
	// TODO: use a type switch with cases for int, string, bool and nil.
	return ""
}
