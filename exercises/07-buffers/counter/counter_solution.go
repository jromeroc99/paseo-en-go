//go:build solution

// Package counter counts the bytes that pass through a Writer.
package counter

import "io"

// CountingWriter writes into W and keeps the number of bytes written so far.
type CountingWriter struct {
	W     io.Writer
	Count int
}

// Write passes p on to W and counts only the bytes W really wrote.
func (c *CountingWriter) Write(p []byte) (int, error) {
	n, err := c.W.Write(p)
	c.Count += n
	return n, err
}
