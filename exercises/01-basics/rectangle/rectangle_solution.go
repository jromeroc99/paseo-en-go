//go:build solution

// Package rectangle computes measures of a rectangle.
package rectangle

// Area returns the area of a rectangle: width times height.
func Area(width, height int) int {
	return width * height
}

// Perimeter returns the length of the border of a rectangle.
func Perimeter(width, height int) int {
	return 2 * (width + height)
}
