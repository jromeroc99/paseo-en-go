package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	src := strings.NewReader("Hola desde un Reader\n")
	n, err := io.Copy(os.Stdout, src)
	fmt.Println(n, "bytes copiados, error:", err)
}
