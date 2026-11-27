//go:build solution

// Package steal simulates work stealing between two processors.
package steal

// P is a processor of the scheduler with its queue of waiting goroutines.
type P struct {
	Queue []string
}

// Steal moves the front half of the victim's queue, rounded up, to the end
// of the thief's queue.
func Steal(thief, victim *P) {
	half := len(victim.Queue) - len(victim.Queue)/2
	thief.Queue = append(thief.Queue, victim.Queue[:half]...)
	victim.Queue = victim.Queue[half:]
}
