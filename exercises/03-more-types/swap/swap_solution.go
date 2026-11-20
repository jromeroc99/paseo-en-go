//go:build solution

// Package swap exchanges two values through pointers.
package swap

// Swap exchanges the integers that a and b point to.
func Swap(a, b *int) {
	*a, *b = *b, *a
}
