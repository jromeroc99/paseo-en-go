//go:build solution

// Package between builds a run of whole numbers.
package between

// Between returns the numbers from from to to, both included, in order.
func Between(from, to int) []int {
	numbers := []int{}
	for n := from; n <= to; n++ {
		numbers = append(numbers, n)
	}
	return numbers
}
