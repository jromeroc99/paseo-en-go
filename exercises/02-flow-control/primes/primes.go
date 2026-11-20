//go:build !solution

// Package primes recognises and counts prime numbers.
package primes

// IsPrime reports whether n is a prime number: greater than 1 and divisible
// only by 1 and by itself.
// Example: IsPrime(17) returns true and IsPrime(25) returns false.
func IsPrime(n int) bool {
	// TODO: rule out numbers below 2, then look for a divisor with a for loop.
	return false
}

// CountPrimes returns how many prime numbers there are from 1 to limit,
// both included.
// Example: CountPrimes(10) returns 4 (2, 3, 5 and 7).
func CountPrimes(limit int) int {
	// TODO: call IsPrime for every number up to limit.
	return 0
}
