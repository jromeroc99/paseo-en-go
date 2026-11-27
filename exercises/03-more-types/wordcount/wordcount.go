//go:build !solution

// Package wordcount counts the words of a text.
//
// Write this function in this file:
//
//	func WordCount(s string) map[string]int
//
// WordCount returns how many times each word appears in s. Words are
// separated by spaces, and "Go" and "go" are different words.
// Example: WordCount("the quick the") returns
// map[string]int{"the": 2, "quick": 1}.
//
// Hint: split s with strings.Fields and add 1 for each word.
package wordcount
