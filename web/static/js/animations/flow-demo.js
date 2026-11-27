// <flow-demo> shows the order in which Go runs the parts of a statement.
// Scenes: "for" (start once, then condition, body and step on every round)
// and "switch" (the cases are checked top-down and only the first match
// runs).
import { Scene, lerp, pointOnPath } from "./scene.js";

const DROP = 6; // units an entering box falls before it settles
const BOX = { width: 32, height: 16 };

// for: the three parts of the loop header, side by side, with i below.
const PARTS = [
  { id: "init", code: "i := 0", label: "inicio", x: 4 },
  { id: "cond", code: "i < 3", label: "condición", x: 56 },
  { id: "post", code: "i++", label: "paso", x: 108 },
];
const PART = { y: 6, width: 48, height: 12 };
const I_BOX = { x: 8, y: 50 };
const FOR_OUTPUT = { x: 84, y: 40, width: 72, height: 52 };

// switch: day on the left, the cases on the right, top to bottom.
const DAY_BOX = { x: 8, y: 22 };
const CASES = ["case 1, 2, 3, 4", "case 6, 7", "default"];
const CASE = { x: 56, width: 100, height: 12, rows: [8, 24, 40] };
const SWITCH_OUTPUT = { x: 4, y: 62, width: 152, height: 30 };

const scenes = {
  for: {
    output: FOR_OUTPUT,
    code: ["for i := 0; i < 3; i++ {", '\tfmt.Println("vuelta", i)', "}", 'fmt.Println("fin")'],
    steps: [
      {
        line: 0,
        caption: "El inicio, `i := 0`, se ejecuta una sola vez: crea `i` con 0.",
        part: "init",
        i: "0",
        enter: true,
        output: [],
      },
      ...round(0),
      ...round(1),
      ...round(2),
      {
        line: 0,
        caption: "3 < 3 es `false`: el bucle termina sin ejecutar el cuerpo otra vez.",
        part: "cond",
        i: "3",
        verdict: false,
        output: printedUpTo(3),
      },
      {
        line: 3,
        caption: "Go sigue después de la llave de cierre e imprime «fin». `i` ya no existe.",
        output: [...printedUpTo(3), "fin"],
        printed: true,
      },
    ],
  },

  switch: {
    output: SWITCH_OUTPUT,
    code: [
      "day := 6",
      "switch day {",
      "case 1, 2, 3, 4:",
      '\tfmt.Println("laborable")',
      "case 6, 7:",
      '\tfmt.Println("fin de semana")',
      "default:",
      '\tfmt.Println("otro")',
      "}",
    ],
    steps: [
      {
        line: 0,
        caption: "`day` vale 6. El `switch` va a comparar ese valor con cada `case`.",
        verdicts: [],
        enter: true,
      },
      {
        line: 2,
        caption: "Empieza por arriba: `case 1, 2, 3, 4`. El 6 no está, así que baja al siguiente.",
        verdicts: ["no"],
        checking: 0,
      },
      {
        line: 4,
        caption: "`case 6, 7` sí tiene el 6: este es el caso que se ejecuta.",
        verdicts: ["no", "sí"],
        checking: 1,
      },
      {
        line: 5,
        caption: "Solo se ejecuta el cuerpo de ese `case`: imprime «fin de semana».",
        verdicts: ["no", "sí"],
        output: ["fin de semana"],
        printed: true,
      },
      {
        line: 8,
        caption: "Y el `switch` termina. `default` ni siquiera se comprueba.",
        verdicts: ["no", "sí"],
        output: ["fin de semana"],
      },
    ],
  },
};

