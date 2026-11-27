package main

import "fmt"

func countdown(n int, ch chan int) {
	for i := n; i > 0; i-- {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go countdown(3, ch)
	for v := range ch {
		fmt.Println(v)
	}
}
