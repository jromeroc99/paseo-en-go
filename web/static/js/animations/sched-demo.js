// <sched-demo> shows who runs your code and when. Scenes: "threads" (each
// process has its own memory and its threads share it), "goroutines" (main
// returning ends every goroutine; a WaitGroup makes it wait), "gmp" (an M
// runs the Gs queued on its P, and an idle P steals half of another queue)
// and "race" (two goroutines lose an update; a Mutex lets only one in).
import { Scene, lerp } from "./scene.js";

const DROP = 6; // units an entering box falls before it settles

const OUTPUT = { x: 4, y: 62, width: 152, height: 30 };

// threads: two processes side by side, each with its memory on top and its
// threads below.
const PROCESSES = {
  A: { x: 4, y: 4, width: 74, height: 88 },
  B: { x: 82, y: 4, width: 74, height: 88 },
};
const THREAD = { width: 30, height: 12 };

// goroutines: one row per goroutine and the WaitGroup's counter on the right.
const ROW = { x: 4, y: 4, width: 100, height: 12, gap: 15 };
const COUNTER = { x: 112, y: 4, width: 44, height: 27 };

// gmp: two processors, each with its thread (M) and the G it runs above
// its queue.
const PROCS = {
  P0: { x: 4, y: 24, width: 74, height: 24 },
  P1: { x: 82, y: 24, width: 74, height: 24 },
};
const CHIP = { width: 12, height: 10, gap: 14 };

// race: the counter on top, the Mutex under it and a goroutine on each side.
const COUNT = { x: 64, y: 12, width: 32, height: 14 };
const MUTEX = { x: 64, y: 40, width: 32, height: 14 };
const GOROUTINES = {
  g1: { x: 4, y: 62, width: 44, height: 26 },
  g2: { x: 112, y: 62, width: 44, height: 26 },
};
const READ = { g1: [[70, 27], [26, 61]], g2: [[90, 27], [134, 61]] };
const WRITE = { g1: [[30, 61], [74, 27]], g2: [[130, 61], [86, 27]] };

