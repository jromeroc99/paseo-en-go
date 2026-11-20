package main

import "fmt"

// discounted applies a discount only when the result stays above a minimum.
func discounted(price, discount, minimum float64) float64 {
	if p := price - discount; p > minimum {
		return p
	}
	// p does not exist here: its scope ended with the if.
	return minimum
}

func main() {
	fmt.Println(discounted(20, 5, 10))
	fmt.Println(discounted(12, 5, 10))
}
