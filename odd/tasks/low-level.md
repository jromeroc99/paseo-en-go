# Feature: low-level modules, more animations and write-it-all exercises

Objective: three new modules on how Go works underneath (memory, buffers, processes and the scheduler) plus the generics the official Go tour covers and a closing module on networks and the web (TCP, HTTP, a small API, templates), many more animations across the course, and exercises where the student writes all the code.
Branch: `feat/low-level`. Delivery strategy: `ask-on-risk` (forecast ≈ 9000 authored lines; asked only when the user wants to publish).
Test runner: `go test ./...`, `go vet ./...` and `cd exercises && go test -tags solution ./...`. Repo rules: `CLAUDE.md` (read it first; it is not repeated here).

## Specs

### S1 — Scope

User: "Quiero que haya muchas más animaciones en las explicaciones y extiendas las lecciones a conceptos de bajo nivel, bffer, memoria, procesos planifica".
Approved syllabus (user: "me cuadra"):

- Part A — about 12 new animations in existing lessons, only where something invisible is shown: `01-basics` (a variable as a box with an address, `int8` overflow, `float → int` losing decimals), `02-flow-control` (`for` step by step, `switch` top-down evaluation), `03-more-types` (contiguous struct fields, copying an array copies everything, `range` gives a copy, map buckets, the variable a closure captures), `04-methods` (value vs pointer receiver, type switch).
- Part B — three new low-level modules, every lesson with its animation, plus `06-generics` (S5) and `09-web` (S9):
  - `05-memory` "La memoria por dentro": bytes and addresses, the stack and call frames, the heap and escape analysis (`-gcflags=-m`), the garbage collector (mark & sweep), struct padding (`unsafe.Sizeof`), strings, bytes and runes (UTF-8).
  - `07-buffers` "Buffers y E/S": what a buffer is, `bytes.Buffer`, `bufio.Writer` and `Flush` (forgetting it is the typical mistake), `bufio.Scanner`, `io.Copy`.
  - `08-concurrency` "Procesos, hilos y goroutines": process vs thread, goroutines, the G-M-P scheduler, unbuffered and buffered channels, `range` and `close`, `select` and `default`, `WaitGroup`, `Mutex` and data races (`-race`).
- Module order (user: "sí, déjalo en 5"): `05-memory`, `06-generics`, `07-buffers`, `08-concurrency`, `09-web` (the course ends there, S9). Generics goes right after memory because it builds on interfaces, as in the tour; buffers then sit next to concurrency, where buffered channels reuse the idea.

### S2 — Exercises: the student writes all the code

User: "Luego los ejercicios, añade mas ejercicios y que el alumno escriba todo el código para mejorar la sintaxis y tenga mas soltura."

- New exercises (modules 05–09): the stub `<name>.go` holds only the build tag, the package clause and a doc comment that states what to write, signatures included. No function skeletons, no bodies. Until the student writes them, `go test` fails with `undefined: …`, which the lesson's "Tu turno" card explains.
- The "Tu turno" card shows the exact signatures to write, as a Go code block.
- Existing exercises (modules 00–04) are converted to the same style in phase 4.
- "Más ejercicios": every lesson of modules 05–09 has its own exercise; in phase 4 the existing lessons without one get one. Assumption: one exercise per lesson (the domain allows one), not several.

### S3 — Phases

| Phase | Work |
|---|---|
| 1 | Pilot: module `05-memory` complete; the user reviews it |
| 2 | Modules `06-generics`, `07-buffers`, `08-concurrency` and `09-web` (S5–S9) |
| 3 | Part A: new animations in existing lessons |
| 4 | Exercises: write-it-all stubs in modules 00–04, plus exercises for lessons without one |

Corrections from the pilot review are logged in `## Log` and folded into the specs before phase 2.

### S4 — Module 05-memory (pilot)

