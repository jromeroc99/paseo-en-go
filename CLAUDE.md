# Paseo en Go

A modern, beginner-friendly tour of Go. The website explains each concept
(with animations for the hard ones); students write code **locally** in
`exercises/` and validate it with `go test`. No database, no accounts:
progress is a manual checkbox stored in the browser's `localStorage`.

## The golden rule: code that explains itself

Code must be so simple that it explains itself just by reading it, as if a
senior engineer wrote it for a junior to maintain.

- Prefer the obvious solution over the clever one. If it needs a comment to be
  understood, first try a better name or a simpler shape.
- Small functions with one job. Descriptive names. Early returns, shallow nesting.
- Comments explain *why*, never *what*. Every package and exported identifier
  has a doc comment.
- No abstraction without a second real use: no interfaces "just in case", no
  generic helpers for one caller, no layers that only forward calls.
- This repo is also teaching material: it should read like idiomatic Go.

## Architecture

One rule from Clean Architecture: **the domain knows nothing about HTTP**.
Dependencies point inward: `cmd` → `site` → `course`.

```
cmd/paseo/          CLI wiring only: `serve` (dev) and `build` (static site)
internal/course/    domain: Course, Module, Lesson, Exercise + Validate (no I/O)
internal/site/      renders pages once; serves them or writes them to disk
content/            the course itself, embedded
  catalog.go        course structure, declared in Go
  <module>/<lesson>.html          lesson body (Go template)
  <module>/examples/<name>/main.go runnable snippets, compiled by `go vet ./...`
web/                embedded templates/ and static/ (css, js)
exercises/          separate Go module: the student's workspace
```

- Standard library only. Adding a dependency requires an explicit decision.
- The site is fully static: `paseo build` output can be hosted on any static
  host (GitHub Pages, Netlify). Never add server-side state.

## Commands

```bash
go run ./cmd/paseo serve                 # http://localhost:8080
go run ./cmd/paseo build -out dist       # static site; use -base /repo/ for subpaths
go test ./...                            # site, domain and content checks
go vet ./...
cd exercises && go test -tags solution ./...   # reference solutions must pass
```

## Content

- Lesson text is in **neutral Spanish** (no regional slang). Code, identifiers,
  comments, tests and commit messages are in **English**.
- Write for someone who is learning to program in Go: explain the problem
  first, then the concept, then the example. One idea per lesson.
- Embed snippets with `{{code "examples/<name>/main.go"}}` instead of pasting
  code into HTML, so every example is compiled and never goes stale.
- Adding a lesson: add it to `content/catalog.go`, create its `.html` body,
  run `go test ./...` (it fails if a body is missing or a slug is invalid).

## Animations

- Pixel-art style: simple, visual and pretty. A small palette, crisp pixels
  (`image-rendering: pixelated`), chunky shapes, short and clear motion.
- Use them only where a picture explains better than text: memory and
  pointers, slices and backing arrays, interfaces, readers, channels.
- Each animation is a vanilla JS Web Component in `web/static/js/animations/`,
  one per file, used as a tag inside the lesson body (e.g. `<pointer-demo>`).
- Always readable without the animation (alt text or a static fallback), and
  respect `prefers-reduced-motion`.

## Exercises

Each exercise lives in `exercises/<module>/<name>/` with three files:

- `<name>.go` — the stub the student completes, tagged `//go:build !solution`,
  with a `TODO` and a doc comment that states the goal.
- `<name>_solution.go` — the reference solution, tagged `//go:build solution`.
- `<name>_test.go` — table-driven tests with clear failure messages; they are
  the student's only feedback, so make them friendly.

Exercises reuse concepts from earlier modules; each module ends with one
exercise that combines everything seen so far.

## Testing

- Table-driven tests with `t.Run`; `httptest` for handlers; `t.TempDir()` for files.
- Test behavior through the public API, not implementation details.
- `gofmt`, `go vet` and `go test ./...` must be clean before committing.

## Git

- Conventional Commits (`feat:`, `fix:`, `docs:`, `test:`, `refactor:`, `chore:`).
- No AI attribution or `Co-Authored-By` lines in commits.
