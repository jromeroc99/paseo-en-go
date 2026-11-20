package main

import "fmt"

func do(i any) {
	switch v := i.(type) {
	case int:
		fmt.Printf("el doble de %v es %v\n", v, v*2)
	case string:
		fmt.Printf("%q tiene %v bytes\n", v, len(v))
	default:
		fmt.Printf("no sé qué hacer con %T\n", v)
	}
}

func main() {
	do(21)
	do("hola")
	do(true)
}
