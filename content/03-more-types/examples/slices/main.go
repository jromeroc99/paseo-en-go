package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// From index 1 up to, but not including, index 4.
	var s []int = primes[1:4]
	fmt.Println(s)
	fmt.Println(s[0], len(s))
}
