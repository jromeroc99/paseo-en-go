// PixelScene is the base of every pixel-art animation in the course.
//
// A component extends it and declares two things:
//
//   static scenes = { basics: { code: ["i := 42", …], steps: [ … ] } };
//   draw(paint, step, t) { … }
//
// A scene is plain data. Each step has at least `line` (the index of the Go
// line it represents) and `caption` (one line of Spanish, with `code` in
// backticks); the rest of the step is whatever the component's draw needs.
// draw renders one step at progress t, which goes from 0 to 1 while the
// step plays. Everything else lives here, so every animation looks and
// behaves the same: canvas, controls, caption, code panel and timing.

const WIDTH = 160;
const HEIGHT = 96;
const FRAMES_PER_STEP = 4; // motion in a few chunky frames, like a retro game
const FRAME_MS = 110;
const HOLD_MS = 2600; // time to read the caption before the next step plays

const PALETTE = ["ink", "paper", "value", "pointer", "highlight", "ok", "error", "dim"];

// A 3×5 pixel font, one row of three pixels per group. Letters have a
// single case, so text is drawn in capitals.
const GLYPHS = {
  A: "010 101 111 101 101", B: "110 101 110 101 110", C: "011 100 100 100 011",
  D: "110 101 101 101 110", E: "111 100 110 100 111", F: "111 100 110 100 100",
  G: "011 100 101 101 011", H: "101 101 111 101 101", I: "111 010 010 010 111",
  J: "001 001 001 101 010", K: "101 101 110 101 101", L: "100 100 100 100 111",
  M: "101 111 111 101 101", N: "110 101 101 101 101", O: "010 101 101 101 010",
  P: "110 101 110 100 100", Q: "010 101 101 110 011", R: "110 101 110 101 101",
  S: "011 100 010 001 110", T: "111 010 010 010 010", U: "101 101 101 101 111",
  V: "101 101 101 101 010", W: "101 101 111 111 101", X: "101 101 010 101 101",
  Y: "101 101 010 010 010", Z: "111 001 010 100 111",
  0: "111 101 101 101 111", 1: "010 110 010 010 111", 2: "110 001 010 100 111",
  3: "110 001 010 001 110", 4: "101 101 111 001 001", 5: "111 100 110 001 110",
  6: "011 100 111 101 111", 7: "111 001 010 010 010", 8: "111 101 111 101 111",
  9: "111 101 111 001 110",
  " ": "000 000 000 000 000", ".": "000 000 000 000 010", ",": "000 000 000 010 100",
  "!": "010 010 010 000 010", "?": "110 001 010 000 010", ":": "000 010 000 010 000",
  ";": "000 010 000 010 100", "=": "000 111 000 111 000", "&": "010 101 010 101 011",
  "*": "000 101 010 101 000", "[": "110 100 100 100 110", "]": "011 001 001 001 011",
  "{": "011 010 110 010 011", "}": "110 010 011 010 110", "(": "001 010 010 010 001",
  ")": "100 010 010 010 100", "-": "000 000 111 000 000", "+": "000 010 111 010 000",
  ">": "100 010 001 010 100", "<": "001 010 100 010 001", _: "000 000 000 000 111",
  '"': "101 101 000 000 000", "'": "010 010 000 000 000", "/": "001 001 010 100 100",
  "#": "101 111 101 111 101", "%": "101 001 010 100 101",
};
const GLYPH_ADVANCE = 4; // 3 pixels of glyph and 1 of space

/** lerp moves from a to b as t goes from 0 to 1, snapped to whole pixels. */
export function lerp(a, b, t) {
  return Math.round(a + (b - a) * t);
}

// Paint draws chunky shapes on the canvas, always on whole pixels and with
// no smoothing. Colours are named by meaning: paint.colors.value, .pointer…
class Paint {
  constructor(context, colors) {
    this.context = context;
    this.colors = colors;
    this.width = WIDTH;
    this.height = HEIGHT;
  }

  clear() {
    this.rect(0, 0, WIDTH, HEIGHT, this.colors.ink);
  }

  rect(x, y, width, height, color) {
    this.context.fillStyle = color;
    this.context.fillRect(Math.round(x), Math.round(y), width, height);
  }

  /** frame draws a 1-pixel outline; a dashed one marks a scope, not a value. */
  frame(x, y, width, height, color, dashed = false) {
    if (!dashed) {
      this.rect(x, y, width, 1, color);
      this.rect(x, y + height - 1, width, 1, color);
      this.rect(x, y, 1, height, color);
      this.rect(x + width - 1, y, 1, height, color);
      return;
    }
    for (let i = 0; i < width; i += 4) {
      this.rect(x + i, y, 2, 1, color);
      this.rect(x + i, y + height - 1, 2, 1, color);
    }
    for (let i = 0; i < height; i += 4) {
      this.rect(x, y + i, 1, 2, color);
      this.rect(x + width - 1, y + i, 1, 2, color);
    }
  }

