//go:build !solution

// Package limit runs many tasks, but only a few of them at the same time.
//
// Write this function in this file:
//
//	func RunLimited(tasks []func(), limit int)
//
// RunLimited runs every task in its own goroutine, but never more than
// limit of them at once, and returns when all of them have finished. limit
// is always at least 1.
//
// Example: with 6 tasks and a limit of 2, two tasks run together; when one
// finishes, the next one starts, until all six are done.
//
// Hint: a buffered channel with room for limit values works as a set of
// tickets. Send a value before starting a task, and receive it back when
// the task finishes: when the buffer is full, the next send waits. A
// sync.WaitGroup tells you when every task is done.
package limit
