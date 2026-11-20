package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// x and y share a type, so it is written once.
func multiply(x, y int) int {
	return x * y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(multiply(6, 7))
}
