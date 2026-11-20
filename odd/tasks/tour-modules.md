# Feature: tour modules (basics → images)

Objective: add the four modules that follow the Go Tour route from "Packages" to "Exercise: Images", with lessons, compiled examples and local exercises.
Branch: `feat/tour-modules`. Delivery strategy: `ask-on-risk` (forecast ≈ 6000 authored lines, asked only when the user wants to publish).
Test runner: `go test ./...` (site) and `cd exercises && go test -tags solution ./...` (exercises). Rules of the repo: `CLAUDE.md` (read it first; it is not repeated here).

## Specs

### S1 — Scope

User request: "quiero que implementes todos los modulos hasta que llegues a Los Reader de imagenes creo que era esta ruta: https://go.dev/tour/methods/26".

Verified against `golang/website/_content/tour/methods.article`: `methods/24` is "Images", `methods/25` is "Exercise: Images", `methods/26` is "Congratulations!" (end of the chapter). So the scope is the tour chapters `basics`, `flowcontrol`, `moretypes` and `methods`, complete. Out of scope: `generics`, `concurrency`, and any change to `internal/`, `cmd/`, `web/templates/` or `00-setup`.

- 39 lessons and 27 exercises in 4 new modules (S7–S10). Related tour slides are merged so that each lesson keeps one idea.
- Assumption: lesson prose is original and follows the tour's route; it is not a translation of the tour.

### S2 — Phases and models

User: "diseñar el plan con fable e implementar con opus medium", "las animaciones las dejamos para el final", "antes de cada fase que se hará del tirón plantea el modelo recomendado".

| Phase | Work | Recommended model |
|---|---|---|
| 0 | This plan | Fable, effort high |
| 1 | Pilot: module `01-basics`, then the user reviews it | Opus, effort medium |
| 2 | Modules `02`, `03`, `04` in one go | Opus, effort medium |
| 3 | Pixel-art animations | Fable medium or Opus high |

Before each phase starts, stop and tell the user the recommended model so they can switch with `/model`. Corrections from the pilot review are appended to `## Log` and folded into S3–S5 before phase 2 starts.

### S3 — Lesson conventions

- One file per lesson: `content/<module>/<lesson>.html`, plus its entry in `content/catalog.go` (slug, title, exercise path and goal exactly as written in S7–S10).
- Section order inside the body, with `<h2>` headings: the motivating problem (opening paragraphs, no heading) → the concept → the example → "El error típico" → a one-paragraph bridge to the exercise when the lesson has one. The exercise box itself is rendered by `lesson.html` from the catalog; do not repeat the command in the body.
- Length: 250–500 words of prose. Short sentences. Spanish rules of `CLAUDE.md` apply to every word the student reads.
- Go terms stay in English, in `<em>` with a one-line Spanish explanation the first time they appear in the course; later lessons use them plainly.
- "Review" lessons (`Repaso: …`) are short: a recap list that links back to the module's lessons, then the bridge to the final exercise.
- Only the tour content listed for the lesson. A modern one-line note is allowed where the spec names it (for example `any`, `for i := range 10`).

### S4 — Example conventions

- Every runnable snippet is `content/<module>/examples/<name>/main.go`, embedded with `{{code "examples/<name>/main.go"}}`. Names are listed per lesson in S7–S10.
- Examples are 10–30 lines, deterministic (no `time.Now`, no `rand`, no ranging over a map with more than one key when the output is shown) and clean under `go vet`.
- Program output is shown right after the example as a plain `<pre><code>` block introduced by "Salida:". Run the example and paste the real output; never write it by hand.
- Code that must not compile cannot live in `examples/`. Paste the minimal broken snippet inline in `<pre><code class="language-go">`, then the compiler message in a plain `<pre><code>`. Reproduce it with the real toolchain in a `mktemp -d` scratch directory and paste the exact message, shortening only the file path to `./prog.go`. The same applies to runtime panics: paste the first line of the real panic.
- Each lesson shows the item named "On purpose" in its spec, and explains how to read the message (file, line, what the compiler expected).

### S5 — Exercise conventions

- Three files per exercise, as in `exercises/00-setup/hello/` (the reference to copy): `<name>.go` (`//go:build !solution`), `<name>_solution.go` (`//go:build solution`), `<name>_test.go`.
- Package name = folder name. `exercises/go.mod` says `go 1.22`: no newer language features, standard library only.
- The stub always compiles: every function and method exists and returns zero values, with a `// TODO:` line. Feedback is always a failing test, never a build error. The stub's tests must fail; the solution's tests must pass.
- The doc comment on each exported identifier states the goal and gives one example, in English (repo rule: code, comments and tests are in English). The catalog `Goal` is the Spanish text the student reads on the site.
- Tests are table-driven with `t.Run`, use only the API in the spec, and print friendly messages in the form `Name(args) = got, want want`. Floats are compared with a tolerance (`1e-9` unless stated).
- The solution uses only concepts taught up to that lesson, and is the obvious one, not the clever one.
- `exercises/00-setup/hello/hello.go` has a local, uncommitted change made by the user. Never stage it and never revert it.

### S6 — Animation placeholders

Animations are phase 3. In phases 1–2, lessons marked "Animation" leave one marker at the place where the animation will go, right after the concept section:

