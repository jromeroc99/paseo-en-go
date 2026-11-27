//go:build !solution

// Package timeout waits for a result, but not forever.
//
// Write this function in this file:
//
//	func WithTimeout(work func() int, limit time.Duration) (int, bool)
//
// WithTimeout runs work in a goroutine and waits for its result. If the
// result arrives within limit, it returns the result and true. If limit
// passes first, it gives up and returns 0 and false, without waiting for
// work any longer.
//
// Example: WithTimeout(func() int { return 42 }, time.Second) returns
// 42 and true.
//
// Hint: select between the channel with the result and time.After(limit).
// Give the result channel a buffer of 1: then the goroutine can always
// deliver its result and finish, even when nobody is waiting for it anymore.
package timeout