| Lesson | Animation (component · scene) | Exercise · goal |
|---|---|---|
| `01-bytes` "Bytes y direcciones" | `memory-demo` · `bytes` | `05-memory/bits` · "Decide cuántos bits necesitas para guardar un número." |
| `02-stack` "El stack y las llamadas" | `stack-demo` · `calls` | `05-memory/factorial` · "Calcula el factorial con una función recursiva." |
| `03-heap` "El heap y el escape analysis" | `stack-demo` · `escape` | `05-memory/linkedlist` · "Construye una lista enlazada con punteros a nodos." |
| `04-gc` "El recolector de basura" | `gc-demo` · `mark-sweep` | `05-memory/reachable` · "Marca qué objetos siguen vivos, como hace el recolector." |
| `05-padding` "Alineación y padding" | `memory-demo` · `padding` | `05-memory/layout` · "Calcula cuánto ocupa un struct, con su padding." |
| `06-strings` "Strings, bytes y runes" | `memory-demo` · `utf8` | `05-memory/runes` · "Cuenta y da la vuelta a los caracteres de un texto UTF-8." |
| `07-review` "Repaso: tu propia memoria" | — | `05-memory/arena` · "Reserva trozos alineados de un bloque de memoria." |

### S5 — Nothing from the official Go tour is missing

User: "sí, déjalo en 5 y sigue con la fase 2, pero ten en cuenta el tour original de go, que no se nos olvide nada"

- Audit (2026-11-27) against `golang/website/_content/tour/*.article`: every page of Welcome, Basics, Flow control, More types and Methods and interfaces is already covered by modules 00–04 (named results, exported names, `if` with a short statement, `switch` order and without a condition, stacked `defer`, nil slices, `make`, map literals and mutation, method values vs functions, nil interface values, type assertions, and the exercises Loops, Slices, Maps, Fibonacci closure, Stringers, Errors, Readers, rot13Reader and Images).
- Missing: Generics (Type parameters, Generic types) → module `06-generics` (S6). In Concurrency: Range and Close, Default Selection, Exercise: Equivalent Binary Trees and Exercise: Web Crawler → added to `08-concurrency` (S8). The tour's `golang.org/x/tour/tree` package is not a dependency: the exercise declares its own `Tree` type.

### S6 — Module 06-generics "Genéricos"

Summary: "Funciones y tipos que sirven para cualquier tipo: parámetros de tipo, restricciones y colecciones genéricas."

| Lesson | Animation (component · scene) | Exercise · goal |
|---|---|---|
| `01-type-parameters` "Parámetros de tipo" | `generic-demo` · `instantiate` (the compiler fills `T` with `int`, then with `string`: two functions from one) | `06-generics/index` · "Busca la posición de un valor en un slice de cualquier tipo." |
| `02-constraints` "Restricciones" | `generic-demo` · `constraints` (each type tries to pass the constraint; `[]int` does not satisfy `comparable`) | `06-generics/sum` · "Suma un slice de números, sean enteros o decimales." |
| `03-generic-types` "Tipos genéricos" | — | `06-generics/stack` · "Construye una pila que guarde valores de cualquier tipo." |
| `04-review` "Repaso: tus propias colecciones" | — | `06-generics/set` · "Crea un conjunto genérico y las funciones Map y Filter para cualquier tipo." |

### S7 — Module 07-buffers "Buffers y E/S"

Summary: "Por qué escribir de golpe es más rápido que byte a byte: bytes.Buffer, bufio y io.Copy."

| Lesson | Animation (component · scene) | Exercise · goal |
|---|---|---|
| `01-buffers` "Qué es un buffer" | `buffer-demo` · `batch` (many small writes fill the buffer; one big write reaches the disk) | `07-buffers/ring` · "Guarda los últimos n valores en un buffer circular." |
| `02-bytes-buffer` "bytes.Buffer" | `buffer-demo` · `bytes` (`Write` appends at the end, `Read` advances from the start) | `07-buffers/table` · "Escribe una tabla de texto en un bytes.Buffer." |
| `03-bufio-writer` "bufio.Writer y Flush" | `buffer-demo` · `flush` (the bytes wait in the buffer; without `Flush` they never arrive) | `07-buffers/report` · "Escribe un informe con bufio.Writer sin olvidar el Flush." |
| `04-scanner` "Leer línea a línea con bufio.Scanner" | `buffer-demo` · `scan` (the scanner fills its buffer and cuts it at each `\n`) | `07-buffers/wc` · "Cuenta las líneas, las palabras y los bytes de un texto, como wc." |
| `05-io-copy` "io.Copy: de un Reader a un Writer" | `buffer-demo` · `copy` (chunks travel through one fixed buffer) | `07-buffers/counter` · "Cuenta cuántos bytes pasan por un Writer mientras copias con io.Copy." |
| `06-review` "Repaso: tu propio uniq" | — | `07-buffers/uniq` · "Lee líneas y quita las que se repiten seguidas, como uniq." |

