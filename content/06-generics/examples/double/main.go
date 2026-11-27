package main

import "fmt"

type Number interface {
	~int | ~float64
}

func Double[T Number](x T) T {
	return x * 2
}

type Celsius float64

func main() {
	fmt.Println(Double(21), Double(1.25), Double(Celsius(18.5)))
}
