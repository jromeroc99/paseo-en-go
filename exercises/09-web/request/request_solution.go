//go:build solution

// Package request reads an HTTP request by hand, line by line.
package request

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

// Request is the part of an HTTP request that comes before the body.
type Request struct {
	Method  string
	Path    string
	Headers map[string]string
}

// Parse reads the request line and the headers up to the blank line.
func Parse(r io.Reader) (Request, error) {
	lines := bufio.NewScanner(r)
	if !lines.Scan() {
		return Request{}, errors.New("empty request")
	}
	parts := strings.Fields(lines.Text())
	if len(parts) != 3 {
		return Request{}, fmt.Errorf("bad request line %q", lines.Text())
	}

	req := Request{Method: parts[0], Path: parts[1], Headers: map[string]string{}}
	for lines.Scan() && lines.Text() != "" {
		name, value, ok := strings.Cut(lines.Text(), ":")
		if !ok {
			return Request{}, fmt.Errorf("bad header %q", lines.Text())
		}
		req.Headers[name] = strings.TrimSpace(value)
	}
	return req, nil
}
