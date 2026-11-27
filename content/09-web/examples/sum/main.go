package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type sumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type sumResponse struct {
	Result int `json:"result"`
}

func sum(w http.ResponseWriter, r *http.Request) {
	var req sumRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON no válido", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sumResponse{Result: req.A + req.B})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /sum", sum)
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
