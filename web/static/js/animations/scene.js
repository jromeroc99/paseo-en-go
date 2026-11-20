// Scene is the base of every animation in the course.
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
// step moves. Everything else lives here, so every animation looks and
// behaves the same: canvas, caption, code panel and the steps the reader
// walks through with the deck's arrows.
//
// Components draw on a 160×96 grid of logical units; the canvas stretches
// that grid to the width of the page and stays sharp on any screen.

const WIDTH = 160;
const HEIGHT = 96;

const MOVE_MS = 450; // the step's motion

const FONT_SIZE = 6.5; // logical units; a character is about 4 units wide
const ARROW_HEAD = 3; // length of an arrow head, in logical units

const PALETTE = ["ink", "paper", "value", "pointer", "highlight", "ok", "error", "dim"];

/** lerp moves from a to b as t goes from 0 to 1. */
export function lerp(a, b, t) {
  return a + (b - a) * t;
}

// Paint draws the shapes of a scene in the site's style: rounded boxes,
// hairline outlines and monospace text. Coordinates are logical units.
// Colours are named by meaning: paint.colors.value, .pointer…
class Paint {
  constructor(context, colors, style) {
    this.context = context;
    this.colors = colors;
    this.width = WIDTH;
    this.height = HEIGHT;
    this.style = style; // { hairline, radius, font }, in logical units
  }

  clear() {
    this.context.clearRect(0, 0, WIDTH, HEIGHT);
  }

  rect(x, y, width, height, color) {
    this.#withGlow(color, () => {
      this.context.fillStyle = color;
      this.context.fillRect(x, y, width, height);
    });
  }

  /** dot draws a small round marker centred on the unit at (x, y). */
  dot(x, y, color) {
    this.#withGlow(color, () => {
      this.context.fillStyle = color;
      this.context.beginPath();
      this.context.arc(x + 0.5, y + 0.5, 1.5, 0, 2 * Math.PI);
      this.context.fill();
    });
  }

  /** frame draws a rounded outline; a dashed one marks a scope, not a value. */
  frame(x, y, width, height, color, dashed = false) {
    const { context, style } = this;
    this.#withGlow(color, () => {
      context.strokeStyle = color;
      context.lineWidth = style.hairline;
      context.setLineDash(dashed ? [2, 2] : []);
      context.beginPath();
      context.roundRect(x + 0.5, y + 0.5, width - 1, height - 1, style.radius);
      context.stroke();
      context.setLineDash([]);
    });
  }

  /** line draws a straight line through the centres of both end units. */
  line(x1, y1, x2, y2, color) {
    const { context } = this;
    context.strokeStyle = color;
    context.lineWidth = this.style.stroke;
    context.beginPath();
    context.moveTo(x1 + 0.5, y1 + 0.5);
    context.lineTo(x2 + 0.5, y2 + 0.5);
    context.stroke();
  }

  /**
   * arrow draws a path of straight segments with a head at its tip. With t
   * below 1 only that fraction of the path is drawn, so an arrow can grow
   * from its start to its target.
   */
  arrow(points, color, t = 1) {
    if (t <= 0) return;
    const path = points.slice(0, 1);
    let left = pathLength(points) * t;
    for (let i = 1; i < points.length && left > 0; i++) {
      const [x1, y1] = points[i - 1];
      const [x2, y2] = points[i];
      const length = Math.hypot(x2 - x1, y2 - y1);
      const part = Math.min(1, left / length);
      path.push([lerp(x1, x2, part), lerp(y1, y2, part)]);
      left -= length;
    }
    const [from, tip] = path.slice(-2);
    const angle = Math.atan2(tip[1] - from[1], tip[0] - from[0]);
    // Stop the line where the head starts, so the tip stays sharp.
    const end = [tip[0] - Math.cos(angle) * ARROW_HEAD, tip[1] - Math.sin(angle) * ARROW_HEAD];

    const { context } = this;
    context.strokeStyle = color;
    context.fillStyle = color;
    context.lineWidth = this.style.stroke;
    context.lineJoin = "round";
    context.beginPath();
    path.slice(0, -1).forEach(([x, y], i) => (i === 0 ? context.moveTo(x + 0.5, y + 0.5) : context.lineTo(x + 0.5, y + 0.5)));
    context.lineTo(end[0] + 0.5, end[1] + 0.5);
    context.stroke();

    context.save();
    context.translate(tip[0] + 0.5, tip[1] + 0.5);
    context.rotate(angle);
    context.beginPath();
    context.moveTo(0, 0);
    context.lineTo(-ARROW_HEAD, -ARROW_HEAD * 0.7);
    context.lineTo(-ARROW_HEAD, ARROW_HEAD * 0.7);
    context.closePath();
    context.fill();
    context.restore();
  }

  /**
   * text draws s with its top-left corner at (x, y) and returns its width.
   * The text box is 5 units tall, so components can lay text out on a grid.
   */
  text(s, x, y, color) {
    const { context } = this;
    context.font = this.style.font;
    context.textBaseline = "middle";
    context.textAlign = "left";
    this.#withGlow(color, () => {
      context.fillStyle = color;
      context.fillText(String(s), x, y + 2.5);
    });
    return this.textWidth(s);
  }

  /** centerText draws s centred inside the span [x, x+width). */
  centerText(s, x, y, width, color) {
    this.text(s, x + (width - this.textWidth(s)) / 2, y, color);
  }

  textWidth(s) {
    this.context.font = this.style.font;
    return this.context.measureText(String(s)).width;
  }

  // The glow marks what changes in this step, as it does on the rest of
  // the site, so it follows the highlight colour and nothing else.
  #withGlow(color, paint) {
    if (color !== this.colors.highlight) {
      paint();
      return;
    }
    const { context } = this;
    context.save();
    context.shadowColor = color;
    context.shadowBlur = 4;
    paint();
    context.restore();
  }
}

