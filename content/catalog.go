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
		{
			Slug:    "01-basics",
			Title:   "Lo básico",
			Summary: "Paquetes, funciones, variables, tipos y constantes: las piezas de cualquier programa.",
			Lessons: []course.Lesson{
				{Slug: "01-packages", Title: "Paquetes e imports"},
				{
					Slug:  "02-functions",
					Title: "Funciones",
					Exercise: &course.Exercise{
						Path: "01-basics/rectangle",
						Goal: "Calcula el área y el perímetro de un rectángulo.",
					},
				},
				{
					Slug:  "03-multiple-results",
					Title: "Varios resultados",
					Exercise: &course.Exercise{
						Path: "01-basics/divmod",
						Goal: "Devuelve el cociente y el resto de una división entera.",
					},
				},
				{Slug: "04-variables", Title: "Variables"},
				{Slug: "05-basic-types", Title: "Tipos básicos"},
				{Slug: "06-zero-values", Title: "Valores cero"},
				{
					Slug:  "07-type-conversions",
					Title: "Conversiones e inferencia de tipos",
					Exercise: &course.Exercise{
						Path: "01-basics/average",
						Goal: "Calcula la media de dos enteros sin perder los decimales.",
					},
				},
				{Slug: "08-constants", Title: "Constantes"},
				{
					Slug:  "09-review",
					Title: "Repaso: el círculo",
					Exercise: &course.Exercise{
						Path: "01-basics/circle",
						Goal: "Calcula el área y el perímetro de un círculo a partir de su radio.",
					},
				},
			},
		},
		{
			Slug:    "02-flow-control",
			Title:   "Control de flujo",
			Summary: "Bucles, condiciones, switch y defer: cómo decide tu programa qué hacer y cuántas veces.",
			Lessons: []course.Lesson{
				{
					Slug:  "01-for",
					Title: "El bucle for",
					Exercise: &course.Exercise{
						Path: "02-flow-control/sum",
						Goal: "Suma todos los números del 1 al n.",
					},
				},
				{
					Slug:  "02-while-forever",
					Title: "for como while y bucles infinitos",
					Exercise: &course.Exercise{
						Path: "02-flow-control/halvings",
						Goal: "Cuenta cuántas veces puedes dividir un número entre dos hasta llegar a cero.",
					},
				},
				{
					Slug:  "03-if",
					Title: "if y else",
					Exercise: &course.Exercise{
						Path: "02-flow-control/fizzbuzz",
						Goal: "Devuelve Fizz, Buzz, FizzBuzz o el número, según sus divisores.",
					},
				},
				{
					Slug:  "04-sqrt",
					Title: "Práctica: tu propia raíz cuadrada",
					Exercise: &course.Exercise{
						Path: "02-flow-control/sqrt",
						Goal: "Calcula la raíz cuadrada con el método de Newton, sin usar math.Sqrt.",
					},
				},
				{
					Slug:  "05-switch",
					Title: "switch",
					Exercise: &course.Exercise{
						Path: "02-flow-control/season",
						Goal: "Indica a qué estación del año pertenece un mes.",
					},
				},
				{Slug: "06-defer", Title: "defer"},
				{
					Slug:  "07-review",
					Title: "Repaso: números primos",
					Exercise: &course.Exercise{
						Path: "02-flow-control/primes",
						Goal: "Decide si un número es primo y cuenta cuántos primos hay hasta un límite.",
					},
				},
			},
		},
	},
}
