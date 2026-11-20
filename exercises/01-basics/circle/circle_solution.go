//go:build solution

// Package circle computes measures of a circle.
package circle

import "math"

// Circle returns the area (π·r²) and the perimeter (2·π·r) of a circle
// with the given radius.
func Circle(radius int) (area, perimeter float64) {
	r := float64(radius)
	return math.Pi * r * r, 2 * math.Pi * r
}
