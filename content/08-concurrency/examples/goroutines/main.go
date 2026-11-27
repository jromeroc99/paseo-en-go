package main

import (
	"fmt"
	"sync"
)

func say(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go say("hola", &wg)
	go say("adiós", &wg)
	wg.Wait()
	fmt.Println("fin")
}
