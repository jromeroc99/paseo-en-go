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
// behaves the same: canvas, caption, code panel and the endless loop.
//
// Components draw on a 160×96 grid of logical units; the canvas stretches
// that grid to the width of the page and stays sharp on any screen.

const WIDTH = 160;
const HEIGHT = 96;

const MOVE_MS = 450; // the step's motion
const HOLD_MS = 2600; // time to read the caption before the next step
const END_HOLD_MS = 4000; // the last step stays a little longer
const FADE_MS = 400; // fade out before the loop starts again

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
 * Scene is the custom element every animation extends. It plays its steps
 * in an endless loop while it is on screen, and pauses while the reader
 * points at it or focuses it, so nobody has to read against the clock.
 */
export class Scene extends HTMLElement {
  #scene = null;
  #paint = null;
  #ui = {};
  #index = 0;
  #phase = "hold"; // "move", then "hold", then "fade" after the last step
  #elapsed = 0;
  #lastFrame = 0;
  #frameRequest = 0;
  #onScreen = false;
  #held = new Set(); // why the reader paused it: "pointer", "focus"
  #observers = [];
  #reducedMotion = matchMedia("(prefers-reduced-motion: reduce)");
  #onVisibility = () => this.#update();

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
    this.#observe();
  }

  disconnectedCallback() {
    for (const observer of this.#observers) observer.disconnect();
    this.#observers = [];
    document.removeEventListener("visibilitychange", this.#onVisibility);
    this.#onScreen = false;
    this.#update();
  }

  /** draw renders one step at progress t (0 to 1). Components override it. */
  draw(paint, step, t) {}

  #build() {
    // The fallback paragraph becomes the description of the whole scene.
    const description = this.textContent.trim().replace(/\s+/g, " ");

    const canvas = document.createElement("canvas");
    canvas.className = "anim-canvas";
    canvas.width = WIDTH; // sets the aspect ratio until fitCanvas measures it
    canvas.height = HEIGHT;
    canvas.setAttribute("role", "img");
    canvas.setAttribute("aria-label", description);

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

    const root = element("div", "anim");
    root.tabIndex = 0;
    root.append(canvas, code, caption);
    this.replaceChildren(root);

    root.addEventListener("pointerenter", () => this.#hold("pointer", true));
    root.addEventListener("pointerleave", () => this.#hold("pointer", false));
    root.addEventListener("focusin", () => this.#hold("focus", true));
    root.addEventListener("focusout", () => this.#hold("focus", false));

    Object.assign(this.#ui, { root, canvas, caption });
    this.#paint = new Paint(canvas.getContext("2d"), readPalette(), {});
  }

  #observe() {
    // CSS gives the canvas its width; the drawing follows it.
    const resize = new ResizeObserver(() => this.#fitCanvas());
    resize.observe(this.#ui.canvas);
    const visible = new IntersectionObserver(([entry]) => {
      this.#onScreen = entry.isIntersecting;
      this.#update();
    });
    visible.observe(this.#ui.root);
    this.#observers = [resize, visible];
    document.addEventListener("visibilitychange", this.#onVisibility);
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

  #hold(reason, on) {
    if (on) this.#held.add(reason);
    else this.#held.delete(reason);
    this.#update();
  }

  #isPlaying() {
    return this.#onScreen && this.#held.size === 0 && !document.hidden;
  }

  // update starts or stops the frame loop to match #isPlaying.
  #update() {
    if (this.#isPlaying() && !this.#frameRequest) {
      this.#lastFrame = performance.now();
      this.#frameRequest = requestAnimationFrame((now) => this.#frame(now));
    }
    if (!this.#isPlaying() && this.#frameRequest) {
      cancelAnimationFrame(this.#frameRequest);
      this.#frameRequest = 0;
    }
  }

  #frame(now) {
    this.#elapsed += now - this.#lastFrame;
    this.#lastFrame = now;
    while (this.#elapsed >= this.#duration()) {
      this.#elapsed -= this.#duration();
      this.#nextPhase();
    }
    this.#render();
    this.#frameRequest = requestAnimationFrame((next) => this.#frame(next));
  }

  #duration() {
    const smooth = !this.#reducedMotion.matches;
    if (this.#phase === "move") return smooth ? MOVE_MS : 0;
    if (this.#phase === "fade") return smooth ? FADE_MS : 0;
    return this.#isLast() ? END_HOLD_MS : HOLD_MS;
  }

  #nextPhase() {
    if (this.#phase === "move") {
      this.#phase = "hold";
    } else if (this.#phase === "hold" && this.#isLast()) {
      this.#phase = "fade";
    } else {
      // After the fade the loop starts again at the first step.
      this.#index = this.#phase === "fade" ? 0 : this.#index + 1;
      this.#phase = "move";
      this.#showText();
    }
  }

  #isLast() {
    return this.#index === this.#scene.steps.length - 1;
  }

  #render() {
    if (!this.#paint.style.font) return; // not measured yet; fitCanvas renders
    const progress = this.#duration() === 0 ? 1 : Math.min(1, this.#elapsed / this.#duration());
    const t = this.#phase === "move" ? easeInOut(progress) : 1;
    // Fade out after the last step, and back in while the first one moves.
    let opacity = 1;
    if (this.#phase === "fade") opacity = 1 - progress;
    if (this.#phase === "move" && this.#index === 0) opacity = progress;
    this.#ui.canvas.style.opacity = opacity;

    this.#paint.clear();
    this.draw(this.#paint, this.#scene.steps[this.#index], t);
  }

  #showText() {
    const step = this.#scene.steps[this.#index];
    const parts = step.caption.split("`");
    this.#ui.caption.replaceChildren(
      ...parts.map((part, i) => (i % 2 === 1 ? element("code", "", part) : document.createTextNode(part))),
    );
    this.#ui.lines.forEach((line, i) => line.classList.toggle("is-current", i === step.line));
  }
}

function element(tag, className = "", text = "") {
  const node = document.createElement(tag);
  if (className) node.className = className;
  if (text) node.textContent = text;
  return node;
}
