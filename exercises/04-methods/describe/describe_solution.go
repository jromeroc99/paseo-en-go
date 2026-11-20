//go:build solution

// Package describe tells what kind of value it receives.
package describe

import "fmt"

// Describe returns the type and the value of v, such as "int: 42".
func Describe(v any) string {
	switch x := v.(type) {
	case int:
		return fmt.Sprintf("int: %d", x)
	case string:
		return "string: " + x
	case bool:
		return fmt.Sprintf("bool: %t", x)
	case nil:
		return "nil"
	default:
		return "unknown"
	}
}
