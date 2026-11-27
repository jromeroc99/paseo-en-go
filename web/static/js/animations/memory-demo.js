// <memory-demo> shows memory as a row of bytes. Scenes: "bytes" (each type
// takes a number of bytes), "padding" (the order of a struct's fields
// changes its size) and "utf8" (a string is bytes; a rune can take several).
import { Scene, lerp } from "./scene.js";

const DROP = 6; // units an entering value falls before it settles

// bytes: two rows of eight bytes, with the address of each row on the left.
const BYTES = { x: 36, cellWidth: 15, cellHeight: 12, rows: [14, 40] };
const ADDRESSES = ["0xc000", "0xc008"];
const A = { id: "a", label: "a · 1 byte", row: 0, start: 0, size: 1, value: "1" };
const B = { id: "b", label: "b · 4 bytes", row: 0, start: 4, size: 4, value: "2" };
const C = { id: "c", label: "c · 8 bytes", row: 1, start: 0, size: 8, value: "3" };

// padding: the same three fields, in two orders, eight bytes per row.
const LOOSE = { title: "Loose", x: 10, cellWidth: 8, cellHeight: 12, rows: [16, 27, 38] };
const TIGHT = { title: "Tight", x: 92, cellWidth: 8, cellHeight: 12, rows: [16, 27, 38] };
const LOOSE_A = { name: "A", offset: 0, size: 1 };
const LOOSE_B = { name: "B", offset: 8, size: 8 };
const LOOSE_C = { name: "C", offset: 16, size: 1 };
const TIGHT_B = { name: "B", offset: 0, size: 8 };
const TIGHT_A = { name: "A", offset: 8, size: 1 };
const TIGHT_C = { name: "C", offset: 9, size: 1 };

// utf8: the four bytes of "año", with the character each group encodes.
const UTF8 = { x: 38, y: 20, cellWidth: 22, cellHeight: 14 };
const UTF8_BYTES = ["61", "c3", "b1", "6f"];
const RUNES = [
  { text: "a", from: 0, to: 0 },
  { text: "ñ", from: 1, to: 2 },
  { text: "o", from: 3, to: 3 },
];

