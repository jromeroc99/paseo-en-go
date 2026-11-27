//go:build solution

// Package shout turns a text into a shout, with the help of package strings.
package shout

import "strings"

// Shout returns s in capital letters with an exclamation mark at the end.
func Shout(s string) string {
	return strings.ToUpper(s) + "!"
}
