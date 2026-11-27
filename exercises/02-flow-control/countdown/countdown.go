//go:build !solution

// Package countdown counts backwards with defer.
//
// Write this function in this file:
//
//	func Countdown(n int)
//
// Countdown prints the numbers from n down to 1, one per line, and then
// "Go!". With n = 0 it prints only "Go!".
// Example: Countdown(3) prints 3, 2, 1 and Go!, each on its own line.
//
// Hint: the deferred calls run when the function ends, the last one first.
// Defer the "Go!" first, then loop from 1 up to n and defer each number.
package countdown