const scenes = {
  bytes: {
    output: { x: 4, y: 62, width: 152, height: 30 },
    code: [
      "var a int8 = 1",
      "var b int32 = 2",
      "var c int64 = 3",
      "fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(c))",
    ],
    steps: [
      {
        line: 0,
        caption: "La memoria es una fila de bytes con direcciones seguidas. Un `int8` ocupa 1 byte.",
        vars: [A],
        enter: "a",
      },
      {
        line: 1,
        caption: "Un `int32` ocupa 4 bytes y empieza en una dirección múltiplo de 4: quedan 3 bytes sin usar.",
        vars: [A, B],
        enter: "b",
      },
      {
        line: 2,
        caption: "Un `int64` ocupa 8 bytes: una fila entera.",
        vars: [A, B, C],
        enter: "c",
      },
      {
        line: 3,
        caption: "`unsafe.Sizeof` dice cuántos bytes ocupa cada variable: 1, 4 y 8.",
        vars: [A, B, C],
        output: ["1 4 8"],
        printed: true,
      },
    ],
  },

  padding: {
    output: { x: 4, y: 58, width: 152, height: 34 },
    code: [
      "type Loose struct {",
      "\tA bool",
      "\tB int64",
      "\tC bool",
      "}",
      "type Tight struct {",
      "\tB int64",
      "\tA bool",
      "\tC bool",
      "}",
      "fmt.Println(unsafe.Sizeof(Loose{}),",
      "\tunsafe.Sizeof(Tight{}))",
    ],
    steps: [
      {
        line: 1,
        caption: "En `Loose`, `A` es un `bool`: ocupa 1 byte, el primero.",
        loose: { fields: [LOOSE_A], padded: 1 },
        tight: { fields: [], padded: 0 },
        enter: "A",
      },
      {
        line: 2,
        caption: "`B` es un `int64` y tiene que empezar en un múltiplo de 8. Go deja 7 bytes de relleno: el padding.",
        loose: { fields: [LOOSE_A, LOOSE_B], padded: 16 },
        tight: { fields: [], padded: 0 },
        enter: "B",
      },
      {
        line: 3,
        caption: "`C` ocupa 1 byte, pero el struct tiene que medir un múltiplo de 8: otros 7 bytes de relleno.",
        loose: { fields: [LOOSE_A, LOOSE_B, LOOSE_C], padded: 24 },
        tight: { fields: [], padded: 0 },
        enter: "C",
      },
      {
        line: 6,
        caption: "En `Tight`, `B` va primero y ocupa la primera fila entera.",
        loose: { fields: [LOOSE_A, LOOSE_B, LOOSE_C], padded: 24 },
        tight: { fields: [TIGHT_B], padded: 8 },
        enterTight: "B",
      },
      {
        line: 7,
        caption: "`A` va justo detrás: un `bool` puede empezar en cualquier byte.",
        loose: { fields: [LOOSE_A, LOOSE_B, LOOSE_C], padded: 24 },
        tight: { fields: [TIGHT_B, TIGHT_A], padded: 9 },
        enterTight: "A",
      },
      {
        line: 8,
        caption: "`C` también. Al final solo quedan 6 bytes de relleno para llegar a 16.",
        loose: { fields: [LOOSE_A, LOOSE_B, LOOSE_C], padded: 24 },
        tight: { fields: [TIGHT_B, TIGHT_A, TIGHT_C], padded: 16 },
        enterTight: "C",
      },
      {
        line: 10,
        caption: "Los mismos campos en otro orden: 24 bytes contra 16.",
        loose: { fields: [LOOSE_A, LOOSE_B, LOOSE_C], padded: 24 },
        tight: { fields: [TIGHT_B, TIGHT_A, TIGHT_C], padded: 16 },
        output: ["24 16"],
        printed: true,
      },
    ],
  },

  utf8: {
    output: { x: 4, y: 50, width: 152, height: 44 },
    code: ['s := "año"', "fmt.Println(len(s))", "for i, r := range s {", "\tfmt.Println(i, string(r))", "}"],
    steps: [
      {
        line: 0,
        caption: "Un string guarda su texto en UTF-8: aquí, 4 bytes. La ñ ocupa dos, c3 y b1.",
        enter: true,
      },
      {
        line: 1,
        caption: "`len(s)` cuenta bytes, no letras: 4.",
        reading: { from: 0, to: 3 },
        output: ["4"],
        printed: true,
      },
      {
        line: 3,
        caption: "`range` lee una letra entera, una rune, cada vez. Empieza en el byte 0: la a.",
        reading: { from: 0, to: 0 },
        rune: 0,
        output: ["4", "0 a"],
        printed: true,
      },
      {
        line: 3,
        caption: "La ñ empieza en el byte 1 y ocupa dos: `range` los lee juntos.",
        reading: { from: 1, to: 2 },
        rune: 1,
        output: ["4", "0 a", "1 ñ"],
        printed: true,
      },
      {
        line: 3,
        caption: "El índice salta del 1 al 3: `i` es la posición en bytes, no el número de letra.",
        reading: { from: 3, to: 3 },
        rune: 2,
        output: ["4", "0 a", "1 ñ", "3 o"],
        printed: true,
      },
    ],
  },
};

class MemoryDemo extends Scene {
  static scenes = scenes;

  // The three scenes share cells and output, but each has its own layout.
  draw(paint, step, t) {
    const name = this.getAttribute("scene");
    const drawScene = { bytes: drawBytes, padding: drawPadding, utf8: drawUtf8 }[name];
    drawScene(paint, step, t);
    if (step.output) drawOutput(paint, scenes[name].output, step.output, step.printed, t);
  }
}

// drawBytes draws two rows of empty bytes and, over them, each variable as
// one frame across the bytes it takes. The new variable drops in.
function drawBytes(paint, step, t) {
  const { colors } = paint;
  BYTES.rows.forEach((y, row) => {
    paint.text(ADDRESSES[row], 4, y + textOffset(BYTES), colors.pointer);
    drawCells(paint, BYTES, y, Array(8).fill({ text: "", kind: "dim" }));
  });
  for (const v of step.vars) {
    const entering = v.id === step.enter;
    const y = BYTES.rows[v.row] - (entering ? lerp(DROP, 0, t) : 0);
    const color = entering ? colors.highlight : colors.value;
    const x = cellX(BYTES, v.start);
    const width = spanWidth(BYTES, v.size);
    paint.text(v.label, x, y - 7, colors.paper);
    paint.frame(x, y, width, BYTES.cellHeight, color);
    paint.centerText(v.value, x, y + textOffset(BYTES), width, color);
  }
}

