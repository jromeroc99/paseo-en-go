package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	buf.WriteString("uno ")
	buf.Write([]byte("dos"))
	fmt.Println(buf.Len())

	b := make([]byte, 4)
	n, _ := buf.Read(b) // it only fails with io.EOF, once it is empty
	fmt.Printf("%d %q\n", n, b[:n])
	fmt.Println(buf.Len(), buf.String())
}
