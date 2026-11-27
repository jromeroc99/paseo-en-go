//go:build !solution

// Package report writes a score report through a bufio.Writer.
//
// Write this type and this function in this file:
//
//	type Score struct {
//		Name   string
//		Points int
//	}
//
//	func WriteReport(w io.Writer, scores []Score) error
//
// WriteReport wraps w in a bufio.Writer and writes the report through it:
// first the line "Informe", then one line "Name: Points" per score, and at
// the end the line "Total: N", where N is the sum of every Points. Every
// line ends in "\n". It returns the error of Flush.
//
// Example: the scores {"Ana", 31} and {"Luis", 7} give
//
//	Informe
//	Ana: 31
//	Luis: 7
//	Total: 38
//
// Hint: fmt.Fprintf(bw, ...) writes into the bufio.Writer bw. The report
// must reach w in one single Write, so do not write to w directly.
package report
