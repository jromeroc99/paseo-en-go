// <slice-demo> shows a slice as a window over an array. Scenes: "view" (two
// slices share one array), "len-cap" (a window has a length and a capacity)
// and "append" (append moves to a bigger array when the current one is full).
import { PixelScene, lerp } from "./pixel-scene.js";

const DROP = 6; // pixels an entering array falls before it settles

// The output strip, where the program prints. It fits four lines.
const OUTPUT = { x: 4, y: 52, width: 152, height: 44 };

// view: the array names, with window a above it and window b below it.
const NAMES = { x: 38, y: 17, cellWidth: 28, cellHeight: 16 };
const A = { id: "a", label: "a", from: 0, to: 2, y: 4, side: "above" };
const B = { id: "b", label: "b", from: 1, to: 3, y: 40, side: "below" };
const NAMES_BEFORE = ["Ana", "Luis", "Marta", "Pablo"];
const NAMES_AFTER = ["Ana", "XXX", "Marta", "Pablo"];

// len-cap: one array; the frame around its cells is the slice s.
const PRIMES = { x: 22, y: 18, cellWidth: 20, cellHeight: 16 };
const PRIMES_VALUES = ["2", "3", "5", "7", "11", "13"];

// append: the slice s on the left and, on the right, one row for each array
// that append creates.
const HEADER = { x: 8, y: 14, width: 28, height: 22 };
const ROW = { x: 52, cellWidth: 12, cellHeight: 13 };
const ARRAYS = [
  { y: 3, cap: 1 },
  { y: 19, cap: 2 },
  { y: 35, cap: 6 },
];

