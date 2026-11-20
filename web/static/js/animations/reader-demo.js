// <reader-demo> shows an io.Reader filling a byte slice one chunk at a time.
// Scene: "chunks" (a 13-byte text read through an 8-byte buffer: the second
// Read leaves stale bytes behind, so only b[:n] counts).
import { Scene, lerp } from "./scene.js";

// Source bytes and buffer bytes share one grid, so the first chunk falls
// straight down and the second visibly jumps back to the start of b.
const CELL_X = 9;
const CELL_WIDTH = 10;
const CELL_HEIGHT = 11;
const PITCH = 11; // cell width plus a 1-unit gap
const SOURCE_Y = 3;
const BUFFER_Y = 29;
const DROP = 6; // units an entering row falls before it settles

// n and err, the two results of Read.
const N_BOX = { label: "n", labelX: 8, x: 14, y: 46, width: 13 };
const ERR_BOX = { label: "err", labelX: 34, x: 47, y: 46, width: 17 };

// The output strip, where fmt.Printf writes.
const OUTPUT = { x: 2, y: 60, width: 156, height: 34 };

const TEXT = "Hola, Reader!";
const ZEROS = "\0".repeat(8); // make fills the slice with zero bytes
const FIRST = "Hola, Re";
const SECOND = "ader! Re"; // "ader!" plus three bytes left from the first Read

const PRINTED = [
  'n = 8 err = <nil> b[:n] = "Hola, Re"',
  'n = 5 err = <nil> b[:n] = "ader!"',
  'n = 0 err = EOF b[:n] = ""',
];

const scenes = {
  chunks: {
    code: [
      'r := strings.NewReader("Hola, Reader!")',
      "b := make([]byte, 8)",
      "for {",
      "\tn, err := r.Read(b)",
      '\tfmt.Printf("n = %v err = %v b[:n] = %q\\n", n, err, b[:n])',
      "\tif err == io.EOF {",
      "\t\tbreak",
      "\t}",
      "}",
    ],
    steps: [
      {
        line: 0,
        caption: "`strings.NewReader` convierte el texto en un `Reader` de 13 bytes. La flecha marca por dónde va la lectura.",
        offset: { value: 0 },
        enter: ["source"],
        output: [],
      },
      {
        line: 1,
        caption: "`make([]byte, 8)` crea un `slice` de 8 bytes, todos a cero. Ahí dejará `Read` lo que lea.",
        offset: { value: 0 },
        buffer: { value: ZEROS },
        enter: ["buffer"],
        output: [],
      },
      {
        line: 3,
        caption: "El primer `Read` copia 8 bytes en `b`, todo lo que cabe. Devuelve `n = 8` y `err = nil`.",
        offset: { value: 8, was: 0 },
        buffer: { value: FIRST, was: ZEROS },
        read: { from: 0, count: 8 },
        window: { value: 8, was: null },
        n: { value: "8", was: null },
        err: { value: "nil", was: null },
        output: [],
      },
      {
        line: 4,
        caption: "`b[:n]` son justo los bytes que ha escrito `Read`. Se imprime `\"Hola, Re\"`.",
        offset: { value: 8 },
        buffer: { value: FIRST },
        window: { value: 8 },
        n: { value: "8" },
        err: { value: "nil" },
        output: PRINTED.slice(0, 1),
        printed: true,
      },
      {
        line: 3,
        caption: "El segundo `Read` solo encuentra 5 bytes y sobrescribe el principio de `b`. Los 3 últimos son restos de antes.",
        offset: { value: 13, was: 8 },
        buffer: { value: SECOND, was: FIRST },
        read: { from: 8, count: 5 },
        window: { value: 5, was: 8 },
        n: { value: "5", was: "8" },
        err: { value: "nil" },
        output: PRINTED.slice(0, 1),
      },
      {
        line: 4,
        caption: "Por eso usas `b[:n]`: se imprime `\"ader!\"`. Con `b` entero saldría `\"ader! Re\"`.",
        offset: { value: 13 },
        buffer: { value: SECOND },
        window: { value: 5 },
        n: { value: "5" },
        err: { value: "nil" },
        output: PRINTED.slice(0, 2),
        printed: true,
      },
      {
        line: 3,
        caption: "El tercer `Read` ya no encuentra nada: `n = 0` y `err = EOF`. `b` no cambia.",
        offset: { value: 13 },
        buffer: { value: SECOND },
        window: { value: 0, was: 5 },
        n: { value: "0", was: "5" },
        err: { value: "EOF", was: "nil" },
        output: PRINTED.slice(0, 2),
      },
      {
        line: 4,
        caption: "Con `n = 0`, `b[:n]` está vacío: se imprime `\"\"`.",
        offset: { value: 13 },
        buffer: { value: SECOND },
        window: { value: 0 },
        n: { value: "0" },
        err: { value: "EOF" },
        output: PRINTED,
        printed: true,
      },
      {
        line: 6,
        caption: "`err == io.EOF`: no quedan datos. No es un fallo, así que el bucle termina con `break`.",
        offset: { value: 13 },
        buffer: { value: SECOND },
        window: { value: 0 },
        n: { value: "0" },
        err: { value: "EOF" },
        output: PRINTED,
      },
    ],
  },
};

class ReaderDemo extends Scene {
  static scenes = scenes;

