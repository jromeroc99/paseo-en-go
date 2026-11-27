//go:build !solution

// Package arena hands out aligned pieces of one block of memory.
//
// Write this type and these functions in this file:
//
//	type Arena struct { ... }
//
//	func New(size int) *Arena
//	func (a *Arena) Alloc(n, align int) (offset int, ok bool)
//	func (a *Arena) Reset()
//
// New returns an empty arena of size bytes. Alloc reserves n bytes and
// returns the offset where they start: the first multiple of align after
// the last piece handed out. If the piece does not fit, Alloc returns
// 0, false and the arena does not change. Reset frees everything at once,
// so the next Alloc starts again at offset 0.
//
// Example: in New(16), Alloc(1, 1) returns 0, true, and then Alloc(8, 8)
// returns 8, true: bytes 1 to 7 are padding.
//
// You decide which fields Arena needs; the tests only use the methods.
package arena