const scenes = {
  threads: {
    code: ["// proceso A", "x := 0   // hilo 1", "x++      // hilo 2", "// proceso B", "y := 5   // hilo 1"],
    steps: [
      {
        line: 0,
        caption: "Un proceso es un programa en ejecución. El sistema le da su propia memoria.",
        processes: ["A"],
        changed: ["A"],
      },
      {
        line: 1,
        caption: "El hilo 1 de A ejecuta `x := 0` y guarda `x` en la memoria de A.",
        processes: ["A"],
        vars: { A: { name: "x", value: "0" } },
        threads: ["A1"],
        uses: ["A1"],
        changed: ["A1"],
      },
      {
        line: 2,
        caption: "El hilo 2 hace `x++` sobre la misma `x`: los hilos de un proceso comparten su memoria.",
        processes: ["A"],
        vars: { A: { name: "x", value: "1" } },
        threads: ["A1", "A2"],
        uses: ["A2"],
        changed: ["A2", "xA"],
      },
      {
        line: 3,
        caption: "Otro proceso, B, tiene otra memoria, separada de la de A.",
        processes: ["A", "B"],
        vars: { A: { name: "x", value: "1" } },
        threads: ["A1", "A2"],
        changed: ["B"],
      },
      {
        line: 4,
        caption: "El hilo de B guarda `y` en su memoria. La `x` de A queda fuera de su alcance.",
        processes: ["A", "B"],
        vars: { A: { name: "x", value: "1" }, B: { name: "y", value: "5" } },
        threads: ["A1", "A2", "B1"],
        uses: ["B1"],
        changed: ["B1"],
        blocked: true,
      },
    ],
  },

  goroutines: {
    code: [
      'go say("hola")',
      'go say("adiós")',
      "return  // main termina",
      "// con un WaitGroup:",
      "wg.Add(2)",
      'go say("hola")   // wg.Done() al acabar',
      'go say("adiós")  // wg.Done() al acabar',
      "wg.Wait()",
    ],
    steps: [
      {
        line: 0,
        caption: '`go say("hola")` arranca una goroutine nueva, y `main` sigue sin esperarla.',
        rows: [run("main", "ejecuta"), run("g1", 'say("hola")')],
        changed: ["g1"],
      },
      {
        line: 1,
        caption: "Otra más. Ahora hay tres goroutines: `main` y dos que acaban de empezar.",
        rows: [run("main", "ejecuta"), run("g1", 'say("hola")'), run("g2", 'say("adiós")')],
        changed: ["g2"],
      },
      {
        line: 2,
        caption: "`main` termina y el programa acaba con ella: las otras dos desaparecen sin imprimir nada.",
        rows: [done("main", "termina"), gone("g1"), gone("g2")],
        changed: ["g1", "g2"],
        output: [],
      },
      {
        line: 4,
        caption: "Con un `sync.WaitGroup`, `wg.Add(2)` apunta que hay dos goroutines pendientes.",
        rows: [run("main", "ejecuta")],
        counter: 2,
        changed: ["wg"],
        output: [],
      },
      {
        line: 5,
        caption: "La primera goroutine arranca. Al acabar llamará a `wg.Done()`.",
        rows: [run("main", "ejecuta"), run("g1", 'say("hola")')],
        counter: 2,
        changed: ["g1"],
        output: [],
      },
      {
        line: 6,
        caption: "La segunda también arranca.",
        rows: [run("main", "ejecuta"), run("g1", 'say("hola")'), run("g2", 'say("adiós")')],
        counter: 2,
        changed: ["g2"],
        output: [],
      },
      {
        line: 7,
        caption: "`wg.Wait()` bloquea `main` hasta que el contador llegue a 0.",
        rows: [wait("main", "espera"), run("g1", 'say("hola")'), run("g2", 'say("adiós")')],
        counter: 2,
        changed: ["main"],
        output: [],
      },
      {
        line: 5,
        caption: "La primera imprime hola y llama a `wg.Done()`: queda una pendiente.",
        rows: [wait("main", "espera"), done("g1", "terminada"), run("g2", 'say("adiós")')],
        counter: 1,
        changed: ["g1", "wg"],
        output: ["hola"],
        printed: true,
      },
      {
        line: 6,
        caption: "La segunda imprime adiós y llama a `wg.Done()`: el contador llega a 0.",
        rows: [wait("main", "espera"), done("g1", "terminada"), done("g2", "terminada")],
        counter: 0,
        changed: ["g2", "wg"],
        output: ["hola", "adiós"],
        printed: true,
      },
      {
        line: 7,
        caption: "`main` sigue y acaba con todo impreso. El orden de hola y adiós puede cambiar.",
        rows: [run("main", "sigue"), done("g1", "terminada"), done("g2", "terminada")],
        counter: 0,
        changed: ["main"],
        output: ["hola", "adiós"],
      },
    ],
  },

  gmp: {
    code: ["for i := range 5 {", "\tgo task(i)", "}", "// cada M ejecuta las G de su P", "// P1 sin G: roba a P0"],
    steps: [
      {
        line: 1,
        caption: "`go task(i)` crea una G, una goroutine, y la pone en la cola de un P.",
        chips: [queued("G0", "P0", 0)],
        changed: ["G0"],
      },
      {
        line: 1,
        caption: "El bucle crea cinco. Todas esperan en la cola de `P0`.",
        chips: [0, 1, 2, 3, 4].map((i) => queued(`G${i}`, "P0", i)),
        changed: ["G1", "G2", "G3", "G4"],
      },
      {
        line: 3,
        caption: "`M0`, un hilo del sistema, coge la primera G de la cola de `P0` y la ejecuta.",
        chips: [
          { id: "G0", from: ["P0", 0], at: ["P0", "run"] },
          ...[1, 2, 3, 4].map((i) => ({ id: `G${i}`, from: ["P0", i], at: ["P0", i - 1] })),
        ],
        changed: ["G0"],
      },
      {
        line: 3,
        caption: "`P1` también tiene su hilo, `M1`, pero su cola está vacía: no tiene nada que hacer.",
        chips: [running("G0", "P0"), ...[1, 2, 3, 4].map((i) => queued(`G${i}`, "P0", i - 1))],
        changed: ["P1"],
      },
      {
        line: 4,
        caption: "`P1` roba la mitad de la cola de `P0`: `G1` y `G2` pasan a la suya.",
        chips: [
          running("G0", "P0"),
          { id: "G1", from: ["P0", 0], at: ["P1", 0] },
          { id: "G2", from: ["P0", 1], at: ["P1", 1] },
          { id: "G3", from: ["P0", 2], at: ["P0", 0] },
          { id: "G4", from: ["P0", 3], at: ["P0", 1] },
        ],
        changed: ["G1", "G2"],
      },
      {
        line: 4,
        caption: "`M1` ejecuta `G1`. Los dos hilos trabajan a la vez y nadie se queda parado.",
        chips: [
          running("G0", "P0"),
          { id: "G1", from: ["P1", 0], at: ["P1", "run"] },
          { id: "G2", from: ["P1", 1], at: ["P1", 0] },
          queued("G3", "P0", 0),
          queued("G4", "P0", 1),
        ],
        changed: ["G1"],
      },
    ],
  },

  race: {
    code: [
      "go func() { count++ }()  // g1",
      "go func() { count++ }()  // g2",
      "// con un Mutex:",
      "mu.Lock()",
      "count++",
      "mu.Unlock()",
    ],
    steps: [
      {
        line: 0,
        caption: "`count++` son tres pasos: leer, sumar y escribir. `g1` lee count: 0.",
        count: "0",
        copies: { g1: "0" },
        arrow: ["read", "g1"],
        changed: ["g1"],
      },
      {
        line: 1,
        caption: "Antes de que `g1` escriba, `g2` también lee count: 0.",
        count: "0",
        copies: { g1: "0", g2: "0" },
        arrow: ["read", "g2"],
        changed: ["g2"],
      },
      {
        line: 0,
        caption: "`g1` suma 1 a su copia y escribe: count vale 1.",
        count: "1",
        copies: { g1: "1", g2: "0" },
        arrow: ["write", "g1"],
        changed: ["count"],
      },
      {
        line: 1,
        caption: "`g2` suma 1 a su 0 y escribe 1. Dos incrementos, y count vale 1: se ha perdido uno.",
        count: "1",
        copies: { g1: "1", g2: "1" },
        arrow: ["write", "g2"],
        lost: true,
      },
      {
        line: 3,
        caption: "Con un `sync.Mutex`, `g1` hace `Lock()` y entra. Solo quien tiene el candado toca count.",
        count: "0",
        copies: { g1: "", g2: "" },
        holder: "g1",
        changed: ["mu"],
      },
      {
        line: 3,
        caption: "`g2` también llama a `Lock()`, pero espera fuera hasta que `g1` salga.",
        count: "0",
        copies: { g1: "", g2: "" },
        holder: "g1",
        waiting: "g2",
        changed: ["g2"],
      },
      {
        line: 4,
        caption: "Dentro, `g1` lee, suma y escribe sin que nadie se cuele: count vale 1.",
        count: "1",
        copies: { g1: "1", g2: "" },
        holder: "g1",
        waiting: "g2",
        arrow: ["write", "g1"],
        changed: ["count"],
      },
      {
        line: 5,
        caption: "`g1` hace `Unlock()` y suelta el candado. Ahora entra `g2`.",
        count: "1",
        copies: { g1: "1", g2: "" },
        holder: "g2",
        changed: ["mu"],
      },
      {
        line: 4,
        caption: "`g2` lee 1 y escribe 2. No se pierde ninguna actualización: count vale 2.",
        count: "2",
        copies: { g1: "1", g2: "2" },
        holder: "g2",
        arrow: ["write", "g2"],
        changed: ["count"],
        safe: true,
      },
    ],
  },
};

