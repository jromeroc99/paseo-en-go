package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	for i := 1; i <= 3; i++ {
		fmt.Fprintln(w, "línea", i)
	}
	w.Flush()
}
