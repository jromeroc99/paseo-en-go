//go:build !solution

// Package areader provides a Reader that never runs out of the letter A.
package areader

// AReader is an io.Reader that emits an endless stream of 'A'.
type AReader struct{}

// Read fills the whole of b with 'A' and returns len(b) and a nil error.
// Example: reading into a 3-byte buffer leaves "AAA" in it and returns (3, nil).
func (AReader) Read(b []byte) (int, error) {
	// TODO: set every byte of b to 'A'.
	return 0, nil
}