// round builds the three steps of one round of the loop: the condition
// lets it in, the body prints and the step adds 1 to i.
function round(n) {
  const i = String(n);
  return [
    {
      line: 0,
      caption: `Antes de cada vuelta se comprueba la condición: ${n} < 3 es \`true\`, así que entra.`,
      part: "cond",
      i,
      verdict: true,
      output: printedUpTo(n),
    },
    {
      line: 1,
      caption: `El cuerpo se ejecuta con \`i\` valiendo ${n}: imprime «vuelta ${n}».`,
      i,
      output: printedUpTo(n + 1),
      printed: true,
    },
    {
      line: 0,
      caption: `Al final de la vuelta se ejecuta el paso, \`i++\`: \`i\` pasa a valer ${n + 1}.`,
      part: "post",
      i: String(n + 1),
      was: i,
      output: printedUpTo(n + 1),
    },
  ];
}

// printedUpTo is what the loop has printed after its first n rounds.
function printedUpTo(n) {
  return Array.from({ length: n }, (_, i) => `vuelta ${i}`);
}

class FlowDemo extends Scene {
  static scenes = scenes;

  draw(paint, step, t) {
    const name = this.getAttribute("scene");
    const drawScene = { for: drawFor, switch: drawSwitch }[name];
    drawScene(paint, step, t);
    if (step.output) drawOutput(paint, scenes[name].output, step.output, step.printed, t);
  }
}

// drawFor draws the loop header with the part that runs now lit, the
// verdict of the condition under it, and i while it exists.
function drawFor(paint, step, t) {
  const { colors } = paint;
  for (const part of PARTS) {
    const color = part.id === step.part ? colors.highlight : colors.dim;
    paint.frame(part.x, PART.y, PART.width, PART.height, color);
    paint.centerText(part.code, part.x, PART.y + 3.5, PART.width, part.id === step.part ? color : colors.paper);
    paint.centerText(part.label, part.x, PART.y + PART.height + 3, PART.width, colors.paper);
  }
  if (step.verdict !== undefined) {
    const cond = PARTS[1];
    const color = step.verdict ? colors.ok : colors.error;
    paint.centerText(String(step.verdict), cond.x, PART.y + PART.height + 12, PART.width, color);
  }
  if (step.i === undefined) return;
  const changing = step.was !== undefined;
  const y = I_BOX.y - (step.enter ? lerp(DROP, 0, t) : 0);
  const value = changing && t < 1 ? step.was : step.i;
  drawBox(paint, "i", value, I_BOX.x, y, step.enter || changing ? colors.highlight : colors.value);
}

// drawSwitch draws day and the cases. The value travels to the case being
// checked, and the verdict appears when it arrives.
function drawSwitch(paint, step, t) {
  const { colors } = paint;
  const y = DAY_BOX.y - (step.enter ? lerp(DROP, 0, t) : 0);
  drawBox(paint, "day", "6", DAY_BOX.x, y, step.enter ? colors.highlight : colors.value);

  CASES.forEach((text, row) => {
    const top = CASE.rows[row];
    const verdict = step.verdicts[row];
    const isChecking = row === step.checking;
    const arrived = !isChecking || t >= 1;
    const color = isChecking ? colors.highlight : verdict ? colors.value : colors.dim;
    paint.frame(CASE.x, top, CASE.width, CASE.height, color);
    paint.text(text, CASE.x + 4, top + 3.5, verdict ? color : colors.paper);
    if (verdict && arrived) {
      const verdictColor = verdict === "sí" ? colors.ok : colors.error;
      paint.text(verdict, CASE.x + CASE.width - 14, top + 3.5, verdictColor);
    }
  });

  if (step.checking !== undefined && t < 1) {
    const from = [DAY_BOX.x + BOX.width + 2, DAY_BOX.y + BOX.height / 2];
    const to = [CASE.x - 2, CASE.rows[step.checking] + CASE.height / 2];
    const [x, dotY] = pointOnPath([from, [from[0] + 6, from[1]], [to[0] - 6, to[1]], to], t);
    paint.dot(x, dotY, colors.highlight);
  }
}

// drawBox draws a variable as a box with its name above it.
function drawBox(paint, label, value, x, y, color) {
  paint.text(label, x, y - 7, paint.colors.paper);
  paint.frame(x, y, BOX.width, BOX.height, color);
  paint.centerText(value, x, y + 5.5, BOX.width, color);
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

customElements.define("flow-demo", FlowDemo);
