package main

import "fmt"

func main() {
	name := "Ana"
	defer fmt.Println("adiós,", name)

	name = "Luis"
	fmt.Println("hola,", name)
}
