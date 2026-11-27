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
				{
					Slug:  "01-packages",
					Title: "Paquetes e imports",
					Exercise: &course.Exercise{
						Path: "01-basics/shout",
						Goal: "Usa el paquete strings para gritar un texto en mayúsculas.",
					},
				},
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
				{
					Slug:  "04-variables",
					Title: "Variables",
					Exercise: &course.Exercise{
						Path: "01-basics/seconds",
						Goal: "Convierte horas, minutos y segundos en segundos.",
					},
				},
				{
					Slug:  "05-basic-types",
					Title: "Tipos básicos",
					Exercise: &course.Exercise{
						Path: "01-basics/discount",
						Goal: "Calcula el precio con descuento usando float64.",
					},
				},
				{
					Slug:  "06-zero-values",
					Title: "Valores cero",
					Exercise: &course.Exercise{
						Path: "01-basics/zero",
						Goal: "Devuelve el valor cero de cada tipo básico.",
					},
				},
				{
					Slug:  "07-type-conversions",
					Title: "Conversiones e inferencia de tipos",
					Exercise: &course.Exercise{
						Path: "01-basics/average",
						Goal: "Calcula la media de dos enteros sin perder los decimales.",
					},
				},
				{
					Slug:  "08-constants",
					Title: "Constantes",
					Exercise: &course.Exercise{
						Path: "01-basics/vat",
						Goal: "Usa una constante para sumar el IVA a un precio.",
					},
				},
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
		{
			Slug:    "03-more-types",
			Title:   "Más tipos",
			Summary: "Punteros, structs, slices, maps y funciones como valores: cómo se organizan los datos.",
			Lessons: []course.Lesson{
				{
					Slug:  "01-pointers",
					Title: "Punteros",
					Exercise: &course.Exercise{
						Path: "03-more-types/swap",
						Goal: "Intercambia dos enteros usando punteros.",
					},
				},
				{
					Slug:  "02-structs",
					Title: "Structs",
					Exercise: &course.Exercise{
						Path: "03-more-types/player",
						Goal: "Crea un jugador y súmale puntos.",
					},
				},
				{Slug: "03-arrays", Title: "Arrays"},
				{Slug: "04-slices", Title: "Slices"},
				{Slug: "05-slice-len-cap", Title: "Longitud y capacidad"},
				{Slug: "06-append", Title: "Añadir elementos con append"},
				{
					Slug:  "07-range",
					Title: "Recorrer con range",
					Exercise: &course.Exercise{
						Path: "03-more-types/evens",
						Goal: "Devuelve solo los números pares de una lista.",
					},
				},
				{
					Slug:  "08-pic",
					Title: "Práctica: slices de slices",
					Exercise: &course.Exercise{
						Path: "03-more-types/pic",
						Goal: "Construye una cuadrícula de dy filas y dx columnas.",
					},
				},
				{
					Slug:  "09-maps",
					Title: "Maps",
					Exercise: &course.Exercise{
						Path: "03-more-types/wordcount",
						Goal: "Cuenta cuántas veces aparece cada palabra en un texto.",
					},
				},
				{
					Slug:  "10-function-values",
					Title: "Funciones como valores",
					Exercise: &course.Exercise{
						Path: "03-more-types/apply",
						Goal: "Aplica una función a cada elemento de una lista.",
					},
				},
				{
					Slug:  "11-closures",
					Title: "Closures",
					Exercise: &course.Exercise{
						Path: "03-more-types/fibonacci",
						Goal: "Devuelve una función que genera la sucesión de Fibonacci.",
					},
				},
				{
					Slug:  "12-review",
					Title: "Repaso: el inventario",
					Exercise: &course.Exercise{
						Path: "03-more-types/inventory",
						Goal: "Gestiona un inventario con structs, slices, maps y funciones.",
					},
				},
			},
		},
		{
			Slug:    "04-methods",
			Title:   "Métodos e interfaces",
			Summary: "Métodos, interfaces y las más usadas de la biblioteca estándar: Stringer, error, Reader e Image.",
			Lessons: []course.Lesson{
				{
					Slug:  "01-methods",
					Title: "Métodos",
					Exercise: &course.Exercise{
						Path: "04-methods/temperature",
						Goal: "Añade métodos a un tipo para convertir temperaturas.",
					},
				},
				{
					Slug:  "02-pointer-receivers",
					Title: "Receptores de puntero",
					Exercise: &course.Exercise{
						Path: "04-methods/counter",
						Goal: "Crea un contador que se incrementa y se reinicia.",
					},
				},
				{
					Slug:  "03-interfaces",
					Title: "Interfaces",
					Exercise: &course.Exercise{
						Path: "04-methods/shapes",
						Goal: "Calcula el área total de varias figuras con una interfaz.",
					},
				},
				{Slug: "04-interface-values", Title: "Valores de interfaz"},
				{Slug: "05-any", Title: "La interfaz vacía y las type assertions"},
				{
					Slug:  "06-type-switches",
					Title: "Type switches",
					Exercise: &course.Exercise{
						Path: "04-methods/describe",
						Goal: "Describe un valor según su tipo.",
					},
				},
				{
					Slug:  "07-stringers",
					Title: "Stringer",
					Exercise: &course.Exercise{
						Path: "04-methods/ipaddr",
						Goal: "Haz que una dirección IP se imprima como 127.0.0.1.",
					},
				},
				{
					Slug:  "08-errors",
					Title: "Errores",
					Exercise: &course.Exercise{
						Path: "04-methods/safesqrt",
						Goal: "Devuelve un error cuando te pidan la raíz de un número negativo.",
					},
				},
				{
					Slug:  "09-readers",
					Title: "Readers",
					Exercise: &course.Exercise{
						Path: "04-methods/areader",
						Goal: "Crea un Reader que emite la letra A sin parar.",
					},
				},
				{
					Slug:  "10-rot13",
					Title: "Práctica: rot13Reader",
					Exercise: &course.Exercise{
						Path: "04-methods/rot13",
						Goal: "Descifra un mensaje con un Reader que aplica ROT13.",
					},
				},
				{
					Slug:  "11-images",
					Title: "Imágenes",
					Exercise: &course.Exercise{
						Path: "04-methods/images",
						Goal: "Crea tu propio tipo de imagen implementando image.Image.",
					},
				},
			},
		},
		{
			Slug:    "05-memory",
			Title:   "La memoria por dentro",
			Summary: "Bytes, direcciones, el stack, el heap y el recolector de basura: dónde viven tus datos y cuánto ocupan.",
			Lessons: []course.Lesson{
				{
					Slug:  "01-bytes",
					Title: "Bytes y direcciones",
					Exercise: &course.Exercise{
						Path: "05-memory/bits",
						Goal: "Decide cuántos bits necesitas para guardar un número.",
					},
				},
				{
					Slug:  "02-stack",
					Title: "El stack y las llamadas",
					Exercise: &course.Exercise{
						Path: "05-memory/factorial",
						Goal: "Calcula el factorial con una función recursiva.",
					},
				},
				{
					Slug:  "03-heap",
					Title: "El heap y el escape analysis",
					Exercise: &course.Exercise{
						Path: "05-memory/linkedlist",
						Goal: "Construye una lista enlazada con punteros a nodos.",
					},
				},
				{
					Slug:  "04-gc",
					Title: "El recolector de basura",
					Exercise: &course.Exercise{
						Path: "05-memory/reachable",
						Goal: "Marca qué objetos siguen vivos, como hace el recolector.",
					},
				},
				{
					Slug:  "05-padding",
					Title: "Alineación y padding",
					Exercise: &course.Exercise{
						Path: "05-memory/layout",
						Goal: "Calcula cuánto ocupa un struct, con su padding.",
					},
				},
				{
					Slug:  "06-strings",
					Title: "Strings, bytes y runes",
					Exercise: &course.Exercise{
						Path: "05-memory/runes",
						Goal: "Cuenta y da la vuelta a los caracteres de un texto UTF-8.",
					},
				},
				{
					Slug:  "07-review",
					Title: "Repaso: tu propia memoria",
					Exercise: &course.Exercise{
						Path: "05-memory/arena",
						Goal: "Reserva trozos alineados de un bloque de memoria.",
					},
				},
			},
		},
		{
			Slug:    "06-generics",
			Title:   "Genéricos",
			Summary: "Funciones y tipos que sirven para cualquier tipo: parámetros de tipo, restricciones y colecciones genéricas.",
			Lessons: []course.Lesson{
				{
					Slug:  "01-type-parameters",
					Title: "Parámetros de tipo",
					Exercise: &course.Exercise{
						Path: "06-generics/index",
						Goal: "Busca la posición de un valor en un slice de cualquier tipo.",
					},
				},
				{
					Slug:  "02-constraints",
					Title: "Restricciones",
					Exercise: &course.Exercise{
						Path: "06-generics/sum",
						Goal: "Suma un slice de números, sean enteros o decimales.",
					},
				},
				{
					Slug:  "03-generic-types",
					Title: "Tipos genéricos",
					Exercise: &course.Exercise{
						Path: "06-generics/stack",
						Goal: "Construye una pila que guarde valores de cualquier tipo.",
					},
				},
				{
					Slug:  "04-review",
					Title: "Repaso: tus propias colecciones",
					Exercise: &course.Exercise{
						Path: "06-generics/set",
						Goal: "Crea un conjunto genérico y las funciones Map y Filter para cualquier tipo.",
					},
				},
			},
		},
		{
			Slug:    "07-buffers",
			Title:   "Buffers y E/S",
			Summary: "Por qué escribir de golpe es más rápido que byte a byte: bytes.Buffer, bufio y io.Copy.",
			Lessons: []course.Lesson{
				{
					Slug:  "01-buffers",
					Title: "Qué es un buffer",
					Exercise: &course.Exercise{
						Path: "07-buffers/ring",
						Goal: "Guarda los últimos n valores en un buffer circular.",
					},
				},
				{
					Slug:  "02-bytes-buffer",
					Title: "bytes.Buffer",
					Exercise: &course.Exercise{
						Path: "07-buffers/table",
						Goal: "Escribe una tabla de texto en un bytes.Buffer.",
					},
				},
				{
					Slug:  "03-bufio-writer",
					Title: "bufio.Writer y Flush",
					Exercise: &course.Exercise{
						Path: "07-buffers/report",
						Goal: "Escribe un informe con bufio.Writer sin olvidar el Flush.",
					},
				},
				{
					Slug:  "04-scanner",
					Title: "Leer línea a línea con bufio.Scanner",
					Exercise: &course.Exercise{
						Path: "07-buffers/wc",
						Goal: "Cuenta las líneas, las palabras y los bytes de un texto, como wc.",
					},
				},
				{
					Slug:  "05-io-copy",
					Title: "io.Copy: de un Reader a un Writer",
					Exercise: &course.Exercise{
						Path: "07-buffers/counter",
						Goal: "Cuenta cuántos bytes pasan por un Writer mientras copias con io.Copy.",
					},
				},
				{
					Slug:  "06-review",
					Title: "Repaso: tu propio uniq",
					Exercise: &course.Exercise{
						Path: "07-buffers/uniq",
						Goal: "Lee líneas y quita las que se repiten seguidas, como uniq.",
					},
				},
			},
		},
		{
			Slug:    "08-concurrency",
			Title:   "Procesos, hilos y goroutines",
			Summary: "Cómo hace Go varias cosas a la vez: procesos, hilos, goroutines, el planificador, canales y Mutex.",
			Lessons: []course.Lesson{
				{
					Slug:  "01-processes",
					Title: "Procesos e hilos",
					Exercise: &course.Exercise{
						Path: "08-concurrency/roundrobin",
						Goal: "Reparte tareas entre varios hilos por turnos, como un planificador.",
					},
				},
				{
					Slug:  "02-goroutines",
					Title: "Goroutines",
					Exercise: &course.Exercise{
						Path: "08-concurrency/squares",
						Goal: "Calcula cuadrados en goroutines y espera a todas con un WaitGroup.",
					},
				},
				{
					Slug:  "03-scheduler",
					Title: "El planificador G-M-P",
					Exercise: &course.Exercise{
						Path: "08-concurrency/steal",
						Goal: "Simula el robo de trabajo: un procesador sin tareas se lleva la mitad de la cola de otro.",
					},
				},
				{
					Slug:  "04-channels",
					Title: "Canales",
					Exercise: &course.Exercise{
						Path: "08-concurrency/splitsum",
						Goal: "Suma las dos mitades de un slice en dos goroutines y junta los resultados por un canal.",
					},
				},
				{
					Slug:  "05-buffered-channels",
					Title: "Canales con buffer",
					Exercise: &course.Exercise{
						Path: "08-concurrency/limit",
						Goal: "Limita cuántas tareas se ejecutan a la vez con un canal con buffer.",
					},
				},
				{
					Slug:  "06-range-close",
					Title: "range y close",
					Exercise: &course.Exercise{
						Path: "08-concurrency/generate",
						Goal: "Envía la sucesión de Fibonacci por un canal y ciérralo al acabar.",
					},
				},
				{
					Slug:  "07-select",
					Title: "select y default",
					Exercise: &course.Exercise{
						Path: "08-concurrency/timeout",
						Goal: "Espera un resultado, pero ríndete si tarda más de un tiempo máximo.",
					},
				},
				{
					Slug:  "08-mutex",
					Title: "Mutex y data races",
					Exercise: &course.Exercise{
						Path: "08-concurrency/safecounter",
						Goal: "Cuenta visitas desde muchas goroutines sin perder ninguna, con un Mutex.",
					},
				},
				{
					Slug:  "09-trees",
					Title: "Práctica: árboles equivalentes",
					Exercise: &course.Exercise{
						Path: "08-concurrency/trees",
						Goal: "Recorre dos árboles binarios en goroutines y decide si guardan los mismos valores.",
					},
				},
				{
					Slug:  "10-review",
					Title: "Repaso: el rastreador web",
					Exercise: &course.Exercise{
						Path: "08-concurrency/crawler",
						Goal: "Rastrea enlaces en paralelo sin visitar dos veces la misma página.",
					},
				},
			},
		},
		{
			Slug:    "09-web",
			Title:   "Redes y la web",
			Summary: "Cómo hablan dos programas por la red: TCP, un socket propio, HTTP por dentro, una pequeña API y plantillas HTML.",
			Lessons: []course.Lesson{
				{
					Slug:  "01-tcp",
					Title: "TCP: dos programas que hablan",
					Exercise: &course.Exercise{
						Path: "09-web/echo",
						Goal: "Escribe un servidor de eco: devuelve cada línea que recibe.",
					},
				},
				{
					Slug:  "02-sockets",
					Title: "Tu propio socket",
					Exercise: &course.Exercise{
						Path: "09-web/kv",
						Goal: "Un pequeño servidor de clave-valor por TCP con las órdenes SET y GET.",
					},
				},
				{
					Slug:  "03-http",
					Title: "HTTP: texto sobre TCP",
					Exercise: &course.Exercise{
						Path: "09-web/request",
						Goal: "Lee una petición HTTP a mano y saca el método, la ruta y las cabeceras.",
					},
				},
				{
					Slug:  "04-net-http",
					Title: "Un servidor con net/http",
					Exercise: &course.Exercise{
						Path: "09-web/greet",
						Goal: "Escribe un handler que salude por el nombre que llega en la ruta.",
					},
				},
				{
					Slug:  "05-json-api",
					Title: "Una pequeña API JSON",
					Exercise: &course.Exercise{
						Path: "09-web/todos",
						Goal: "Una API de tareas: lista, crea y borra tareas en JSON con los códigos de estado correctos.",
					},
				},
				{
					Slug:  "06-templates",
					Title: "Plantillas con html/template",
					Exercise: &course.Exercise{
						Path: "09-web/page",
						Goal: "Genera una página HTML con una lista a partir de una plantilla.",
					},
				},
				{
					Slug:  "07-review",
					Title: "Repaso: tu pequeña web",
					Exercise: &course.Exercise{
						Path: "09-web/notes",
						Goal: "Una web de notas: una API JSON y una página HTML que las muestra.",
					},
				},
			},
		},
	},
}
