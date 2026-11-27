// <net-demo> shows what travels between two programs on a network. Scenes:
// "handshake" (TCP opens a connection with three messages), "stream" (TCP
// carries bytes, not messages, so lines mark where each one ends),
// "request" (an HTTP request and its response are lines of text) and
// "template" (html/template fills a page with data and escapes it).
import { Scene, lerp } from "./scene.js";

// handshake and stream: the client on the left, the server on the right
// and the connection between them.
const CLIENT = { x: 4, y: 12, width: 38, height: 76 };
const SERVER = { x: 118, y: 12, width: 38, height: 76 };
const GAP = { from: 44, to: 116 }; // the space between the two boxes
const HANDSHAKE_LANES = [50, 62, 74]; // one row per handshake message
const TUBE = { y: 28, height: 12 }; // the open connection
const SERVER_STATUS = { listen: ":8080", waiting: "esperando", conn: "conn" };

// stream: one cell per byte inside the tube, in the order they were sent,
// and what the server takes out of it.
const CELL = 9;
const SLIDE = 6; // units a new byte slides in from the client's side
const BYTES = ["u", "n", "o", "\\n", "d", "o", "s", "\\n"];
const RESULT = { x: 46, valueX: 70, y: 52, row: 12 };

// request: the request on top, the response below, one line per row.
const REQUEST_SHEET = { x: 4, y: 2, width: 152, height: 44, firstRow: 7 };
const RESPONSE_SHEET = { x: 4, y: 50, width: 152, height: 44, firstRow: 55 };
const SHEET_ROW = 7;

// template: the data on top and the HTML it produces below.
const DATA = { x: 4, y: 2, width: 152, height: 24 };
const OUTPUT = { x: 4, y: 30, width: 152, height: 64 };

const REQUEST_ROWS = [
  { text: "POST /notes HTTP/1.1", label: "método y ruta" },
  { text: "Host: localhost:8080", label: "cabecera" },
  { text: "Content-Length: 4", label: "cabecera" },
  { text: "\\r\\n", label: "línea en blanco", blank: true },
  { text: "hola", label: "cuerpo" },
];
const RESPONSE_ROWS = [
  { text: "HTTP/1.1 201 Created", label: "estado" },
  { text: "Content-Length: 2", label: "cabecera" },
  { text: "\\r\\n", label: "línea en blanco", blank: true },
  { text: "ok", label: "cuerpo" },
];

const PAGE_ROWS = ["<h1>Compra</h1>", "<ul>", "  <li>pan</li>", "  <li>&lt;b&gt;sal&lt;/b&gt;</li>", "</ul>"];

