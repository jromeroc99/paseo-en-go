//go:build solution

// Package shapes adds up the areas of different shapes through an interface.
package shapes

import "math"

// Shape is anything that knows its own area.
type Shape interface {
	Area() float64
}

// Rectangle is a shape with a width and a height.
type Rectangle struct {
	Width, Height float64
}

// Area returns the area of r.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle is a shape with a radius.
type Circle struct {
	Radius float64
}

// Area returns the area of c.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// TotalArea returns the sum of the areas of all the shapes.
func TotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}
