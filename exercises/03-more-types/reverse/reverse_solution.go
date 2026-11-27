//go:build solution

// Package reverse turns an array around.
package reverse

// Reverse returns an array with the elements of a in the opposite order.
func Reverse(a [5]int) [5]int {
	var reversed [5]int
	for i := 0; i < len(a); i++ {
		reversed[len(a)-1-i] = a[i]
	}
	return reversed
}