const scenes = {
  handshake: {
    code: [
      "// servidor",
      'ln, err := net.Listen("tcp", ":8080")',
      "conn, err := ln.Accept()",
      "// cliente",
      'conn, err := net.Dial("tcp", "localhost:8080")',
      'fmt.Fprintln(conn, "hola")',
    ],
    steps: [
      {
        line: 1,
        caption: "`net.Listen` reserva el puerto 8080 del ordenador para este servidor.",
        server: "listen",
        lit: "server",
      },
      {
        line: 2,
        caption: "`Accept` se queda esperando hasta que llame un cliente.",
        server: "waiting",
        lit: "server",
      },
      {
        line: 4,
        caption: "`Dial` llama al puerto 8080. Su primer mensaje es SYN: «¿hablamos?».",
        server: "waiting",
        packets: 1,
      },
      {
        line: 4,
        caption: "El servidor contesta SYN-ACK: «vale, ¿me oyes?».",
        server: "waiting",
        packets: 2,
      },
      {
        line: 4,
        caption: "El cliente responde ACK: «te oigo». Tres mensajes: el saludo de TCP.",
        server: "waiting",
        packets: 3,
      },
      {
        line: 2,
        caption: "La conexión está abierta: `Dial` y `Accept` devuelven cada uno su `conn`.",
        server: "conn",
        connected: true,
      },
      {
        line: 5,
        caption: "Ahora los bytes de `hola\\n` viajan por ella, en orden y sin perderse.",
        server: "conn",
        connected: true,
        sending: "hola\\n",
      },
    ],
  },

  stream: {
    code: [
      'fmt.Fprint(conn, "uno\\n")',
      'fmt.Fprint(conn, "dos\\n")',
      "// al otro lado:",
      "n, err := conn.Read(buf)",
      "sc := bufio.NewScanner(conn)",
      "sc.Scan() // sc.Text()",
    ],
    steps: [
      {
        line: 0,
        caption: "El cliente escribe `uno\\n`: cuatro bytes entran en la conexión.",
        bytes: 4,
        entering: [0, 4],
      },
      {
        line: 1,
        caption: "Luego escribe `dos\\n`. Dentro no queda ninguna separación: solo una fila de bytes.",
        bytes: 8,
        entering: [4, 8],
      },
      {
        line: 3,
        caption: "`Read` te da los bytes que hayan llegado, como `uno\\nd`: un trozo, no un mensaje.",
        bytes: 8,
        read: 5, // how many bytes the Read call got
      },
      {
        line: 4,
        caption: "Por eso los dos programas acuerdan una frontera: cada mensaje acaba en `\\n`.",
        bytes: 8,
        boundaries: true,
      },
      {
        line: 5,
        caption: "`Scan` lee hasta el `\\n` y te da `uno`, ya sin el salto de línea.",
        bytes: 8,
        lines: 1,
      },
      {
        line: 5,
        caption: "La siguiente llamada a `Scan` te da `dos`. Cada línea es un mensaje.",
        bytes: 8,
        lines: 2,
      },
    ],
  },

  request: {
    code: [
      "POST /notes HTTP/1.1",
      "Host: localhost:8080",
      "Content-Length: 4",
      "",
      "hola",
      "HTTP/1.1 201 Created",
      "Content-Length: 2",
      "",
      "ok",
    ],
    steps: [
      {
        line: 0,
        caption: "Una petición HTTP es texto. La primera línea dice el método, la ruta y la versión.",
        request: 1,
      },
      {
        line: 1,
        caption: "Después van las cabeceras, una por línea: `Nombre: valor`.",
        request: 2,
      },
      {
        line: 2,
        caption: "`Content-Length` dice cuántos bytes ocupa el cuerpo: 4.",
        request: 3,
      },
      {
        line: 3,
        caption: "Una línea en blanco cierra las cabeceras. Así el servidor sabe dónde acaban.",
        request: 4,
      },
      {
        line: 4,
        caption: "Luego va el cuerpo, `hola`. Todo viaja por una conexión TCP.",
        request: 5,
      },
      {
        line: 5,
        caption: "La respuesta tiene la misma forma. Su primera línea trae el código de estado.",
        request: 5,
        response: 1,
        lit: [0],
      },
      {
        line: 6,
        caption: "Cabeceras, línea en blanco y cuerpo, igual que en la petición.",
        request: 5,
        response: 4,
        lit: [1, 2, 3],
      },
      {
        line: 5,
        caption: "2xx: todo bien. 4xx: la petición tiene un error. 5xx: ha fallado el servidor.",
        request: 5,
        response: 4,
        lit: [0],
      },
    ],
  },

  template: {
    code: ["<h1>{{.Title}}</h1>", "<ul>", "{{range .Items}}", "  <li>{{.}}</li>", "{{end}}", "</ul>"],
    steps: [
      {
        line: 0,
        caption: "`{{.Title}}` escribe el campo `Title` de los datos: Compra.",
        data: "title",
        rows: 1,
      },
      {
        line: 1,
        caption: "El texto que está fuera de `{{ }}` sale tal cual.",
        rows: 2,
      },
      {
        line: 2,
        caption: "`{{range .Items}}` repite su bloque una vez por cada elemento de `Items`.",
        data: "items",
        rows: 2,
      },
      {
        line: 3,
        caption: "Primera vuelta: dentro del `range`, el punto `.` es el elemento, pan.",
        data: "pan",
        rows: 3,
      },
      {
        line: 3,
        caption: "Segunda vuelta: `<b>sal</b>`. `html/template` lo escapa y se verá como texto.",
        data: "sal",
        rows: 4,
      },
      {
        line: 5,
        caption: "`{{end}}` cierra el `range`, y la página acaba con `</ul>`.",
        rows: 5,
      },
    ],
  },
};

