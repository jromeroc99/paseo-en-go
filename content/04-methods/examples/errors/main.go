package main

import (
	"fmt"
	"strconv"
)

func main() {
	for _, s := range []string{"42", "cuarenta"} {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		fmt.Println("número:", n)
	}
}
