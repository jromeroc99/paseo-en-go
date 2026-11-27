// <stack-demo> shows the stack as a pile of call frames. Scenes: "calls"
// (each call stacks a frame and each return takes it off) and "escape" (a
// value whose address leaves its function is kept on the heap instead).
import { Scene, lerp, pointOnPath } from "./scene.js";

const DROP = 6; // units an entering frame falls before it settles

// The stack grows upwards from its bottom edge; frames share their borders.
const STACK = { x: 4, bottom: 90, width: 70, frameHeight: 14 };

// calls: the stack on the left and the output on the right.
const CALLS_OUTPUT = { x: 86, y: 52, width: 70, height: 40 };

// escape: the heap at the top right, with p inside it, and the output
// below it. q's arrow climbs between the stack and the heap.
const HEAP = { x: 88, y: 4, width: 68, height: 40 };
const P_BOX = { x: 100, y: 20, width: 40, height: 14 };
const ESCAPE_OUTPUT = { x: 88, y: 52, width: 68, height: 40 };
const Q_TO_P = [
  [74, 84],
  [80, 84],
  [80, 27],
  [99, 27],
];

const MAIN = { name: "main" };
const MAIN_Q = { name: "main", vars: "q", pointer: true }; // q holds an address

const scenes = {
  calls: {
    code: [
      "func square(n int) int {",
      "\tr := n * n",
      "\treturn r",
      "}",
      "func sumSquares(a, b int) int {",
      "\treturn square(a) + square(b)",
      "}",
      "func main() {",
      "\tfmt.Println(sumSquares(2, 3))",
      "}",
    ],
    steps: [
      {
        line: 7,
        caption: "Al arrancar, `main` tiene su frame en el stack: su trozo de memoria para variables.",
        frames: [MAIN],
        enter: "main",
      },
      {
        line: 4,
        caption: "`sumSquares(2, 3)` apila un frame encima, con sus parámetros `a` y `b`.",
        frames: [MAIN, { name: "sumSquares", vars: "a=2 b=3" }],
        enter: "sumSquares",
      },
      {
        line: 1,
        caption: "`square(a)` apila otro frame. Dentro, `n` vale 2 y `r` vale 4.",
        frames: [MAIN, { name: "sumSquares", vars: "a=2 b=3" }, { name: "square", vars: "n=2 r=4" }],
        enter: "square",
      },
      {
        line: 2,
        caption: "`return r` devuelve 4 y el frame de `square` se va. Sus variables dejan de existir.",
        frames: [MAIN, { name: "sumSquares", vars: "a=2 b=3" }, { name: "square", vars: "n=2 r=4", gone: true }],
      },
      {
        line: 1,
        caption: "`square(b)` usa el mismo hueco del stack, ahora con `n` igual a 3.",
        frames: [MAIN, { name: "sumSquares", vars: "a=2 b=3" }, { name: "square", vars: "n=3 r=9" }],
        enter: "square",
      },
      {
        line: 2,
        caption: "Devuelve 9 y su frame se va. `sumSquares` suma 4 + 9.",
        frames: [MAIN, { name: "sumSquares", vars: "a=2 b=3" }, { name: "square", vars: "n=3 r=9", gone: true }],
      },
      {
        line: 8,
        caption: "`sumSquares` devuelve 13 y su frame también se va. `main` imprime 13.",
        frames: [MAIN, { name: "sumSquares", vars: "a=2 b=3", gone: true }],
        output: ["13"],
        printed: true,
      },
    ],
  },

  escape: {
    code: [
      "func newPoint(x, y int) *Point {",
      "\tp := Point{x, y}",
      "\treturn &p",
      "}",
      "func main() {",
      "\tq := newPoint(1, 2)",
      "\tfmt.Println(*q)",
      "}",
    ],
    steps: [
      {
        line: 5,
        caption: "`main` va a llamar a `newPoint`. Su frame guarda `q`, que aún no apunta a nada.",
        frames: [MAIN_Q],
        enter: "main",
      },
      {
        line: 0,
        caption: "`newPoint(1, 2)` apila su frame, con `x` y `y`.",
        frames: [MAIN_Q, { name: "newPoint", vars: "x=1 y=2" }],
        enter: "newPoint",
      },
      {
        line: 1,
        caption: "El compilador ve que la dirección de `p` va a salir de la función. Por eso guarda `p` en el heap.",
        frames: [MAIN_Q, { name: "newPoint", vars: "x=1 y=2" }],
        point: "entering",
      },
      {
        line: 2,
        caption: "`return &p` devuelve esa dirección: `q` apunta a `p`, en el heap.",
        frames: [MAIN_Q, { name: "newPoint", vars: "x=1 y=2" }],
        point: "there",
        arrow: "growing",
      },
      {
        line: 5,
        caption: "El frame de `newPoint` se va, pero `p` sigue en el heap: `q` aún la necesita.",
        frames: [MAIN_Q, { name: "newPoint", vars: "x=1 y=2", gone: true }],
        point: "there",
        arrow: "there",
      },
      {
        line: 6,
        caption: "`*q` sigue la flecha hasta el heap e imprime {1 2}.",
        frames: [MAIN_Q],
        point: "there",
        arrow: "followed",
        output: ["{1 2}"],
        printed: true,
      },
    ],
  },
};

