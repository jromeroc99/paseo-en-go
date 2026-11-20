package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := 3, 4
	// math.Sqrt needs a float64, so the ints are converted first.
	f := math.Sqrt(float64(x*x + y*y))
	z := int(f)
	fmt.Println(x, y, z)
}
