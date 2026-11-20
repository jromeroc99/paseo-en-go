//go:build solution

// Package wordcount counts the words of a text.
package wordcount

import "strings"

// WordCount returns how many times each word appears in s.
func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	for _, word := range strings.Fields(s) {
		counts[word]++
	}
	return counts
}
