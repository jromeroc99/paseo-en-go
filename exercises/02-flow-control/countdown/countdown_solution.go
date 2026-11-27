//go:build solution

// Package countdown counts backwards with defer.
package countdown

import "fmt"

// Countdown prints the numbers from n down to 1, one per line, and then
// "Go!".
func Countdown(n int) {
	defer fmt.Println("Go!")
	for i := 1; i <= n; i++ {
		defer fmt.Println(i)
	}
}