class NetDemo extends Scene {
  static scenes = scenes;

  draw(paint, step, t) {
    const scene = this.getAttribute("scene");
    if (scene === "handshake") drawHandshake(paint, step, t);
    if (scene === "stream") drawStream(paint, step, t);
    if (scene === "request") drawRequest(paint, step, t);
    if (scene === "template") drawTemplate(paint, step, t);
  }
}

// drawHost draws a program as a tall box with its name and a status line.
function drawHost(paint, box, name, status, statusColor) {
  const { colors } = paint;
  paint.frame(box.x, box.y, box.width, box.height, colors.paper);
  paint.centerText(name, box.x, box.y + 4, box.width, colors.paper);
  if (status) paint.centerText(status, box.x, box.y + 50, box.width, statusColor);
}

// drawHandshake draws the client, the server, the three opening messages
// and, once they are done, the connection with the bytes it carries.
function drawHandshake(paint, step, t) {
  const { colors } = paint;
  const serverLit = step.lit === "server" || step.lit === "conns";
  const clientLit = step.lit === "conns";
  drawHost(paint, CLIENT, "cliente", step.connected ? "conn" : "", clientLit ? colors.highlight : colors.value);
  drawHost(paint, SERVER, "servidor", SERVER_STATUS[step.server], serverLit ? colors.highlight : colors.value);

  if (step.connected) {
    drawTube(paint, clientLit ? colors.highlight : colors.paper);
    if (step.sending) {
      const x = lerp(GAP.from, GAP.to - paint.textWidth(step.sending), t);
      paint.text(step.sending, x, TUBE.y + 3.5, colors.highlight);
    }
    return;
  }
  const messages = ["SYN", "SYN-ACK", "ACK"];
  for (let i = 0; i < (step.packets || 0); i++) {
    const isNew = i === step.packets - 1;
    const y = HANDSHAKE_LANES[i];
    const toServer = i !== 1;
    const path = toServer
      ? [[GAP.from, y], [GAP.to, y]]
      : [[GAP.to, y], [GAP.from, y]];
    const color = isNew ? colors.highlight : colors.dim;
    paint.arrow(path, color, isNew ? t : 1);
    paint.centerText(messages[i], GAP.from, y - 7, GAP.to - GAP.from, color);
  }
}

// drawTube draws the open connection as a pipe between the two boxes.
function drawTube(paint, color) {
  paint.frame(GAP.from - 2, TUBE.y, GAP.to - GAP.from + 4, TUBE.height, color);
}

// drawStream draws the bytes inside the connection and what the server
// takes out of it: one Read chunk or one line per Scan.
function drawStream(paint, step, t) {
  const { colors } = paint;
  drawHost(paint, CLIENT, "cliente", "", colors.value);
  drawHost(paint, SERVER, "servidor", "", colors.value);
  drawTube(paint, colors.paper);

  for (let i = 0; i < step.bytes; i++) {
    const entering = step.entering && i >= step.entering[0] && i < step.entering[1];
    const home = GAP.from + i * CELL;
    const x = entering ? lerp(home - SLIDE, home, t) : home;
    drawByte(paint, BYTES[i], x, byteColor(paint, step, i, entering));
  }

  if (step.read) {
    paint.text("Read", RESULT.x, RESULT.y, colors.paper);
    if (t === 1) paint.text('"uno\\nd"', RESULT.valueX, RESULT.y, colors.error);
  }
  const found = ['"uno"', '"dos"'];
  for (let i = 0; i < (step.lines || 0); i++) {
    const y = RESULT.y + i * RESULT.row;
    const isNew = i === step.lines - 1;
    paint.text("Scan", RESULT.x, y, colors.paper);
    if (!isNew || t === 1) paint.text(found[i], RESULT.valueX, y, isNew ? colors.highlight : colors.ok);
  }
}