// The rows of the goroutines scene: a goroutine and what it is doing.
function run(name, status) {
  return { name, status, state: "run" };
}
function wait(name, status) {
  return { name, status, state: "wait" };
}
function done(name, status) {
  return { name, status, state: "done" };
}
function gone(name) {
  return { name, status: "desaparece", state: "gone" };
}

// The chips of the gmp scene: a G waiting in a queue or running on a P.
function queued(id, proc, i) {
  return { id, at: [proc, i] };
}
function running(id, proc) {
  return { id, at: [proc, "run"] };
}

class SchedDemo extends Scene {
  static scenes = scenes;

  draw(paint, step, t) {
    const drawScene = {
      threads: drawThreads,
      goroutines: drawGoroutines,
      gmp: drawGmp,
      race: drawRace,
    }[this.getAttribute("scene")];
    drawScene(paint, step, t);
  }
}

// changeColor is the colour of something this step changes: it glows while
// it moves and settles into its normal colour.
function changeColor(paint, step, id, t, normal) {
  return step.changed?.includes(id) && t < 1 ? paint.colors.highlight : normal;
}

// drop is how far above its place something this step adds still is.
function drop(step, id, t) {
  return step.changed?.includes(id) ? lerp(DROP, 0, t) : 0;
}

// drawThreads draws each process with its memory and its threads, and an
// arrow from each thread that uses a variable to it.
function drawThreads(paint, step, t) {
  const { colors } = paint;
  for (const name of step.processes) {
    const p = PROCESSES[name];
    const y = p.y - drop(step, name, t);
    const color = changeColor(paint, step, name, t, colors.paper);
    paint.frame(p.x, y, p.width, p.height, color);
    paint.text(`proceso ${name}`, p.x + 4, y + 4, color);
    paint.frame(p.x + 4, y + 14, p.width - 8, 30, colors.dim, true);
    paint.text("memoria", p.x + 8, y + 18, colors.dim);

    const variable = step.vars?.[name];
    if (variable) {
      const valueColor = changeColor(paint, step, `x${name}`, t, colors.value);
      paint.text(variable.name, p.x + 12, y + 30, colors.paper);
      paint.frame(p.x + 22, y + 27, 20, 12, valueColor);
      paint.centerText(variable.value, p.x + 22, y + 30.5, 20, valueColor);
    }
  }

  for (const id of step.threads ?? []) {
    const p = PROCESSES[id[0]];
    const x = p.x + 4 + (Number(id[1]) - 1) * 36;
    const y = p.y + 60 - drop(step, id, t);
    const color = changeColor(paint, step, id, t, colors.value);
    paint.frame(x, y, THREAD.width, THREAD.height, color);
    paint.centerText(`hilo ${id[1]}`, x, y + 3.5, THREAD.width, color);
    if (step.uses?.includes(id)) {
      paint.arrow([[x + THREAD.width / 2, p.y + 59], [p.x + 32, p.y + 40]], colors.pointer, t);
    }
  }

  // B's thread cannot reach A's memory: the line stops at B's edge.
  if (step.blocked && t === 1) {
    paint.line(94, 63, 84, 52, colors.error);
    paint.centerText("✕", 76, 47, 8, colors.error);
  }
}