const scenes = {
  view: {
    code: [
      'names := [4]string{"Ana", "Luis", "Marta", "Pablo"}',
      "a := names[0:2]",
      "b := names[1:3]",
      "fmt.Println(a, b)",
      'b[0] = "XXX"',
      "fmt.Println(a, b)",
      "fmt.Println(names)",
    ],
    steps: [
      {
        line: 0,
        caption: "`names` es un array de 4 cadenas. Ocupa 4 casillas seguidas en la memoria.",
        values: NAMES_BEFORE,
        enter: ["names"],
      },
      {
        line: 1,
        caption: "`a := names[0:2]` es una ventana sobre las casillas 0 y 1. No copia nada.",
        values: NAMES_BEFORE,
        windows: [A],
        enter: ["a"],
      },
      {
        line: 2,
        caption: "`b := names[1:3]` es otra ventana, sobre las casillas 1 y 2. La casilla 1 está en las dos.",
        values: NAMES_BEFORE,
        windows: [A, B],
        enter: ["b"],
      },
      {
        line: 3,
        caption: "Cada slice muestra lo que hay en su trozo del array.",
        values: NAMES_BEFORE,
        windows: [A, B],
        output: ["[Ana Luis] [Luis Marta]"],
        printed: true,
      },
      {
        line: 4,
        caption: '`b[0] = "XXX"` escribe en la casilla 1 del array. Es la primera de `b` y también la segunda de `a`.',
        values: NAMES_AFTER,
        change: { index: 1, was: "Luis" },
        windows: [A, B],
        output: ["[Ana Luis] [Luis Marta]"],
      },
      {
        line: 5,
        caption: "Los dos slices ven el cambio, porque miran la misma casilla.",
        values: NAMES_AFTER,
        windows: [A, B],
        output: ["[Ana Luis] [Luis Marta]", "[Ana XXX] [XXX Marta]"],
        printed: true,
      },
      {
        line: 6,
        caption: "El array también ha cambiado: los slices no guardan datos, solo miran el array.",
        values: NAMES_AFTER,
        windows: [A, B],
        output: ["[Ana Luis] [Luis Marta]", "[Ana XXX] [XXX Marta]", "[Ana XXX Marta Pablo]"],
        printed: true,
      },
    ],
  },

  "len-cap": {
    code: [
      "s := []int{2, 3, 5, 7, 11, 13}",
      "printSlice(s)",
      "s = s[:0]",
      "printSlice(s)",
      "s = s[:4]",
      "printSlice(s)",
      "s = s[2:]",
      "printSlice(s)",
    ],
    steps: [
      {
        line: 0,
        caption: "El literal crea un array de 6 elementos y un slice `s` que lo cubre entero.",
        window: { start: 0, len: 6, cap: 6 },
        enter: true,
      },
      {
        line: 1,
        caption: "`len(s)` es 6: lo que ves. `cap(s)` también es 6: el array no tiene más sitio.",
        window: { start: 0, len: 6, cap: 6 },
        output: ["len=6 cap=6 [2 3 5 7 11 13]"],
        printed: true,
      },
      {
        line: 2,
        caption: "`s[:0]` cierra la ventana: la longitud pasa a 0. El array sigue ahí, con sus 6 elementos.",
        window: { start: 0, len: 0, cap: 6 },
        was: { start: 0, len: 6, cap: 6 },
        output: ["len=6 cap=6 [2 3 5 7 11 13]"],
      },
      {
        line: 3,
        caption: "Ahora `s` no muestra nada, pero su capacidad sigue siendo 6.",
        window: { start: 0, len: 0, cap: 6 },
        output: ["len=6 cap=6 [2 3 5 7 11 13]", "len=0 cap=6 []"],
        printed: true,
      },
      {
        line: 4,
        caption: "`s[:4]` vuelve a abrir la ventana hasta 4. Puedes hacerlo porque 4 no pasa de la capacidad.",
        window: { start: 0, len: 4, cap: 6 },
        was: { start: 0, len: 0, cap: 6 },
        output: ["len=6 cap=6 [2 3 5 7 11 13]", "len=0 cap=6 []"],
      },
      {
        line: 5,
        caption: "Los valores no se habían perdido: `s` vuelve a ver 2, 3, 5 y 7.",
        window: { start: 0, len: 4, cap: 6 },
        output: ["len=6 cap=6 [2 3 5 7 11 13]", "len=0 cap=6 []", "len=4 cap=6 [2 3 5 7]"],
        printed: true,
      },
      {
        line: 6,
        caption: "`s[2:]` mueve el inicio de la ventana dos casillas. El final no se mueve, así que la capacidad baja a 4.",
        window: { start: 2, len: 2, cap: 4 },
        was: { start: 0, len: 4, cap: 6 },
        output: ["len=6 cap=6 [2 3 5 7 11 13]", "len=0 cap=6 []", "len=4 cap=6 [2 3 5 7]"],
      },
      {
        line: 7,
        caption: "Los dos primeros elementos quedan fuera: desde `s` ya no puedes volver a ellos.",
        window: { start: 2, len: 2, cap: 4 },
        output: ["len=6 cap=6 [2 3 5 7 11 13]", "len=0 cap=6 []", "len=4 cap=6 [2 3 5 7]", "len=2 cap=4 [5 7]"],
        printed: true,
      },
    ],
  },

  append: {
    code: [
      "var s []int",
      "printSlice(s)",
      "s = append(s, 0)",
      "printSlice(s)",
      "s = append(s, 1)",
      "printSlice(s)",
      "s = append(s, 2, 3, 4)",
      "printSlice(s)",
    ],
    steps: [
      {
        line: 0,
        caption: "`var s []int` crea un slice `nil`: no tiene array debajo, y su longitud y su capacidad son 0.",
        header: { len: 0, cap: 0 },
        enter: true,
      },
      {
        line: 1,
        caption: "Un slice `nil` se comporta como uno vacío: `printSlice` muestra `[]`.",
        header: { len: 0, cap: 0 },
        output: ["len=0 cap=0 []"],
        printed: true,
      },
      {
        line: 2,
        caption: "No hay sitio para el 0. `append` crea un array de capacidad 1, escribe el 0 y `s` pasa a apuntar a él.",
        header: { len: 1, cap: 1, was: { len: 0, cap: 0 } },
        arrays: [{ row: 0, values: ["0"], lands: [0], enter: true }],
        target: 0,
        retarget: true,
        output: ["len=0 cap=0 []"],
      },
      {
        line: 3,
        caption: "Ahora `s` tiene longitud 1 y capacidad 1: el array está lleno.",
        header: { len: 1, cap: 1 },
        arrays: [{ row: 0, values: ["0"] }],
        target: 0,
        output: ["len=0 cap=0 []", "len=1 cap=1 [0]"],
        printed: true,
      },
      {
        line: 4,
        caption: "El 1 no cabe. `append` reserva otro array, de capacidad 2, y copia el 0.",
        header: { len: 1, cap: 1 },
        arrays: [
          { row: 0, values: ["0"] },
          { row: 1, values: ["0"], copyFrom: 0, enter: true },
        ],
        target: 0,
        output: ["len=0 cap=0 []", "len=1 cap=1 [0]"],
      },
      {
        line: 4,
        caption: "Escribe el 1 en el array nuevo y `s` pasa a apuntar a él. El viejo ya no se usa.",
        header: { len: 2, cap: 2, was: { len: 1, cap: 1 } },
        arrays: [
          { row: 0, values: ["0"], gone: true },
          { row: 1, values: ["0", "1"], lands: [1] },
        ],
        target: 1,
        retarget: true,
        output: ["len=0 cap=0 []", "len=1 cap=1 [0]"],
      },
      {
        line: 5,
        caption: "Longitud 2 y capacidad 2: otra vez lleno.",
        header: { len: 2, cap: 2 },
        arrays: [
          { row: 0, values: ["0"], gone: true },
          { row: 1, values: ["0", "1"] },
        ],
        target: 1,
        output: ["len=0 cap=0 []", "len=1 cap=1 [0]", "len=2 cap=2 [0 1]"],
        printed: true,
      },
      {
        line: 6,
        caption: "No caben tres más. `append` reserva un array de capacidad 6 y copia el 0 y el 1.",
        header: { len: 2, cap: 2 },
        arrays: [
          { row: 0, values: ["0"], gone: true },
          { row: 1, values: ["0", "1"] },
          { row: 2, values: ["0", "1"], copyFrom: 1, enter: true },
        ],
        target: 1,
        output: ["len=0 cap=0 []", "len=1 cap=1 [0]", "len=2 cap=2 [0 1]"],
      },
      {
        line: 6,
        caption: "Escribe 2, 3 y 4, y `s` pasa a apuntar al array nuevo. Los arrays viejos quedan sin usar.",
        header: { len: 5, cap: 6, was: { len: 2, cap: 2 } },
        arrays: [
          { row: 0, values: ["0"], gone: true },
          { row: 1, values: ["0", "1"], gone: true },
          { row: 2, values: ["0", "1", "2", "3", "4"], lands: [2, 3, 4] },
        ],
        target: 2,
        retarget: true,
        output: ["len=0 cap=0 []", "len=1 cap=1 [0]", "len=2 cap=2 [0 1]"],
      },
      {
        line: 7,
        caption: "Longitud 5 y capacidad 6: queda una casilla libre para el próximo `append`.",
        header: { len: 5, cap: 6 },
        arrays: [
          { row: 0, values: ["0"], gone: true },
          { row: 1, values: ["0", "1"], gone: true },
          { row: 2, values: ["0", "1", "2", "3", "4"] },
        ],
        target: 2,
        output: ["len=0 cap=0 []", "len=1 cap=1 [0]", "len=2 cap=2 [0 1]", "len=5 cap=6 [0 1 2 3 4]"],
        printed: true,
      },
    ],
  },
};

