//go:build !solution

// Package uniq removes repeated lines that come one after another, like
// the uniq command.
//
// Write this function in this file:
//
//	func Uniq(r io.Reader, w io.Writer) error
//
// Uniq reads r line by line with a bufio.Scanner and writes each line into
// w followed by "\n", except when it is the same as the line just before
// it. Write through a bufio.Writer and call Flush at the end. Return the
// error of scanner.Err() if there is one, or else the error of Flush.
//
// Example: "a\na\nb\na\n" gives "a\nb\na\n". The last "a" stays, because
// the line before it is "b".
//
// Careful: an empty line is a line too. "\n\nx\n" gives "\nx\n".
//
// This exercise puts together everything in the module: a Scanner to
// read, a bufio.Writer to write and Flush so nothing is left behind.
package uniq
