// <buffer-demo> shows bytes waiting in a buffer before they move on.
// Scenes: "batch" (many small writes fill a buffer; one big write reaches
// the disk), "bytes" (a bytes.Buffer: Write appends at the end, Read
// advances from the start), "flush" (a bufio.Writer keeps the lines until
// Flush sends them), "scan" (a bufio.Scanner fills its buffer and cuts it
// at each newline) and "copy" (io.Copy moves chunks through one fixed
// buffer).
import { Scene, lerp } from "./scene.js";

const DROP = 6; // units an entering byte falls before it settles

// A byte is a small box; neighbouring boxes share their border.
const CELL = { width: 7, height: 10, step: 6 };

const scenes = {
  batch: {
    code: [
      "for _, word := range words {",
      "\tslow.Write([]byte(word))",
      "}",
      "",
      "var buf []byte",
      "for _, word := range words {",
      "\tbuf = append(buf, word...)",
      "}",
      "fast.Write(buf)",
      "fmt.Println(slow.writes, fast.writes)",
    ],
    steps: [
      { line: 1, caption: "Sin buffer, cada `Write` lleva una palabra hasta el disco.", slow: 1, trip: "uno " },
      { line: 1, caption: "La segunda palabra hace otro viaje. Cada viaje cuesta tiempo, aunque lleve pocos bytes.", slow: 2, trip: "dos " },
      { line: 1, caption: "Tres palabras, tres viajes al disco.", slow: 3, trip: "tres " },
      { line: 6, caption: "Con un buffer, `append` deja la palabra en memoria, cerca del programa.", slow: 3, buffer: "uno ", from: 0 },
      { line: 6, caption: "La segunda palabra se pone detrás. El disco aún no ha recibido nada.", slow: 3, buffer: "uno dos ", from: 4 },
      { line: 6, caption: "El buffer ya tiene las tres palabras: 13 bytes esperando.", slow: 3, buffer: "uno dos tres ", from: 8 },
      { line: 8, caption: "Un solo `Write` lleva los 13 bytes de golpe: un viaje en vez de tres.", slow: 3, fast: 1, flushed: "uno dos tres " },
      { line: 9, caption: "Por eso el programa imprime `3 1`.", slow: 3, fast: 1, printed: "3 1" },
    ],
  },

  bytes: {
    code: [
      "var buf bytes.Buffer",
      'buf.WriteString("uno ")',
      'buf.Write([]byte("dos"))',
      "b := make([]byte, 4)",
      "n, _ := buf.Read(b)",
      "fmt.Println(buf.Len(), buf.String())",
    ],
    steps: [
      { line: 0, caption: "Un `bytes.Buffer` vacío: un slice de bytes que crece cuando escribes.", buffer: "", read: 0 },
      { line: 1, caption: "`WriteString` añade los bytes al final, donde apunta «escribe».", buffer: "uno ", read: 0, from: 0 },
      { line: 2, caption: "`Write` también añade al final. Ahora `Len()` vale 7.", buffer: "uno dos", read: 0, from: 4 },
      { line: 3, caption: "`b` es un slice de 4 bytes. `Read` va a copiar en él.", buffer: "uno dos", read: 0, target: "", newTarget: true },
      { line: 4, caption: "`Read` copia desde el principio y avanza «lee»: esos 4 bytes ya están leídos.", buffer: "uno dos", read: 4, target: "uno ", reading: true },
      { line: 5, caption: "Solo cuenta lo que queda sin leer: `Len()` vale 3 y `String()` es \"dos\".", buffer: "uno dos", read: 4, target: "uno ", printed: "3 dos" },
    ],
  },

  flush: {
    code: [
      "w := bufio.NewWriter(os.Stdout)",
      "for i := 1; i <= 3; i++ {",
      '\tfmt.Fprintln(w, "línea", i)',
      "}",
      "w.Flush()",
    ],
    steps: [
      { line: 0, caption: "`bufio.NewWriter` pone un buffer de 4096 bytes delante de `os.Stdout`.", lines: 0 },
      { line: 2, caption: "`Fprintln` escribe en el buffer, no en la pantalla. La salida sigue vacía.", lines: 1, entering: true },
      { line: 2, caption: "La segunda línea también espera. El buffer está lejos de llenarse.", lines: 2, entering: true },
      { line: 2, caption: "Si `main` terminara aquí, estas tres líneas se perderían sin llegar a la pantalla.", lines: 3, entering: true },
      { line: 4, caption: "`Flush` vacía el buffer en `os.Stdout`: ahora sí aparecen las tres líneas.", lines: 3, flushing: true },
    ],
  },

  scan: {
    code: [
      'r := strings.NewReader("uno\\ndos\\ntres\\n")',
      "scanner := bufio.NewScanner(r)",
      "for scanner.Scan() {",
      '\tfmt.Println(">", scanner.Text())',
      "}",
    ],
    steps: [
      { line: 1, caption: "El `Scanner` envuelve al `Reader`. Todavía no ha leído nada.", output: [] },
      { line: 2, caption: "El primer `Scan` llena su buffer con un `Read`: todo el texto cabe.", filled: true, filling: true, output: [] },
      { line: 2, caption: "Busca el primer salto de línea y corta ahí. `Text()` es la línea sin el salto.", filled: true, token: [0, 3], output: [] },
      { line: 3, caption: "El programa imprime la primera línea.", filled: true, token: [0, 3], output: ["> uno"], printed: true },
      { line: 2, caption: "El siguiente `Scan` sigue donde lo dejó, sin volver a leer del `Reader`.", filled: true, token: [4, 7], output: ["> uno"] },
      { line: 3, caption: "Imprime la segunda línea.", filled: true, token: [4, 7], output: ["> uno", "> dos"], printed: true },
      { line: 2, caption: "Otro corte, en el último salto de línea.", filled: true, token: [8, 12], output: ["> uno", "> dos"] },
      { line: 3, caption: "Imprime la tercera.", filled: true, token: [8, 12], output: ["> uno", "> dos", "> tres"], printed: true },
      { line: 2, caption: "No quedan líneas: `Scan` devuelve false y el bucle termina.", filled: true, done: true, output: ["> uno", "> dos", "> tres"] },
    ],
  },

  copy: {
    code: [
      "buf := make([]byte, 4)",
      "for {",
      "\tn, err := src.Read(buf)",
      "\tdst.Write(buf[:n])",
      "\tif err == io.EOF {",
      "\t\treturn",
      "\t}",
      "}",
    ],
    steps: [
      { line: 0, caption: "Por dentro, `io.Copy` hace algo así. Su buffer suele tener 32 KB; este, 4 bytes.", read: 0, written: 0, newBuffer: true },
      { line: 2, caption: "`Read` llena el buffer con los 4 primeros bytes.", read: 4, written: 0, move: "read" },
      { line: 3, caption: "`Write` los lleva al destino. El buffer queda libre para el siguiente trozo.", read: 4, written: 4, move: "write" },
      { line: 2, caption: "El mismo buffer recibe los 4 bytes siguientes.", read: 8, written: 4, move: "read" },
      { line: 3, caption: "Y otra vez al destino.", read: 8, written: 8, move: "write" },
      { line: 2, caption: "El último trozo.", read: 12, written: 8, move: "read" },
      { line: 3, caption: "Ya están los 12 bytes en el destino, y nunca hubo más de 4 en memoria.", read: 12, written: 12, move: "write" },
      { line: 4, caption: "El siguiente `Read` no encuentra nada: devuelve `io.EOF` y la copia termina.", read: 12, written: 12, eof: true },
    ],
  },
};

