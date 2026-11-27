//go:build solution

// Package ring keeps the last values it is given in a circular buffer.
package ring

// Ring holds the last values put into it in a slice of fixed size that it
// reuses from the start once it is full.
type Ring struct {
	values []int
	next   int  // where the next value goes
	full   bool // every slot holds a value, so next is also the oldest
}

// NewRing returns an empty ring with room for size values.
func NewRing(size int) *Ring {
	return &Ring{values: make([]int, size)}
}

// Put stores v, in the place of the oldest value once the ring is full.
func (r *Ring) Put(v int) {
	r.values[r.next] = v
	r.next = (r.next + 1) % len(r.values)
	if r.next == 0 {
		r.full = true
	}
}

// Values returns a copy of the values, from the oldest to the newest.
func (r *Ring) Values() []int {
	if !r.full {
		return append([]int{}, r.values[:r.next]...)
	}
	oldest := r.values[r.next:]
	newest := r.values[:r.next]
	return append(append([]int{}, oldest...), newest...)
}
