//go:build solution

// Package layout works out how many bytes a struct takes, padding included.
package layout

// Size returns the size of a struct whose fields have the given sizes.
func Size(fields []int) int {
	offset := 0
	largest := 1
	for _, size := range fields {
		offset = roundUp(offset, size)
		offset += size
		largest = max(largest, size)
	}
	return roundUp(offset, largest)
}

// roundUp returns the first multiple of align that is n or more.
func roundUp(n, align int) int {
	return (n + align - 1) / align * align
}