// Each scene has its own picture, drawn by its own function below.
const painters = { batch: drawBatch, bytes: drawBytes, flush: drawFlush, scan: drawScan, copy: drawCopy };

class BufferDemo extends Scene {
  static scenes = scenes;

  draw(paint, step, t) {
    painters[this.getAttribute("scene")](paint, step, t);
  }
}

// batch: the slow writes on top, the buffered ones below, a disk for each.
const SLOW_DISK = { x: 110, y: 6, width: 46, height: 28 };
const FAST_DISK = { x: 110, y: 52, width: 46, height: 28 };
const BATCH_BUFFER = { x: 4, y: 62, size: 16 };

function drawBatch(paint, step, t) {
  const { colors } = paint;
  paint.text("sin buffer", 4, 4, colors.paper);
  paint.text("con buffer", 4, 50, colors.paper);
  paint.arrow([[4, 30], [SLOW_DISK.x - 3, 30]], colors.dim);

  const slowTrips = step.trip && t < 1 ? step.slow - 1 : step.slow;
  drawDisk(paint, SLOW_DISK, slowTrips);
  if (step.trip && t < 1) {
    drawPacket(paint, step.trip, lerp(4, SLOW_DISK.x - 26, t), 16, colors.highlight);
  }

  const flushing = step.flushed && t < 1;
  drawDisk(paint, FAST_DISK, flushing ? 0 : step.fast || 0);
  const buffer = step.buffer || "";
  drawRow(paint, BATCH_BUFFER, BATCH_BUFFER.size, (i) => {
    if (i >= buffer.length) return { ch: "", color: colors.dim };
    const entering = step.from !== undefined && i >= step.from;
    return { ch: buffer[i], color: entering ? colors.highlight : colors.value, drop: entering ? 1 - t : 0 };
  });
  if (flushing) {
    drawPacket(paint, step.flushed, lerp(4, FAST_DISK.x - 58, t), BATCH_BUFFER.y, colors.highlight);
  }

  if (step.printed && t >= 1) {
    const x = 4 + paint.text("salida: ", 4, 88, colors.dim);
    paint.text(step.printed, x, 88, colors.highlight);
  }
}

