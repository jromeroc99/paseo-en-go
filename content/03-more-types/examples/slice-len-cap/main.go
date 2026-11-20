package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0] // length 0, but the array is still there
	printSlice(s)

	s = s[:4] // grow back, up to the capacity
	printSlice(s)

	s = s[2:] // dropping from the front lowers the capacity
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