// drawGoroutines draws one row per goroutine, the WaitGroup's counter
// and what the program printed.
function drawGoroutines(paint, step, t) {
  const { colors } = paint;
  step.rows.forEach((row, i) => {
    const y = ROW.y + i * ROW.gap - (row.state === "gone" ? 0 : drop(step, row.name, t));
    const color = rowColor(paint, step, row, t);
    paint.frame(ROW.x, y, ROW.width, ROW.height, color, row.state === "wait");
    paint.text(row.name, ROW.x + 4, y + 3.5, color);
    paint.text(row.status, ROW.x + 28, y + 3.5, color);
  });

  if (step.counter !== undefined) {
    const color = changeColor(paint, step, "wg", t, colors.value);
    paint.frame(COUNTER.x, COUNTER.y, COUNTER.width, COUNTER.height, color);
    paint.text("wg", COUNTER.x + 4, COUNTER.y + 4, colors.paper);
    paint.centerText(step.counter, COUNTER.x, COUNTER.y + 15, COUNTER.width, color);
  }
  if (step.output) drawOutput(paint, step.output, step.printed, t);
}

// rowColor gives each state of a goroutine its colour. A goroutine that
// vanishes before it could finish is an error.
function rowColor(paint, step, row, t) {
  const { colors } = paint;
  if (row.state === "gone") return colors.error;
  const normal = { run: colors.value, wait: colors.pointer, done: colors.dim }[row.state];
  return changeColor(paint, step, row.name, t, normal);
}

