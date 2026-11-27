//go:build !solution

// Package table writes text tables into a bytes.Buffer.
//
// Write this function in this file:
//
//	func WriteTable(buf *bytes.Buffer, rows [][]string)
//
// WriteTable writes one line per row at the end of buf, each line ending in
// "\n". The cells of a row are joined by " | ". Every cell except the last
// one of its row is padded with spaces up to the width of the widest cell
// in its column, so the columns line up. Every row has the same number of
// cells. Measure widths with len: the tests only use letters without
// accents.
//
// Example: the rows {"Nombre", "Edad"}, {"Ana", "31"} and {"Luis", "7"}
// give
//
//	Nombre | Edad
//	Ana    | 31
//	Luis   | 7
//
// Hint: first find the width of each column, then write the rows.
// strings.Repeat(" ", n) gives you n spaces.
package table
