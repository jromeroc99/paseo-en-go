package main

import "fmt"

func main() {
	fmt.Println("contando")
	for i := 1; i <= 3; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("hecho")
}