// drawDisk draws a disk with the number of writes that reached it.
function drawDisk(paint, disk, trips) {
  const { colors } = paint;
  paint.frame(disk.x, disk.y, disk.width, disk.height, colors.paper);
  paint.text("disco", disk.x + 4, disk.y + 4, colors.paper);
  paint.text(`viajes: ${trips}`, disk.x + 4, disk.y + 16, colors.value);
}

// bytes: the buffer with its two marks, and b below it.
const BYTES_BUFFER = { x: 20, y: 24, size: 8 };
const TARGET = { x: 20, y: 66, size: 4 };
const BYTES_OUTPUT = { x: 84, y: 54, width: 72, height: 26 };

function drawBytes(paint, step, t) {
  const { colors } = paint;
  const moving = step.reading && t < 1;
  const written = step.from !== undefined && t < 1 ? step.from : step.buffer.length;
  const read = moving ? 0 : step.read;

  paint.text("buf", 4, BYTES_BUFFER.y + 2.5, colors.paper);
  drawRow(paint, BYTES_BUFFER, BYTES_BUFFER.size, (i) => {
    if (i >= step.buffer.length) return { ch: "", color: colors.dim };
    const entering = step.from !== undefined && i >= step.from;
    if (entering) return { ch: step.buffer[i], color: colors.highlight, drop: 1 - t };
    return { ch: step.buffer[i], color: i < read ? colors.dim : colors.value };
  });
  drawMark(paint, cellX(BYTES_BUFFER, written), BYTES_BUFFER.y, "escribe", "above");
  drawMark(paint, cellX(BYTES_BUFFER, read), BYTES_BUFFER.y + CELL.height, "lee", "below");
  paint.text(`Len() = ${written - read}`, 84, BYTES_BUFFER.y + 2.5, colors.value);

  if (step.target !== undefined) drawTarget(paint, step, t);
  if (step.printed) drawOutput(paint, BYTES_OUTPUT, [step.printed], true, t);
}

// drawTarget draws b and, while Read runs, the bytes travelling into it.
function drawTarget(paint, step, t) {
  const { colors } = paint;
  const moving = step.reading && t < 1;
  paint.text("b", 4, TARGET.y + 2.5, colors.paper);
  drawRow(paint, TARGET, TARGET.size, (i) => {
    if (step.newTarget) return { ch: "", color: t < 1 ? colors.highlight : colors.dim, drop: 1 - t };
    if (moving || i >= step.target.length) return { ch: "", color: colors.dim };
    return { ch: step.target[i], color: colors.value };
  });
  if (moving) {
    [...step.target].forEach((ch, i) => {
      drawCell(paint, cellX(TARGET, i), lerp(BYTES_BUFFER.y, TARGET.y, t), ch, colors.highlight);
    });
  }
  if (step.target && !moving) paint.text(`n = ${step.target.length}`, 52, TARGET.y + 2.5, colors.value);
}

