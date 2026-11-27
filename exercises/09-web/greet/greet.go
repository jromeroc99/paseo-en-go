//go:build !solution

// Package greet is a web server that says hello by name.
//
// Write this function in this file:
//
//	func NewHandler() http.Handler
//
// NewHandler returns a handler that answers GET /hello/{name} with status
// 200 and the body "Hola, <name>" followed by a newline.
//
// Example: GET /hello/Ana answers "Hola, Ana\n".
//
// Use an http.ServeMux with the pattern "GET /hello/{name}" and read the
// name with r.PathValue. The mux already answers 404 for other paths and
// 405 for other methods, so you do not have to.
package greet
