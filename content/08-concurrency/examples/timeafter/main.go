package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string) // nadie envía nada
	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("tiempo agotado")
	}
}
