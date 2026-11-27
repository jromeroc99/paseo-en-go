//go:build !solution

// Package generate sends the Fibonacci sequence through a channel.
//
// Write this function in this file:
//
//	func Fibonacci(n int, ch chan int)
//
// Fibonacci sends the first n numbers of the Fibonacci sequence through ch
// and then closes ch, so whoever reads it with range knows there is nothing
// else to come. The sequence starts with 0 and 1, and each number after them
// is the sum of the two before it.
//
// Example: with n = 6 it sends 0, 1, 1, 2, 3 and 5, and closes the channel.
// With n = 0 it sends nothing, but it still closes the channel.
package generate
