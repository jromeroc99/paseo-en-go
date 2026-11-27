//go:build !solution

// Package kv is a tiny key-value server that talks a line protocol over TCP.
//
// Write this type and these functions in this file:
//
//	type Store struct { ... }
//
//	func NewStore() *Store
//	func (s *Store) Handle(conn net.Conn)
//
// NewStore returns an empty store. Handle reads conn line by line, answers
// each line with one line, and closes conn when the client stops sending:
//
//	SET clave valor   ->  OK          (the value is the rest of the line)
//	GET clave         ->  the value, or NOT FOUND
//	anything else     ->  ERROR
//
// Example: "SET saludo hola mundo" answers "OK", and then "GET saludo"
// answers "hola mundo".
//
// Many connections share one Store at the same time, so protect the data
// with a sync.Mutex.
package kv