// byteColor says what a byte cell means in this step: arriving, taken by
// Read as a loose chunk, a line boundary, or part of a line Scan returned.
function byteColor(paint, step, i, entering) {
  const { colors } = paint;
  if (entering) return colors.highlight;
  if (step.read) return i < step.read ? colors.error : colors.dim;
  if (step.boundaries) return BYTES[i] === "\\n" ? colors.highlight : colors.value;
  if (step.lines) {
    const line = i < 4 ? 1 : 2;
    if (line === step.lines) return colors.highlight;
    return line < step.lines ? colors.ok : colors.value;
  }
  return colors.value;
}

function drawByte(paint, char, x, color) {
  paint.frame(x, TUBE.y + 1.5, CELL, CELL, color);
  paint.centerText(char, x, TUBE.y + 3.5, CELL, color);
}

// drawRequest draws the request and the response as sheets of text. While
// only the request is shown, its newest line is the one lit.
function drawRequest(paint, step) {
  const requestLit = step.response ? [] : [step.request - 1];
  drawSheet(paint, REQUEST_SHEET, REQUEST_ROWS, step.request, requestLit);
  if (step.response) drawSheet(paint, RESPONSE_SHEET, RESPONSE_ROWS, step.response, step.lit);
}

// drawSheet draws the first count rows, one line of text per row with what
// that line is on the right. The rows in lit stand out.
function drawSheet(paint, sheet, rows, count, lit) {
  const { colors } = paint;
  paint.frame(sheet.x, sheet.y, sheet.width, sheet.height, colors.paper);
  const right = sheet.x + sheet.width - 4;
  rows.slice(0, count).forEach((row, i) => {
    const y = sheet.firstRow + i * SHEET_ROW;
    let color = row.blank ? colors.dim : colors.value;
    if (lit.includes(i)) color = colors.highlight;
    paint.text(row.text, sheet.x + 4, y, color);
    paint.text(row.label, right - paint.textWidth(row.label), y, colors.dim);
  });
}

// drawTemplate draws the data the template reads and the HTML it has
// written so far. The row written in this step is the highlighted one.
function drawTemplate(paint, step, t) {
  const { colors } = paint;
  paint.frame(DATA.x, DATA.y, DATA.width, DATA.height, colors.dim);
  paint.text("datos", DATA.x + 4, DATA.y + 3, colors.dim);
  drawDataRow(paint, ["Title: ", '"Compra"'], DATA.y + 10, step.data === "title");
  const items = paint.text("Items: [", DATA.x + 4, DATA.y + 17, colors.paper);
  let x = DATA.x + 4 + items;
  x += paint.text('"pan"', x, DATA.y + 17, step.data === "pan" || step.data === "items" ? colors.highlight : colors.value);
  x += paint.text(", ", x, DATA.y + 17, colors.paper);
  x += paint.text('"<b>sal</b>"', x, DATA.y + 17, step.data === "sal" || step.data === "items" ? colors.highlight : colors.value);
  paint.text("]", x, DATA.y + 17, colors.paper);

  paint.frame(OUTPUT.x, OUTPUT.y, OUTPUT.width, OUTPUT.height, colors.dim);
  paint.text("salida", OUTPUT.x + 4, OUTPUT.y + 3, colors.dim);
  PAGE_ROWS.slice(0, step.rows).forEach((row, i) => {
    const isNew = i === step.rows - 1 && step.data !== "items";
    if (isNew && t < 1) return;
    paint.text(row, OUTPUT.x + 4, OUTPUT.y + 12 + i * 9, isNew ? colors.highlight : colors.paper);
  });
}

function drawDataRow(paint, [name, value], y, lit) {
  const { colors } = paint;
  const width = paint.text(name, DATA.x + 4, y, colors.paper);
  paint.text(value, DATA.x + 4 + width, y, lit ? colors.highlight : colors.value);
}

customElements.define("net-demo", NetDemo);
