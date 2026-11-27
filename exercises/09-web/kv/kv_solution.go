//go:build solution

// Package kv is a tiny key-value server that talks a line protocol over TCP.
package kv

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

// Store keeps the values that every connection shares.
type Store struct {
	mu     sync.Mutex
	values map[string]string
}

// NewStore returns an empty store.
func NewStore() *Store {
	return &Store{values: map[string]string{}}
}

// Handle answers each line from conn with one line, then closes conn.
func (s *Store) Handle(conn net.Conn) {
	defer conn.Close()
	lines := bufio.NewScanner(conn)
	for lines.Scan() {
		fmt.Fprintln(conn, s.run(lines.Text()))
	}
}

// run carries out one command and returns its answer.
func (s *Store) run(line string) string {
	parts := strings.SplitN(line, " ", 3)
	switch {
	case len(parts) == 3 && parts[0] == "SET":
		s.set(parts[1], parts[2])
		return "OK"
	case len(parts) == 2 && parts[0] == "GET":
		value, ok := s.get(parts[1])
		if !ok {
			return "NOT FOUND"
		}
		return value
	default:
		return "ERROR"
	}
}

func (s *Store) set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[key] = value
}

func (s *Store) get(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	value, ok := s.values[key]
	return value, ok
}
