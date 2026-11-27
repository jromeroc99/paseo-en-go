package main

import "fmt"

func square(n int) int {
	r := n * n
	return r
}

func sumSquares(a, b int) int {
	return square(a) + square(b)
}

func main() {
	fmt.Println(sumSquares(2, 3))
}
