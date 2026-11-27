//go:build solution

// Package generate sends the Fibonacci sequence through a channel.
package generate

// Fibonacci sends the first n Fibonacci numbers through ch and closes it.
func Fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for range n {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
