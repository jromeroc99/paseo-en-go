//go:build solution

// Package index finds a value in a slice of any comparable type.
package index

// Index returns the position of the first x in s, or -1 if s has no x.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}
