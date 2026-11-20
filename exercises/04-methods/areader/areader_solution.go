//go:build solution

// Package areader provides a Reader that never runs out of the letter A.
package areader

// AReader is an io.Reader that emits an endless stream of 'A'.
type AReader struct{}

// Read fills the whole of b with 'A' and returns len(b) and a nil error.
func (AReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}
