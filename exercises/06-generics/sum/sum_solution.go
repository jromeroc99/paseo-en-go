//go:build solution

// Package sum adds up a slice of numbers, whole or decimal.
package sum

// Number is every type whose underlying type is int or float64.
type Number interface {
	~int | ~float64
}

// Sum returns the sum of nums; the sum of no numbers is 0.
func Sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}
