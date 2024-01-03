import fs from "fs";
const file = fs.readFileSync(process.argv[2], { encoding: "utf8" }).trim();
const lines = file.split("\n");

const instructions = {};
lines.forEach((line) => {
  let [opr, tid] = line.split(" -> ");
  opr = opr.split(" ");
  instructions[tid] = [
    opr.length === 2 ? "NOT" : opr.length === 3 ? opr[1] : "",
    opr.length === 2 ? opr[1] : opr[0],
    opr.length === 3 ? opr[2] : -1,
  ];
});
const getWireVal = (wire, values) =>
  !isNaN(wire) ? Number.parseInt(wire) : values?.[wire] ?? -1;

const OPERATORS = {
  AND: (v1, v2) => v1 & v2,
  OR: (v1, v2) => v1 | v2,
  LSHIFT: (v1, v2) => (v1 << v2) & 0xffff,
  RSHIFT: (v1, v2) => v1 >> v2,
};

const findVal = (instructions, target, br) => {
  const values = {};
  while (true) {
    let resolved = true;
    for (const tid in instructions) {
      if (values[tid] !== undefined) {
        continue;
      }
      if (values[target] !== undefined) {
        resolved = true;
        break;
      }
      const [gate, op1, op2] = instructions[tid];
      const v1 = getWireVal(op1, values);
      if (v1 === -1) {
        resolved = false;
        continue;
      }
      if (gate == "") {
        values[tid] = v1;
      } else if (gate == "NOT") {
        values[tid] = 65535 - v1;
      } else {
        const v2 = getWireVal(op2, values);
        if (v2 === -1) {
          resolved = false;
          continue;
        }

        values[tid] = OPERATORS[gate](v1, v2);
      }
    }

    if (br) {
      console.log(values);
      break;
    }

    if (resolved) break;
  }

  return values[target];
};

// FOR PART 1
const part1 = findVal(instructions, "a");
console.log("Part 1", part1);
// FOR PART 2
instructions["b"][1] = part1;
const part2 = findVal(instructions, "a");
console.log("Part 2", part2);
