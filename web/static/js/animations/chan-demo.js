// <chan-demo> shows values travelling through a channel. Scenes:
// "unbuffered" (sender and receiver wait for each other), "buffered" (the
// sender only blocks when the buffer is full), "close" (range ends when the
// channel is closed and empty) and "select" (the first ready case wins, and
// default runs when none is ready).
import { Scene, lerp, pointOnPath } from "./scene.js";

const DROP = 6; // units an entering value falls before it settles

const OUTPUT = { x: 4, y: 62, width: 152, height: 30 };

// The sender on the left, the channel in the middle, the receiver on the
// right; a value is a small box that moves between them.
const SENDER = { x: 4, y: 22, width: 36, height: 20 };
const RECEIVER = { x: 120, y: 22, width: 36, height: 20 };
const PIPE = { x: 46, y: 24, width: 68, height: 16 };
const VALUE = { width: 12, height: 10 };
const SLOT = { width: 16, height: 14 };

// Where a value can be: its top-left corner. slot0 is the buffer's head,
// the next value a receiver gets.
const PLACES = {
  sender: [16, 27],
  middle: [74, 27],
  slot0: [88, 27],
  slot1: [62, 27],
  receiver: [132, 27],
};

// select: two channels on the left, the select with its cases on the right.
const CHANNELS = {
  a: { x: 16, y: 6, width: 72, height: 14 },
  b: { x: 16, y: 30, width: 72, height: 14 },
};
const SELECT = { x: 100, y: 4, width: 56, height: 42 };
const CASES = ["case <-a", "case <-b", "default"];
const SELECT_PLACES = { b: [20, 32], select: [142, 25] };

