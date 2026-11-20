# Feature: lessons as slides

Objective: every lesson is a deck of cards that fit the screen with no scroll; arrow keys or the on-screen arrows move through them, and an animation advances one step per press, its code panel side by side with the canvas.
Branch: `feat/lesson-slides`, from `feat/tour-modules`. Delivery strategy: `ask-on-risk` (forecast ≈ 2000 authored lines, most of it lesson markup).
Test runner: `go test ./...`, plus a headless fit check (S7). Rules of the repo: `CLAUDE.md` (read it first; it is not repeated here).

## Specs

### S1 — Scope

User request: "quiero que el frontend sean cards rollo diapositivas, las lecciones, sin scrolls, dandole a la fecha, las animaciones secuencia de código en paralelo a la animación, siguiendo los clicks de la diapositiva, haz el plan".

Read as ("fecha" taken as "flecha"):

- "cards rollo diapositivas, las lecciones": each lesson page becomes a deck; one card on screen at a time.
- "sin scrolls": a card always fits the viewport; the page itself never scrolls in deck mode.
- "dandole a la flecha": → and ← move between cards (keyboard and on-screen buttons).
- "secuencia de código en paralelo a la animación": the animation's code panel sits beside the canvas, not under it.
- "siguiendo los clicks de la diapositiva": an animation no longer loops on a timer; each → press plays the next step.

Out of scope: the index page (it stays the scrolling *walk*, assumption A1), lesson prose rewrites beyond what a card needs to fit (S6), new animations, server-side state.

### S2 — The deck

- Lesson body markup: one `<section class="slide">` per card, authored in `content/<module>/<lesson>.html` (decision D1). `lesson.html` adds two cards of its own: a title card first (module breadcrumb + `<h1>`) and a closing card last (exercise box, "Marcar como completada", previous/next lesson links).
- One new script, `web/static/js/deck.js`, vanilla, no libraries. It turns `.deck` into deck mode by adding a class on `<html>`; without JS the sections stack and the page scrolls as today (progressive enhancement).
- Card size: the card fills the space between the site header and the deck bar (`100dvh`), max width about 960 px, centred, `--radius`, 1 px `--line` hairline, `--surface` background. Text keeps its normal size; nothing is scaled to fit.
- Deck bar under the card: ‹ and › buttons, and a mini route with one dot per card in the index's style; the current dot glows (`--glow`), visited ones use `--accent`. It also reads "3 / 9" for screen readers (`aria-label`).
- Position in the URL: `#3` opens card 3; `history.replaceState`, so the browser's back button still leaves the lesson instead of walking every card.

### S3 — Navigation

