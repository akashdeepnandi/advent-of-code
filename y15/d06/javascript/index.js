import fs from "fs";
const file = fs.readFileSync(process.argv[2], { encoding: "utf8" }).trim();
const lines = file.split("\n");

const parseInstruction = (s) => {
  let mode = -1,
    si = 2,
    ei = 4;
  if (s.length == 4) {
    (mode = 0), (si = 1), (ei = 3);
  } else if (s[1] == "on") {
    mode = 1;
  } else {
    mode = 2;
  }

  const r1 = s[si].split(",").map((n) => parseInt(n));
  const r2 = s[ei].split(",").map((n) => parseInt(n));

  return [mode, ...r1, ...r2];
};
const operateLight = (v, mode) => {
  switch (mode) {
    case 0:
      return v == 1 ? 0 : 1;
    case 1:
      return 1;
    case 2:
      return 0;
  }
};

const operateBrightness = (v, mode) => {
  switch (mode) {
    case 0:
      return v + 2;
    case 1:
      return v + 1;
    case 2:
      return Math.max(0, v - 1);
  }
};

const operateDecoration = (lines, operate) => {
  let grid = [];
  for (let i = 0; i < 1000; i++) {
    grid.push(Array(1000).fill(0));
  }

  lines.forEach((line) => {
    const [mode, startR, startC, endR, endC] = parseInstruction(
      line.split(" "),
    );
    for (let r = startR; r <= endR; r++) {
      for (let c = startC; c <= endC; c++) {
        grid[r][c] = operate(grid[r][c], mode);
      }
    }
  });

  let count = grid.reduce(
    (rsum, row) => rsum + row.reduce((csum, col) => csum + col, 0),
    0,
  );
  return count;
};

console.log("Part 1:", operateDecoration(lines, operateLight));
console.log("Part 1:", operateDecoration(lines, operateBrightness));
