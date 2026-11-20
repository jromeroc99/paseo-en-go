package main

import "fmt"

// MyFloat is a type of our own, so it can have methods.
type MyFloat float64

// Abs returns the absolute value of f.
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-2.5)
	fmt.Println(f.Abs())
}
