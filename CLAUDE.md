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

- Code, identifiers, comments, tests and commit messages are in **English**.
- Everything the student reads (lessons, UI text, exercise goals) is in
  **neutral Spanish from Spain**:
  - Address the reader as `tú`: "abre", "ejecuta", "fíjate". Never voseo
    ("abrí", "corré") and never `usted`.
  - Spain vocabulary: "ordenador" (not "computadora"), "archivo", "carpeta",
    "ejecutar un programa/test" (not "correr").
  - No regional slang or idioms. Clear, warm and direct, short sentences.
  - Keep Go terms in English when that is what the reader will see in the
    docs and errors (`slice`, `map`, `struct`, `goroutine`), introduced in
    *cursiva* the first time with a one-line Spanish explanation.
- Write for someone who is learning to program in Go: explain the problem
  first, then the concept, then the example. One idea per lesson.
- Lesson structure: a short motivating problem, the concept, a compiled
  example, the common mistake to avoid, and the exercise (if any).
- A lesson body is a deck: a list of `<section class="slide">` cards, one
  idea each. The template adds the title card and the closing card (the
  exercise, the "done" checkbox and the pager), so never write them.
- One card must fit a 1280×720 screen without scrolling: at most about 90
  words of prose and one code block. A program and its output may share a
  card. If a card does not fit, split it; never shrink the text.
- Cut a lesson into cards in this order: the problem (no heading), the
  concept, the animation (alone, with at most a `<h2>`), the example, its
  output and the "fíjate" note, "Un error a propósito" (the code and its
  message), "Cómo leer el error" (the message again and how to read it),
  "El error típico" and "Tu turno". A card with a heading holds at most
  17 lines of code; a longer program is read across cards with
  `{{codeLines "examples/<name>/main.go" first last}}`, cut at a blank line
  or between functions. A program and its output share a card only while
  both fit in 13 lines; otherwise the output goes to the next card. An
  error program that does not fit is cut the same way, and its message
  goes with the second half.
- `TestSlideBudget` checks the budget: at most 90 words outside headings,
  code and animations, one Go program and one output block per card, at
  most 17 lines per snippet and 13 for a snippet with its output. It is a
  proxy; the real proof is that the card fits on screen.
- Show errors on purpose: when a concept has a typical compiler error, show
  the exact message and explain how to read it.
- Embed snippets with `{{code "examples/<name>/main.go"}}` instead of pasting
  code into HTML, so every example is compiled and never goes stale.
- Adding a lesson: add it to `content/catalog.go`, create its `.html` body,
  run `go test ./...` (it fails if a body is missing or a slug is invalid).

## Visual style

- **Dark theme only.** Deep blue-ink background, light text, one neon accent:
  Gopher Blue (`--accent`). All colours are CSS tokens on `:root` in
  `web/static/css/style.css`; never hard-code a colour anywhere else.
- **Modern and minimalist.** Rounded boxes (`--radius`, `--radius-sm`), 1 px
  hairlines, generous space. Every code block gets a copy button (`copy.js`).
  The neon glow (`--glow`) is used sparingly: the course route, focus, active
  state and the exercise box. If everything glows, nothing does.
- The index is a *walk*: a neon route with one stop per lesson that lights up
  when the lesson is done. Animations share this same look.
- A lesson is a deck of cards, one on screen at a time and no page scroll.
  `deck.js` moves through it with the arrow keys, Space, the bar under the
  card or a swipe; → on the last card opens the next lesson. The bar repeats
  the walk: a mini route with one stop per card. Without JS the cards stack
  and the page scrolls like an article.
- Body text in a readable system font; code in monospace.
- Accessible contrast (WCAG AA), visible keyboard focus, works on a phone.

## Animations

Simple, visual and pretty, in the site's own style. Their only job is to make
an invisible concept visible (memory, pointers, backing arrays, interface
values, buffers, channels). If text explains it just as well, do not animate it.

- **Look:** the page's style, not pixel art. Rounded boxes (`--radius-sm`),
  1 CSS px hairlines, monospace text, the glow only on what changes in the
  step. Colours come only from the `--anim-*` tokens in `style.css`; each one
  has one meaning across the whole course (a value, a pointer, a highlight,
  an error).
- **Drawing:** a `<canvas>` with 160×96 logical units, stretched to the frame's
  width (max 640 px) and drawn at `devicePixelRatio`, so it stays sharp.
  Components draw through `Paint` (`rect`, `dot`, `frame`, `line`, `arrow`,
  `text`, `centerText`) and never touch the canvas directly.
- **Motion:** the reader sets the pace. Nothing moves on its own: each → in
  the deck (or a click on the canvas) plays the next step with a short ease,
  ← shows the previous one at once, and only after the last step does → move
  to the next card. Arriving forwards starts at the first step; coming back
  shows the last. One thing moves at a time, so the eye knows where to look.
- **Layout:** the code on the left, the canvas on the right and the caption
  under both; on a narrow screen the code goes above. An animation card holds
  the animation and at most a `<h2>`.
- **Captions:** each step shows a one-line Spanish caption and highlights the
  line of Go code it represents. Lines already run read normally; lines still
  to come are dimmed, so the reader sees how much is left.
- **Code:** one vanilla JS Web Component per file in
  `web/static/js/animations/<name>.js`, extending `Scene` from `scene.js` and
  used as a tag in the lesson body (`<pointer-demo>`). No libraries. The scene
  is plain data (a list of steps); `draw` just renders the current step.
- **Accessibility:** the caption is `aria-live="polite"`, so a screen reader
  announces each step the reader asks for. Respect `prefers-reduced-motion`
  (no tweening). The canvas is an image labelled with the text fallback inside
  the tag, so the lesson still makes sense without JS.

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