  /** line draws a horizontal or vertical line, both ends included. */
  line(x1, y1, x2, y2, color) {
    this.rect(Math.min(x1, x2), Math.min(y1, y2), Math.abs(x2 - x1) + 1, Math.abs(y2 - y1) + 1, color);
  }

  /**
   * arrow draws a path of horizontal and vertical segments with a head at
   * its tip. With t below 1 only that fraction of the path is drawn, so an
   * arrow can grow from its start to its target.
   */
  arrow(points, color, t = 1) {
    let left = Math.round(pathLength(points) * t);
    let tip = null;
    for (let i = 1; i < points.length && left > 0; i++) {
      const [x1, y1] = points[i - 1];
      const [x2, y2] = points[i];
      const used = Math.min(left, Math.abs(x2 - x1) + Math.abs(y2 - y1));
      const dx = Math.sign(x2 - x1);
      const dy = Math.sign(y2 - y1);
      tip = { x: x1 + dx * used, y: y1 + dy * used, dx, dy };
      this.line(x1, y1, tip.x, tip.y, color);
      left -= used;
    }
    if (tip) this.#arrowHead(tip, color);
  }

  #arrowHead({ x, y, dx, dy }, color) {
    // Each row behind the tip is one pixel wider on both sides.
    for (let back = 1; back <= 2; back++) {
      for (let side = -back; side <= back; side++) {
        this.rect(x - dx * back + Math.abs(dy) * side, y - dy * back + Math.abs(dx) * side, 1, 1, color);
      }
    }
  }

  /** text draws s with the pixel font and returns its width in pixels. */
  text(s, x, y, color) {
    const chars = [...String(s).toUpperCase()];
    chars.forEach((char, i) => {
      const rows = (GLYPHS[char] ?? GLYPHS["?"]).split(" ");
      rows.forEach((row, dy) => {
        [...row].forEach((bit, dx) => {
          if (bit === "1") this.rect(x + i * GLYPH_ADVANCE + dx, y + dy, 1, 1, color);
        });
      });
    });
    return textWidth(s);
  }

  /** centerText draws s centred inside the span [x, x+width). */
  centerText(s, x, y, width, color) {
    this.text(s, x + Math.floor((width - textWidth(s)) / 2), y, color);
  }
}

/** pointOnPath returns the point a fraction t of the way along a path. */
export function pointOnPath(points, t) {
  let left = pathLength(points) * t;
  for (let i = 1; i < points.length; i++) {
    const [x1, y1] = points[i - 1];
    const [x2, y2] = points[i];
    const length = Math.abs(x2 - x1) + Math.abs(y2 - y1);
    if (left <= length) {
      const part = length === 0 ? 0 : left / length;
      return [lerp(x1, x2, part), lerp(y1, y2, part)];
    }
    left -= length;
  }
  return points[points.length - 1];
}

function textWidth(s) {
  const length = [...String(s)].length;
  return length === 0 ? 0 : length * GLYPH_ADVANCE - 1;
}

function pathLength(points) {
  let length = 0;
  for (let i = 1; i < points.length; i++) {
    length += Math.abs(points[i][0] - points[i - 1][0]) + Math.abs(points[i][1] - points[i - 1][1]);
  }
  return length;
}

function readPalette() {
  const style = getComputedStyle(document.documentElement);
  return Object.fromEntries(PALETTE.map((name) => [name, style.getPropertyValue(`--px-${name}`).trim()]));
}

/** PixelScene is the custom element every animation extends. */
export class PixelScene extends HTMLElement {
  #scene = null;
  #paint = null;
  #ui = {};
  #index = 0;
  #frame = FRAMES_PER_STEP;
  #playing = false;
  #timer = 0;
  #reducedMotion = matchMedia("(prefers-reduced-motion: reduce)");

  connectedCallback() {
    if (this.#paint) return; // already built; the element was only moved
    const name = this.getAttribute("scene");
    this.#scene = this.constructor.scenes?.[name];
    if (!this.#scene) {
      // Keep the text fallback: the lesson still reads well without the scene.
      console.error(`<${this.localName}>: unknown scene "${name}"`);
      return;
    }
    this.#build();
    this.#goTo(0, false);
  }

  disconnectedCallback() {
    this.#playing = false;
    clearTimeout(this.#timer);
  }

  /** draw renders one step at progress t (0 to 1). Components override it. */
  draw(paint, step, t) {}

  #build() {
    const canvas = element("canvas", "pixel-canvas");
    canvas.width = WIDTH;
    canvas.height = HEIGHT;
    canvas.setAttribute("role", "img");

    const code = element("pre", "anim-code");
    const lines = element("code");
    this.#ui.lines = this.#scene.code.map((text) => {
      const line = element("span", "anim-line");
      line.textContent = text;
      lines.append(line);
      return line;
    });
    code.append(lines);

