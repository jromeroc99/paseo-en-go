//go:build !solution

// Package pic builds a grid of pixel values.
//
// Write this function in this file:
//
//	func Pic(dx, dy int) [][]uint8
//
// Pic returns a grid of dy rows and dx columns. The cell [y][x] holds
// uint8((x + y) / 2). Each row is its own slice.
// Example: Pic(3, 2) returns [][]uint8{{0, 0, 1}, {0, 1, 1}}.
//
// Hint: make the outer slice, then make and fill each row.
package pic