// drawGmp draws both processors, their threads and every G, moving the
// ones this step takes from one place to another.
function drawGmp(paint, step, t) {
  const { colors } = paint;
  for (const [name, p] of Object.entries(PROCS)) {
    const color = changeColor(paint, step, name, t, colors.paper);
    paint.frame(p.x, p.y, p.width, p.height, color);
    paint.text(name, p.x + 4, p.y + 3, color);
    const m = `M${name[1]}`;
    paint.frame(p.x + 4, 4, 30, 12, colors.pointer);
    paint.centerText(m, p.x + 4, 7.5, 30, colors.pointer);
    paint.line(p.x + 19, 16, p.x + 19, p.y, colors.pointer);
    paint.frame(p.x + 40, 3, 30, 14, colors.dim, true);
  }

  for (const chip of step.chips) {
    const to = chipPosition(chip.at);
    const from = chip.from ? chipPosition(chip.from) : to;
    const x = lerp(from.x, to.x, t);
    const y = lerp(from.y, to.y, t) - (chip.from ? 0 : drop(step, chip.id, t));
    const color = changeColor(paint, step, chip.id, t, colors.value);
    paint.frame(x, y, CHIP.width, CHIP.height, color);
    paint.centerText(chip.id, x, y + 2.5, CHIP.width, color);
  }

  paint.text("G  una goroutine", 4, 60, colors.dim);
  paint.text("M  un hilo del sistema", 4, 70, colors.dim);
  paint.text("P  un procesador lógico", 4, 80, colors.dim);
}

// chipPosition is where a G sits: its place in a queue, or "run" for the
// slot next to the thread that runs it.
function chipPosition([proc, place]) {
  const p = PROCS[proc];
  if (place === "run") return { x: p.x + 49, y: 5 };
  return { x: p.x + 4 + place * CHIP.gap, y: p.y + 11 };
}

// drawRace draws the shared counter, each goroutine's copy of it and, once
// there is one, the Mutex with the goroutine that holds it.
function drawRace(paint, step, t) {
  const { colors } = paint;
  let countColor = changeColor(paint, step, "count", t, colors.value);
  if (step.lost && t === 1) countColor = colors.error;
  if (step.safe && t === 1) countColor = colors.ok;
  paint.centerText("count", COUNT.x, COUNT.y - 8, COUNT.width, colors.paper);
  paint.frame(COUNT.x, COUNT.y, COUNT.width, COUNT.height, countColor);
  paint.centerText(step.count, COUNT.x, COUNT.y + 4.5, COUNT.width, countColor);

  if (step.holder) {
    const color = changeColor(paint, step, "mu", t, colors.pointer);
    paint.frame(MUTEX.x, MUTEX.y, MUTEX.width, MUTEX.height, color);
    paint.centerText(`mu: ${step.holder}`, MUTEX.x, MUTEX.y + 4.5, MUTEX.width, color);
  }

  for (const [name, g] of Object.entries(GOROUTINES)) {
    const waiting = step.waiting === name;
    const color = changeColor(paint, step, name, t, waiting ? colors.pointer : colors.value);
    paint.frame(g.x, g.y, g.width, g.height, color, waiting);
    paint.text(name, g.x + 4, g.y + 4, color);
    const copy = waiting ? "espera" : step.copies[name] && `copia: ${step.copies[name]}`;
    if (copy) paint.text(copy, g.x + 4, g.y + 15, color);
  }

  if (step.arrow) {
    const [kind, name] = step.arrow;
    const points = kind === "read" ? READ[name] : WRITE[name];
    paint.arrow(points, kind === "read" ? colors.value : colors.pointer, t);
  }
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

customElements.define("sched-demo", SchedDemo);
