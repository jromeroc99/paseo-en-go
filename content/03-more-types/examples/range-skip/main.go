package main

import "fmt"

func main() {
	pow := make([]int, 5)

	// Only the index: fill the slice.
	for i := range pow {
		pow[i] = 1 << i // 2 to the power of i
	}

	// Only the value: skip the index with _.
	for _, value := range pow {
		fmt.Println(value)
	}
}
