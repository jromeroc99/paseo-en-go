// <pointer-demo> shows variables as boxes in memory and pointers as arrows
// to those boxes. Scenes: "variable" (a variable is a box with an address),
// "convert" (int(f) copies the value and drops its decimals), "basics" (&,
// * and nil), "closure" (each closure keeps its own captured variable) and
// "receivers" (a value receiver gets a copy, a pointer receiver gets the
// original).
import { Scene, lerp, pointOnPath } from "./scene.js";

const BOX_WIDTH = 40;
const BOX_HEIGHT = 16;
const DROP = 6; // units an entering box falls before it settles
const FALL = 20; // units a discarded piece of a value falls as it is lost

// The output strip, where fmt.Println writes.
const OUTPUT = { x: 84, y: 52, width: 72, height: 40 };

// variable: two ints, eight bytes apart in memory.
const COUNT = { id: "count", label: "count", x: 16, y: 16, addr: "0xc0" };
const OTHER = { id: "other", label: "other", x: 100, y: 16, addr: "0xc8" };

// convert: the float64 on the left, the int made from it on the right. The
// decimals start next to the 9 inside euros and fall away.
const PRICE = { id: "price", label: "price · float64", x: 8, y: 20, width: 64, value: "9.75" };
const EUROS = { id: "euros", label: "euros · int", x: 88, y: 20, width: 64 };
const DECIMALS = { text: ".75", x: 123, y: 26 };

// basics: p points to i; q is nil.
const I = { id: "i", label: "i", x: 100, y: 16, addr: "0xc0" };
const P = { id: "p", label: "p", x: 20, y: 16, kind: "pointer", value: "0xc0" };
const Q = { id: "q", label: "q", x: 20, y: 60, kind: "nil", value: "nil" };
const P_TO_I = { from: "p", points: [[61, 23], [98, 23]] };

// closure: each function adder returns points to its own sum.
const POS = { id: "pos", label: "pos", x: 8, y: 14, kind: "pointer", value: "func" };
const NEG = { id: "neg", label: "neg", x: 8, y: 60, kind: "pointer", value: "func" };
const POS_SUM = { id: "pos-sum", label: "sum", x: 100, y: 14 };
const NEG_SUM = { id: "neg-sum", label: "sum", x: 100, y: 60 };
const POS_TO_SUM = { from: "pos", points: [[49, 21], [98, 21]] };
const NEG_TO_SUM = { from: "neg", points: [[49, 67], [98, 67]] };

// receivers: the method's own v lives inside the method's scope.
const V = { id: "v", label: "v", x: 8, y: 24, width: 56, addr: "0xc0" };
const METHOD_V = { id: "method-v", label: "v", x: 88, y: 24, width: 60 };
const METHOD_TO_V = { from: "method-v", points: [[87, 31], [65, 31]] };
const SCOPE = { x: 80, y: 4, width: 76, height: 44 };