`<!-- animation: <tag-name> scene=<scene> -->`

The prose must stand without the animation (it is also the text fallback later). `html/template` drops HTML comments, so the marker never reaches the page.

### S7 — Module `01-basics`

Title: "Lo básico". Summary: "Paquetes, funciones, variables, tipos y constantes: las piezas de cualquier programa."

1. `01-packages` — "Paquetes e imports"
   - Tour: Packages; Imports; Exported names.
   - Teach: every file belongs to a package; `main` is where programs start; the package name is the last element of the import path; grouped imports; a name is exported when it starts with a capital letter.
   - Examples: `packages` (uses `fmt` and `math.Sqrt`), `exported-names` (`math.Pi`).
   - On purpose: `math.pi` (unexported name). Common mistake: an imported package that is not used is also a compile error.
2. `02-functions` — "Funciones"
   - Tour: Functions; Functions continued.
   - Teach: the type goes after the name; `x, y int` shares a type; the return type goes last.
   - Examples: `functions`.
   - On purpose: a function that declares `int` and returns a string. Common mistake: writing the type before the name, as in C or Java.
   - Exercise `01-basics/rectangle` — Goal: "Calcula el área y el perímetro de un rectángulo." API: `Area(width, height int) int`, `Perimeter(width, height int) int`.
3. `03-multiple-results` — "Varios resultados"
   - Tour: Multiple results; Named return values.
   - Teach: a function can return several values; named results and the bare `return`.
   - Examples: `multiple-results` (swap two strings), `named-results`.
   - On purpose: assigning two results to one variable (`assignment mismatch`). Common mistake: a bare `return` in a long function hides what is returned; use it only in short ones.
   - Exercise `01-basics/divmod` — Goal: "Devuelve el cociente y el resto de una división entera." API: `DivMod(a, b int) (quotient, remainder int)`. Cases: (7,2)→(3,1), (10,5)→(2,0), (3,7)→(0,3), (-7,2)→(-3,-1).
4. `04-variables` — "Variables"
   - Tour: Variables; Variables with initializers; Short variable declarations.
   - Teach: `var` at package and function level; initializers; `:=` only inside functions.
   - Examples: `variables`, `short-declarations`.
   - On purpose: `:=` outside a function. Common mistake: a declared variable that is never used does not compile.
5. `05-basic-types` — "Tipos básicos"
   - Tour: Basic types.
   - Teach: `bool`, `string`, the `int` and `uint` families, `byte`, `rune`, `float64`; mention `complex128` in one sentence; use `int` unless there is a reason not to.
   - Examples: `basic-types` (prints type and value with `%T` and `%v`).
   - On purpose: adding an `int` and a `float64` (`mismatched types`). Common mistake: integer division, `7 / 2` is `3`.
6. `06-zero-values` — "Valores cero"
   - Tour: Zero values.
   - Teach: a variable without an initial value holds the zero value of its type: `0`, `false`, `""`.
   - Examples: `zero-values`.
   - Common mistake: expecting an "undefined" or random value; in Go a variable is always usable. No compile error in this lesson.
7. `07-type-conversions` — "Conversiones e inferencia de tipos"
   - Tour: Type conversions; Type inference.
   - Teach: `T(v)` converts; Go never converts on its own; `:=` infers the type from the value on the right.
   - Examples: `type-conversions` (hypotenuse with `math.Sqrt`), `type-inference`.
   - On purpose: assigning a `float64` to an `int` variable without converting. Common mistake: `int(f)` truncates, it does not round.
   - Exercise `01-basics/average` — Goal: "Calcula la media de dos enteros sin perder los decimales." API: `Average(a, b int) float64`. Cases: (1,2)→1.5, (2,2)→2, (-1,2)→0.5, (0,0)→0.
8. `08-constants` — "Constantes"
   - Tour: Constants; Numeric constants.
   - Teach: `const`; constants cannot use `:=`; untyped numeric constants take the type the context needs.
   - Examples: `constants`, `numeric-constants` (`Big` and `Small` as in the tour).
   - On purpose: passing a constant that overflows `int`. Common mistake: trying to assign to a constant.
9. `09-review` — "Repaso: el círculo"
   - Exercise `01-basics/circle` — Goal: "Calcula el área y el perímetro de un círculo a partir de su radio." API: `Circle(radius int) (area, perimeter float64)` with `math.Pi`. Cases: 0→(0,0), 1→(π,2π), 2→(4π,4π), 10→(100π,20π).

### S8 — Module `02-flow-control`

Title: "Control de flujo". Summary: "Bucles, condiciones, switch y defer: cómo decide tu programa qué hacer y cuántas veces."

1. `01-for` — "El bucle for"
   - Tour: For; For continued.
   - Teach: Go has one loop; its three parts; no parentheses, braces always; the init and post parts are optional.
   - Examples: `for`.
   - On purpose: a `for` with parentheses around its three parts. Common mistake: off by one, `<` against `<=`.
   - Exercise `02-flow-control/sum` — Goal: "Suma todos los números del 1 al n." API: `SumTo(n int) int`. Cases: 0→0, 1→1, 10→55, 100→5050, -5→0.
