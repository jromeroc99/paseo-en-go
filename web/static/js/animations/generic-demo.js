// <generic-demo> shows what the compiler does with a type parameter.
// Scenes: "instantiate" (one generic function is filled with int, then with
// string: two functions from one) and "constraints" (each type tries to
// pass the constraint; []int does not satisfy comparable).
import { Scene, lerp } from "./scene.js";

const DROP = 6; // units an entering box falls before it settles

const OUTPUT = { x: 4, y: 62, width: 152, height: 30 };

// instantiate: the generic function on the left, with a slot for T, and
// the functions the compiler writes from it on the right.
const MOLD = { x: 6, y: 8, width: 60, height: 44 };
const SLOT = { x: 12, y: 30, width: 48, height: 14 };
const INSTANCE = { x: 92, width: 64, height: 14 };
const INT_FUNC = { id: "int", name: "Index[int]", y: 12 };
const STRING_FUNC = { id: "string", name: "Index[string]", y: 34 };

// constraints: the types wait on the left, the gate of the constraint
// stands in the middle and the types that pass it line up on the right.
const GATE = { x: 55, y: 4, width: 50, height: 54 };
const CHIP = { width: 30, height: 9, waitX: 4, stopX: 24, passX: 124 };
const TYPES = [
  { id: "int", y: 17 },
  { id: "string", y: 28 },
  { id: "bool", y: 39 },
  { id: "[]int", y: 50 },
];

const scenes = {
  instantiate: {
    code: [
      "func Index[T comparable](s []T, x T) int",
      "fmt.Println(Index([]int{10, 20, 30}, 20))",
      'fmt.Println(Index([]string{"a", "b", "c"}, "c"))',
    ],
    steps: [
      {
        line: 0,
        caption: "`Index` es un molde: `T` es un hueco para cualquier tipo que cumpla `comparable`.",
        funcs: [],
      },
      {
        line: 1,
        caption: "Le pasas un `[]int`, así que el compilador deduce que `T` es `int`.",
        type: "int",
        enterType: true,
        funcs: [],
      },
      {
        line: 1,
        caption: "Con `T = int` escribe una función nueva, `Index[int]`, y la ejecuta: imprime 1.",
        type: "int",
        funcs: [INT_FUNC],
        enter: "int",
        output: ["1"],
        printed: true,
      },
      {
        line: 2,
        caption: "Ahora le pasas un `[]string`: esta vez `T` es `string`.",
        type: "string",
        enterType: true,
        funcs: [INT_FUNC],
        output: ["1"],
      },
      {
        line: 2,
        caption: "Sale otra función, `Index[string]`, que imprime 2. Un solo molde, dos funciones.",
        type: "string",
        funcs: [INT_FUNC, STRING_FUNC],
        enter: "string",
        output: ["1", "2"],
        printed: true,
      },
    ],
  },

  constraints: {
    code: [
      "func Index[T comparable](s []T, x T) int",
      "Index([]int{1, 2}, 2)",
      'Index([]string{"a"}, "a")',
      "Index([]bool{true}, true)",
      "Index([][]int{{1}}, []int{1})",
    ],
    steps: [
      {
        line: 0,
        caption: "`comparable` es una restricción: solo deja pasar los tipos que se comparan con `==`.",
        passed: [],
      },
      {
        line: 1,
        caption: "Dos `int` se comparan con `==`: `int` cumple `comparable`.",
        passed: ["int"],
        moving: "int",
      },
      {
        line: 2,
        caption: "Dos `string` también se comparan: `string` cumple `comparable`.",
        passed: ["int", "string"],
        moving: "string",
      },
      {
        line: 3,
        caption: "Y dos `bool`, también: `bool` cumple `comparable`.",
        passed: ["int", "string", "bool"],
        moving: "bool",
      },
      {
        line: 4,
        caption: "Dos slices no se comparan con `==`: `[]int` no cumple `comparable` y no compila.",
        passed: ["int", "string", "bool"],
        moving: "[]int",
        rejected: "[]int",
        output: ["[]int does not satisfy comparable"],
        printed: true,
        compileError: true,
      },
    ],
  },
};

class GenericDemo extends Scene {
  static scenes = scenes;

