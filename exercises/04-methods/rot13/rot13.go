//go:build !solution

// Package rot13 decodes text with a Reader that applies the ROT13 cipher.
package rot13

import "io"

// Rot13Reader reads from R and replaces each letter with the letter
// 13 places further on in the alphabet, wrapping around after z.
// Other bytes pass through unchanged.
// Example: reading Rot13Reader{R: strings.NewReader("Uryyb")} gives "Hello".
type Rot13Reader struct {
	R io.Reader
}

// Read reads from r.R into b and applies ROT13 to the bytes it got.
func (r Rot13Reader) Read(b []byte) (int, error) {
	// TODO: call r.R.Read(b), then change only the first n bytes.
	return 0, io.EOF
}
