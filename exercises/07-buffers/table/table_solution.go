//go:build solution

// Package table writes text tables into a bytes.Buffer.
package table

import (
	"bytes"
	"strings"
)

// WriteTable writes rows at the end of buf, one line each, with the
// columns lined up and the cells joined by " | ".
func WriteTable(buf *bytes.Buffer, rows [][]string) {
	widths := columnWidths(rows)
	for _, row := range rows {
		for i, cell := range row {
			if i > 0 {
				buf.WriteString(" | ")
			}
			buf.WriteString(cell)
			if i < len(row)-1 {
				buf.WriteString(strings.Repeat(" ", widths[i]-len(cell)))
			}
		}
		buf.WriteByte('\n')
	}
}

// columnWidths returns the width of the widest cell of each column.
func columnWidths(rows [][]string) []int {
	var widths []int
	for _, row := range rows {
		for i, cell := range row {
			if i == len(widths) {
				widths = append(widths, 0)
			}
			widths[i] = max(widths[i], len(cell))
		}
	}
	return widths
}
