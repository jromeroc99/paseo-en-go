package main

import "fmt"

func Contains[T comparable](s []T, x T) bool {
	for _, v := range s {
		if v == x {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(Contains([]int{10, 20, 30}, 20))
	fmt.Println(Contains([]string{"hola", "adiós"}, "buenas"))
}
