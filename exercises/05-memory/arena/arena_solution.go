//go:build solution

// Package arena hands out aligned pieces of one block of memory.
package arena

// Arena is a block of memory that is handed out from the start to the end.
type Arena struct {
	size int
	next int // where the free part starts
}

// New returns an empty arena of size bytes.
func New(size int) *Arena {
	return &Arena{size: size}
}

// Alloc reserves n bytes aligned to align and returns where they start.
// It returns ok false, and changes nothing, when they do not fit.
func (a *Arena) Alloc(n, align int) (offset int, ok bool) {
	start := (a.next + align - 1) / align * align
	if start+n > a.size {
		return 0, false
	}
	a.next = start + n
	return start, true
}

// Reset frees every piece at once.
func (a *Arena) Reset() {
	a.next = 0
}
