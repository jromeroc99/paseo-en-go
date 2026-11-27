//go:build !solution

// Package reachable finds the live objects, as the garbage collector does.
//
// Write this function in this file:
//
//	func Reachable(refs map[string][]string, roots []string) []string
//
// refs says which objects each object points to: refs["a"] = []string{"b"}
// means a points to b. roots are the objects your variables point to.
// Reachable returns every object you can reach from the roots by following
// the pointers, sorted alphabetically and without repeats. Everything else
// is garbage. Watch out for cycles: a can point to b and b back to a.
//
// Example: with refs {"a": {"b"}, "c": {"a"}} and roots ["a"], the result
// is [a b]; c is garbage, because nothing points to it.
//
// Hint: keep a map[string]bool of the objects you have already marked.
package reachable
