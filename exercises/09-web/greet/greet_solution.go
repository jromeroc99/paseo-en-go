//go:build solution

// Package greet is a web server that says hello by name.
package greet

import (
	"fmt"
	"net/http"
)

// NewHandler returns a handler that greets the name in the path.
func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello/{name}", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola, %s\n", r.PathValue("name"))
}
