//go:build solution

// Package reachable finds the live objects, as the garbage collector does.
package reachable

import "slices"

// Reachable returns, sorted, every object reachable from roots.
func Reachable(refs map[string][]string, roots []string) []string {
	marked := map[string]bool{}
	pending := slices.Clone(roots)
	for len(pending) > 0 {
		obj := pending[len(pending)-1]
		pending = pending[:len(pending)-1]
		if marked[obj] {
			continue
		}
		marked[obj] = true
		pending = append(pending, refs[obj]...)
	}

	live := []string{}
	for obj := range marked {
		live = append(live, obj)
	}
	slices.Sort(live)
	return live
}
