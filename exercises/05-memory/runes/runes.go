//go:build !solution

// Package runes counts and reverses the characters of a UTF-8 text.
//
// Write these two functions in this file:
//
//	func CountRunes(s string) int
//	func Reverse(s string) string
//
// CountRunes returns how many characters (runes) s has. It is not len(s):
// "año" has 4 bytes but 3 runes. Reverse returns s with its characters in
// the opposite order, so Reverse("año") is "oña". Reversing the bytes
// would break the ñ, which takes two bytes.
//
// Hint: range over a string gives you one rune at a time, and []rune(s)
// turns a string into a slice of runes.
package runes