- Next: →, Space, PageDown, the › button. Previous: ←, PageUp, the ‹ button. Home and End jump to the first and last card. Touch: a horizontal swipe.
- Keys with a modifier (Alt+← is the browser's back) and keys typed into a form control are ignored.
- → on the last card opens the next lesson at its first card; ← on the first card opens the previous lesson at its last card (assumption A2).
- Accessibility: hidden cards are `inert`, so Tab never reaches them; after a move, focus goes to the new card (`tabindex="-1"`, `aria-roledescription="diapositiva"`), so a screen reader reads it. With `prefers-reduced-motion` cards switch with no slide transition; otherwise a short fade/slide (about 200 ms).

### S4 — Animations follow the deck

Replaces the playback rules of S13 in `tour-modules.md`; the look stays.

- No timer, no loop, no fade-out, no pause on hover or focus. Each → on an animation card plays the next step (the step's ease-in, about 450 ms); ← shows the previous step at once. Only after the last step does → move to the next card.
- Entering the card forwards starts at step 1; coming back from the next card shows the last step, so going backwards reads naturally.
- `Scene` exposes two methods the deck calls: `next()` and `prev()`, each returning `false` when there is no step left in that direction. `deck.js` asks the current card's scene first, then moves the card. Clicking the canvas also calls `next()`.
- Layout: a grid with the code panel on the left and the canvas on the right, the caption under both. Code lines already run show normal, the current line is highlighted (`--accent` border, as now), lines still to come are dimmed (`--muted` at lower opacity), so the reader sees how many steps remain. Under about 720 px wide the code panel goes above the canvas.
- The caption is `aria-live="polite"` again: it changes only when the reader presses a key, so it no longer repeats forever.
- `IntersectionObserver`, `visibilitychange` and the hold logic leave `scene.js`. Components (`pointer-demo`, `slice-demo`, `interface-demo`, `reader-demo`, `defer-demo`) keep their scene data and `draw`; they change only if the narrower canvas breaks their layout.
- An animation card holds the animation and at most a `<h2>`; nothing else (it needs the whole card).

### S5 — Content rules for cards

- Typical cards of a lesson, in order: the problem; the concept (one or two cards); the animation (if any); the example code; its output and the "fíjate" note; "Un error a propósito"; "El error típico"; "Tu turno". The lesson structure of `CLAUDE.md` stays the same, only cut into cards.
- Budget per card (testable proxy for "fits"): at most about 90 words of prose, and at most one code block; a program and its output count as one (D2). An example over about 20 lines gets a card of its own with no prose besides its heading.
- Split, do not rewrite: existing sentences move into cards unchanged; a sentence is shortened only when a card does not fit at the sizes of S7. Spanish rules of `CLAUDE.md` apply to any changed word.

### S6 — Tests (Go, standard library only)

- `TestEveryLessonIsSlides` (in `content/`): every lesson body is made only of top-level `<section class="slide">` elements (parsed with `encoding/xml` in non-strict mode with `xml.HTMLEntity` and `AutoClose`, no new dependency).
- `TestSlideBudget`: every card stays within the budget of S5, and an animation card holds only the animation and an optional `<h2>`. Failure messages name the lesson and the card number.
- `internal/site` tests: the lesson page carries the title card, the closing card and `deck.js`; the build test pins `static/js/deck.js`; the base-path test pins its prefixed `src`.
- `TestEveryAnimationHasComponent` and `TestNoAnimationMarkerLeft` stay as they are.

### S7 — Checks

- `gofmt`, `go vet ./...`, `go test ./...`, `cd exercises && go test -tags solution ./...`, `node --check` on every JS file, no hex colours in JS.
- Headless fit check over every card of every lesson at 1280×720, 1024×768 and 390×844 (DPR 3): the card's `scrollHeight` is not greater than its `clientHeight` and the page does not scroll. Desktop and tablet sizes are strict. On the phone a card that still overflows may scroll inside itself (assumption A3), and the check lists those cards so they can be split.
- Headless walk: → through a lesson with an animation reaches every step in order, the highlighted code line matches the step, ← goes back, `#n` deep links work, → on the last card opens the next lesson, no console errors.

### S8 — Docs

- `CLAUDE.md`: "Visual style" (a lesson is a deck of cards), "Content" (lesson structure cut into cards, the budget of S5, how to add a card), "Animations" (stepped by the deck, code beside the canvas; drop the loop and the WCAG 2.2.2 pause, which no longer applies because nothing moves on its own).

### Decisions and assumptions

- D1 — Authored `<section>` cards rather than splitting at every `<h2>` in JS. Automatic splitting changes no content, but it cannot promise "sin scrolls": "El concepto" in `01-pointers` holds three paragraphs, a list and the animation, and the longest examples are 30–32 lines. Authored cards let the author put the animation alone and keep each card in budget, and a Go test can check it.
- A1 — The index stays a scrolling page; only lessons become decks.
- A2 — The deck continues across lessons at both ends.
- A3 — "Sin scrolls" is strict from 1024×768 up; on a phone, a card that cannot fit scrolls inside itself as a fallback, never the page.

Risk: item 6 (no automated test catches a layout overflow or a broken step order), covered by the headless fit check and walk of S7. Not item 3: lesson URLs stay the same, `#n` is only added.

## Tasks

- [x] T1 — Deck: `deck.js`, `lesson.html` title and closing cards, CSS tokens and classes, S3 navigation, site tests (S6), `CLAUDE.md` "Visual style" and "Content". Pilot content: `00-setup/01-welcome` and `01-basics/01-packages` cut into cards. Route: inline (the base must exist before any lesson is converted). Recommended model: Opus 5.5, effort high. Commit `73c4a2f`; RED/GREEN in L4.
- [x] T2 — Stepped animations (S4): `scene.js` `next`/`prev`, the side-by-side layout, `CLAUDE.md` "Animations"; `03-more-types/01-pointers` cut into cards as the animation pilot. Route: inline. Recommended model: Opus 5.5, effort high. Then stop for the user's review of the three pilot lessons; corrections go to `## Log` and into S2–S5. Commit `1c244f4`; checks in L5.
- [ ] T3 — Content tests `TestEveryLessonIsSlides` and `TestSlideBudget` (RED on every unconverted lesson), then convert the rest: `01-basics`, `02-flow-control`, `03-more-types`, `04-methods`. Route: delegated, four parallel writers (Writer trigger (a): four independent units with disjoint surfaces `content/<module>/*.html`, each well over 100 lines); the parent owns the tests, `CLAUDE.md` and git, and runs the seam check. Recommended model: Opus 5.5, effort medium.
- [ ] T4 — Fit check and walk of S7 over all 40 lessons; split the cards it reports; adapt any demo whose layout breaks on the narrower canvas. Route: inline. Recommended model: Opus 5.5, effort high.

Commits: T1 and T2 one each (`feat(web): …`); T3 one per module (`refactor(content): cut <module> lessons into slides`); T4 one (`fix(content): …` or `fix(web): …`).

## Log

- L1 (2026-11-20, user, original request): "quiero que el frontend sean cards rollo diapositivas, las lecciones, sin scrolls, dandole a la fecha, las animaciones secuencia de código en paralelo a la animación, siguiendo los clicks de la diapositiva, haz el plan"
- L2 (2026-11-20): plan written (S1–S8, T1–T4). Explored: `lesson.html`, `layout.html`, `scene.js` (loop, hold, code panel under the canvas), `style.css` `.anim*`, `content/catalog_test.go`. Facts behind D1: 40 lessons with 166 `<h2>` sections, 8 lessons with an animation, longest examples 30–32 lines. Next: user approval, then T1.
- L3 (2026-11-20, user): "Apruebo el plan, comitea, limpio y arranca". Plan committed on `feat/lesson-slides` with no attribution lines; T1 starts.
- L4 (2026-11-20): T1 done in `73c4a2f`. RED: `TestPages` (deck, title and closing cards, `deck.js` src), `TestBasePathPrefixesLinks` and `TestBuildWritesEveryPage` failed on the missing markup and `static/js/deck.js`; GREEN after the template, layout and script. Checks: `gofmt` clean, `go vet ./...`, `go test ./...`, exercises with `-tags solution`, `node --check deck.js`. Headless Chrome (CDP driver in the session scratchpad, kept for T4): both pilots fit at 1280×720, 1024×768 and 390×844 with no page scroll and no clipped `pre`; walk verified deep link `#2`, ←/→ with focus on the new card, ← on the first card opens the previous lesson at `#last`, → on the last opens the next lesson; no console errors.
  - D2: a program and its output may share a card (S5 updated).
  - D3: an error gets two cards, the code with its message, then the message again beside how to read or fix it; packages prose moved unchanged, only `<h2>` headings added.
  - D4: `#last` in the URL opens a lesson at its final card; ← on a first card follows `a[rel=prev]` with it.
  - Fixes found by the checks: the incoming card turns visible at once (a `visibility` transition starts hidden, so `focus()` was dropped); cards use `gap` with zero child margins and `flex-shrink: 0` (flexbox does not collapse margins and squashed `pre` into a scrolling sliver); the first fit check missed clipped `pre`, so it now measures them too.
  - Next: T2.
- L5 (2026-11-20): T2 done in `1c244f4`. `Scene` has `next()`/`prev()` (false at the ends), `rewind()` and `skipToEnd()`; the loop, hold, fade, `IntersectionObserver` and `visibilitychange` are gone (only the `ResizeObserver` stays). `deck.js` asks the card's scene first and rewinds or skips to the end on entering a card. `.anim` is a grid, code left and canvas right, stacked under 45rem. `01-pointers` cut into 10 cards; the animation card holds only `<h2>En la memoria</h2>` and the demo. `CLAUDE.md` "Animations" rewritten (Motion, Layout, Captions, Accessibility). No RED test: the stepping is browser behaviour with no Go-side observable, so the headless walk is the proof (S7).
  - Checks: `gofmt` clean, `go vet ./...`, `go test ./...`, exercises with `-tags solution`, `node --check` on every JS file, no hex colours in JS. Fit: the three pilots pass at 1280×720, 1024×768 and 390×844. Walk on `01-pointers` card 5: → plays steps 0–6 with the highlighted line equal to the step and the dimmed lines counting down 6→0; the next → moves to card 6; ← from card 6 shows the last step; ← steps back; a click on the canvas plays the next step; caption `aria-live="polite"`; no console errors.
  - D5: lines still to come are those below the furthest line any shown step reached, because `defer-demo` and `reader-demo` jump back up in a loop.
  - Next: the user's review of the three pilots; then T3.
- L6 (2026-11-20, user): "Esta genial, limpio contexto, preparame el prompt para seguir hasta el final de lo que llevamos, que modelo me recomiedas?". The three pilots are accepted with no corrections; S2–S5 stay as they are. Next: T3, then T4, in a fresh session.
