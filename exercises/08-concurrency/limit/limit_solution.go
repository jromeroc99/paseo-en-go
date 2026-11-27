//go:build solution

// Package limit runs many tasks, but only a few of them at the same time.
package limit

import "sync"

// RunLimited runs every task in its own goroutine, at most limit at once,
// and waits for all of them.
func RunLimited(tasks []func(), limit int) {
	// Each value in the buffer is a running task; a full buffer makes the
	// next send wait until one of them finishes.
	tickets := make(chan struct{}, limit)
	var wg sync.WaitGroup
	for _, task := range tasks {
		tickets <- struct{}{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			task()
			<-tickets
		}()
	}
	wg.Wait()
}