// drawMark draws a labelled arrow that points at the border between two
// bytes, from above or from below the row.
function drawMark(paint, x, edge, label, side) {
  const { colors } = paint;
  const above = side === "above";
  const tail = above ? edge - 10 : edge + 10;
  const tip = above ? edge - 1 : edge + 1;
  paint.arrow([[x, tail], [x, tip]], colors.pointer);
  const labelY = above ? tail - 6 : tail + 1;
  paint.centerText(label, x - 20, labelY, 40, colors.pointer);
}

// flush: the bufio.Writer's buffer on the left, the screen on the right.
const FLUSH_BUFFER = { x: 4, y: 10, width: 70, height: 46 };
const SCREEN = { x: 86, y: 10, width: 70, height: 46 };
const LINE_BYTES = 9; // "línea 1\n": the í takes two bytes

function drawFlush(paint, step, t) {
  const { colors } = paint;
  const flushing = step.flushing && t < 1;
  const flushed = step.flushing && t >= 1;

  paint.frame(FLUSH_BUFFER.x, FLUSH_BUFFER.y, FLUSH_BUFFER.width, FLUSH_BUFFER.height, colors.paper, true);
  paint.text("buffer", FLUSH_BUFFER.x + 4, FLUSH_BUFFER.y + 4, colors.paper);
  paint.frame(SCREEN.x, SCREEN.y, SCREEN.width, SCREEN.height, colors.dim);
  paint.text("pantalla", SCREEN.x + 4, SCREEN.y + 4, colors.dim);
  paint.arrow([[FLUSH_BUFFER.x + FLUSH_BUFFER.width + 1, 33], [SCREEN.x - 2, 33]], colors.dim);

  for (let i = 0; i < step.lines; i++) {
    const text = `línea ${i + 1}`;
    const y = FLUSH_BUFFER.y + 14 + i * 9;
    if (flushing) {
      paint.text(text, lerp(FLUSH_BUFFER.x + 6, SCREEN.x + 6, t), y, colors.highlight);
    } else if (flushed) {
      paint.text(text, SCREEN.x + 6, y, colors.highlight);
    } else {
      const entering = step.entering && i === step.lines - 1;
      const dropY = entering ? y - lerp(DROP, 0, t) : y;
      paint.text(text, FLUSH_BUFFER.x + 6, dropY, entering ? colors.highlight : colors.value);
    }
  }

  const waiting = step.flushing ? (flushed ? 0 : step.lines) : step.entering && t < 1 ? step.lines - 1 : step.lines;
  paint.text(`en el buffer: ${waiting * LINE_BYTES} de 4096 bytes`, 4, 68, colors.value);
}

// scan: the Reader's text on top, the Scanner's buffer below it, the line
// it cut under both and the program's output on the right.
const SCAN_TEXT = "uno\ndos\ntres\n";
const SOURCE = { x: 4, y: 10, size: SCAN_TEXT.length };
const SCAN_BUFFER = { x: 4, y: 36, size: SCAN_TEXT.length };
const SCAN_OUTPUT = { x: 100, y: 4, width: 56, height: 40 };

function drawScan(paint, step, t) {
  const { colors } = paint;
  const filling = step.filling && t < 1;
  const inBuffer = step.filled && !filling;

  paint.text("Reader", 4, SOURCE.y - 7, colors.paper);
  drawRow(paint, SOURCE, SOURCE.size, (i) => ({ ch: SCAN_TEXT[i], color: step.filled ? colors.dim : colors.value }));

  paint.text("buffer del Scanner", 4, SCAN_BUFFER.y - 7, colors.paper);
  drawRow(paint, SCAN_BUFFER, SCAN_BUFFER.size, (i) => {
    if (!inBuffer) return { ch: "", color: colors.dim };
    return { ch: SCAN_TEXT[i], color: scanColor(colors, step, i) };
  });
  if (filling) {
    [...SCAN_TEXT].forEach((ch, i) => {
      drawCell(paint, cellX(SOURCE, i), lerp(SOURCE.y, SCAN_BUFFER.y, t), ch, colors.highlight);
    });
  }

  if (step.token) {
    const [start, end] = step.token;
    const cutting = !step.printed;
    const y = 62 - (cutting ? lerp(DROP, 0, t) : 0);
    const x = 4 + paint.text("Text()", 4, y + 2.5, colors.paper) + 4;
    drawPacket(paint, `"${SCAN_TEXT.slice(start, end)}"`, x, y, cutting ? colors.highlight : colors.value);
  }
  drawOutput(paint, SCAN_OUTPUT, step.output, step.printed, t);
}