/** pointOnPath returns the point a fraction t of the way along a path. */
export function pointOnPath(points, t) {
  let left = pathLength(points) * t;
  for (let i = 1; i < points.length; i++) {
    const [x1, y1] = points[i - 1];
    const [x2, y2] = points[i];
    const length = Math.hypot(x2 - x1, y2 - y1);
    if (left <= length) {
      const part = length === 0 ? 0 : left / length;
      return [lerp(x1, x2, part), lerp(y1, y2, part)];
    }
    left -= length;
  }
  return points[points.length - 1];
}

function pathLength(points) {
  let length = 0;
  for (let i = 1; i < points.length; i++) {
    length += Math.hypot(points[i][0] - points[i - 1][0], points[i][1] - points[i - 1][1]);
  }
  return length;
}

// easeInOut starts and ends a motion slowly, like the site's transitions.
function easeInOut(t) {
  return t < 0.5 ? 2 * t * t : 1 - (-2 * t + 2) ** 2 / 2;
}

function readPalette() {
  const style = getComputedStyle(document.documentElement);
  return Object.fromEntries(PALETTE.map((name) => [name, style.getPropertyValue(`--anim-${name}`).trim()]));
}

/**
 * Scene is the custom element every animation extends. Nothing moves on its
 * own: the deck calls next() and prev() on every arrow press, and a click on
 * the canvas plays the next step, so the reader sets the pace.
 */
export class Scene extends HTMLElement {
  #scene = null;
  #paint = null;
  #ui = {};
  #index = 0;
  #reached = 0; // the furthest step shown; the code after it is still to come
  #progress = 1; // how far the current step has moved, from 0 to 1
  #moveStart = 0;
  #frameRequest = 0;
  #resize = null;
  #reducedMotion = matchMedia("(prefers-reduced-motion: reduce)");

  connectedCallback() {
    if (!this.#paint) {
      const name = this.getAttribute("scene");
      this.#scene = this.constructor.scenes?.[name];
      if (!this.#scene) {
        // Keep the text fallback: the lesson still reads well without the scene.
        console.error(`<${this.localName}>: unknown scene "${name}"`);
        return;
      }
      this.#build();
      this.#showText();
    }
    // CSS gives the canvas its width; the drawing follows it.
    this.#resize = new ResizeObserver(() => this.#fitCanvas());
    this.#resize.observe(this.#ui.canvas);
  }

  disconnectedCallback() {
    this.#resize?.disconnect();
    this.#stopMoving();
  }

  /** draw renders one step at progress t (0 to 1). Components override it. */
  draw(paint, step, t) {}

