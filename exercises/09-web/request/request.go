//go:build !solution

// Package request reads an HTTP request by hand, line by line.
//
// Write this type and this function in this file:
//
//	type Request struct {
//		Method  string
//		Path    string
//		Headers map[string]string
//	}
//
//	func Parse(r io.Reader) (Request, error)
//
// Parse reads the request line, "GET /notes HTTP/1.1", and takes the
// method and the path from it. Then it reads one header per line,
// "Host: localhost:8080", until the blank line, and stores each one in
// Headers with the name as the key and the value without spaces around it.
// Headers is never nil, even when there are no headers.
//
// Parse returns an error when the input is empty, when the request line
// does not have exactly three parts, or when a header has no ":".
//
// Tip: bufio.Scanner splits lines and also drops the "\r" before "\n".
package request