const scenes = {
  variable: {
    code: ["var count int", "count = 42", "other := count", "other = 7", "fmt.Println(count, other)"],
    steps: [
      {
        line: 0,
        caption: "`var count int` reserva una caja en la memoria, con su dirección. Aún sin valor, guarda 0.",
        boxes: [{ ...COUNT, value: "0" }],
        enter: ["count"],
      },
      {
        line: 1,
        caption: "`count = 42` escribe 42 dentro de la misma caja. La dirección no cambia.",
        boxes: [{ ...COUNT, value: "42", was: "0" }],
      },
      {
        line: 2,
        caption: "`other := count` crea otra caja, en otra dirección, y copia dentro el 42.",
        boxes: [{ ...COUNT, value: "42" }, { ...OTHER, value: "42" }],
        enter: ["other"],
      },
      {
        line: 3,
        caption: "`other = 7` solo cambia su propia caja. `count` sigue valiendo 42.",
        boxes: [{ ...COUNT, value: "42" }, { ...OTHER, value: "7", was: "42" }],
      },
      {
        line: 4,
        caption: "Cada variable es su propia caja: `fmt.Println` imprime 42 y 7.",
        boxes: [{ ...COUNT, value: "42" }, { ...OTHER, value: "7" }],
        output: ["42 7"],
        printed: true,
      },
    ],
  },

  convert: {
    code: ["price := 9.75", "euros := int(price)", "fmt.Println(price, euros)"],
    steps: [
      {
        line: 0,
        caption: "`price := 9.75` es un `float64`: guarda la parte entera y los decimales.",
        boxes: [PRICE],
        enter: ["price"],
      },
      {
        line: 1,
        caption: "`int(price)` crea un `int`, que no tiene sitio para decimales: el .75 se pierde. No redondea.",
        boxes: [PRICE, { ...EUROS, value: "9" }],
        enter: ["euros"],
        discard: DECIMALS,
      },
      {
        line: 2,
        caption: "Convertir hace una copia: `price` sigue valiendo 9.75. Se imprime 9.75 y 9.",
        boxes: [PRICE, { ...EUROS, value: "9" }],
        output: ["9.75 9"],
        printed: true,
      },
    ],
  },

  basics: {
    code: ["i := 42", "p := &i", "fmt.Println(*p)", "*p = 21", "fmt.Println(i)", "var q *int", "fmt.Println(q == nil)"],
    steps: [
      {
        line: 0,
        caption: "`i := 42` guarda 42 en una caja de la memoria. La caja tiene una dirección; aquí la abreviamos como 0xc0.",
        boxes: [{ ...I, value: "42" }],
        enter: ["i"],
      },
      {
        line: 1,
        caption: "`p := &i` guarda la dirección de `i`. Ahora `p` apunta a `i`.",
        boxes: [{ ...I, value: "42" }, P],
        arrows: [P_TO_I],
        enter: ["p"],
      },
      {
        line: 2,
        caption: "`*p` sigue la flecha y lee lo que hay en la caja de `i`: 42.",
        boxes: [{ ...I, value: "42" }, P],
        arrows: [P_TO_I],
        follow: P_TO_I,
        output: ["42"],
        printed: true,
      },
      {
        line: 3,
        caption: "`*p = 21` sigue la flecha y escribe 21 en la caja de `i`.",
        boxes: [{ ...I, value: "21", was: "42" }, P],
        arrows: [P_TO_I],
        follow: P_TO_I,
        output: ["42"],
      },
      {
        line: 4,
        caption: "`i` vale 21, aunque nunca has escrito `i = 21`.",
        boxes: [{ ...I, value: "21" }, P],
        arrows: [P_TO_I],
        output: ["42", "21"],
        printed: true,
      },
      {
        line: 5,
        caption: "`var q *int` crea un puntero sin dirección. Vale `nil`: no apunta a nada.",
        boxes: [{ ...I, value: "21" }, P, Q],
        arrows: [P_TO_I],
        enter: ["q"],
        output: ["42", "21"],
      },
      {
        line: 6,
        caption: "`q == nil` es `true`. Seguir `q` con `*q` haría fallar el programa.",
        boxes: [{ ...I, value: "21" }, P, Q],
        arrows: [P_TO_I],
        output: ["42", "21", "true"],
        printed: true,
      },
    ],
  },

  closure: {
    code: [
      "func adder() func(int) int {",
      "\tsum := 0",
      "\treturn func(x int) int {",
      "\t\tsum += x",
      "\t\treturn sum",
      "\t}",
      "}",
      "pos, neg := adder(), adder()",
      "pos(1)",
      "pos(2)",
      "neg(-2)",
    ],
    steps: [
      {
        line: 7,
        caption: "`adder()` crea una variable `sum` que vale 0 y devuelve una función que la usa: `pos`.",
        boxes: [POS, { ...POS_SUM, value: "0" }],
        arrows: [POS_TO_SUM],
        enter: ["pos", "pos-sum"],
      },
      {
        line: 7,
        caption: "La segunda llamada a `adder` crea otra `sum`, nueva. La función `neg` usa esa.",
        boxes: [POS, { ...POS_SUM, value: "0" }, NEG, { ...NEG_SUM, value: "0" }],
        arrows: [POS_TO_SUM, NEG_TO_SUM],
        enter: ["neg", "neg-sum"],
      },
      {
        line: 8,
        caption: "`pos(1)` sigue su flecha y suma 1 a su `sum`: ahora vale 1.",
        boxes: [POS, { ...POS_SUM, value: "1", was: "0" }, NEG, { ...NEG_SUM, value: "0" }],
        arrows: [POS_TO_SUM, NEG_TO_SUM],
        follow: POS_TO_SUM,
      },
      {
        line: 9,
        caption: "`adder` ya terminó, pero su `sum` sigue viva: `pos(2)` la deja en 3.",
        boxes: [POS, { ...POS_SUM, value: "3", was: "1" }, NEG, { ...NEG_SUM, value: "0" }],
        arrows: [POS_TO_SUM, NEG_TO_SUM],
        follow: POS_TO_SUM,
      },
      {
        line: 10,
        caption: "`neg(-2)` cambia solo su propia `sum`. La de `pos` sigue valiendo 3.",
        boxes: [POS, { ...POS_SUM, value: "3" }, NEG, { ...NEG_SUM, value: "-2", was: "0" }],
        arrows: [POS_TO_SUM, NEG_TO_SUM],
        follow: NEG_TO_SUM,
      },
    ],
  },

  receivers: {
    code: ["v := Vertex{3, 4}", "v.ScaleCopy(10)", "fmt.Println(v)", "v.Scale(10)", "fmt.Println(v)"],
    steps: [
      {
        line: 0,
        caption: "`v := Vertex{3, 4}` guarda el vértice en una caja de la memoria.",
        boxes: [{ ...V, value: "{3 4}" }],
        enter: ["v"],
      },
      {
        line: 1,
        caption: "`ScaleCopy` tiene receptor de valor: el método recibe una copia de `v`.",
        scope: { ...SCOPE, title: "ScaleCopy" },
        boxes: [{ ...V, value: "{3 4}" }, { ...METHOD_V, value: "{3 4}" }],
        enter: ["method-v"],
      },
      {
        line: 1,
        caption: "Dentro del método, la copia pasa a valer {30 40}. El original no se entera.",
        scope: { ...SCOPE, title: "ScaleCopy" },
        boxes: [{ ...V, value: "{3 4}" }, { ...METHOD_V, value: "{30 40}", was: "{3 4}" }],
      },
      {
        line: 2,
        caption: "Al terminar el método, la copia se descarta. `fmt.Println(v)` imprime {3 4}.",
        scope: { ...SCOPE, title: "ScaleCopy", gone: true },
        boxes: [{ ...V, value: "{3 4}" }, { ...METHOD_V, value: "{30 40}", gone: true }],
        output: ["{3 4}"],
        printed: true,
      },
      {
        line: 3,
        caption: "`Scale` tiene receptor de puntero: Go le pasa `&v`, la dirección del original.",
        scope: { ...SCOPE, title: "Scale" },
        boxes: [{ ...V, value: "{3 4}" }, { ...METHOD_V, kind: "pointer", value: "0xc0" }],
        arrows: [METHOD_TO_V],
        enter: ["method-v"],
        output: ["{3 4}"],
      },
      {
        line: 3,
        caption: "El método sigue el puntero y cambia el original: `v` pasa a valer {30 40}.",
        scope: { ...SCOPE, title: "Scale" },
        boxes: [{ ...V, value: "{30 40}", was: "{3 4}" }, { ...METHOD_V, kind: "pointer", value: "0xc0" }],
        arrows: [METHOD_TO_V],
        follow: METHOD_TO_V,
        output: ["{3 4}"],
      },
      {
        line: 4,
        caption: "Al terminar `Scale`, el cambio se queda. `fmt.Println(v)` imprime {30 40}.",
        scope: { ...SCOPE, title: "Scale", gone: true },
        boxes: [{ ...V, value: "{30 40}" }, { ...METHOD_V, kind: "pointer", value: "0xc0", gone: true }],
        output: ["{3 4}", "{30 40}"],
        printed: true,
      },
    ],
  },
};