### S8 — Module 08-concurrency "Procesos, hilos y goroutines"

Summary: "Cómo hace Go varias cosas a la vez: procesos, hilos, goroutines, el planificador, canales y Mutex."

| Lesson | Animation (component · scene) | Exercise · goal |
|---|---|---|
| `01-processes` "Procesos e hilos" | `sched-demo` · `threads` (each process has its own memory; its threads share it) | `08-concurrency/roundrobin` · "Reparte tareas entre varios hilos por turnos, como un planificador." |
| `02-goroutines` "Goroutines" | `sched-demo` · `goroutines` (when `main` returns the other goroutines vanish; a `WaitGroup` makes it wait) | `08-concurrency/squares` · "Calcula cuadrados en goroutines y espera a todas con un WaitGroup." |
| `03-scheduler` "El planificador G-M-P" | `sched-demo` · `gmp` (Gs queue on a P, an M runs them; an idle P steals half of another queue) | `08-concurrency/steal` · "Simula el robo de trabajo: un procesador sin tareas se lleva la mitad de la cola de otro." |
| `04-channels` "Canales" | `chan-demo` · `unbuffered` (sender and receiver wait for each other) | `08-concurrency/splitsum` · "Suma las dos mitades de un slice en dos goroutines y junta los resultados por un canal." |
| `05-buffered-channels` "Canales con buffer" | `chan-demo` · `buffered` (the sender only blocks when the buffer is full) | `08-concurrency/limit` · "Limita cuántas tareas se ejecutan a la vez con un canal con buffer." |
| `06-range-close` "range y close" | `chan-demo` · `close` (`range` ends when the channel is closed and empty) | `08-concurrency/generate` · "Envía la sucesión de Fibonacci por un canal y ciérralo al acabar." |
| `07-select` "select y default" | `chan-demo` · `select` (the first ready case wins; `default` when none is ready) | `08-concurrency/timeout` · "Espera un resultado, pero ríndete si tarda más de un tiempo máximo." |
| `08-mutex` "Mutex y data races" | `sched-demo` · `race` (two goroutines read, add and write the same counter and one update is lost; the `Mutex` lets only one in) | `08-concurrency/safecounter` · "Cuenta visitas desde muchas goroutines sin perder ninguna, con un Mutex." |
| `09-trees` "Práctica: árboles equivalentes" | — | `08-concurrency/trees` · "Recorre dos árboles binarios en goroutines y decide si guardan los mismos valores." |
| `10-review` "Repaso: el rastreador web" | — | `08-concurrency/crawler` · "Rastrea enlaces en paralelo sin visitar dos veces la misma página." |

### S9 — Module 09-web "Redes y la web" (the course ends here)

User: "Me gustaría añadir al plan acabar el curso con una pequeña api, un pequeño socket, y go templates (Explicando un poco TCP, HTTP, etc)"

Summary: "Cómo hablan dos programas por la red: TCP, un socket propio, HTTP por dentro, una pequeña API y plantillas HTML."

