// Turns a lesson into a deck of cards: one card on screen at a time, moved
// with the arrow keys, the buttons under the card or a swipe. Without this
// script the cards simply stack and the page scrolls, so nothing is lost.

const NEXT_KEYS = new Set(["ArrowRight", " ", "PageDown"]);
const PREV_KEYS = new Set(["ArrowLeft", "PageUp"]);
const SWIPE_PX = 50;
// Walking forwards opens the next lesson on its first card; walking back
// from it lands on this lesson's final card.
const FIRST_HASH = "#1";
const LAST_HASH = "#last";

// Without .is-deck the reader chose the article view: the cards stay stacked.
const deck = document.querySelector(".deck");
if (deck && document.documentElement.classList.contains("is-deck")) start(deck);

function start(deck) {
  const slides = [...deck.querySelectorAll(":scope > .slide")];
  const bar = buildBar(slides.length);
  const visited = new Set();
  let current = -1;

  function show(index, { focus = true } = {}) {
    // Arriving forwards plays the card's animation from its start; coming
    // back shows where it ended, so reading backwards makes sense.
    if (index !== current) {
      const scene = sceneOf(slides[index]);
      if (index > current) scene?.rewind();
      else scene?.skipToEnd();
    }
    current = index;
    visited.add(index);
    slides.forEach((slide, i) => {
      slide.dataset.pos = i < index ? "before" : i > index ? "after" : "current";
      slide.inert = i !== index;
    });
    bar.update(index, visited, { hasPrev: index > 0 || hasLink("prev"), hasNext: index < slides.length - 1 || hasLink("next") });
    // replaceState, not pushState: the back button leaves the lesson
    // instead of walking back through every card.
    history.replaceState(null, "", `#${index + 1}`);
    if (focus) slides[index].focus({ preventScroll: true });
  }

  // next and prev let the card's animation take the step first; the card
  // moves only when the animation has no step left in that direction.
  function next() {
    if (sceneOf(slides[current])?.next()) return;
    if (current < slides.length - 1) show(current + 1);
    else follow("next", FIRST_HASH);
  }

  function prev() {
    if (sceneOf(slides[current])?.prev()) return;
    if (current > 0) show(current - 1);
    else follow("prev", LAST_HASH);
  }

  function hasLink(rel) {
    return deck.querySelector(`a[rel="${rel}"]`) !== null;
  }

  function follow(rel, hash = "") {
    const link = deck.querySelector(`a[rel="${rel}"]`);
    if (link) location.href = link.href + hash;
  }

  function indexFromHash() {
    if (location.hash === LAST_HASH) return slides.length - 1;
    const n = Number.parseInt(location.hash.slice(1), 10);
    if (!Number.isInteger(n)) return 0;
    return Math.min(Math.max(n, 1), slides.length) - 1;
  }

  function onKey(event) {
    if (event.altKey || event.ctrlKey || event.metaKey || event.shiftKey) return;
    if (event.target.closest("input, textarea, select, [contenteditable]")) return;
    // Space already presses the focused button or link.
    if (event.key === " " && event.target.closest("button, a")) return;

    if (NEXT_KEYS.has(event.key)) next();
    else if (PREV_KEYS.has(event.key)) prev();
    else if (event.key === "Home") show(0);
    else if (event.key === "End") show(slides.length - 1);
    else return;
    event.preventDefault();
  }

  label(deck, slides);
  deck.after(bar.element);

  bar.onPrev(prev);
  bar.onNext(next);
  bar.onDot(show);
  document.addEventListener("keydown", onKey);
  window.addEventListener("hashchange", () => show(indexFromHash()));
  onSwipe(deck, { left: next, right: prev });

  // A hash means the reader came from another card (or a link to one), so
  // the keyboard focus follows them onto the card; a plain visit leaves the
  // focus at the top of the page.
  show(indexFromHash(), { focus: location.hash !== "" });
}

// sceneOf returns the animation on a card, once its component has loaded;
// until then the tag is only its text fallback.
function sceneOf(slide) {
  const scene = slide.querySelector("[scene]");
  return typeof scene?.next === "function" ? scene : null;
}

// label gives the deck and every card the names a screen reader announces.
function label(deck, slides) {
  deck.setAttribute("role", "region");
  deck.setAttribute("aria-roledescription", "presentación");
  deck.setAttribute("aria-label", deck.querySelector("h1")?.textContent ?? "");
  slides.forEach((slide, i) => {
    slide.setAttribute("role", "group");
    slide.setAttribute("aria-roledescription", "diapositiva");
    slide.setAttribute("aria-label", `${i + 1} de ${slides.length}`);
    slide.tabIndex = -1;
  });
}

// buildBar draws the bar under the card: the two arrows, a mini route with
// one stop per card and the "3 / 9" counter.
function buildBar(count) {
  const element = document.createElement("nav");
  element.className = "deck-bar";
  element.setAttribute("aria-label", "Diapositivas");

  const prevButton = arrow("‹", "Diapositiva anterior");
  const nextButton = arrow("›", "Diapositiva siguiente");
  const route = document.createElement("ol");
  route.className = "deck-route";
  const dots = Array.from({ length: count }, (_, i) => {
    const dot = document.createElement("button");
    dot.type = "button";
    dot.className = "deck-dot";
    dot.setAttribute("aria-label", `Ir a la diapositiva ${i + 1}`);
    const stop = document.createElement("li");
    stop.append(dot);
    route.append(stop);
    return dot;
  });
  const counter = document.createElement("p");
  counter.className = "deck-count";

  element.append(prevButton, route, counter, nextButton);

  return {
    element,
    onPrev: (handler) => prevButton.addEventListener("click", handler),
    onNext: (handler) => nextButton.addEventListener("click", handler),
    onDot: (handler) => dots.forEach((dot, i) => dot.addEventListener("click", () => handler(i))),
    update(index, visited, { hasPrev, hasNext }) {
      dots.forEach((dot, i) => {
        dot.classList.toggle("is-visited", visited.has(i));
        if (i === index) dot.setAttribute("aria-current", "step");
        else dot.removeAttribute("aria-current");
      });
      counter.textContent = `${index + 1} / ${count}`;
      setEnabled(prevButton, hasPrev, nextButton);
      setEnabled(nextButton, hasNext, prevButton);
    },
  };
}

function arrow(text, name) {
  const button = document.createElement("button");
  button.type = "button";
  button.className = "deck-arrow";
  button.textContent = text;
  button.setAttribute("aria-label", name);
  return button;
}

// setEnabled turns an arrow on or off. A disabled button drops focus, so the
// focus moves to the other arrow instead of getting lost.
function setEnabled(button, enabled, other) {
  const hadFocus = document.activeElement === button;
  button.disabled = !enabled;
  if (hadFocus && !enabled) other.focus();
}

// onSwipe calls left or right after a horizontal swipe on a touch screen.
// A swipe that starts on a code block scrolls the code instead.
function onSwipe(element, { left, right }) {
  let start = null;
  element.addEventListener("pointerdown", (event) => {
    if (event.pointerType !== "touch" || event.target.closest("pre")) return;
    start = { x: event.clientX, y: event.clientY };
  });
  element.addEventListener("pointercancel", () => (start = null));
  element.addEventListener("pointerup", (event) => {
    if (!start) return;
    const dx = event.clientX - start.x;
    const dy = event.clientY - start.y;
    start = null;
    if (Math.abs(dx) < SWIPE_PX || Math.abs(dx) < Math.abs(dy)) return;
    if (dx < 0) left();
    else right();
  });
}
