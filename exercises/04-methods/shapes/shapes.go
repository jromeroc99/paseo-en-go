//go:build !solution

// Package shapes adds up the areas of different shapes through an interface.
package shapes

// Shape is anything that knows its own area.
type Shape interface {
	Area() float64
}

// Rectangle is a shape with a width and a height.
type Rectangle struct {
	Width, Height float64
}

// Area returns the area of r.
// Example: Rectangle{Width: 3, Height: 4}.Area() returns 12.
func (r Rectangle) Area() float64 {
	// TODO: multiply the width by the height.
	return 0
}

// Circle is a shape with a radius.
type Circle struct {
	Radius float64
}

// Area returns the area of c.
// Example: Circle{Radius: 1}.Area() returns math.Pi.
func (c Circle) Area() float64 {
	// TODO: use math.Pi times the radius squared.
	return 0
}

// TotalArea returns the sum of the areas of all the shapes.
// Example: TotalArea([]Shape{Rectangle{Width: 2, Height: 5}}) returns 10.
func TotalArea(shapes []Shape) float64 {
	// TODO: add up Area() for every shape.
	return 0
}