| Lesson | Animation (component · scene) | Exercise · goal |
|---|---|---|
| `01-tcp` "TCP: dos programas que hablan" | `net-demo` · `handshake` (IP and port, SYN / SYN-ACK / ACK, then a two-way byte stream) | `09-web/echo` · "Escribe un servidor de eco: devuelve cada línea que recibe." |
| `02-sockets` "Tu propio socket" | `net-demo` · `stream` (TCP sends bytes, not messages: one write can arrive in two pieces, so the protocol marks where each message ends) | `09-web/kv` · "Un pequeño servidor de clave-valor por TCP con las órdenes SET y GET." |
| `03-http` "HTTP: texto sobre TCP" | `net-demo` · `request` (request line, headers, blank line, body; the response comes back with its status code) | `09-web/request` · "Lee una petición HTTP a mano y saca el método, la ruta y las cabeceras." |
| `04-net-http` "Un servidor con net/http" | — | `09-web/greet` · "Escribe un handler que salude por el nombre que llega en la ruta." |
| `05-json-api` "Una pequeña API JSON" | — | `09-web/todos` · "Una API de tareas: lista, crea y borra tareas en JSON con los códigos de estado correctos." |
| `06-templates` "Plantillas con html/template" | `net-demo` · `template` (data plus template give HTML; `<` in the data is escaped) | `09-web/page` · "Genera una página HTML con una lista a partir de una plantilla." |
| `07-review` "Repaso: tu pequeña web" | — | `09-web/notes` · "Una web de notas: una API JSON y una página HTML que las muestra." |

- Standard library only: `net`, `bufio`, `net/http` (Go 1.22 patterns such as `GET /todos/{id}`), `encoding/json`, `html/template`, `net/http/httptest`.
- One goroutine per connection reuses `08-concurrency`; reading lines reuses `bufio.Scanner` from `07-buffers`. The `todos` and `notes` stores protect shared state with a `Mutex`.
- Tests never use fixed ports: TCP exercises listen on `127.0.0.1:0` or take a `net.Conn` (`net.Pipe`); HTTP exercises are tested with `httptest`. Tests pass under `-race`.
- Lessons show real terminal output (for example `nc localhost 8080` or `curl -i`), run on the machine. The site stays fully static: the servers live only in examples and exercises.

## Tasks

| ID | Specs | Route | Status | Commit |
|---|---|---|---|---|
| T1 | S1, S2, S4 | inline | done, approved by the user (L5) | `f2807db`, `311cfd1`, `1424039` |
| T2a | S2, S5, S6 (`06-generics`) | delegated, worktree | done | `ebc3dad`, `41e9605`, `a268ac0` |
| T2b | S1, S2, S7 (`07-buffers`) | delegated, worktree | done | `8d80f3b`, `303ef33`, `dd33c5a` |
| T2c | S1, S2, S5, S8 (`08-concurrency`) | delegated, worktree | done | `a2969ea`, `0d7759e`, `5007677` |
| T2d | S1, S2, S9 (`09-web`) | delegated, worktree | done | `ba7d208`, `77d5fb6`, `492b337` |
| T3 | S1 Part A | delegated, worktree | done | `0543ccf`, `3b14775`, `dbda53e`, `4f50a1d` |
| T4 | S2 (modules 00–04) | delegated, worktree | done | `059f7fc`, `75492eb`, `fe1578c`, `4ca9fd8` |

Route evidence: T1 is one module written in sequence; no writer reason (no parallel units launched together), so inline. T2a–T2d are four independent modules with disjoint edit surfaces, each hours of work, launched together in one turn: Writer trigger (a). They all add an entry to `content/catalog.go`, so each runs in its own worktree and the parent merges the four branches into `feat/low-level`. T3 and T4 are two independent units launched together in one turn (Writer trigger (a)); both edit lesson bodies of modules 01–04, so each runs in its own worktree and the parent merges them.

## Log