const scenes = {
  unbuffered: {
    actors: { sender: "goroutine", receiver: "main" },
    code: ["ch := make(chan string)", 'go func() { ch <- "ping" }()', "msg := <-ch", "fmt.Println(msg)"],
    steps: [
      {
        line: 0,
        caption: "`make(chan string)` crea un canal sin buffer: no guarda ningún valor.",
        changed: ["ch"],
      },
      {
        line: 1,
        caption: "`go func()` arranca la goroutine que va a enviar.",
        sender: "run",
        changed: ["sender"],
      },
      {
        line: 2,
        caption: "`main` llega a `<-ch` y espera: nadie ha enviado nada todavía.",
        sender: "run",
        receiver: "wait",
        changed: ["receiver"],
      },
      {
        line: 1,
        caption: 'La goroutine envía `"ping"`. Emisor y receptor se encuentran, y el valor pasa de uno a otro.',
        sender: "done",
        receiver: "run",
        values: [{ v: "ping", path: ["sender", "middle", "receiver"], got: true }],
      },
      {
        line: 3,
        caption: "`main` imprime ping. En un canal sin buffer, el primero que llega espera al otro.",
        sender: "done",
        receiver: "run",
        values: [{ v: "ping", path: ["receiver"], got: true }],
        output: ["ping"],
        printed: true,
      },
    ],
  },

  buffered: {
    actors: { sender: "main", receiver: "goroutine" },
    slots: 2,
    code: ["ch := make(chan int, 2)", "ch <- 1", "ch <- 2", "ch <- 3  // no cabe", "<-ch     // otra goroutine"],
    steps: [
      {
        line: 0,
        caption: "`make(chan int, 2)` crea un canal con un buffer de dos huecos.",
        sender: "run",
        changed: ["ch"],
      },
      {
        line: 1,
        caption: "`ch <- 1`: hay sitio, así que el 1 entra en el buffer y `main` sigue sin esperar.",
        sender: "run",
        values: [{ v: "1", path: ["sender", "slot0"] }],
      },
      {
        line: 2,
        caption: "`ch <- 2` también cabe. Ahora el buffer está lleno.",
        sender: "run",
        values: [{ v: "1", path: ["slot0"] }, { v: "2", path: ["sender", "slot1"] }],
      },
      {
        line: 3,
        caption: "`ch <- 3` no cabe: `main` se bloquea hasta que alguien saque un valor.",
        sender: "wait",
        values: [{ v: "1", path: ["slot0"] }, { v: "2", path: ["slot1"] }, { v: "3", path: ["sender"] }],
        changed: ["sender", "3"],
      },
      {
        line: 4,
        caption: "Otra goroutine recibe y sale el 1, el primero que entró.",
        sender: "wait",
        receiver: "run",
        values: [
          { v: "1", path: ["slot0", "receiver"], got: true },
          { v: "2", path: ["slot1"] },
          { v: "3", path: ["sender"] },
        ],
      },
      {
        line: 3,
        caption: "Queda un hueco: el 2 avanza, el 3 entra y `main` sigue.",
        sender: "run",
        receiver: "run",
        values: [
          { v: "1", path: ["receiver"], got: true },
          { v: "2", path: ["slot1", "slot0"] },
          { v: "3", path: ["sender", "slot1"] },
        ],
      },
    ],
  },

  close: {
    actors: { sender: "goroutine", receiver: "main" },
    code: ["go func() {", "\tch <- 1", "\tch <- 2", "\tclose(ch)", "}()", "for v := range ch {", "\tfmt.Println(v)", "}"],
    steps: [
      {
        line: 5,
        caption: "`for v := range ch` espera el siguiente valor del canal.",
        sender: "run",
        receiver: "wait",
        changed: ["receiver"],
      },
      {
        line: 1,
        caption: "La goroutine envía 1. `range` lo recibe y el bucle lo imprime.",
        sender: "run",
        receiver: "wait",
        values: [{ v: "1", path: ["sender", "middle", "receiver"], got: true }],
        output: ["1"],
        printed: true,
      },
      {
        line: 2,
        caption: "Lo mismo con el 2. Después, `range` vuelve a esperar.",
        sender: "run",
        receiver: "wait",
        values: [{ v: "2", path: ["sender", "middle", "receiver"], got: true }],
        output: ["1", "2"],
        printed: true,
      },
      {
        line: 3,
        caption: "`close(ch)` cierra el canal: ya no llegarán más valores. Solo cierra quien envía.",
        sender: "done",
        receiver: "wait",
        closed: true,
        changed: ["ch"],
        output: ["1", "2"],
      },
      {
        line: 7,
        caption: "El canal está cerrado y vacío, así que `range` termina y `main` sigue.",
        sender: "done",
        receiver: "run",
        closed: true,
        changed: ["receiver"],
        output: ["1", "2"],
      },
    ],
  },

  select: {
    code: [
      "select {",
      "case v := <-a:",
      '\tfmt.Println("a", v)',
      "case v := <-b:",
      '\tfmt.Println("b", v)',
      "default:",
      '\tfmt.Println("nada listo")',
      "}",
    ],
    steps: [
      {
        line: 0,
        caption: "`select` mira varios canales a la vez y elige el primero que esté listo.",
        changed: ["select"],
      },
      {
        line: 5,
        caption: "Ningún canal tiene un valor, así que se ejecuta `default` sin esperar.",
        chosen: "default",
        output: ["nada listo"],
        printed: true,
      },
      {
        line: 3,
        caption: "Más tarde, alguien envía un 7 por `b`. Ahora uno de los casos está listo.",
        values: [{ v: "7", path: ["b"] }],
        changed: ["7"],
        output: ["nada listo"],
      },
      {
        line: 4,
        caption: "Gana el caso listo: `select` recibe de `b` e imprime b 7. Si hay varios, elige uno al azar.",
        chosen: "case <-b",
        values: [{ v: "7", path: ["b", "select"], got: true }],
        output: ["nada listo", "b 7"],
        printed: true,
      },
    ],
  },
};

class ChanDemo extends Scene {
  static scenes = scenes;

  draw(paint, step, t) {
    const name = this.getAttribute("scene");
    if (name === "select") {
      drawSelect(paint, step, t);
    } else {
      drawChannel(paint, scenes[name], step, t);
    }
    if (step.output) drawOutput(paint, step.output, step.printed, t);
  }
}

// changeColor is the colour of something this step changes: it glows while
// it moves and settles into its normal colour.
function changeColor(paint, step, id, t, normal) {
  return step.changed?.includes(id) && t < 1 ? paint.colors.highlight : normal;
}

