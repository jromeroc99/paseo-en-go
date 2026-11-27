//go:build !solution

// Package ring keeps the last values it is given in a circular buffer.
//
// Write this type and these functions in this file:
//
//	type Ring struct { ... }
//
//	func NewRing(size int) *Ring
//	func (r *Ring) Put(v int)
//	func (r *Ring) Values() []int
//
// NewRing returns an empty ring with room for size values. Put stores v;
// once the ring is full, each new value takes the place of the oldest one.
// Values returns the values the ring holds, from the oldest to the newest.
//
// Example: in NewRing(3), after Put(1), Put(2), Put(3) and Put(4),
// Values() is [2 3 4]: the 4 took the place of the 1.
//
// Hint: keep a slice of size values and the index where the next value
// goes. When the index reaches the end of the slice, it goes back to 0.
//
// You decide which fields Ring needs; the tests only use the methods.
package ring
