package main

import "fmt"

func main() {
	// A slice literal builds the array and the slice in one step.
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	fmt.Println(q[1:4])
	fmt.Println(q[:2]) // low defaults to 0
	fmt.Println(q[4:]) // high defaults to len(q)
	fmt.Println(q[:])  // the whole slice
}
