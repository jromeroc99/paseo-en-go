//go:build solution

// Package room measures the free space of a slice.
package room

// Room returns how many more elements fit in s without a new array.
func Room(s []int) int {
	return cap(s) - len(s)
}

// Extend returns s grown by n elements, using the room it already has, or
// s unchanged when there is not enough room.
func Extend(s []int, n int) []int {
	if n > Room(s) {
		return s
	}
	return s[:len(s)+n]
}