// drawChannel draws the sender, the channel with its buffer, the receiver
// and every value on its way.
function drawChannel(paint, scene, step, t) {
  const { colors } = paint;
  drawActor(paint, step, "sender", SENDER, scene.actors.sender, t);
  drawActor(paint, step, "receiver", RECEIVER, scene.actors.receiver, t);

  const settle = step.changed?.includes("ch") && !step.closed ? lerp(DROP, 0, t) : 0;
  const pipeColor = step.closed ? colors.dim : changeColor(paint, step, "ch", t, colors.paper);
  const note = step.closed ? "cerrado" : scene.slots ? `buffer ${scene.slots}` : "sin buffer";
  paint.frame(PIPE.x, PIPE.y - settle, PIPE.width, PIPE.height, pipeColor);
  paint.text(`ch: ${note}`, PIPE.x, PIPE.y - settle - 9, pipeColor);
  for (const slot of ["slot0", "slot1"].slice(0, scene.slots ?? 0)) {
    const [x, y] = PLACES[slot];
    paint.frame(x - 2, y - 2 - settle, SLOT.width, SLOT.height, colors.dim, true);
  }
  if (step.closed) {
    const color = changeColor(paint, step, "ch", t, colors.paper);
    paint.rect(PIPE.x, PIPE.y, 2, PIPE.height, color);
  }

  for (const value of step.values ?? []) drawValue(paint, step, value, PLACES, t);
}

// drawActor draws a goroutine at one end of the channel. A goroutine that
// waits is dashed; one that has finished is dim.
function drawActor(paint, step, role, box, name, t) {
  const state = step[role];
  if (!state) return;
  const { colors } = paint;
  const normal = { run: colors.value, wait: colors.pointer, done: colors.dim }[state];
  const color = changeColor(paint, step, role, t, normal);
  paint.text(name, box.x, box.y - 8, color);
  paint.frame(box.x, box.y, box.width, box.height, color, state === "wait");
  if (state === "wait") paint.text("espera", box.x, box.y + box.height + 3, color);
}

// drawValue draws a value at its place, or along its path while it moves.
// A value that arrives at a receiver turns green.
function drawValue(paint, step, value, places, t) {
  const { colors } = paint;
  const moving = value.path.length > 1;
  const points = value.path.map((place) => places[place]);
  let [x, y] = moving ? pointOnPath(points, t) : points[0];
  if (!moving && step.changed?.includes(value.v)) y -= lerp(DROP, 0, t);

  let color = value.got ? colors.ok : colors.value;
  if ((moving || step.changed?.includes(value.v)) && t < 1) color = colors.highlight;
  const width = Math.max(VALUE.width, paint.textWidth(value.v) + 4);
  paint.frame(x - (width - VALUE.width) / 2, y, width, VALUE.height, color);
  paint.centerText(value.v, x - (width - VALUE.width) / 2, y + 2.5, width, color);
}

// drawSelect draws both channels, the select with its cases and the value
// that makes one of them ready.
function drawSelect(paint, step, t) {
  const { colors } = paint;
  for (const [name, c] of Object.entries(CHANNELS)) {
    paint.text(name, 6, c.y + 4.5, colors.paper);
    paint.frame(c.x, c.y, c.width, c.height, colors.paper);
    paint.line(c.x + c.width, c.y + c.height / 2, SELECT.x, SELECT.y + 14 + (name === "a" ? 0 : 9), colors.dim);
  }

  const settle = step.changed?.includes("select") ? lerp(DROP, 0, t) : 0;
  const boxColor = changeColor(paint, step, "select", t, colors.paper);
  paint.frame(SELECT.x, SELECT.y - settle, SELECT.width, SELECT.height, boxColor);
  paint.text("select", SELECT.x + 4, SELECT.y + 3 - settle, boxColor);
  CASES.forEach((label, i) => {
    const chosen = step.chosen === label && t === 1;
    paint.text(label, SELECT.x + 4, SELECT.y + 13 + i * 9 - settle, chosen ? colors.highlight : colors.value);
  });

  for (const value of step.values ?? []) drawValue(paint, step, value, SELECT_PLACES, t);
}

// drawOutput draws what the program has printed so far. When this step
// printed, its line is the last one: it appears when the step's motion
// ends, in the highlight colour.
function drawOutput(paint, lines, printed, t) {
  const { colors } = paint;
  const { x, y, width, height } = OUTPUT;
  paint.frame(x, y, width, height, colors.dim);
  paint.text("salida", x + 4, y + 4, colors.dim);
  // Only the last two lines fit; older ones scroll out, like a terminal.
  const shown = lines.slice(-2);
  shown.forEach((line, i) => {
    const isNew = printed && i === shown.length - 1;
    if (isNew && t < 1) return;
    paint.text(line, x + 4, y + 12 + i * 8, isNew ? colors.highlight : colors.paper);
  });
}

customElements.define("chan-demo", ChanDemo);
