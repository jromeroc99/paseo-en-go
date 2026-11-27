//go:build solution

// Package uniq removes repeated lines that come one after another, like
// the uniq command.
package uniq

import (
	"bufio"
	"io"
)

// Uniq copies the lines of r into w, dropping each line that repeats the
// one just before it.
func Uniq(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	bw := bufio.NewWriter(w)
	previous := ""
	first := true
	for scanner.Scan() {
		line := scanner.Text()
		if !first && line == previous {
			continue
		}
		bw.WriteString(line + "\n")
		previous = line
		first = false
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return bw.Flush()
}