  draw(paint, step, t) {
    const name = this.getAttribute("scene");
    const drawScene = { instantiate: drawInstantiate, constraints: drawConstraints }[name];
    drawScene(paint, step, t);
    if (step.output) drawOutput(paint, step, t);
  }
}

// drawInstantiate plays a step in two halves: the type drops into the slot,
// or the new function drops in, and then its arrow from the mould grows.
function drawInstantiate(paint, step, t) {
  const { colors } = paint;
  const settle = Math.min(1, t * 2);
  const grow = Math.max(0, t * 2 - 1);

  paint.frame(MOLD.x, MOLD.y, MOLD.width, MOLD.height, colors.paper);
  paint.centerText("Index[T]", MOLD.x, MOLD.y + 5, MOLD.width, colors.paper);
  paint.centerText("comparable", MOLD.x, MOLD.y + 14, MOLD.width, colors.dim);
  drawSlot(paint, step, settle);

  for (const func of step.funcs) {
    const entering = func.id === step.enter;
    const y = func.y - (entering ? lerp(DROP, 0, settle) : 0);
    const color = entering ? colors.highlight : colors.value;
    paint.frame(INSTANCE.x, y, INSTANCE.width, INSTANCE.height, color);
    paint.centerText(func.name, INSTANCE.x, y + 4.5, INSTANCE.width, color);
    const tip = [INSTANCE.x - 1, func.y + INSTANCE.height / 2];
    paint.arrow([[MOLD.x + MOLD.width, MOLD.y + MOLD.height / 2], tip], colors.pointer, entering ? grow : 1);
  }
}

// drawSlot draws the hole for T: dashed and empty until the compiler knows
// the type, then filled with it.
function drawSlot(paint, step, settle) {
  const { colors } = paint;
  if (!step.type) {
    paint.frame(SLOT.x, SLOT.y, SLOT.width, SLOT.height, colors.dim, true);
    paint.centerText("T = ?", SLOT.x, SLOT.y + 4.5, SLOT.width, colors.dim);
    return;
  }
  const y = SLOT.y - (step.enterType ? lerp(DROP, 0, settle) : 0);
  const color = step.enterType ? colors.highlight : colors.value;
  paint.frame(SLOT.x, y, SLOT.width, SLOT.height, color);
  paint.centerText(`T = ${step.type}`, SLOT.x, y + 4.5, SLOT.width, color);
}

// drawConstraints draws the gate and every type. The moving type slides
// from the queue through the gate, or up to it and no further when the
// constraint rejects it.
function drawConstraints(paint, step, t) {
  const { colors } = paint;
  paint.frame(GATE.x, GATE.y, GATE.width, GATE.height, colors.paper, true);
  paint.centerText("comparable", GATE.x, GATE.y + 3, GATE.width, colors.paper);

  for (const type of TYPES) {
    const moving = type.id === step.moving;
    const passed = step.passed.includes(type.id);
    const rejected = type.id === step.rejected;
    let x = CHIP.waitX;
    if (passed) x = moving ? lerp(CHIP.waitX, CHIP.passX, t) : CHIP.passX;
    if (rejected) x = lerp(CHIP.waitX, CHIP.stopX, t);

    let color = colors.value;
    if (passed && !moving) color = colors.ok;
    if (moving) color = t < 1 ? colors.highlight : passed ? colors.ok : colors.error;
    paint.frame(x, type.y, CHIP.width, CHIP.height, color);
    paint.centerText(type.id, x, type.y + 2, CHIP.width, color);
  }
}

// drawOutput draws what the program has printed, or what the compiler
// answered when the program does not compile. When this step printed, its
// line is the last one: it appears when the step's motion ends, in the
// highlight colour, or in the error colour for a compiler error.
function drawOutput(paint, step, t) {
  const { colors } = paint;
  const { x, y, width, height } = OUTPUT;
  const newColor = step.compileError ? colors.error : colors.highlight;
  paint.frame(x, y, width, height, colors.dim);
  paint.text(step.compileError ? "compilador" : "salida", x + 4, y + 4, colors.dim);
  step.output.forEach((line, i) => {
    const isNew = step.printed && i === step.output.length - 1;
    if (isNew && t < 1) return;
    paint.text(line, x + 4, y + 12 + i * 8, isNew ? newColor : colors.paper);
  });
}

customElements.define("generic-demo", GenericDemo);
