package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	for range 2 {
		select {
		case v := <-ch:
			fmt.Println("recibido", v)
		default:
			fmt.Println("nada listo")
			ch <- 7
		}
	}
}