// drawPadding draws Loose and Tight side by side.
function drawPadding(paint, step, t) {
  drawStruct(paint, LOOSE, step.loose, step.enter, t);
  drawStruct(paint, TIGHT, step.tight, step.enterTight, t);
}

// drawStruct draws three rows of eight bytes. Bytes the struct already
// takes but no field uses are padding, marked with a dot.
function drawStruct(paint, grid, layout, entering, t) {
  const { colors } = paint;
  paint.text(grid.title, grid.x, 6, colors.paper);
  grid.rows.forEach((y, row) => {
    drawCells(paint, grid, y, Array(8).fill({ text: "", kind: "dim" }));
    for (let col = 0; col < 8; col++) {
      const offset = row * 8 + col;
      if (offset >= layout.padded || layout.fields.some((f) => covers(f, offset))) continue;
      paint.dot(cellX(grid, col) + grid.cellWidth / 2 - 1, y + grid.cellHeight / 2 - 1, colors.dim);
    }
  });
  for (const field of layout.fields) {
    const isNew = field.name === entering;
    const y = grid.rows[Math.floor(field.offset / 8)] - (isNew ? lerp(DROP, 0, t) : 0);
    const x = cellX(grid, field.offset % 8);
    const width = spanWidth(grid, field.size);
    const color = isNew ? colors.highlight : colors.value;
    paint.frame(x, y, width, grid.cellHeight, color);
    paint.centerText(field.name, x, y + textOffset(grid), width, color);
  }
}

// covers reports whether a field takes the byte at offset.
function covers(field, offset) {
  return offset >= field.offset && offset < field.offset + field.size;
}

// drawUtf8 draws the bytes of the string with their index above and the
// character they encode below. The bytes the step reads glow.
function drawUtf8(paint, step, t) {
  const { colors } = paint;
  const y = UTF8.y - (step.enter ? lerp(DROP, 0, t) : 0);
  const reading = step.reading;

  paint.text("s", 20, y + textOffset(UTF8), colors.paper);
  const cells = UTF8_BYTES.map((text, i) => {
    const isRead = reading && i >= reading.from && i <= reading.to;
    paint.centerText(String(i), cellX(UTF8, i), y - 7, UTF8.cellWidth, colors.paper);
    return { text, kind: isRead ? "highlight" : "value" };
  });
  drawCells(paint, UTF8, y, cells);

  RUNES.forEach((rune, i) => {
    const color = step.rune === i ? colors.highlight : colors.paper;
    const left = cellX(UTF8, rune.from) + 2;
    const right = cellX(UTF8, rune.to) + UTF8.cellWidth - 3;
    const below = y + UTF8.cellHeight + 2;
    paint.line(left, below, right, below, colors.dim);
    paint.centerText(rune.text, left, below + 3, right - left, color);
  });
}

// Brighter cells are drawn last, so the borders they share with their
// neighbours keep the brighter colour.
const LAYERS = ["dim", "value", "highlight"];

// drawCells draws a row of cells that share their borders. Each cell is
// { text, kind }, where kind names its colour.
function drawCells(paint, row, y, cells) {
  const order = [...cells.keys()].sort((i, j) => LAYERS.indexOf(cells[i].kind) - LAYERS.indexOf(cells[j].kind));
  for (const i of order) {
    const color = paint.colors[cells[i].kind];
    paint.frame(cellX(row, i), y, row.cellWidth, row.cellHeight, color);
    paint.centerText(cells[i].text, cellX(row, i), y + textOffset(row), row.cellWidth, color);
  }
}

// cellX returns the left edge of cell i; neighbours share one border.
function cellX(row, i) {
  return row.x + i * (row.cellWidth - 1);
}

// spanWidth is the width of n neighbouring cells, borders shared.
function spanWidth(row, n) {
  return n * (row.cellWidth - 1) + 1;
}

// textOffset is how far below a cell's top edge a line of text starts so
// that it sits in the middle of the cell.
function textOffset(row) {
  return (row.cellHeight - 5) / 2;
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

customElements.define("memory-demo", MemoryDemo);
