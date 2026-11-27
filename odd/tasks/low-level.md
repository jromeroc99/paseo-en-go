# Feature: low-level modules, more animations and write-it-all exercises

Objective: three new modules on how Go works underneath (memory, buffers, processes and the scheduler), many more animations across the course, and exercises where the student writes all the code.
Branch: `feat/low-level`. Delivery strategy: `ask-on-risk` (forecast ≈ 9000 authored lines; asked only when the user wants to publish).
Test runner: `go test ./...`, `go vet ./...` and `cd exercises && go test -tags solution ./...`. Repo rules: `CLAUDE.md` (read it first; it is not repeated here).

## Specs

### S1 — Scope

User: "Quiero que haya muchas más animaciones en las explicaciones y extiendas las lecciones a conceptos de bajo nivel, bffer, memoria, procesos planifica".
Approved syllabus (user: "me cuadra"):

- Part A — about 12 new animations in existing lessons, only where something invisible is shown: `01-basics` (a variable as a box with an address, `int8` overflow, `float → int` losing decimals), `02-flow-control` (`for` step by step, `switch` top-down evaluation), `03-more-types` (contiguous struct fields, copying an array copies everything, `range` gives a copy, map buckets, the variable a closure captures), `04-methods` (value vs pointer receiver, type switch).
- Part B — three new modules, every lesson with its animation:
  - `05-memory` "La memoria por dentro": bytes and addresses, the stack and call frames, the heap and escape analysis (`-gcflags=-m`), the garbage collector (mark & sweep), struct padding (`unsafe.Sizeof`), strings, bytes and runes (UTF-8).
  - `06-buffers` "Buffers y E/S": what a buffer is, `bytes.Buffer`, `bufio.Writer` and `Flush` (forgetting it is the typical mistake), `bufio.Scanner`, `io.Copy`.
  - `07-concurrency` "Procesos, hilos y goroutines": process vs thread, goroutines, the G-M-P scheduler, unbuffered and buffered channels, `select`, `WaitGroup`, `Mutex` and data races (`-race`).

### S2 — Exercises: the student writes all the code

User: "Luego los ejercicios, añade mas ejercicios y que el alumno escriba todo el código para mejorar la sintaxis y tenga mas soltura."

- New exercises (modules 05–07): the stub `<name>.go` holds only the build tag, the package clause and a doc comment that states what to write, signatures included. No function skeletons, no bodies. Until the student writes them, `go test` fails with `undefined: …`, which the lesson's "Tu turno" card explains.
- The "Tu turno" card shows the exact signatures to write, as a Go code block.
- Existing exercises (modules 00–04) are converted to the same style in phase 4.
- "Más ejercicios": every lesson of modules 05–07 has its own exercise; in phase 4 the existing lessons without one get one. Assumption: one exercise per lesson (the domain allows one), not several.

### S3 — Phases

| Phase | Work |
|---|---|
| 1 | Pilot: module `05-memory` complete; the user reviews it |
| 2 | Modules `06-buffers` and `07-concurrency` |
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

## Tasks

| ID | Specs | Route | Status | Commit |
|---|---|---|---|---|
| T1 | S1, S2, S4 | inline | done, waiting for the user's review | `f2807db`, `311cfd1`, `1424039` |
| T2 | S1, S2 (06-buffers, 07-concurrency) | — | pending | |
| T3 | S1 Part A | — | pending | |
| T4 | S2 (modules 00–04) | — | pending | |

Route evidence: T1 is one module written in sequence; no writer reason (no parallel units launched together), so inline.

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
