//go:build !solution

// Package steal simulates work stealing between two processors.
//
// Write this type and this function in this file:
//
//	type P struct {
//		Queue []string
//	}
//
//	func Steal(thief, victim *P)
//
// A P is a processor of the Go scheduler, and Queue holds the goroutines
// waiting to run on it, the first one in front. Steal moves half of the
// victim's queue to the end of the thief's queue. The thief takes the tasks
// from the front of the victim's queue, in order, and rounds up: with 3
// tasks it takes 2, and with 1 task it takes it. If the victim has nothing,
// nothing changes.
//
// Example: if thief.Queue is [] and victim.Queue is [a b c d e], after
// Steal(thief, victim) thief.Queue is [a b c] and victim.Queue is [d e].
//
// Hint: half rounded up is n - n/2.
package steal