  /** next plays the next step, or returns false when there is none left. */
  next() {
    if (!this.#scene || this.#index === this.#lastIndex()) return false;
    this.#show(this.#index + 1, { move: true });
    return true;
  }

  /** prev shows the previous step at once, or returns false on the first. */
  prev() {
    if (!this.#scene || this.#index === 0) return false;
    this.#show(this.#index - 1);
    return true;
  }

  /** rewind plays the scene from its first step, for a reader arriving forwards. */
  rewind() {
    if (!this.#scene) return;
    this.#reached = 0;
    this.#show(0, { move: true });
  }

  /** skipToEnd shows the last step, for a reader coming back from the next card. */
  skipToEnd() {
    if (!this.#scene) return;
    this.#show(this.#lastIndex());
  }

  #build() {
    // The fallback paragraph becomes the description of the whole scene.
    const description = this.textContent.trim().replace(/\s+/g, " ");

    const canvas = document.createElement("canvas");
    canvas.className = "anim-canvas";
    canvas.width = WIDTH; // sets the aspect ratio until fitCanvas measures it
    canvas.height = HEIGHT;
    canvas.setAttribute("role", "img");
    canvas.setAttribute("aria-label", description);
    canvas.addEventListener("click", () => this.next());

    const code = element("pre", "anim-code");
    const lines = element("code");
    this.#ui.lines = this.#scene.code.map((text) => {
      const line = element("span", "anim-line");
      line.textContent = text;
      lines.append(line);
      return line;
    });
    code.append(lines);

    // The caption changes only when the reader moves, so it is announced.
    const caption = element("p", "anim-caption");
    caption.setAttribute("aria-live", "polite");

    const root = element("div", "anim");
    root.append(code, canvas, caption);
    this.replaceChildren(root);

    Object.assign(this.#ui, { root, canvas, caption });
    this.#paint = new Paint(canvas.getContext("2d"), readPalette(), {});
  }

  // fitCanvas sizes the canvas to its frame and to the screen's pixel
  // density, so lines and text stay sharp.
  #fitCanvas() {
    const { canvas, root } = this.#ui;
    const cssWidth = canvas.clientWidth;
    if (cssWidth === 0) return;
    const scale = cssWidth / WIDTH; // CSS pixels per logical unit
    const density = window.devicePixelRatio || 1;
    canvas.width = Math.round(cssWidth * density);
    canvas.height = Math.round(HEIGHT * scale * density);
    this.#paint.context.setTransform(scale * density, 0, 0, scale * density, 0, 0);

    const radius = parseFloat(getComputedStyle(root).getPropertyValue("--radius-sm")) || 6;
    const mono = getComputedStyle(root).getPropertyValue("--mono").trim() || "monospace";
    this.#paint.style = {
      hairline: 1 / scale, // one CSS pixel, like every border on the site
      stroke: 1.5 / scale,
      radius: radius / scale,
      font: `${FONT_SIZE}px ${mono}`,
    };
    this.#render();
  }

  #lastIndex() {
    return this.#scene.steps.length - 1;
  }

  // show moves to a step. Going forwards the step plays its motion; going
  // back, or with reduced motion, it appears already finished.
  #show(index, { move = false } = {}) {
    this.#stopMoving();
    this.#index = index;
    this.#reached = Math.max(this.#reached, index);
    this.#showText();
    if (move && !this.#reducedMotion.matches) {
      this.#progress = 0;
      this.#moveStart = performance.now();
      this.#frameRequest = requestAnimationFrame((now) => this.#frame(now));
    } else {
      this.#progress = 1;
      this.#render();
    }
  }

  #frame(now) {
    this.#progress = Math.min(1, (now - this.#moveStart) / MOVE_MS);
    this.#render();
    this.#frameRequest = this.#progress < 1 ? requestAnimationFrame((next) => this.#frame(next)) : 0;
  }

  #stopMoving() {
    cancelAnimationFrame(this.#frameRequest);
    this.#frameRequest = 0;
  }

  #render() {
    if (!this.#paint.style.font) return; // not measured yet; fitCanvas renders
    // The first step fades in, so the scene starts from an empty frame.
    this.#ui.canvas.style.opacity = this.#index === 0 ? this.#progress : 1;
    this.#paint.clear();
    this.draw(this.#paint, this.#scene.steps[this.#index], easeInOut(this.#progress));
  }

  #showText() {
    const steps = this.#scene.steps;
    const step = steps[this.#index];
    const parts = step.caption.split("`");
    this.#ui.caption.replaceChildren(
      ...parts.map((part, i) => (i % 2 === 1 ? element("code", "", part) : document.createTextNode(part))),
    );
    // A loop in the Go code jumps back up, so the code still to come is what
    // lies below the furthest line any step shown so far has reached.
    const lastRun = Math.max(...steps.slice(0, this.#reached + 1).map((s) => s.line));
    this.#ui.lines.forEach((line, i) => {
      line.classList.toggle("is-current", i === step.line);
      line.classList.toggle("is-ahead", i > lastRun);
    });
  }
}

function element(tag, className = "", text = "") {
  const node = document.createElement(tag);
  if (className) node.className = className;
  if (text) node.textContent = text;
  return node;
}