  // A Read step moves its chunk from the text into b while t runs; the
  // results (bytes, n, err and the b[:n] window) land together at t = 1.
  // A step that adds a row lets that row drop in instead.
  draw(paint, step, t) {
    const entering = new Set(step.enter ?? []);
    const drop = (row) => (entering.has(row) ? lerp(DROP, 0, t) : 0);

    drawSource(paint, step, t, drop("source"));
    if (step.buffer) drawBuffer(paint, step, t, drop("buffer"));
    if (step.read && t < 1) drawChunk(paint, step.read, t);
    drawField(paint, N_BOX, step.n, t);
    drawField(paint, ERR_BOX, step.err, t);
    drawOutput(paint, step.output, step.printed, t);
  }
}

// at returns what a changing field shows: its old value until the step's
// motion ends, then the new one.
function at(field, t) {
  if (!field) return null;
  return field.was !== undefined && t < 1 ? field.was : field.value;
}

function cellX(i) {
  return CELL_X + i * PITCH;
}

// drawSource draws the text as one cell per byte. Bytes already read are
// dim, and the arrow below points at the next byte Read will take.
function drawSource(paint, step, t, drop) {
  const { colors } = paint;
  const y = SOURCE_Y - drop;
  const offset = at(step.offset, t);
  const { read } = step;

  paint.text("r", 1, y + 3, colors.paper);
  [...TEXT].forEach((char, i) => {
    let color = i < offset ? colors.dim : colors.value;
    if (read && t < 1 && i >= read.from && i < read.from + read.count) color = colors.highlight;
    drawCell(paint, char, cellX(i), y, color);
  });

  const x = cellX(offset) + Math.floor(CELL_WIDTH / 2);
  paint.arrow([[x, y + 19], [x, y + CELL_HEIGHT + 1]], colors.pointer);
}

// drawBuffer draws b. Bytes inside the b[:n] window are the ones that
// count; the rest are stale (dim), whatever they hold.
function drawBuffer(paint, step, t, drop) {
  const { colors } = paint;
  const y = BUFFER_Y - drop;
  const bytes = at(step.buffer, t);
  const window = at(step.window, t);
  const landed = step.read && t === 1 ? step.read.count : 0;

  paint.text("b", 1, y + 3, colors.paper);
  [...bytes].forEach((char, i) => {
    let color = colors.dim;
    if (window !== null && i < window) color = colors.value;
    if (i < landed) color = colors.highlight;
    drawCell(paint, char, cellX(i), y, color);
  });

  if (window === null) return;
  const left = cellX(0) - 2;
  paint.frame(left, y - 2, window * PITCH + 3, CELL_HEIGHT + 4, colors.pointer);
  paint.text("b[:n]", left, y - 9, colors.pointer);
}

// drawChunk draws the bytes a Read is copying, on their way from the text
// to the start of b.
function drawChunk(paint, { from, count }, t) {
  for (let k = 0; k < count; k++) {
    const x = lerp(cellX(from + k), cellX(k), t);
    const y = lerp(SOURCE_Y, BUFFER_Y, t);
    drawCell(paint, TEXT[from + k], x, y, paint.colors.highlight);
  }
}

// drawCell draws one byte. A zero byte is an empty cell; a space gets a
// visible mark, or it would look like a missing byte.
function drawCell(paint, char, x, y, color) {
  paint.frame(x, y, CELL_WIDTH, CELL_HEIGHT, color);
  if (char === "\0") return;
  if (char === " ") {
    drawSpaceMark(paint, x + CELL_WIDTH / 2, y + 7, color);
    return;
  }
  paint.centerText(char, x, y + (CELL_HEIGHT - 5) / 2, CELL_WIDTH, color);
}

// drawSpaceMark draws the open box (␣) that marks a space, centred on x.
function drawSpaceMark(paint, x, y, color) {
  paint.line(x - 2, y, x - 2, y + 1.5, color);
  paint.line(x - 2, y + 1.5, x + 2, y + 1.5, color);
  paint.line(x + 2, y + 1.5, x + 2, y, color);
}

// drawField draws one result of Read as a labelled box. The box lights up
// in the step that changes it; the value keeps the colour of its meaning.
function drawField(paint, box, field, t) {
  const value = at(field, t);
  if (value === null) return;
  const { colors } = paint;
  const changing = field.was !== undefined;
  const color = value === "nil" || value === "EOF" ? colors.error : colors.value;

  paint.text(box.label, box.labelX, box.y + 3, colors.paper);
  paint.frame(box.x, box.y, box.width, CELL_HEIGHT, changing ? colors.highlight : colors.value);
  paint.centerText(value, box.x, box.y + 3, box.width, color);
}

// drawOutput draws what the program has printed so far. When this step
// printed, its line is the last one: it appears when the step's motion
// ends, in the highlight colour.
function drawOutput(paint, lines, printed, t) {
  const { colors } = paint;
  const { x, y, width, height } = OUTPUT;
  paint.frame(x, y, width, height, colors.dim);
  paint.text("salida", x + 4, y + 3, colors.dim);
  lines.forEach((line, i) => {
    const isNew = printed && i === lines.length - 1;
    if (isNew && t < 1) return;
    paint.text(line, x + 4, y + 11 + i * 8, isNew ? colors.highlight : colors.paper);
  });
}

customElements.define("reader-demo", ReaderDemo);
