//go:build solution

// Package runes counts and reverses the characters of a UTF-8 text.
package runes

// CountRunes returns the number of runes in s.
func CountRunes(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

// Reverse returns s with its runes in the opposite order.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