class PointerDemo extends Scene {
  static scenes = scenes;

  // A step plays in two halves: new boxes drop in, then new arrows grow and
  // discarded pieces fall. A step that changes a value or follows an arrow
  // uses all of t.
  draw(paint, step, t) {
    const entering = new Set(step.enter ?? []);
    const settle = Math.min(1, t * 2);
    const grow = Math.max(0, t * 2 - 1);

    if (step.scope) drawScope(paint, step.scope);
    for (const box of step.boxes) {
      const drop = entering.has(box.id) ? lerp(DROP, 0, settle) : 0;
      drawBox(paint, box, drop, t);
    }
    for (const arrow of step.arrows ?? []) {
      paint.arrow(arrow.points, paint.colors.pointer, entering.has(arrow.from) ? grow : 1);
    }
    if (step.follow && t < 1) {
      const [x, y] = pointOnPath(step.follow.points, t);
      paint.dot(x, y, paint.colors.highlight);
    }
    if (step.discard && t < 1) {
      // The piece rides down with its box, then falls on its own.
      const y = step.discard.y - lerp(DROP, 0, settle) + lerp(0, FALL, grow);
      paint.text(step.discard.text, step.discard.x, y, paint.colors.error);
    }
    if (step.output) drawOutput(paint, step.output, step.printed, t);
  }
}

