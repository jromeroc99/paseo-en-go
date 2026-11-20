package main

import "fmt"

func main() {
	var words [2]string
	words[0] = "hola"
	words[1] = "mundo"
	fmt.Println(words[0], words[1])
	fmt.Println(words)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes, len(primes))

	copied := primes // assigning an array copies every element
	copied[0] = 100
	fmt.Println(primes[0], copied[0])
}
