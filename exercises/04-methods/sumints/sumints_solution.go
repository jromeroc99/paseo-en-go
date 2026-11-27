//go:build solution

// Package sumints adds up the integers hidden in a list of any values.
package sumints

// SumInts returns the sum of the values that hold an int and skips all
// the others.
func SumInts(values []any) int {
	total := 0
	for _, v := range values {
		if n, ok := v.(int); ok {
			total += n
		}
	}
	return total
}
