package main

import (
	"html/template"
	"os"
)

const page = `<h1>{{.Title}}</h1>
<ul>
{{range .Items}}  <li>{{.}}</li>
{{end}}</ul>
`

type List struct {
	Title string
	Items []string
}

func main() {
	tmpl := template.Must(template.New("page").Parse(page))
	list := List{Title: "Compra", Items: []string{"pan", "<b>sal</b>"}}
	tmpl.Execute(os.Stdout, list)
}
