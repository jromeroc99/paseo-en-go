//go:build !solution

// Package shapes adds up the areas of different shapes through an interface.
//
// Write these types, their methods and this function in this file:
//
//	type Shape interface {
//		Area() float64
//	}
//
//	type Rectangle struct {
//		Width, Height float64
//	}
//
//	func (r Rectangle) Area() float64
//
//	type Circle struct {
//		Radius float64
//	}
//
//	func (c Circle) Area() float64
//
//	func TotalArea(shapes []Shape) float64
//
// Shape is anything that knows its own area.
//
// Rectangle's Area multiplies the width by the height.
// Example: Rectangle{Width: 3, Height: 4}.Area() returns 12.
//
// Circle's Area is math.Pi times the radius squared.
// Example: Circle{Radius: 1}.Area() returns math.Pi.
//
// TotalArea returns the sum of the areas of all the shapes.
// Example: TotalArea([]Shape{Rectangle{Width: 2, Height: 5}}) returns 10.
package shapes
