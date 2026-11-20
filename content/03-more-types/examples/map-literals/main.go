package main

import "fmt"

// Vertex is a position on a map: latitude and longitude.
type Vertex struct {
	Lat, Long float64
}

func main() {
	// Inside the literal, the type of each value can be left out.
	places := map[string]Vertex{
		"Madrid":  {40.4168, -3.7038},
		"Sevilla": {37.3891, -5.9845},
	}
	fmt.Println(places["Madrid"])
	fmt.Println(places)
}