    const caption = element("p", "anim-caption");
    caption.setAttribute("aria-live", "polite");

    const play = button("anim-play", () => (this.#playing ? this.#pause() : this.#play()));
    const step = button("anim-step", () => this.#stepForward(), "→", "Paso");
    const restart = button("anim-restart", () => this.#restart(), "↺", "Reiniciar");
    const counter = element("span", "anim-counter");
    const controls = element("div", "anim-controls");
    controls.append(play, step, restart, counter);

    const root = element("div", "anim");
    root.append(canvas, code, caption, controls);
    this.replaceChildren(root);

    Object.assign(this.#ui, { root, canvas, caption, play, step, counter });
    this.#paint = new Paint(canvas.getContext("2d"), readPalette());
    // The custom element itself is inline and has no width; the wrapper has.
    new ResizeObserver(() => this.#fitCanvas()).observe(root);
  }

  // Scale by a whole number, so every logical pixel is a square block.
  #fitCanvas() {
    const scale = Math.max(1, Math.floor(this.#ui.root.clientWidth / WIDTH));
    this.#ui.canvas.style.width = `${WIDTH * scale}px`;
  }

  #play() {
    this.#playing = true;
    if (this.#isLast()) {
      this.#goTo(0, false);
    } else {
      this.#goTo(this.#index + 1, true);
    }
  }

  #pause() {
    this.#playing = false;
    // A step that is still moving finishes its frames; only the wait stops.
    if (this.#frame === FRAMES_PER_STEP) clearTimeout(this.#timer);
    this.#showControls();
  }

  #stepForward() {
    this.#pause();
    if (!this.#isLast()) this.#goTo(this.#index + 1, true);
  }

  #restart() {
    this.#pause();
    this.#goTo(0, false);
  }

  #isLast() {
    return this.#index === this.#scene.steps.length - 1;
  }

  #goTo(index, animate) {
    clearTimeout(this.#timer);
    this.#index = index;
    this.#frame = animate && !this.#reducedMotion.matches ? 0 : FRAMES_PER_STEP;
    this.#showText();
    this.#showControls();
    this.#tick();
  }

  // tick draws the current frame, then schedules the next frame, the next
  // step (while playing) or nothing.
  #tick() {
    this.#paint.clear();
    this.draw(this.#paint, this.#scene.steps[this.#index], this.#frame / FRAMES_PER_STEP);

    if (this.#frame < FRAMES_PER_STEP) {
      this.#timer = setTimeout(() => {
        this.#frame++;
        this.#tick();
      }, FRAME_MS);
      return;
    }
    if (!this.#playing) return;
    if (this.#isLast()) {
      this.#pause();
      return;
    }
    this.#timer = setTimeout(() => this.#goTo(this.#index + 1, true), HOLD_MS);
  }

  #showText() {
    const step = this.#scene.steps[this.#index];
    const parts = step.caption.split("`");
    this.#ui.caption.replaceChildren(
      ...parts.map((part, i) => (i % 2 === 1 ? element("code", "", part) : document.createTextNode(part))),
    );
    this.#ui.canvas.setAttribute("aria-label", parts.join(""));
    this.#ui.lines.forEach((line, i) => line.classList.toggle("is-current", i === step.line));
  }

  #showControls() {
    const { play, step, counter } = this.#ui;
    setLabel(play, this.#playing ? "❚❚" : "▶", this.#playing ? "Pausa" : "Reproducir");
    play.classList.toggle("is-active", this.#playing);
    step.disabled = this.#isLast();
    counter.textContent = `Paso ${this.#index + 1} de ${this.#scene.steps.length}`;
  }
}

function element(tag, className = "", text = "") {
  const node = document.createElement(tag);
  if (className) node.className = className;
  if (text) node.textContent = text;
  return node;
}

function button(className, onClick, icon = "", label = "") {
  const node = element("button", className);
  node.type = "button";
  node.addEventListener("click", onClick);
  if (label) setLabel(node, icon, label);
  return node;
}

// The icon is decoration; screen readers hear only the label.
function setLabel(node, icon, label) {
  const mark = element("span", "anim-icon", icon);
  mark.setAttribute("aria-hidden", "true");
  node.replaceChildren(mark, ` ${label}`);
}