2. `02-while-forever` — "for como while y bucles infinitos"
   - Tour: For is Go's "while"; Forever.
   - Teach: `for condition {}`; `for {}`; `break` leaves the loop.
   - Examples: `for-while`, `forever` (must end, with `break`).
   - Common mistake: a condition that never changes; how to stop a program with Ctrl+C. No compile error in this lesson.
   - Exercise `02-flow-control/halvings` — Goal: "Cuenta cuántas veces puedes dividir un número entre dos hasta llegar a cero." API: `Halvings(n int) int`. Cases: 0→0, 1→1, 8→4, 1000→10.
3. `03-if` — "if y else"
   - Tour: If; If with a short statement; If and else.
   - Teach: `if` without parentheses; the short statement and its scope; `else` and `else if`. Introduce `strconv.Itoa` in the example, the exercise needs it.
   - Examples: `if`, `if-short-statement`.
   - On purpose: `else` written on its own line after `}`. Common mistake: using the short-statement variable after the `if` ends.
   - Exercise `02-flow-control/fizzbuzz` — Goal: "Devuelve Fizz, Buzz, FizzBuzz o el número, según sus divisores." API: `FizzBuzz(n int) string`. Cases: 3→"Fizz", 5→"Buzz", 15→"FizzBuzz", 7→"7", 30→"FizzBuzz", 1→"1".
4. `04-sqrt` — "Práctica: tu propia raíz cuadrada"
   - Tour: Exercise: Loops and Functions.
   - Teach: Newton's method in plain words; the step `z -= (z*z - x) / (2*z)` shown inline as a formula; stop when the value barely changes. No compiled example: it would be the solution.
   - Common mistake: comparing floats with `==`.
   - Exercise `02-flow-control/sqrt` — Goal: "Calcula la raíz cuadrada con el método de Newton, sin usar math.Sqrt." API: `Sqrt(x float64) float64`. Cases, against `math.Sqrt` with tolerance `1e-9`: 0, 1, 2, 4, 9, 100, 0.25, 2e6.
5. `05-switch` — "switch"
   - Tour: Switch; Switch evaluation order; Switch with no condition.
   - Teach: only the matching case runs, there is no automatic fallthrough; several values per case; cases are tried top to bottom; `switch {}` replaces a long `if`/`else` chain.
   - Examples: `switch`, `switch-no-condition`.
   - On purpose: the same constant in two cases (`duplicate case`). Common mistake: expecting the next case to run, as in C.
   - Exercise `02-flow-control/season` — Goal: "Indica a qué estación del año pertenece un mes." API: `Season(month int) string`, northern hemisphere. Cases: 12, 1, 2→"winter"; 3–5→"spring"; 6–8→"summer"; 9–11→"autumn"; 0 and 13→"unknown".
6. `06-defer` — "defer"
   - Tour: Defer; Stacking defers.
   - Teach: a deferred call runs when the surrounding function returns; its arguments are evaluated at the `defer` line; several defers run last in, first out. One sentence on the real use: closing what you opened.
   - Examples: `defer`, `defer-stack`.
   - Common mistake: thinking the arguments are evaluated when the call finally runs. Ask the student to predict the output before reading it. No compile error in this lesson.
   - Animation (optional candidate, decided in phase 3): `defer-demo scene=stack`.
7. `07-review` — "Repaso: números primos"
   - Exercise `02-flow-control/primes` — Goal: "Decide si un número es primo y cuenta cuántos primos hay hasta un límite." API: `IsPrime(n int) bool`, `CountPrimes(limit int) int` (primes ≤ limit). Cases: IsPrime -3, 0, 1, 4, 25→false; 2, 3, 17, 97→true. CountPrimes 1→0, 2→1, 10→4, 100→25.

### S9 — Module `03-more-types`

Title: "Más tipos". Summary: "Punteros, structs, slices, maps y funciones como valores: cómo se organizan los datos."

1. `01-pointers` — "Punteros"
   - Tour: Pointers.
   - Teach: a pointer holds the address of a value; `&` takes it, `*` follows it; the zero value is `nil`; there is no pointer arithmetic.
   - Examples: `pointers`.
   - On purpose: the panic of following a `nil` pointer. Common mistake: changing a copy and expecting the original to change.
   - Animation: `pointer-demo scene=basics`.
   - Exercise `03-more-types/swap` — Goal: "Intercambia dos enteros usando punteros." API: `Swap(a, b *int)`.
2. `02-structs` — "Structs"
   - Tour: Structs; Struct Fields; Pointers to structs; Struct Literals.
   - Teach: a struct groups fields; the dot; `p.X` works through a pointer; literals with field names; `&T{}`.
   - Examples: `structs`, `struct-pointers`, `struct-literals`.
   - On purpose: a field that does not exist. Common mistake: passing a struct by value and expecting the function to change it.
   - Exercise `03-more-types/player` — Goal: "Crea un jugador y súmale puntos." API: `type Player struct { Name string; Score int }`, `NewPlayer(name string) *Player`, `AddPoints(p *Player, points int)`.
3. `03-arrays` — "Arrays"
   - Tour: Arrays.
   - Teach: fixed length; the length is part of the type; arrays are values, assigning copies them.
   - Examples: `arrays`.
   - On purpose: a constant index out of range. Common mistake: `[3]int` and `[5]int` are different types.
