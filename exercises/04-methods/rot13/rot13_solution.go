//go:build solution

// Package rot13 decodes text with a Reader that applies the ROT13 cipher.
package rot13

import "io"

// Rot13Reader reads from R and replaces each letter with the letter
// 13 places further on in the alphabet, wrapping around after z.
type Rot13Reader struct {
	R io.Reader
}

// Read reads from r.R into b and applies ROT13 to the bytes it got.
func (r Rot13Reader) Read(b []byte) (int, error) {
	n, err := r.R.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

// rot13 moves a letter 13 places and leaves any other byte as it is.
func rot13(c byte) byte {
	switch {
	case c >= 'a' && c <= 'z':
		return 'a' + (c-'a'+13)%26
	case c >= 'A' && c <= 'Z':
		return 'A' + (c-'A'+13)%26
	default:
		return c
	}
}
