//go:build !solution

// Package layout works out how many bytes a struct takes, padding included.
//
// Write this function in this file:
//
//	func Size(fields []int) int
//
// fields holds the size in bytes of each field, in order; every size is
// 1, 2, 4 or 8, and a field must start at an offset that is a multiple of
// its own size. Size places the fields one after another, leaving padding
// before a field when it needs it, and then adds padding at the end until
// the total is a multiple of the largest field. With no fields, it is 0.
//
// Example: Size([]int{1, 8, 1}) is 24, like struct{ A bool; B int64; C bool },
// and Size([]int{8, 1, 1}) is 16.
package layout