- L1 (2026-11-27) user: "Quiero que haya muchas más animaciones en las explicaciones y extiendas las lecciones a conceptos de bajo nivel, bffer, memoria, procesos planifica"
- L2 (2026-11-27) syllabus proposed (S1); user: "Luego los ejercicios, añade mas ejercicios y que el alumno escriba todo el código para mejorar la sintaxis y tenga mas soltura. pero sí , me cuadra"
- L3 Branch `feat/low-level` created from `main` at 4508718. Pilot T1 started.
- L4 (2026-11-27) T1 done. Commits: `f2807db` (animations `memory-demo` with scenes `bytes`, `padding` and `utf8`; `stack-demo` with `calls` and `escape`; `gc-demo` with `mark-sweep`), `311cfd1` (7 write-it-all exercises), `1424039` (catalog, 7 lesson decks, 7 compiled examples). The compiler and runtime messages shown in the lessons were checked by running them: `int8` overflow, stack overflow, `moved to heap: p` with `-gcflags='-m -l'`, and `cannot assign to s[0]`.
  - RED/GREEN: content behaviour has no unit-level RED. `TestSlideBudget` and `TestEveryLessonIsSlides` are the content gate. Each stub fails with `undefined: <Name>` (for example `05-memory/bits/bits_test.go:25:14: undefined: BitsNeeded`) and passes with `-tags solution`.
  - Checks: `gofmt -l` empty, `go vet ./...`, `go test ./...`, `cd exercises && go test -tags solution ./05-memory/...`, `node --check` on every animation, no hex colours in JS.
  - Headless check (CDP driver in the session scratchpad): 0 vertical overflows at 1280×720, 1024×768 and 390×844. The walk stepped 35 animation steps across 6 scenes with the expected captions. The only console error was a missing `favicon.ico`.
  - The exact `int8` overflow message scrolls sideways inside its `pre` at 1280, as an existing lesson's message already does in `01-basics/08-constants`.
  - Next: the user reviews the pilot; corrections go into the specs before T2.
- L5 (2026-11-27) user, on module order: "sí, déjalo en 5 y sigue con la fase 2, pero ten en cuenta el tour original de go, que no se nos olvide nada". Read as approval of the pilot. Tour audit in S5: modules 00–04 are complete; generics and four concurrency pages were missing. Added `06-generics` (S6); buffers and concurrency renumbered to `07` and `08` (S7, S8). T2 split into T2a–T2c, launched in parallel worktrees.
- L6 (2026-11-27) First launch of T2a–T2c stopped before writing: the worktrees were created from `main` (4508718), not `feat/low-level`. They were relaunched with an instruction to reset the clean worktree to `6417ebb` (this plan commit) first.
- L7 (2026-11-27) user: "Me gustaría añadir al plan acabar el curso con una pequeña api, un pequeño socket, y go templates (Explicando un poco TCP, HTTP, etc)". Added `09-web` as the final module (S9) and task T2d, launched with T2a–T2c.
- L8 (2026-11-27) T2a returned: `ebc3dad` (generic-demo), `0e4487d` (exercises), `5294b98` (content); all checks green, headless walk 0 overflows at 1280×720 and 390×844. Its diff stays inside the declared surface. Brief defect found: it asked for Spanish test failure messages, but CLAUDE.md puts tests in English and every exercise uses `Func(args) = got, want want`. All four writers were told to follow the repo form; T2a amends its exercise commit.
- L9 (2026-11-27) T2a amended: `ebc3dad`, `41e9605`, `a268ac0` (English test names and messages). T2b returned: `8d80f3b` (exercises), `303ef33` (buffer-demo), `dd33c5a` (content); diff inside its surface, test messages already in English, all checks green, 46/46 cards fit at 1280×720. Note for the seam check: plain `go vet ./...` in `exercises` always fails on the empty stubs by design; the gate there is `go vet -tags solution ./...`. T2b reports a 9 px sideways scroll at 390 px that also shows on existing lessons (pre-existing, not from 07).
- L10 (2026-11-27) T2d returned: `ba7d208` (net-demo), `77d5fb6` (exercises), `492b337` (content); diff inside its surface, English test messages, `-race -count=3` green, headless 0 overflows at 1280×720, 1024×768 and 390×844. Every `nc`/`curl` session shown was run on the machine. Merge follow-up for the parent: `content/04-methods/11-images.html` still ends with "Fin del paseo… Con esta lección termina la ruta"; now that `09-web/07-review` closes the course, that card must go. T2d's check script ran `pkill -f "user-data-dir=/tmp/cdp-"`, which may have killed T2c's headless browser; T2c's browser evidence must be re-checked.
- L11 (2026-11-27) T2c returned: `a2969ea` (sched-demo, chan-demo), `0d7759e` (exercises), `5007677` (content); diff inside its surface; its headless walk finished (0 overflows at 1280×720, 1024×768 and 390×844 after moving 7 outputs to their own cards), so the `pkill` in L10 did not cut it short. Every Go tour Concurrency page maps to a lesson (02, 04–10).
  - Merged into `feat/low-level` in order (`196310a`, the 08 merge, `779f054`); each `content/catalog.go` conflict resolved to modules 00–09 in order. `fix(content): move the course ending to 09-web` removes the "Fin del paseo" card from `04-methods/11-images.html`.
  - Seam check on the merged branch: `gofmt -l` empty (root and exercises), `go vet ./...` ok, `go test -count=1 ./...` ok (content, internal/course, internal/site), `cd exercises && go vet -tags solution ./...` ok, `go test -count=1 -tags solution ./...` 62 packages ok and 0 failing, `go test -race -tags solution ./08-concurrency/... ./09-web/...` ok, `node --check` on every animation ok.
  - Pending: a headless walk of the merged site was not run by the parent; each module's own walk passed on its branch. The 9 px sideways scroll at 390 px (L9) is pre-existing and not fixed here.
  - Next: T3 (Part A animations), then T4.
