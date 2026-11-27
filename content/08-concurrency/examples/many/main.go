package main

import (
	"fmt"
	"runtime"
)

func main() {
	block := make(chan struct{})
	for range 100_000 {
		go func() { <-block }()
	}
	fmt.Println("goroutines:", runtime.NumGoroutine())
	fmt.Println("P:", runtime.GOMAXPROCS(0))
}
