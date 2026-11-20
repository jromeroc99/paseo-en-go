package main

import "fmt"

func main() {
	var (
		ready   bool    = true
		name    string  = "Gopher"
		count   int     = 42
		max     uint8   = 255
		letter  rune    = 'G'
		raw     byte    = 'g'
		balance float64 = 19.99
	)
	fmt.Printf("%T: %v\n", ready, ready)
	fmt.Printf("%T: %v\n", name, name)
	fmt.Printf("%T: %v\n", count, count)
	fmt.Printf("%T: %v\n", max, max)
	fmt.Printf("%T: %v\n", letter, letter)
	fmt.Printf("%T: %v\n", raw, raw)
	fmt.Printf("%T: %v\n", balance, balance)
}
