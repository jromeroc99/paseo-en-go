// <interface-demo> shows an interface value as what it really is: a pair of
// a concrete type and a value. Scene: "pair" (an empty interface is nil; one
// that holds a nil pointer still has a type, so it is not nil).
import { PixelScene, lerp } from "./pixel-scene.js";

const BOX_HEIGHT = 16;
const DROP = 6; // pixels an entering box falls before it settles

// d, the interface value: one box split into a type cell and a value cell.
// The value cell is wider because "GOPHER" is the longest thing it holds.
const D = { id: "d", label: "d", x: 8, y: 20, typeWidth: 36, valueWidth: 48 };

// b, a plain *Box variable, drawn like a pointer box in <pointer-demo>.
const B = { id: "b", label: "b", x: 8, y: 68, width: 40 };

// The output strip, where fmt.Println writes.
const OUTPUT = { x: 104, y: 8, width: 52, height: 32 };

const EMPTY = { type: "nil", value: "nil" };
const TEMP = { type: "Temp", value: "21.5" };
const WORD = { type: "Word", value: '"gopher"' };
const NIL_BOX = { type: "*Box", value: "nil" };

const scenes = {
  pair: {
    code: [
      "var d Describer",
      "fmt.Println(d == nil)",
      "d = Temp(21.5)",
      'd = Word("gopher")',
      "var b *Box",
      "d = b",
      "fmt.Println(d == nil)",
    ],
    steps: [
      {
        line: 0,
        caption: "`var d Describer` crea un valor de interfaz: una pareja de tipo y valor. Aún no guarda nada: (nil, nil).",
        pair: EMPTY,
        enter: ["d"],
      },
      {
        line: 1,
        caption: "Sin tipo ni valor, la interfaz es `nil`. Por eso `d == nil` es `true`.",
        pair: EMPTY,
        output: ["true"],
        printed: true,
      },
      {
        line: 2,
        caption: "`d = Temp(21.5)` guarda la pareja (Temp, 21.5). El tipo dice qué `Describe` se ejecuta.",
        pair: { ...TEMP, was: EMPTY },
        output: ["true"],
      },
      {
        line: 3,
        caption: "`d = Word(\"gopher\")` cambia la pareja entera: ahora es (Word, \"gopher\").",
        pair: { ...WORD, was: TEMP },
        output: ["true"],
      },
      {
        line: 4,
        caption: "`var b *Box` crea un puntero a `Box` que vale `nil`: no apunta a nada.",
        pair: WORD,
        box: true,
        enter: ["b"],
        output: ["true"],
      },
      {
        line: 5,
        caption: "`d = b` guarda la pareja (*Box, nil). El valor es `nil`, pero el tipo ya está.",
        pair: { ...NIL_BOX, was: WORD },
        box: true,
        output: ["true"],
      },
      {
        line: 6,
        caption: "`d == nil` es `false`: la pareja tiene tipo. Una interfaz solo es `nil` sin tipo ni valor.",
        pair: NIL_BOX,
        box: true,
        output: ["true", "false"],
        printed: true,
      },
    ],
  },
};

class InterfaceDemo extends PixelScene {
  static scenes = scenes;

  // An entering box falls into place; a changing pair swaps when t reaches 1.
  draw(paint, step, t) {
    const entering = new Set(step.enter ?? []);
    const dropFor = (id) => (entering.has(id) ? lerp(DROP, 0, t) : 0);

    drawPair(paint, step.pair, dropFor("d"), t);
    if (step.box) drawNilPointer(paint, B, dropFor("b"));
    if (step.output) drawOutput(paint, step.output, step.printed, t);
  }
}

// drawPair draws the interface value: its name above, the type and value
// cells inside, and a label under each cell so the reader knows which is
// which.
function drawPair(paint, pair, drop, t) {
  const { colors } = paint;
  const { x, typeWidth, valueWidth } = D;
  const y = D.y - drop;
  const changing = pair.was !== undefined;
  // The new pair lands when the step's motion ends.
  const shown = changing && t < 1 ? pair.was : pair;

  paint.text(D.label, x, y - 7, colors.paper);
  const frameColor = changing ? colors.highlight : colors.value;
  paint.frame(x, y, typeWidth + valueWidth, BOX_HEIGHT, frameColor);
  paint.line(x + typeWidth, y, x + typeWidth, y + BOX_HEIGHT - 1, frameColor);

  const cellColor = (text, normal) => (changing ? colors.highlight : text === "nil" ? colors.error : normal);
  paint.centerText(shown.type, x, y + 6, typeWidth, cellColor(shown.type, colors.paper));
  paint.centerText(shown.value, x + typeWidth, y + 6, valueWidth, cellColor(shown.value, colors.value));

  paint.centerText("tipo", x, y + BOX_HEIGHT + 3, typeWidth, colors.paper);
  paint.centerText("valor", x + typeWidth, y + BOX_HEIGHT + 3, valueWidth, colors.paper);
}

// drawNilPointer draws a pointer variable that holds nil.
function drawNilPointer(paint, box, drop) {
  const { colors } = paint;
  const y = box.y - drop;
  paint.text(box.label, box.x, y - 7, colors.paper);
  paint.frame(box.x, y, box.width, BOX_HEIGHT, colors.pointer);
  paint.centerText("nil", box.x, y + 6, box.width, colors.error);
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

customElements.define("interface-demo", InterfaceDemo);
