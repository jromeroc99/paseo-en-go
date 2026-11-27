//go:build !solution

// Package notes is a small website for notes: a JSON API and an HTML page.
//
// Write this type and this function in this file:
//
//	type Note struct {
//		ID   int    `json:"id"`
//		Text string `json:"text"`
//	}
//
//	func NewHandler() http.Handler
//
// NewHandler returns a handler with no notes that answers:
//
//	GET /api/notes   200 and the notes as a JSON array; [] when there are none
//	POST /api/notes  201 and the new note, read from {"text": "..."}
//	GET /            200 and an HTML page with the notes
//
// The first note gets ID 1, the next one 2, and so on. A POST with bad
// JSON or an empty text answers 400. JSON answers have the header
// Content-Type: application/json, and the page has text/html.
//
// The page shows one <li> per note, with its text, inside a <ul>. When
// there are no notes it shows <p>Todavía no hay notas.</p> instead. Build
// it with html/template, so a note such as "<b>hola</b>" is shown as text.
//
// Requests can arrive at the same time, so protect the notes with a
// sync.Mutex. Tip: the pattern "GET /{$}" matches only the path "/".
package notes
