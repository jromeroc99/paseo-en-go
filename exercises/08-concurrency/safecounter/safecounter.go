//go:build !solution

// Package safecounter counts visits from many goroutines without losing any.
//
// Write this type and these functions in this file:
//
//	type SafeCounter struct {
//		mu     sync.Mutex
//		counts map[string]int
//	}
//
//	func NewSafeCounter() *SafeCounter
//	func (c *SafeCounter) Inc(key string)
//	func (c *SafeCounter) Value(key string) int
//
// NewSafeCounter returns a counter with its map ready to use. Inc adds one
// visit to key, and Value returns how many visits key has; a key nobody
// visited has 0. Many goroutines call Inc and Value at the same time, so
// both lock mu before they touch the map and unlock it when they are done.
//
// Example: after 100 goroutines call Inc("home") once each,
// Value("home") is 100.
//
// Hint: defer c.mu.Unlock() right after c.mu.Lock() makes sure you never
// forget to unlock.
package safecounter
