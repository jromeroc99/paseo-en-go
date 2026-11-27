//go:build solution

// Package page builds an HTML page from a template.
package page

import (
	"html/template"
	"io"
)

var pageTemplate = template.Must(template.New("page").Parse(`<h1>{{.Title}}</h1>
{{if .Items}}<ul>
{{range .Items}}  <li>{{.}}</li>
{{end}}</ul>
{{else}}<p>No hay nada.</p>
{{end}}`))

// pageData is what the template reads.
type pageData struct {
	Title string
	Items []string
}

// Render writes the page with the title and the list of items to w.
func Render(w io.Writer, title string, items []string) error {
	return pageTemplate.Execute(w, pageData{Title: title, Items: items})
}
