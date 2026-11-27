//go:build solution

// Package notes is a small website for notes: a JSON API and an HTML page.
package notes

import (
	"encoding/json"
	"html/template"
	"net/http"
	"sync"
)

// Note is one note, as the API sends it.
type Note struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

var pageTemplate = template.Must(template.New("page").Parse(`<!doctype html>
<title>Notas</title>
<h1>Notas</h1>
{{if .}}<ul>
{{range .}}  <li>{{.Text}}</li>
{{end}}</ul>
{{else}}<p>Todavía no hay notas.</p>
{{end}}`))

// board holds the notes that every request shares.
type board struct {
	mu     sync.Mutex
	notes  []Note
	lastID int
}

// NewHandler returns the website for a new board with no notes.
func NewHandler() http.Handler {
	b := &board{notes: []Note{}}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/notes", b.handleList)
	mux.HandleFunc("POST /api/notes", b.handleCreate)
	mux.HandleFunc("GET /{$}", b.handlePage)
	return mux
}

func (b *board) handleList(w http.ResponseWriter, r *http.Request) {
	b.mu.Lock()
	defer b.mu.Unlock()
	writeJSON(w, http.StatusOK, b.notes)
}

func (b *board) handleCreate(w http.ResponseWriter, r *http.Request) {
	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil || note.Text == "" {
		http.Error(w, "send a JSON object with a text", http.StatusBadRequest)
		return
	}

	b.mu.Lock()
	defer b.mu.Unlock()
	b.lastID++
	note.ID = b.lastID
	b.notes = append(b.notes, note)
	writeJSON(w, http.StatusCreated, note)
}

func (b *board) handlePage(w http.ResponseWriter, r *http.Request) {
	b.mu.Lock()
	defer b.mu.Unlock()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	pageTemplate.Execute(w, b.notes)
}

// writeJSON sends v as JSON with the given status code. The header must
// be set before WriteHeader, because WriteHeader sends the headers.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