// scanColor colours a byte of the Scanner's buffer: the line just cut
// lights up and the bytes the Scanner has already gone past fade.
function scanColor(colors, step, i) {
  if (step.done) return colors.dim;
  if (!step.token) return colors.value;
  const [start, end] = step.token;
  if (i < start) return colors.dim;
  if (i < end && !step.printed) return colors.highlight;
  return colors.value;
}

// copy: the source, the fixed buffer and the destination, one row each.
const COPY_TEXT = "Hola, Gopher";
const COPY_SOURCE = { x: 28, y: 10, size: COPY_TEXT.length };
const COPY_BUFFER = { x: 28, y: 40, size: 4 };
const COPY_DEST = { x: 28, y: 70, size: COPY_TEXT.length };

function drawCopy(paint, step, t) {
  const { colors } = paint;
  const moving = t < 1 ? step.move : undefined;
  const chunkStart = Math.max(0, step.read - COPY_BUFFER.size);
  const chunk = COPY_TEXT.slice(chunkStart, step.read);

  paint.text("src", 4, COPY_SOURCE.y + 2.5, colors.paper);
  paint.text("buf", 4, COPY_BUFFER.y + 2.5, colors.paper);
  paint.text("dst", 4, COPY_DEST.y + 2.5, colors.paper);

  const taken = moving === "read" ? chunkStart : step.read;
  drawRow(paint, COPY_SOURCE, COPY_SOURCE.size, (i) => ({ ch: COPY_TEXT[i], color: i < taken ? colors.dim : colors.value }));

  drawRow(paint, COPY_BUFFER, COPY_BUFFER.size, (i) => {
    if (step.newBuffer) return { ch: "", color: t < 1 ? colors.highlight : colors.dim, drop: 1 - t };
    if (moving === "read" || moving === "write" || !chunk) return { ch: "", color: colors.dim };
    const waiting = step.read > step.written; // read but not written yet
    return { ch: chunk[i], color: waiting ? colors.value : colors.dim };
  });

  const arrived = moving === "write" ? chunkStart : step.written;
  drawRow(paint, COPY_DEST, COPY_DEST.size, (i) => {
    if (i >= arrived) return { ch: "", color: colors.dim };
    return { ch: COPY_TEXT[i], color: colors.value };
  });

  if (moving) {
    const [from, to] = moving === "read" ? [COPY_SOURCE, COPY_BUFFER] : [COPY_BUFFER, COPY_DEST];
    [...chunk].forEach((ch, i) => {
      const fromX = cellX(from, from === COPY_SOURCE ? chunkStart + i : i);
      const toX = cellX(to, to === COPY_DEST ? chunkStart + i : i);
      drawCell(paint, lerp(fromX, toX, t), lerp(from.y, to.y, t), ch, colors.highlight);
    });
  }

  if (step.eof) {
    const y = COPY_SOURCE.y + 2.5 - lerp(DROP, 0, t);
    paint.text("io.EOF", cellX(COPY_SOURCE, COPY_SOURCE.size) + 4, y, colors.highlight);
  }
}

// cellX is the left edge of the i-th byte of a row.
function cellX(row, i) {
  return row.x + i * CELL.step;
}

// drawRow draws a row of size bytes. look(i) says what the i-th byte holds:
// its character, its colour and how far it still has to drop (0 to 1).
function drawRow(paint, row, size, look) {
  for (let i = 0; i < size; i++) {
    const { ch, color, drop = 0 } = look(i);
    drawCell(paint, cellX(row, i), row.y - drop * DROP, ch, color);
  }
}

// drawCell draws one byte. Spaces and newlines get a visible mark.
function drawCell(paint, x, y, ch, color) {
  paint.frame(x, y, CELL.width, CELL.height, color);
  if (!ch) return;
  const mark = ch === " " ? "·" : ch === "\n" ? "↵" : ch;
  paint.centerText(mark, x, y + 2.5, CELL.width, color);
}

// drawPacket draws a group of bytes travelling together, boxed.
function drawPacket(paint, text, x, y, color) {
  const width = paint.textWidth(text) + 6;
  paint.frame(x, y, width, CELL.height, color);
  paint.text(text, x + 3, y + 2.5, color);
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

customElements.define("buffer-demo", BufferDemo);