class SliceDemo extends PixelScene {
  static scenes = scenes;

  // The three scenes share cells and output, but each has its own layout.
  draw(paint, step, t) {
    const drawScene = { view: drawView, "len-cap": drawLenCap, append: drawAppend }[this.getAttribute("scene")];
    drawScene(paint, step, t);
    if (step.output) drawOutput(paint, step.output, step.printed, t);
  }
}

// drawView draws the array names and the windows a and b over it. A cell
// that changes shows its old value until the step's motion ends.
function drawView(paint, step, t) {
  const { colors } = paint;
  const entering = new Set(step.enter ?? []);
  const y = NAMES.y - (entering.has("names") ? lerp(DROP, 0, t) : 0);

  paint.text("names", 14, y + 6, colors.paper);
  step.values.forEach((_, i) => {
    paint.centerText(String(i), cellX(NAMES, i), y - 7, NAMES.cellWidth, colors.paper);
  });
  const cells = step.values.map((text, i) => {
    if (step.change?.index !== i) return { text, kind: "value" };
    return { text: t < 1 ? step.change.was : text, kind: "highlight" };
  });
  drawCells(paint, NAMES, y, cells);

  for (const window of step.windows ?? []) {
    drawWindow(paint, window, entering.has(window.id) ? t : 1);
  }
}

