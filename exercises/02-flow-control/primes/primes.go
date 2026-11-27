//go:build !solution

// Package primes recognises and counts prime numbers.
//
// Write these two functions in this file:
//
//	func IsPrime(n int) bool
//	func CountPrimes(limit int) int
//
// IsPrime reports whether n is a prime number: greater than 1 and divisible
// only by 1 and by itself.
// Example: IsPrime(17) returns true and IsPrime(25) returns false.
//
// CountPrimes returns how many prime numbers there are from 1 to limit,
// both included.
// Example: CountPrimes(10) returns 4 (2, 3, 5 and 7).
//
// Hint: in IsPrime, rule out numbers below 2, then look for a divisor with
// a for loop. CountPrimes calls IsPrime for every number up to limit.
package primes
