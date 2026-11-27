//go:build solution

// Package todos is a small JSON API for a to-do list.
package todos

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

// Todo is one task on the list.
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// list holds the tasks that every request shares.
type list struct {
	mu     sync.Mutex
	todos  []Todo
	lastID int
}

// NewHandler returns the API for a new, empty list.
func NewHandler() http.Handler {
	l := &list{todos: []Todo{}}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos", l.handleList)
	mux.HandleFunc("POST /todos", l.handleCreate)
	mux.HandleFunc("DELETE /todos/{id}", l.handleDelete)
	return mux
}

func (l *list) handleList(w http.ResponseWriter, r *http.Request) {
	l.mu.Lock()
	defer l.mu.Unlock()
	writeJSON(w, http.StatusOK, l.todos)
}

func (l *list) handleCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil || todo.Title == "" {
		http.Error(w, "send a JSON object with a title", http.StatusBadRequest)
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()
	l.lastID++
	todo.ID = l.lastID
	l.todos = append(l.todos, todo)
	writeJSON(w, http.StatusCreated, todo)
}

func (l *list) handleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "the id must be a number", http.StatusBadRequest)
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()
	for i, todo := range l.todos {
		if todo.ID == id {
			l.todos = append(l.todos[:i], l.todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}

// writeJSON sends v as JSON with the given status code. The header must
// be set before WriteHeader, because WriteHeader sends the headers.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
