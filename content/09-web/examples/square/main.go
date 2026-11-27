package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func square(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.PathValue("n"))
	if err != nil {
		http.Error(w, "n no es un número", http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, n*n)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /square/{n}", square)
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
