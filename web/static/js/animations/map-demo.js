// <map-demo> shows how a map finds where a key lives. Scene: "buckets" (the
// hash of a key picks a bucket, so writing and reading go straight there,
// and the order of the keys is not the order they were added in).
import { Scene, lerp, pointOnPath } from "./scene.js";

const DROP = 6; // units an entering value falls before it settles

// The keys wait on the left, the hash function sits in the middle and the
// buckets of the map are stacked on the right.
const KEY_X = 6;
const HASH = { x: 48, y: 22, width: 32, height: 14 };
const BUCKET = { x: 100, width: 56, height: 13, rows: [4, 19, 34, 49] };
const OUTPUT = { x: 4, y: 64, width: 84, height: 28 };

const ANA = { key: "Ana", value: "31", bucket: 2, y: 10 };
const LUIS = { key: "Luis", value: "27", bucket: 0, y: 44 };

const scenes = {
  buckets: {
    code: [
      "ages := make(map[string]int)",
      'ages["Ana"] = 31',
      'ages["Luis"] = 27',
      'fmt.Println(ages["Ana"])',
    ],
    steps: [
      {
        line: 0,
        caption: "`make` crea un map vacío. Por dentro, sus pares se guardan en casillas llamadas buckets.",
        stored: [],
        enter: true,
      },
      {
        line: 1,
        caption: "Para guardar \"Ana\", Go calcula su hash, un número que sale de la clave. Ese número elige el bucket 2.",
        stored: [ANA],
        moving: ANA,
      },
      {
        line: 2,
        caption: "\"Luis\" tiene otro hash y cae en el bucket 0. El orden no es el de llegada.",
        stored: [ANA, LUIS],
        moving: LUIS,
      },
      {
        line: 3,
        caption: "Para leer, Go calcula el mismo hash y va directo al bucket 2, sin mirar los demás. Imprime 31.",
        stored: [ANA, LUIS],
        reading: ANA,
        output: ["31"],
        printed: true,
      },
    ],
  },
};

class MapDemo extends Scene {
  static scenes = scenes;

  // A step plays in two halves: the key slides into the hash box, then the
  // hash travels to the bucket it picks. The pair shows up there at the end.
  draw(paint, step, t) {
    const { colors } = paint;
    const drop = step.enter ? lerp(DROP, 0, t) : 0;
    const active = step.moving ?? step.reading;
    const toHash = Math.min(1, t * 2);
    const toBucket = Math.max(0, t * 2 - 1);
    drawHash(paint, active && t < 1);
    drawBuckets(paint, step, drop, t);

    if (active && toHash < 1) {
      const text = step.moving ? `${active.key}: ${active.value}` : active.key;
      const endX = HASH.x - paint.textWidth(text) - 2;
      const hashY = HASH.y + (HASH.height - 5) / 2;
      paint.text(text, lerp(KEY_X, endX, toHash), lerp(active.y, hashY, toHash), colors.highlight);
    }
    if (active && toBucket > 0 && t < 1) {
      const [x, y] = pointOnPath(hashPath(active.bucket), toBucket);
      paint.dot(x, y, colors.highlight);
    }
    if (step.output) drawOutput(paint, step.output, step.printed, t);
  }
}

// drawHash draws the hash function as a box the key passes through.
function drawHash(paint, busy) {
  const { colors } = paint;
  const color = busy ? colors.highlight : colors.dim;
  paint.frame(HASH.x, HASH.y, HASH.width, HASH.height, color);
  paint.centerText("hash", HASH.x, HASH.y + 4.5, HASH.width, busy ? color : colors.paper);
}

// drawBuckets draws the buckets with their number and the pairs stored so
// far. The pair moving now appears in its bucket when it arrives.
function drawBuckets(paint, step, drop, t) {
  const { colors } = paint;
  BUCKET.rows.forEach((rowY, bucket) => {
    const y = rowY - drop;
    const pair = step.stored.find((p) => p.bucket === bucket);
    const isTarget = (step.moving ?? step.reading)?.bucket === bucket;
    const arrived = !isTarget || t >= 1;
    const color = isTarget ? colors.highlight : pair ? colors.value : colors.dim;
    paint.text(String(bucket), BUCKET.x - 6, y + 4, colors.paper);
    paint.frame(BUCKET.x, y, BUCKET.width, BUCKET.height, color);
    const isNew = pair === step.moving;
    if (pair && (arrived || !isNew)) {
      paint.centerText(`${pair.key}: ${pair.value}`, BUCKET.x, y + 4, BUCKET.width, isTarget ? color : colors.value);
    }
  });
}

// hashPath leads from the hash box to the left edge of a bucket.
function hashPath(bucket) {
  const fromX = HASH.x + HASH.width + 1;
  const fromY = HASH.y + HASH.height / 2;
  const toY = BUCKET.rows[bucket] + BUCKET.height / 2;
  const turnX = fromX + 6;
  return [[fromX, fromY], [turnX, fromY], [turnX, toY], [BUCKET.x - 8, toY]];
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

customElements.define("map-demo", MapDemo);
