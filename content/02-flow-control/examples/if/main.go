package main

import (
	"fmt"
	"strconv"
)

// describe says whether n is negative, zero or positive.
func describe(n int) string {
	if n < 0 {
		return strconv.Itoa(n) + " es negativo"
	} else if n == 0 {
		return "cero"
	} else {
		return strconv.Itoa(n) + " es positivo"
	}
}

func main() {
	fmt.Println(describe(-4))
	fmt.Println(describe(0))
	fmt.Println(describe(7))
}
