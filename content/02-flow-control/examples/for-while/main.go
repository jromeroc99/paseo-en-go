package main

import "fmt"

func main() {
	// Double a number until it reaches 1000.
	n := 1
	steps := 0
	for n < 1000 {
		n *= 2
		steps++
	}
	fmt.Println(n, "en", steps, "pasos")
}
