package main

import "fmt"

func main() {
	stock := make(map[string]int)

	stock["manzanas"] = 42 // insert
	fmt.Println("manzanas:", stock["manzanas"])

	stock["manzanas"] = 48 // update
	fmt.Println("manzanas:", stock["manzanas"])

	delete(stock, "manzanas") // remove
	fmt.Println("manzanas:", stock["manzanas"])

	v, ok := stock["manzanas"] // ask if the key is there
	fmt.Println("manzanas:", v, "¿existe?", ok)
}