- L12 (2026-11-27) user: "sí, lanza T3 y T4 en paralelo, o lo hago yo, dile que puede o bien por worktree o con cuidado, por los commits". T3 and T4 launched in parallel, each in its own worktree reset to this plan commit, because both touch `content/0[1-4]-*/*.html`.
- L13 (2026-11-27) T4 returned: `059f7fc`, `75492eb`, `fe1578c`, `4ca9fd8`; diff inside its surface (106 files, +1667/−408), no AI attribution. 28 exercises in 00–04 converted to write-it-all stubs (`hello` included, with a hint), 12 new exercises for the lessons that had none, every "Tu turno" card in 00–04 shows the signatures in a Go block. RED: each new test failed first with `undefined: <Name>`; 94 `undefined` lines across the 00–04 stubs. Merged as `c95be80`; parent spot check `cd exercises && go test -count=1 -tags solution ./...`: 74 ok. Pending: no browser check that the new Tu turno cards fit (longest: 03-interfaces with 15 code lines); it goes into the headless walk after T3.
- L14 (2026-11-27) T3 returned: `0543ccf`, `3b14775`, `dbda53e`, `4f50a1d`; diff inside its surface (17 files, +1056/−30), no AI attribution. 11 new scenes: pointer-demo `variable`, `convert`, `closure`; memory-demo `int8`, `struct`; slice-demo `array-copy`, `range-copy`; new flow-demo `for`, `switch`; new map-demo `buckets`; interface-demo `type-switch`. Value vs pointer receiver was already the pointer-demo `receivers` scene in `04-methods/02-pointer-receivers`, so it was not duplicated. interface-demo was refactored to pick its drawing by scene; its existing `pair` scene was not walked by T3. Merged as `6568151` with no conflicts against T4.
  - Seam check on the merged branch: `gofmt -l` empty (root and exercises), `go vet ./...` ok, `go test -count=1 ./...` ok, `cd exercises && go vet -tags solution ./...` ok, `go test -count=1 -tags solution ./...` 74 ok and 0 failing, `go test -race -tags solution ./08-concurrency/... ./09-web/...` ok, `node --check` on every animation ok.
  - Known cosmetic point from T3: in `array-copy` the copied numbers overlap the originals for about 100 ms at the start of their slide.
- L15 (2026-11-27) Headless walk of the merged site (all 74 lessons, every card at 1280×720 and 390×844, every animation stepped forwards to its end and back at 1280): 1786 cards, 44 scenes, 238 steps; each scene went back as many steps as it went forward; no page errors besides `favicon.ico`. One real overflow: the T4 "Tu turno" card of `04-methods/03-interfaces` (119 px at 1280, 48 px at 390), split into "Tu turno" and "Lo que escribes" as in `05-memory/07-review` (`fix(content): split the interfaces exercise card so it fits on screen`); re-walked: no overflow. The remaining finding is the pre-existing 9 px sideways scroll at 390 px on the title card of `01-basics/09-review`, `03-more-types/10-function-values`, `04-methods/03-interfaces` and `09-web/04-net-http` (L9), not fixed. All tasks T1–T4 are done.
  - Next: the user reviews phases 3 and 4; push and PR are the user's decision (delivery strategy `ask-on-risk`, the branch is well above 400 lines).