// drawBox draws a variable: its name above, its value inside and, for
// values that something points to, its address below.
function drawBox(paint, box, drop, t) {
  const { colors } = paint;
  const width = box.width ?? BOX_WIDTH;
  const y = box.y - drop;
  const changing = box.was !== undefined;

  let color = box.kind === "pointer" || box.kind === "nil" ? colors.pointer : colors.value;
  let valueColor = box.kind === "pointer" ? colors.pointer : box.kind === "nil" ? colors.error : colors.value;
  if (changing) color = valueColor = colors.highlight;
  if (box.gone) color = valueColor = colors.dim;

  paint.text(box.label, box.x, y - 7, box.gone ? colors.dim : colors.paper);
  paint.frame(box.x, y, width, BOX_HEIGHT, color);
  // The new value lands when the step's motion ends.
  const value = changing && t < 1 ? box.was : box.value;
  paint.centerText(value, box.x, y + 6, width, valueColor);
  if (box.addr) paint.centerText(box.addr, box.x, y + BOX_HEIGHT + 3, width, colors.pointer);
}

// drawScope draws a method call as a dashed frame: what is inside exists
// only while the method runs.
function drawScope(paint, scope) {
  const color = scope.gone ? paint.colors.dim : paint.colors.paper;
  paint.frame(scope.x, scope.y, scope.width, scope.height, color, true);
  paint.text(scope.title, scope.x + 4, scope.y + 4, color);
}

// drawOutput draws what the program has printed so far. When this step
// printed, its line is the last one: it appears when the step's motion
// ends, in the highlight colour.
function drawOutput(paint, lines, printed, t) {
  const { colors } = paint;
  const { x, y, width, height } = OUTPUT;
  paint.frame(x, y, width, height, colors.dim);
  paint.text("salida", x + 4, y + 4, colors.dim);
  lines.forEach((line, i) => {
    const isNew = printed && i === lines.length - 1;
    if (isNew && t < 1) return;
    paint.text(line, x + 4, y + 12 + i * 8, isNew ? colors.highlight : colors.paper);
  });
}

customElements.define("pointer-demo", PointerDemo);
