// Package content holds the course: its structure (declared below) and the
// lesson bodies and code examples, embedded into the binary.
//
// Layout: <module-slug>/<lesson-slug>.html for lesson bodies and
// <module-slug>/examples/<name>/main.go for runnable snippets.
package content

import (
	"embed"

	"paseo-en-go/internal/course"
)

// FS contains every lesson body and example.
//
//go:embed */*.html */examples
var FS embed.FS

// Course is the tour, in reading order.
var Course = course.Course{
	Title: "Paseo en Go",
	Modules: []course.Module{
		{
			Slug:    "00-setup",
			Title:   "Preparación",
			Summary: "Instala Go, descarga los ejercicios y ejecuta tu primer test.",
			Lessons: []course.Lesson{
				{
					Slug:  "01-welcome",
					Title: "Bienvenida",
					Exercise: &course.Exercise{
						Path: "00-setup/hello",
						Goal: "Haz que Greeting devuelva un saludo.",
					},
				},
			},
		},
	},
}
