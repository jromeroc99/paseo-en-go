//go:build solution

// Package fizzbuzz plays the classic counting game.
package fizzbuzz

import "strconv"

// FizzBuzz returns "Fizz" when n is divisible by 3, "Buzz" when it is
// divisible by 5, "FizzBuzz" when it is divisible by both, and the number
// itself as text otherwise.
func FizzBuzz(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}
