//go:build !solution

// Package todos is a small JSON API for a to-do list.
//
// Write this type and this function in this file:
//
//	type Todo struct {
//		ID    int    `json:"id"`
//		Title string `json:"title"`
//	}
//
//	func NewHandler() http.Handler
//
// NewHandler returns a handler with an empty list that answers:
//
//	GET /todos          200 and the list as a JSON array; [] when it is empty
//	POST /todos         201 and the new task, read from {"title": "..."}
//	DELETE /todos/{id}  204 and no body
//
// The first task gets ID 1, the next one 2, and so on. Errors:
//
//	POST with bad JSON or an empty title    400
//	DELETE with an id that is not a number  400
//	DELETE with an id that does not exist   404
//
// Every JSON answer has the header Content-Type: application/json.
// Requests can arrive at the same time, so protect the list with a
// sync.Mutex.
package todos
