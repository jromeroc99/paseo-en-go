package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(strings.NewReader("uno\ndos\ntres\n"))
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error:", err)
	}
}
