//go:build !solution

// Package room measures the free space of a slice.
//
// Write these two functions in this file:
//
//	func Room(s []int) int
//	func Extend(s []int, n int) []int
//
// Room returns how many more elements fit in s without a new array: its
// capacity minus its length.
// Example: Room(make([]int, 2, 5)) returns 3.
//
// Extend returns s grown by n elements, using the room it already has.
// When there is not enough room, it returns s unchanged.
// Example: with s := make([]int, 2, 5), Extend(s, 2) has length 4 and
// Extend(s, 4) has length 2.
//
// Hint: you can reslice s up to its capacity with s[:newLength]. Extend
// can call Room to know whether n fits.
package room
