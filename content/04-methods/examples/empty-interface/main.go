package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hola"
	describe(i)

	var a any = 3.5 // any is the same as interface{}
	describe(a)
}
