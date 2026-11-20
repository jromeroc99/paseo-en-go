//go:build solution

// Package divmod splits an integer division into its two parts.
package divmod

// DivMod returns the quotient and the remainder of a divided by b.
func DivMod(a, b int) (quotient, remainder int) {
	return a / b, a % b
}
