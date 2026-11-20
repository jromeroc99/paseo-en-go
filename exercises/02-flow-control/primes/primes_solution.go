//go:build solution

// Package primes recognises and counts prime numbers.
package primes

// IsPrime reports whether n is a prime number: greater than 1 and divisible
// only by 1 and by itself.
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for d := 2; d < n; d++ {
		if n%d == 0 {
			return false
		}
	}
	return true
}

// CountPrimes returns how many prime numbers there are from 1 to limit,
// both included.
func CountPrimes(limit int) int {
	count := 0
	for n := 2; n <= limit; n++ {
		if IsPrime(n) {
			count++
		}
	}
	return count
}