class StackDemo extends Scene {
  static scenes = scenes;

  draw(paint, step, t) {
    paint.text("stack", STACK.x, 4, paint.colors.paper);
    step.frames.forEach((frame, i) => drawFrame(paint, frame, i, frame.name === step.enter ? t : 1));
    const escape = this.getAttribute("scene") === "escape";
    if (escape) drawHeap(paint, step, t);
    if (step.output) drawOutput(paint, escape ? ESCAPE_OUTPUT : CALLS_OUTPUT, step.output, step.printed, t);
  }
}

// drawFrame draws the call frame at height i of the stack: its function on
// the left and its variables on the right. A frame on its way out is dim.
function drawFrame(paint, frame, i, settle) {
  const { colors } = paint;
  const y = STACK.bottom - (i + 1) * (STACK.frameHeight - 1) - lerp(DROP, 0, settle);
  let color = settle < 1 ? colors.highlight : colors.paper;
  let varsColor = frame.pointer ? colors.pointer : colors.value;
  if (frame.gone) color = varsColor = colors.dim;

  paint.frame(STACK.x, y, STACK.width, STACK.frameHeight, color);
  paint.text(frame.name, STACK.x + 3, y + 4.5, color);
  if (frame.vars) {
    const width = paint.textWidth(frame.vars);
    paint.text(frame.vars, STACK.x + STACK.width - 4 - width, y + 4.5, varsColor);
  }
}

// drawHeap draws the heap, p inside it once it exists and q's arrow to p.
function drawHeap(paint, step, t) {
  const { colors } = paint;
  paint.frame(HEAP.x, HEAP.y, HEAP.width, HEAP.height, colors.paper, true);
  paint.text("heap", HEAP.x + 4, HEAP.y + 4, colors.paper);
  if (!step.point) return;

  const entering = step.point === "entering";
  const y = P_BOX.y - (entering ? lerp(DROP, 0, t) : 0);
  const color = entering ? colors.highlight : colors.value;
  paint.text("p", P_BOX.x, y - 7, colors.paper);
  paint.frame(P_BOX.x, y, P_BOX.width, P_BOX.height, color);
  paint.centerText("{1 2}", P_BOX.x, y + 4.5, P_BOX.width, color);

  if (!step.arrow) return;
  paint.arrow(Q_TO_P, colors.pointer, step.arrow === "growing" ? t : 1);
  if (step.arrow === "followed" && t < 1) {
    const [x, dotY] = pointOnPath(Q_TO_P, t);
    paint.dot(x, dotY, colors.highlight);
  }
}

// drawOutput draws what the program has printed so far. When this step
// printed, its line is the last one: it appears when the step's motion
// ends, in the highlight colour.
function drawOutput(paint, strip, lines, printed, t) {
  const { colors } = paint;
  const { x, y, width, height } = strip;
  paint.frame(x, y, width, height, colors.dim);
  paint.text("salida", x + 4, y + 4, colors.dim);
  lines.forEach((line, i) => {
    const isNew = printed && i === lines.length - 1;
    if (isNew && t < 1) return;
    paint.text(line, x + 4, y + 12 + i * 8, isNew ? colors.highlight : colors.paper);
  });
}

customElements.define("stack-demo", StackDemo);
