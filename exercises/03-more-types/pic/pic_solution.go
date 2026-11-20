//go:build solution

// Package pic builds a grid of pixel values.
package pic

// Pic returns a grid of dy rows and dx columns. The cell [y][x] holds
// uint8((x + y) / 2). Each row is its own slice.
func Pic(dx, dy int) [][]uint8 {
	grid := make([][]uint8, dy)
	for y := range grid {
		grid[y] = make([]uint8, dx)
		for x := range grid[y] {
			grid[y][x] = uint8((x + y) / 2)
		}
	}
	return grid
}