4. `04-slices` — "Slices"
   - Tour: Slices; Slices are like references to arrays; Slice literals; Slice defaults.
   - Teach: `a[low:high]`; a slice stores no data, it is a window over an array; two slices can share the same array; literals; default bounds.
   - Examples: `slices`, `slices-share-array`, `slice-literals`.
   - On purpose: the panic of an index out of range. Common mistake: changing one slice changes every slice that shares its array.
   - Animation: `slice-demo scene=view`.
5. `05-slice-len-cap` — "Longitud y capacidad"
   - Tour: Slice length and capacity; Nil slices; Creating a slice with make.
   - Teach: `len` and `cap`; re-slicing up to the capacity; the `nil` slice; `make([]T, len, cap)`.
   - Examples: `slice-len-cap`, `nil-slices`, `making-slices`.
   - On purpose: the panic of re-slicing beyond the capacity. Common mistake: treating a `nil` slice as an error; it has length 0 and works with `append`.
   - Animation: `slice-demo scene=len-cap`.
6. `06-append` — "Añadir elementos con append"
   - Tour: Appending to a slice.
   - Teach: `append` returns a slice; when the capacity runs out, Go allocates a bigger array and copies.
   - Examples: `append`.
   - On purpose: calling `append` without using its result. Common mistake: forgetting `s = append(s, x)`.
   - Animation: `slice-demo scene=append`.
7. `07-range` — "Recorrer con range"
   - Tour: Range; Range continued.
   - Teach: `for i, v := range s`; `_` to skip one; the value is a copy. Modern note in one line: `for i := range 10`.
   - Examples: `range`, `range-skip`.
   - On purpose: an index that is declared and not used, fixed with `_`. Common mistake: changing `v` does not change the slice.
   - Exercise `03-more-types/evens` — Goal: "Devuelve solo los números pares de una lista." API: `Evens(nums []int) []int`. Cases: nil→length 0, [1,2,3,4]→[2,4], [1,3]→length 0, [-2,0,5]→[-2,0]; the input is not modified.
8. `08-pic` — "Práctica: slices de slices"
   - Tour: Slices of slices; Exercise: Slices.
   - Teach: a slice whose elements are slices; each row is created separately.
   - Examples: `slices-of-slices` (a tic-tac-toe board).
   - Common mistake: allocating the outer slice and forgetting the rows (index out of range).
   - Exercise `03-more-types/pic` — Goal: "Construye una cuadrícula de dy filas y dx columnas." API: `Pic(dx, dy int) [][]uint8`, where cell `[y][x]` is `uint8((x + y) / 2)`. Cases: (0,0)→length 0; (3,2)→2 rows of 3 with the right values; (256,256) checks a few cells; changing one row must not change another.
9. `09-maps` — "Maps"
   - Tour: Maps; Map literals; Map literals continued; Mutating Maps.
   - Teach: key → value; `make` and literals; insert, read, `delete`; `v, ok := m[k]`.
   - Examples: `maps`, `map-literals`, `mutating-maps`.
   - On purpose: the panic of writing to a `nil` map. Common mistake: a missing key returns the zero value; ask with `ok`. One sentence: the iteration order is not fixed.
   - Exercise `03-more-types/wordcount` — Goal: "Cuenta cuántas veces aparece cada palabra en un texto." API: `WordCount(s string) map[string]int` with `strings.Fields` (introduced in the lesson). Cases: ""→length 0; "I am learning Go"→1 each; "the quick the lazy the"→the:3; repeated spaces; "Go go"→two different words.
10. `10-function-values` — "Funciones como valores"
    - Tour: Function values.
    - Teach: a function is a value; it can be stored in a variable and passed as an argument.
    - Examples: `function-values`.
    - On purpose: passing a function with the wrong signature. Common mistake: `f()` calls the function, `f` passes it.
    - Exercise `03-more-types/apply` — Goal: "Aplica una función a cada elemento de una lista." API: `Apply(nums []int, f func(int) int) []int`. Cases: double, square, empty input; the input is not modified.
11. `11-closures` — "Closures"
    - Tour: Function closures; Exercise: Fibonacci closure.
    - Teach: a function that remembers variables from where it was created; each call to the outer function creates new, independent variables.
    - Examples: `closures` (the adder).
    - Common mistake: believing two closures share state when each has its own. No compile error in this lesson.
    - Animation (optional candidate, decided in phase 3): `closure-demo scene=adder`.
    - Exercise `03-more-types/fibonacci` — Goal: "Devuelve una función que genera la sucesión de Fibonacci." API: `Fibonacci() func() int`. Cases: first ten calls→0, 1, 1, 2, 3, 5, 8, 13, 21, 34; two generators are independent.
12. `12-review` — "Repaso: el inventario"
    - Exercise `03-more-types/inventory` — Goal: "Gestiona un inventario con structs, slices, maps y funciones." API: `type Item struct { Name string; Quantity int }`, `Total(items []Item) int`, `Restock(item *Item, amount int)`, `ByName(items []Item) map[string]int` (adds up repeated names), `Filter(items []Item, keep func(Item) bool) []Item`.

### S10 — Module `04-methods`

Title: "Métodos e interfaces". Summary: "Métodos, interfaces y las más usadas de la biblioteca estándar: Stringer, error, Reader e Image."

