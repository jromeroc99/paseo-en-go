//go:build !solution

// Package areader provides a Reader that never runs out of the letter A.
//
// Write this type and this method in this file:
//
//	type AReader struct{}
//
//	func (AReader) Read(b []byte) (int, error)
//
// AReader is an io.Reader that emits an endless stream of 'A'.
//
// Read fills the whole of b with 'A' and returns len(b) and a nil error.
// Example: reading into a 3-byte buffer leaves "AAA" in it and returns
// (3, nil).
//
// Hint: set every byte of b to 'A' with a for loop.
package areader
