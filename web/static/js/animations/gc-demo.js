// <gc-demo> shows how the garbage collector frees memory. Scene:
// "mark-sweep" (it marks what the variables can reach and frees the rest).
import { Scene, lerp, pointOnPath } from "./scene.js";

const DROP = 6; // units an entering node falls before it settles

const OUTPUT = { x: 4, y: 62, width: 152, height: 30 };
const HEAP = { x: 36, y: 4, width: 120, height: 52 };
const A_BOX = { x: 6, y: 22, width: 20, height: 14 };
const NODE = { width: 24, height: 14 };

// The three nodes in the heap and the pointers between them.
const N1 = { id: "n1", value: "1", x: 48, y: 22 };
const N2 = { id: "n2", value: "2", x: 104, y: 10 };
const N3 = { id: "n3", value: "3", x: 104, y: 36 };
const A_TO_N1 = { from: "a", points: [[26, 29], [47, 29]] };
const N1_TO_N2 = { from: "n2", points: [[72, 26], [103, 17]] };
const N1_TO_N3 = { from: "n3", points: [[72, 32], [103, 43]] };

// The path the collector walks while it marks: from a, through every node
// it can reach.
const MARK_PATH = [[26, 29], [47, 29], [60, 29], [72, 32], [103, 43]];

const scenes = {
  "mark-sweep": {
    code: [
      "a := &Node{Value: 1}",
      "a.Next = &Node{Value: 2}",
      "a.Next = &Node{Value: 3}",
      "runtime.GC()",
      "fmt.Println(a.Value, a.Next.Value)",
    ],
    steps: [
      {
        line: 0,
        caption: "`&Node{...}` reserva un nodo en el heap, y la variable `a` apunta a él.",
        nodes: [N1],
        arrows: [A_TO_N1],
        enter: "n1",
      },
      {
        line: 1,
        caption: "`a.Next` apunta a un nodo nuevo, con el 2.",
        nodes: [N1, N2],
        arrows: [A_TO_N1, N1_TO_N2],
        enter: "n2",
      },
      {
        line: 2,
        caption: "Ahora `a.Next` apunta a otro nodo, con el 3. Ya nadie apunta al nodo 2.",
        nodes: [N1, N2, N3],
        arrows: [A_TO_N1, N1_TO_N3],
        enter: "n3",
      },
      {
        line: 3,
        caption: "Mark: el recolector parte de tus variables y marca todo lo que alcanza siguiendo punteros.",
        nodes: [N1, N2, N3],
        arrows: [A_TO_N1, N1_TO_N3],
        marked: ["n1", "n3"],
        walk: true,
      },
      {
        line: 3,
        caption: "Sweep: el nodo 2 no está marcado. Es basura, y el recolector libera su memoria.",
        nodes: [N1, { ...N2, garbage: true }, N3],
        arrows: [A_TO_N1, N1_TO_N3],
        marked: ["n1", "n3"],
      },
      {
        line: 4,
        caption: "Ese hueco queda libre para lo próximo que reserves. El programa sigue e imprime 1 3.",
        nodes: [N1, N3],
        arrows: [A_TO_N1, N1_TO_N3],
        output: ["1 3"],
        printed: true,
      },
    ],
  },
};

class GcDemo extends Scene {
  static scenes = scenes;

  // A step plays in two halves: a new node drops in, then its arrow grows.
  // Marking walks the whole step; marks appear when the walk ends.
  draw(paint, step, t) {
    const { colors } = paint;
    const settle = Math.min(1, t * 2);
    const grow = Math.max(0, t * 2 - 1);
    const marked = new Set(step.walk && t < 1 ? [] : step.marked ?? []);

    paint.frame(HEAP.x, HEAP.y, HEAP.width, HEAP.height, colors.paper, true);
    paint.text("heap", HEAP.x + 4, HEAP.y + 4, colors.paper);

    paint.text("a", A_BOX.x, A_BOX.y - 7, colors.paper);
    paint.frame(A_BOX.x, A_BOX.y, A_BOX.width, A_BOX.height, colors.pointer);
    paint.centerText("•", A_BOX.x, A_BOX.y + 4.5, A_BOX.width, colors.pointer);

    for (const node of step.nodes) {
      const drop = node.id === step.enter ? lerp(DROP, 0, settle) : 0;
      drawNode(paint, node, drop, node.id === step.enter, marked.has(node.id), t);
    }
    for (const arrow of step.arrows) {
      paint.arrow(arrow.points, colors.pointer, arrow.from === step.enter ? grow : 1);
    }
    if (step.walk && t < 1) {
      const [x, y] = pointOnPath(MARK_PATH, t);
      paint.dot(x, y, colors.highlight);
    }
    if (step.output) drawOutput(paint, step.output, step.printed, t);
  }
}

// drawNode draws a node of the list. A new node glows, a marked one is
// live (green) and garbage turns red as the sweep frees it.
function drawNode(paint, node, drop, entering, marked, t) {
  const { colors } = paint;
  let color = colors.value;
  if (entering) color = colors.highlight;
  if (marked) color = colors.ok;
  if (node.garbage) color = t < 1 ? colors.error : colors.dim;
  const y = node.y - drop;
  paint.frame(node.x, y, NODE.width, NODE.height, color);
  paint.centerText(node.value, node.x, y + 4.5, NODE.width, color);
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

customElements.define("gc-demo", GcDemo);
