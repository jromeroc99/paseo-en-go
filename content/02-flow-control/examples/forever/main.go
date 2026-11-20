package main

import "fmt"

func main() {
	countdown := 3
	for {
		if countdown == 0 {
			fmt.Println("¡despegue!")
			break
		}
		fmt.Println(countdown)
		countdown--
	}
}
