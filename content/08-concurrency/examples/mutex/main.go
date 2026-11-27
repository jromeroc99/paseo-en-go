package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