1. `01-methods` — "Métodos"
   - Tour: Methods; Methods are functions; Methods continued.
   - Teach: a method is a function with a receiver; any type defined in your package can have methods, not only structs.
   - Examples: `methods`, `methods-on-any-type` (`MyFloat`).
   - On purpose: declaring a method on `int`. Common mistake: looking for classes; Go has types with methods.
   - Exercise `04-methods/temperature` — Goal: "Añade métodos a un tipo para convertir temperaturas." API: `type Celsius float64`, `(c Celsius) Fahrenheit() float64`, `(c Celsius) IsFreezing() bool` (true at 0 or below). Cases: 0→32, 100→212, -40→-40, 37→98.6.
2. `02-pointer-receivers` — "Receptores de puntero"
   - Tour: Pointer receivers; Pointers and functions; Methods and pointer indirection (1 and 2); Choosing a value or pointer receiver.
   - Teach: a value receiver works on a copy, a pointer receiver on the original; Go adds `&` or `*` for you in method calls; pick one kind of receiver per type.
   - Examples: `pointer-receivers`, `indirection`.
   - On purpose: calling a pointer method on a value that has no address (`Vertex{3, 4}.Scale(2)`). Common mistake: a value receiver that "changes" a field and nothing happens.
   - Animation (optional candidate, decided in phase 3): `pointer-demo scene=receivers`.
   - Exercise `04-methods/counter` — Goal: "Crea un contador que se incrementa y se reinicia." API: `type Counter struct { count int }`, `(c *Counter) Increment()`, `(c *Counter) Reset()`, `(c Counter) Value() int`.
3. `03-interfaces` — "Interfaces"
   - Tour: Interfaces; Interfaces are implemented implicitly.
   - Teach: an interface is a set of methods; a type implements it just by having them, there is no `implements`.
   - Examples: `interfaces`, `implicit-interfaces`.
   - On purpose: a value that does not implement the interface because the method has a pointer receiver. Common mistake: declaring the interface next to the type "just in case".
   - Exercise `04-methods/shapes` — Goal: "Calcula el área total de varias figuras con una interfaz." API: `type Shape interface { Area() float64 }`, `Rectangle{Width, Height float64}`, `Circle{Radius float64}`, `TotalArea(shapes []Shape) float64`.
4. `04-interface-values` — "Valores de interfaz"
   - Tour: Interface values; Interface values with nil underlying values; Nil interface values.
   - Teach: an interface value is a pair (concrete type, value); the value inside can be `nil`; an interface with nothing inside is `nil`.
   - Examples: `interface-values`, `nil-underlying-value`.
   - On purpose: the panic of calling a method on a `nil` interface. Common mistake: an interface that holds a `nil` pointer is not itself `nil`.
   - Animation: `interface-demo scene=pair`.
5. `05-any` — "La interfaz vacía y las type assertions"
   - Tour: The empty interface; Type assertions.
   - Teach: `interface{}` holds any value and `any` is its modern name; `t, ok := i.(T)` gets the concrete value back.
   - Examples: `empty-interface`, `type-assertions`.
   - On purpose: the panic of an assertion to the wrong type. Common mistake: asserting without `ok`.
6. `06-type-switches` — "Type switches"
   - Tour: Type switches.
   - Teach: `switch v := i.(type)` tries several types in order.
   - Examples: `type-switches`.
   - On purpose: `i.(type)` outside a `switch`. Common mistake: reaching for `any` when a concrete type would do.
   - Exercise `04-methods/describe` — Goal: "Describe un valor según su tipo." API: `Describe(v any) string`. Cases: 42→"int: 42", "hi"→"string: hi", true→"bool: true", nil→"nil", 3.5→"unknown".
7. `07-stringers` — "Stringer"
   - Tour: Stringers; Exercise: Stringers.
   - Teach: `fmt.Stringer`; `fmt` calls `String()` when a type has it.
   - Examples: `stringer`.
   - Common mistake: `String()` with a pointer receiver is not used when you print the value. No compile error in this lesson.
   - Exercise `04-methods/ipaddr` — Goal: "Haz que una dirección IP se imprima como 127.0.0.1." API: `type IPAddr [4]byte`, `(ip IPAddr) String() string`. Tests go through `fmt.Sprint`.
8. `08-errors` — "Errores"
   - Tour: Errors; Exercise: Errors.
   - Teach: `error` is an interface; `if err != nil`; `errors.New` and `fmt.Errorf`; your own error type.
   - Examples: `errors` (`strconv.Atoi` failing), `custom-error`.
   - On purpose: an error that is declared and not used. Common mistake: `fmt.Sprint(e)` inside `Error()` calls itself forever; convert first.
   - Exercise `04-methods/safesqrt` — Goal: "Devuelve un error cuando te pidan la raíz de un número negativo." API: `type ErrNegativeSqrt float64` with `Error()` returning "cannot Sqrt negative number: -2", and `Sqrt(x float64) (float64, error)` (`math.Sqrt` allowed). Cases: 4→(2, nil); -2→(0, ErrNegativeSqrt(-2)).
