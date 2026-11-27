//go:build solution

// Package report writes a score report through a bufio.Writer.
package report

import (
	"bufio"
	"fmt"
	"io"
)

// Score is the number of points one person got.
type Score struct {
	Name   string
	Points int
}

// WriteReport writes one line per score and the total into w. The lines
// wait in a bufio.Writer and reach w together when it is flushed.
func WriteReport(w io.Writer, scores []Score) error {
	bw := bufio.NewWriter(w)
	fmt.Fprintln(bw, "Informe")
	total := 0
	for _, s := range scores {
		fmt.Fprintf(bw, "%s: %d\n", s.Name, s.Points)
		total += s.Points
	}
	fmt.Fprintf(bw, "Total: %d\n", total)
	return bw.Flush()
}