// drawWindow draws a slice of names as a bracket whose ticks point at the
// cells it covers. While it enters, it grows from left to right.
function drawWindow(paint, window, grow) {
  const { colors } = paint;
  const left = cellX(NAMES, window.from) + 2;
  const right = lerp(left, cellX(NAMES, window.to) - 2, grow);
  const toward = window.side === "above" ? 1 : -1;

  paint.text(window.label, left - 6, window.y - 2, colors.paper);
  paint.line(left, window.y, right, window.y, colors.pointer);
  paint.line(left, window.y, left, window.y + 3 * toward, colors.pointer);
  paint.line(right, window.y, right, window.y + 3 * toward, colors.pointer);
}

// drawLenCap draws the slice s as a frame around the cells it shows (its
// length) and a bracket below up to where it could grow (its capacity).
// The edges slide from the previous window to the new one; cells light up
// as the frame passes over them.
function drawLenCap(paint, step, t) {
  const { colors } = paint;
  const { window, was = window } = step;
  const y = PRIMES.y - (step.enter ? lerp(DROP, 0, t) : 0);
  const left = lerp(cellX(PRIMES, was.start), cellX(PRIMES, window.start), t);
  const right = lerp(cellX(PRIMES, was.start + was.len), cellX(PRIMES, window.start + window.len), t);
  const capRight = lerp(cellX(PRIMES, was.start + was.cap), cellX(PRIMES, window.start + window.cap), t);

  const cells = PRIMES_VALUES.map((text, i) => {
    const visible = cellX(PRIMES, i) >= left && cellX(PRIMES, i + 1) <= right;
    return { text, kind: visible ? "value" : "dim" };
  });
  drawCells(paint, PRIMES, y, cells);
  paint.frame(left - 2, y - 2, right - left + 5, PRIMES.cellHeight + 4, colors.pointer);

  const capY = y + PRIMES.cellHeight + 8;
  paint.line(left, capY, capRight, capY, colors.pointer);
  paint.line(left, capY - 3, left, capY, colors.pointer);
  paint.line(capRight, capY - 3, capRight, capY, colors.pointer);

  // The new numbers land when the motion ends.
  const len = t < 1 ? was.len : window.len;
  const cap = t < 1 ? was.cap : window.cap;
  const lenColor = was.len !== window.len ? colors.highlight : colors.paper;
  const capColor = was.cap !== window.cap ? colors.highlight : colors.paper;
  paint.centerText(`len ${len}`, left - 2, y - 10, right - left + 5, lenColor);
  paint.centerText(`cap ${cap}`, left, capY + 3, capRight - left + 1, capColor);
}

// drawAppend draws the slice s and the arrays append has created. A step
// plays in two halves: a new array drops in, then the copied values slide
// into it and the arrow from s grows to its new array.
function drawAppend(paint, step, t) {
  const { colors } = paint;
  const settle = Math.min(1, t * 2);
  const slide = Math.max(0, t * 2 - 1);

  drawHeader(paint, step.header, step.enter ? lerp(DROP, 0, t) : 0, t);
  for (const array of step.arrays ?? []) drawArray(paint, array, settle, slide, t);

  if (step.target === undefined) {
    paint.text("nil", HEADER.x + HEADER.width + 4, HEADER.y + 9, colors.error);
    return;
  }
  paint.arrow(arrowTo(step.target), colors.pointer, step.retarget ? slide : 1);
}