9. `09-readers` — "Readers"
   - Tour: Readers; Exercise: Readers.
   - Teach: `io.Reader`; `Read` fills the slice it receives and says how many bytes it wrote; `io.EOF` means the data ended.
   - Examples: `reader` (`strings.NewReader` read into an 8-byte buffer in a loop).
   - Common mistake: using `b` where `b[:n]` is needed; treating `io.EOF` as a failure. No compile error in this lesson.
   - Animation: `reader-demo scene=chunks`.
   - Exercise `04-methods/areader` — Goal: "Crea un Reader que emite la letra A sin parar." API: `type AReader struct{}`, `(AReader) Read(b []byte) (int, error)`. Cases: buffers of 1, 8 and 1024 bytes are filled with 'A' and return (len(b), nil); an empty buffer returns 0.
10. `10-rot13` — "Práctica: rot13Reader"
    - Tour: Exercise: rot13Reader.
    - Teach: a reader that wraps another reader and changes the bytes as they pass.
    - Examples: `upper-reader` (the same pattern, turning text to upper case).
    - Common mistake: transforming the whole buffer when only the first `n` bytes were read.
    - Exercise `04-methods/rot13` — Goal: "Descifra un mensaje con un Reader que aplica ROT13." API: `type Rot13Reader struct { R io.Reader }` with `Read`. Cases: "Lbh penpxrq gur pbqr!"→"You cracked the code!"; digits and punctuation unchanged; upper and lower case; works one byte at a time (`testing/iotest.OneByteReader`); `io.EOF` is passed through.
11. `11-images` — "Imágenes"
    - Tour: Images; Exercise: Images.
    - Teach: `image.Image` is an interface with three methods: `ColorModel`, `Bounds`, `At`. One closing paragraph tells the student the tour route ends here.
    - Examples: `images` (`image.NewRGBA`, its bounds and one pixel).
    - On purpose: a type that lacks one of the three methods, used as `image.Image` (`missing method`).
    - Exercise `04-methods/images` — Goal: "Crea tu propio tipo de imagen implementando image.Image." API: `type Image struct { Width, Height int }` with `ColorModel()` → `color.RGBAModel`, `Bounds()` → `image.Rect(0, 0, Width, Height)`, `At(x, y)` → `color.RGBA{uint8(x), uint8(y), 255, 255}`. Tests: it satisfies `image.Image`, bounds, several pixels, and `png.Encode` into a buffer returns no error.

### S11 — Verification

Run from the repository root at the end of every task; all must be clean.

```bash
gofmt -l .                       # prints nothing
go vet ./...
go test ./...
go run ./cmd/paseo build -out "$(mktemp -d)"
(cd exercises && go vet ./... && go vet -tags solution ./... && go test -tags solution ./...)
(cd exercises && go test ./<module>/...)   # every new package must FAIL: stubs are not solutions
grep -rnE 'computadora|\bcorr(e|er|é)\b|abrí|ejecutá|fijate|mirá|probá|tenés|podés|usá\b|\busted\b' content/   # review every hit
```

Test-first: add the catalog entry before the body, so `TestEveryLessonHasBody` fails first; write each exercise's test before its solution.

### S12 — Animations (phase 3)

User: "Es el paso de las animaciones, arranca". Rules: `CLAUDE.md`, section "Animations" (read it; it is not repeated here).

Decision on the optional candidates: `pointer-demo scene=receivers` and `defer-demo scene=stack` are in (a copy that vanishes and a LIFO stack are things prose cannot show); `closure-demo` is out, the two-adder example already shows independent state. Its marker is removed.

Infrastructure (shared by every component):

- `web/static/js/animations.js`, loaded once from `layout.html` as a module: for each custom element on the page that is not defined yet, it imports `animations/<tag>.js`. A lesson downloads only its own animations, and it works under any `-base`.
- `web/static/js/animations/pixel-scene.js`: the base class `PixelScene` every component extends. It owns the canvas (160×96 logical pixels, scaled by the largest integer factor that fits, `image-rendering: pixelated`), the controls (`Reproducir`/`Pausa`, `Paso`, `Reiniciar`, real `<button>`s), the caption (`aria-live="polite"`), the code panel with the current line highlighted, the `--px-*` palette read from CSS, a 3×5 pixel font and drawing helpers. A component only declares `scenes` (plain data: `code` lines and `steps`, each `{ line, caption, … }`) and a `draw(paint, step, t)` that renders one step at progress `t`.
- Motion: each step plays a few frames (`t` goes 0→1 in 4 frames) and then holds. Starts paused on step 1. Play advances one step at a time until the end. With `prefers-reduced-motion`, `t` is always 1 (no tweening).
- Markup in the lesson, replacing the marker: `<pointer-demo scene="basics"><p>…text fallback…</p></pointer-demo>`. The fallback is a short paragraph that says what the animation shows; the component replaces it when it starts.
- Test: a content test fails when a lesson uses a `<name-demo>` tag that has no `web/static/js/animations/<name-demo>.js`, or when an `<!-- animation:` marker is left.

