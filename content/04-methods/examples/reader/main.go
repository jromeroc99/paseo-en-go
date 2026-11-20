package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hola, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b[:n] = %q\n", n, err, b[:n])
		if err == io.EOF {
			break
		}
	}
}