// drawHeader draws s as the box Go keeps for every slice: an arrow to its
// array, its length and its capacity.
function drawHeader(paint, header, drop, t) {
  const { colors } = paint;
  const { x, width, height } = HEADER;
  const y = HEADER.y - drop;
  const was = header.was ?? header;
  const color = header.was ? colors.highlight : colors.paper;

  paint.text("s", x, y - 7, colors.paper);
  paint.frame(x, y, width, height, colors.pointer);
  paint.centerText(`len ${t < 1 ? was.len : header.len}`, x, y + 5, width, color);
  paint.centerText(`cap ${t < 1 ? was.cap : header.cap}`, x, y + 13, width, color);
}

// drawArray draws one of append's arrays. Cells past the values are memory
// the array has but the slice does not use yet. Copied values slide down
// from the array they come from; written values appear when the motion ends.
function drawArray(paint, array, settle, slide, t) {
  const { colors } = paint;
  const { cap } = ARRAYS[array.row];
  const y = ARRAYS[array.row].y - (array.enter ? lerp(DROP, 0, settle) : 0);
  const lands = new Set(array.lands ?? []);
  const copying = array.copyFrom !== undefined;

  const cells = [];
  for (let i = 0; i < cap; i++) {
    const text = array.values[i] ?? "";
    if (array.gone) cells.push({ text, kind: "dim" });
    else if (text === "") cells.push({ text, kind: "dim" });
    else if (lands.has(i)) cells.push({ text: t < 1 ? "" : text, kind: "highlight" });
    else if (copying) cells.push({ text: t < 1 ? "" : text, kind: "highlight" });
    else cells.push({ text, kind: "value" });
  }
  drawCells(paint, ROW, y, cells);
  paint.text(`cap ${cap}`, cellX(ROW, cap) + 5, y + 4, array.gone ? colors.dim : colors.paper);

  if (copying && slide > 0 && t < 1) {
    const textY = lerp(ARRAYS[array.copyFrom].y, y, slide) + textOffset(ROW);
    array.values.forEach((text, i) => paint.centerText(text, cellX(ROW, i), textY, ROW.cellWidth, colors.highlight));
  }
}

// arrowTo returns the path from s to the middle of an array's row.
function arrowTo(row) {
  const startX = HEADER.x + HEADER.width;
  const startY = HEADER.y + HEADER.height / 2;
  const endY = ARRAYS[row].y + Math.floor(ROW.cellHeight / 2);
  const tipX = ROW.x - 2;
  if (endY === startY) return [[startX, startY], [tipX, startY]];
  const turnX = startX + 7;
  return [[startX, startY], [turnX, startY], [turnX, endY], [tipX, endY]];
}

// Brighter cells are drawn last, so the borders they share with their
// neighbours keep the brighter colour.
const LAYERS = ["dim", "value", "highlight"];

// drawCells draws an array as a row of cells that share their borders.
// Each cell is { text, kind }, where kind names its colour.
function drawCells(paint, row, y, cells) {
  const order = [...cells.keys()].sort((i, j) => LAYERS.indexOf(cells[i].kind) - LAYERS.indexOf(cells[j].kind));
  for (const i of order) {
    const color = paint.colors[cells[i].kind];
    paint.frame(cellX(row, i), y, row.cellWidth, row.cellHeight, color);
    paint.centerText(cells[i].text, cellX(row, i), y + textOffset(row), row.cellWidth, color);
  }
}

// cellX returns the left edge of cell i; neighbours share one border pixel.
function cellX(row, i) {
  return row.x + i * (row.cellWidth - 1);
}

// textOffset centres a 5-pixel glyph inside a cell's border.
function textOffset(row) {
  return Math.floor((row.cellHeight - 4) / 2);
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

customElements.define("slice-demo", SliceDemo);