Scenes (captions in Spanish, one per step, next to the line of Go they represent; values taken from the real output of each lesson's example):

1. `pointer-demo scene=basics` (`03-more-types/01-pointers`, example `pointers`): `i := 42` puts 42 in a box with an address → `p := &i` creates `p` holding that address and an arrow grows to `i` → `*p` follows the arrow and reads 42 → `*p = 21` writes 21 into `i` → `fmt.Println(i)` prints 21 → `var q *int` is a box with `nil` and no arrow.
2. `pointer-demo scene=receivers` (`04-methods/02-pointer-receivers`, example `pointer-receivers`): `v := Vertex{3, 4}` → `v.ScaleCopy(10)` copies `v` into the method and the copy becomes `{30 40}` → the method returns, the copy vanishes, `v` is still `{3 4}` → `v.Scale(10)` receives `&v` (an arrow to `v`) → the original becomes `{30 40}`.
3. `slice-demo scene=view` (`03-more-types/04-slices`, example `slices-share-array`): the array `names` with 4 cells → `a := names[0:2]` is a window over cells 0–1 → `b := names[1:3]` a second window over 1–2 → `b[0] = "XXX"` changes cell 1, and both windows show it.
4. `slice-demo scene=len-cap` (`03-more-types/05-slice-len-cap`, example `slice-len-cap`): len 6 cap 6 → `s[:0]` len 0 cap 6 → `s[:4]` len 4 cap 6 → `s[2:]` len 2 cap 4 (the start moves, the end does not).
5. `slice-demo scene=append` (`03-more-types/06-append`, example `append`): `var s []int` is nil, len 0 cap 0 → append 0: a new array of cap 1 → append 1: full, Go allocates cap 2 and copies → append 2, 3, 4: len 5 cap 6, copied again; old arrays fade to `--px-dim`.
6. `interface-demo scene=pair` (`04-methods/04-interface-values`, examples `interface-values` and `nil-underlying-value`): an interface value drawn as a pair (type, value) → `var d Describer` is (nil, nil), `d == nil` → `d = Temp(21.5)` is (Temp, 21.5) → `d = Word("gopher")` is (Word, "gopher") → `d = b` with a nil `*Box` is (*Box, nil), and `d == nil` is false.
7. `reader-demo scene=chunks` (`04-methods/09-readers`, example `reader`): the string "Hola, Reader!" as 13 bytes and a buffer of 8 → first `Read`: 8 bytes move in, `n = 8` → second `Read`: 5 bytes overwrite the start, the last 3 cells keep old bytes and only `b[:n]` is lit → third `Read`: `n = 0`, `err = EOF`.
8. `defer-demo scene=stack` (`02-flow-control/06-defer`, example `defer-stack`): prints "contando" → each loop turn pushes a deferred `Println(i)` card onto a stack (1, 2, 3) → prints "hecho" → the function returns and the cards pop 3, 2, 1 into the output.

Colours keep one meaning (`style.css`): value `--px-value`, address/arrow/window `--px-pointer`, the thing that changes in this step `--px-highlight`, nil/EOF `--px-error`, discarded memory `--px-dim`, text `--px-paper`.

## Tasks

- [x] T0 — Plan (S1–S12). Route: inline. Commit: `docs: plan the tour modules feature`.
- [x] T1 — Module `01-basics` (S3–S7, S11), 9 lessons, 4 exercises. Route: inline (single unit, no writer reason). Commits: `c6d225a`, `feat(content): add basics lessons on types, conversions and constants`. Reviewed and accepted by the user (L6).
- [x] T2 — Module `02-flow-control` (S3–S6, S8, S11), 7 lessons, 6 exercises. Route: delegated (parallel writer, trigger (a)). Commits: `874b20e`, `c7f1c3a`.
- [x] T3 — Module `03-more-types` (S3–S6, S9, S11), 12 lessons, 8 exercises. Route: delegated (parallel writer, trigger (a)). Commits: `1e26e5f`, `244ff85`, `6b4c8eb`.
- [x] T4 — Module `04-methods` (S3–S6, S10, S11), 11 lessons, 9 exercises. Route: delegated (parallel writer, trigger (a)). Commits: `6a3628d`, `1f48ea6`, the commit `feat(content): add lessons on errors, readers and images`.
- [x] T5 — Animations (S6, S12). Split in two:
  - [x] T5a — Infrastructure + `pointer-demo` (scenes `basics`, `receivers`) as the pilot. Route: inline (the shared base must exist before any other component). Commit: `feat(web): add pixel-art animation base and pointer demo` (see L9).
  - [x] T5b — `slice-demo`, `interface-demo`, `reader-demo`, `defer-demo`. Route: delegated, four parallel writers (Writer trigger (a): four independent units with disjoint surfaces, each well over 100 lines). Parent owns `pixel-scene.js`, `content/catalog_test.go`, the closure marker and git. Commits: `ee84755`, `2b00be2`, `51e0d7e`, `5e8bd77`, `test(content): forbid leftover animation markers`. Evidence: L10.

T2–T4 route evidence: three independent units with disjoint edit surfaces (`content/<module>/**`, `exercises/<module>/**`), each well over 100 lines, launched together (Writer trigger, parallelism). `content/catalog.go` is shared: the parent adds the three modules' entries before launching and runs the seam check (S11) at the end. Fallback: inline, one module after another.

Commits: slices of 3–5 lessons with their examples, exercises and catalog entries, so every commit passes S11. Message form: `feat(content): add <module> lessons on <topics>`.

## Log

- L1 (2026-11-20, user, original request): "quiero que implementes todos los modulos hasta que llegues a Los Reader de imagenes creo que era esta ruta: https://go.dev/tour/methods/26, quiero hacerlo con fablo que me recomiendas diseñar el plan con fable e implementar con opus medium?"
- L2 (2026-11-20, user): "a que effort cambio? las animaciones las dejamos para el final, si te parece para cambiar de modelo, antes de cada fase que se hará del tirón plantea el modelo recomendado". Then: "Arranca a diseñar el plan".
- L3 (2026-11-20): plan written on branch `feat/tour-modules`. Decisions: tour slides merged into 39 one-idea lessons; each module ends with a review lesson that holds the combined exercise (in `04-methods` the images exercise plays that role); stubs always compile, so feedback is a failing test; program output is pasted from a real run rather than checked by a test (a checked `output.txt` per example was considered and left out to avoid new infrastructure; revisit if outputs go stale). Risk: none of items 1–5; prose quality has no automated check, covered by the pilot review after T1.
- L4 (2026-11-20, user): "Apruebo el plan".
- L5 (2026-11-20, user): "Arranca con el plan". T1 done inline on Opus medium. RED: `TestEveryLessonHasBody` failed for the 9 new slugs; every stub test failed and every solution passed. S11 clean (gofmt, vet, tests, build, voseo grep with no hits). Compiler messages reproduced with go1.26.8. Findings: `int(4.9)` with a constant does not compile (`cannot convert 4.9 (untyped float constant) to type int`), so conversion examples use variables; lessons add a "Un error a propósito" `<h2>` between the example and "El error típico", and a "Tu turno" `<h2>` for the bridge; `06-zero-values` gained a paragraph on `var` versus `:=` to reach the length in S3. Subtests are unnamed (`#00`); the failure message carries the arguments.
- L6 (2026-11-20, user): "Dale todo lo que use opus 5.5 del tirón". Read as: pilot accepted with no corrections; run phase 2 (T2–T4) in one go. Conventions from L5 (extra `<h2>` sections, variables in conversion examples) apply to T2–T4. The parent added the three modules to `content/catalog.go` before launching the writers.
- L7 (2026-11-20, phase 2 evidence): three parallel Opus writers, disjoint surfaces, no writes outside them. Seam check S11 clean on the whole repo: gofmt 0 files, `go vet ./...` ok, `go test ./...` ok, build wrote 41 pages (40 lessons + index), exercises vet ok with and without the solution tag, every solution passes, all 27 stub packages fail, voseo grep 0 hits. Each commit slice was verified alone (catalog trimmed to its lessons). Findings: `0.1+0.2 == 0.3` with constants is true, so `04-sqrt` speaks of variables; Newton on floats oscillates, which motivates a tolerance stop; `&Vertex{3, 4}.Scale(2)` parses as `&(…)`, so the fix stores the value first; `i.(type)` outside a switch reports `invalid syntax tree: use of .(type) outside type switch`; a constant out-of-range slice index on a literal panics at run time. Two inline fragments in `03-more-types` show the message without a position because they are not full files. Review lessons (`07-review`, `12-review`) are about 200 words by design.
- L8 (2026-11-20, user): "Es el paso de las animaciones, arranca". Running on Opus 5.5, effort high, which is the model S2 recommends for phase 3. S12 detailed (above); T5 split into T5a (pilot, inline) and T5b.
- L9 (2026-11-20, T5a evidence): RED `TestEveryAnimationHasComponent` failed on `<pointer-demo>` before `pointer-demo.js` existed, then went GREEN. S11 clean: gofmt 0 files, `go vet ./...` ok, `go test ./...` ok, every solution passes, `node --check` ok on the three JS files. Build test now pins `static/js/animations.js` and `static/js/animations/pixel-scene.js`; base-path test pins `src="/paseo/static/js/animations.js"`. Visual check in headless Chrome (the Chrome extension was not connected), with a CDP script that clicks `Paso` and builds a contact sheet: both scenes, 7 steps each, captions and current code line match, no console errors; 900 px window gives scale 4 (640 px), 390 px gives scale 2 (320 px). Findings: a custom element is inline, so the base measures its block wrapper; `{`/`}` glyphs needed a middle bump so they do not read as `<`/`>`; output lines carry an explicit `printed` flag so an old line is not highlighted again. Port 8080 was held by another running `paseo serve`, left untouched; the check used 127.0.0.1:8091.
- L10 (2026-11-20, T5b evidence): four parallel writers on disjoint surfaces; `git status` showed no writes outside them. Each one observed RED on `TestEveryAnimationHasComponent` for its own tag before creating the JS file, then GREEN. The parent added `TestNoAnimationMarkerLeft`: RED with 7 lessons, then the `closure-demo` marker was removed (S12 drops it) and the writers' tags replaced the rest. Seam S11 clean: gofmt 0 files, `go vet ./...` ok, `go test ./...` ok, solutions ok, `node --check` ok on every animation file, no hex colours in JS. A headless visual check of every scene at 900 px (canvas 640 px) and 390 px (canvas 320 px) found no console errors and no `?` glyphs. Fix: the shared `*` glyph read as `x`, so it was redrawn in `pixel-scene.js` and the writer's local workaround was removed. Steps per scene: view 7, len-cap 8, append 10, pair 7, chunks 9, stack 8. Not observed: mid-motion frames, since the contact sheet captures each step at t=1; defer was the exception, its writer captured them at 280 ms. Next: user review, then push/PR (the user's decision).
