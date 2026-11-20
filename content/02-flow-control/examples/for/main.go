package main

import "fmt"

func main() {
	// init; condition; post
	for i := 0; i < 3; i++ {
		fmt.Println("vuelta", i)
	}

	// The init and post parts are optional.
	total := 1
	for ; total < 100; total += total {
	}
	fmt.Println("total:", total)
}
