//go:build !solution

// Package roundrobin shares tasks among threads by turns, like a scheduler.
//
// Write this function in this file:
//
//	func RoundRobin(tasks []string, threads int) [][]string
//
// RoundRobin returns one list of tasks per thread, so the result always has
// threads lists. The tasks are dealt like cards: the first goes to thread 0,
// the second to thread 1, and so on; after the last thread it starts again
// at thread 0. A thread that gets no task has an empty list. threads is
// always at least 1.
//
// Example: RoundRobin([]string{"a", "b", "c", "d", "e"}, 2) is
// [[a c e] [b d]].
//
// Hint: task number i goes to thread i % threads.
package roundrobin
