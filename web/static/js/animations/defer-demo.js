// <defer-demo> shows deferred calls as cards on a stack. Scene: "stack"
// (each defer drops a card on top; when the function returns, the cards
// leave top-first, so the last one deferred is the first one to run).
import { Scene, lerp, pointOnPath } from "./scene.js";

const CARD_WIDTH = 62;
const CARD_HEIGHT = 12;
const CARD_PITCH = 13; // one pixel of gap between stacked cards

// The stack is an open-top box: cards drop in and leave through the top.
const PILE = { x: 6, top: 32, floor: 84, width: 66 };
const CARD_X = PILE.x + 2;
const DROP_FROM = 4; // a new card falls from the top of the canvas

// The output strip, where fmt.Println writes.
const OUTPUT = { x: 84, y: 34, width: 72, height: 58 };
const OUTPUT_LINE = 8;

// A leaving card rises above the pile, then slides over to the output.
const LIFT_Y = 16;

const scenes = {
  stack: {
    code: [
      "func main() {",
      '\tfmt.Println("contando")',
      "\tfor i := 1; i <= 3; i++ {",
      "\t\tdefer fmt.Println(i)",
      "\t}",
      '\tfmt.Println("hecho")',
      "}",
    ],
    steps: [
      {
        line: 1,
        caption: "Esta llamada no lleva `defer`: se ejecuta en el momento e imprime «contando».",
        cards: [],
        output: ["contando"],
        printed: true,
      },
      {
        line: 3,
        caption: "Primera vuelta, `i` vale 1. `defer` no imprime: apila `fmt.Println(1)` con el 1 ya calculado.",
        cards: ["1"],
        push: true,
        output: ["contando"],
      },
      {
        line: 3,
        caption: "Segunda vuelta: `fmt.Println(2)` cae encima de la anterior.",
        cards: ["1", "2"],
        push: true,
        output: ["contando"],
      },
      {
        line: 3,
        caption: "Tercera vuelta: `fmt.Println(3)` queda arriba del todo. El bucle termina.",
        cards: ["1", "2", "3"],
        push: true,
        output: ["contando"],
      },
      {
        line: 5,
        caption: "Esta llamada se ejecuta ya e imprime «hecho». Las aplazadas siguen esperando en la pila.",
        cards: ["1", "2", "3"],
        output: ["contando", "hecho"],
        printed: true,
      },
      {
        line: 6,
        caption: "`main` termina y Go ejecuta lo aplazado desde arriba: `fmt.Println(3)` imprime 3.",
        cards: ["1", "2"],
        pop: "3",
        output: ["contando", "hecho", "3"],
        printed: true,
      },
      {
        line: 6,
        caption: "Sale la siguiente tarjeta de la pila: `fmt.Println(2)` imprime 2.",
        cards: ["1"],
        pop: "2",
        output: ["contando", "hecho", "3", "2"],
        printed: true,
      },
      {
        line: 6,
        caption: "La última en salir es la primera que entró: `fmt.Println(1)` imprime 1.",
        cards: [],
        pop: "1",
        output: ["contando", "hecho", "3", "2", "1"],
        printed: true,
      },
    ],
  },
};

class DeferDemo extends Scene {
  static scenes = scenes;

  // Only one card moves per step: the one being pushed falls into its
  // slot, or the one being popped flies to the output and becomes a line.
  draw(paint, step, t) {
    drawPile(paint);
    step.cards.forEach((value, slot) => {
      const isNew = step.push && slot === step.cards.length - 1;
      const y = isNew ? lerp(DROP_FROM, slotY(slot), t) : slotY(slot);
      drawCard(paint, value, CARD_X, y, isNew ? paint.colors.highlight : paint.colors.value);
    });
    if (step.pop && t < 1) {
      const [x, y] = pointOnPath(popPath(step.cards.length, step.output.length - 1), t);
      drawCard(paint, step.pop, x, y, paint.colors.highlight);
    }
    drawOutput(paint, step.output, step.printed, t);
  }
}

// slotY is the top of the card at a slot, counted from the floor up.
function slotY(slot) {
  return PILE.floor - (slot + 1) * CARD_PITCH;
}

// popPath leads the top card up out of the pile, across, and down to the
// output line it is about to print.
function popPath(slot, outputLine) {
  const outputX = OUTPUT.x + Math.floor((OUTPUT.width - CARD_WIDTH) / 2);
  const lineY = OUTPUT.y + 12 + outputLine * OUTPUT_LINE - 4;
  return [[CARD_X, slotY(slot)], [CARD_X, LIFT_Y], [outputX, LIFT_Y], [outputX, lineY]];
}

// drawPile draws the stack as an open-top box with its name below.
function drawPile(paint) {
  const { colors } = paint;
  const right = PILE.x + PILE.width - 1;
  paint.line(PILE.x, PILE.top, PILE.x, PILE.floor, colors.dim);
  paint.line(right, PILE.top, right, PILE.floor, colors.dim);
  paint.line(PILE.x, PILE.floor, right, PILE.floor, colors.dim);
  paint.centerText("pila de defer", PILE.x, PILE.floor + 4, PILE.width, colors.paper);
}

// drawCard draws one deferred call with the value its argument had when
// the defer line ran. The ink fill keeps a moving card readable over others.
function drawCard(paint, value, x, y, color) {
  paint.rect(x, y, CARD_WIDTH, CARD_HEIGHT, paint.colors.ink);
  paint.frame(x, y, CARD_WIDTH, CARD_HEIGHT, color);
  paint.centerText(`fmt.Println(${value})`, x, y + 4, CARD_WIDTH, color);
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
    paint.text(line, x + 4, y + 12 + i * OUTPUT_LINE, isNew ? colors.highlight : colors.paper);
  });
}

customElements.define("defer-demo", DeferDemo);
