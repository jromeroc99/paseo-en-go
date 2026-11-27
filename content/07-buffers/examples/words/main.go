package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(strings.NewReader("hola,  mundo\nde Go\n"))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Printf("%q\n", scanner.Text())
	}
}
