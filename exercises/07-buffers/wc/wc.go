//go:build !solution

// Package wc counts the lines, words and bytes of a text, like the wc
// command.
//
// Write this type and this function in this file:
//
//	type Counts struct {
//		Lines, Words, Bytes int
//	}
//
//	func Count(r io.Reader) (Counts, error)
//
// Count reads r line by line with a bufio.Scanner. Lines is the number of
// lines, Words the number of words in them (strings.Fields tells you the
// words of a line) and Bytes the size of the whole text. Every line of the
// text ends in "\n", but scanner.Text() gives it to you without it, so each
// line adds len(line)+1 bytes. Count returns the error of scanner.Err().
//
// Example: "hola mundo\nde Go\n" gives Counts{Lines: 2, Words: 4, Bytes: 17}.
package wc
