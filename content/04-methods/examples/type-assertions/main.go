package main

import "fmt"

func main() {
	var i any = "hola"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	n, ok := i.(int)
	fmt.Println(n, ok)
}
