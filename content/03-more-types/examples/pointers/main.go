package main

import "fmt"

func main() {
	i := 42
	p := &i // p holds the address of i

	fmt.Println(*p) // read i through the pointer
	*p = 21         // write i through the pointer
	fmt.Println(i)

	var q *int // a pointer that points nowhere yet
	fmt.Println(q == nil)
}
