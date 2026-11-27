//go:build solution

// Package wc counts the lines, words and bytes of a text, like the wc
// command.
package wc

import (
	"bufio"
	"io"
	"strings"
)

// Counts holds what Count found in a text.
type Counts struct {
	Lines, Words, Bytes int
}

// Count reads r line by line and counts its lines, words and bytes.
func Count(r io.Reader) (Counts, error) {
	var c Counts
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		c.Lines++
		c.Words += len(strings.Fields(line))
		c.Bytes += len(line) + 1 // the scanner drops the "\n"
	}
	return c, scanner.Err()
}
