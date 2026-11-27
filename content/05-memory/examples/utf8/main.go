package main

import "fmt"

func main() {
	s := "año"
	fmt.Println(len(s))
	for i, r := range s {
		fmt.Println(i, string(r))
	}
}
