package main

import (
	"io"
	"os"
	"strings"
)

// upperReader wraps another reader and turns letters to upper case.
type upperReader struct {
	r io.Reader
}

func (u upperReader) Read(b []byte) (int, error) {
	n, err := u.r.Read(b)
	for i := 0; i < n; i++ {
		if b[i] >= 'a' && b[i] <= 'z' {
			b[i] = b[i] - 'a' + 'A'
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("hola, gopher\n")
	io.Copy(os.Stdout, upperReader{s})
}
