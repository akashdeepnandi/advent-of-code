import fs from "fs";
const file = fs.readFileSync(process.argv[2], { encoding: "utf8" }).trim();
const lines = file.split("\n");

const p1 = lines
  .map((original) => original.length - eval(original).length)
  .reduce((acc, len) => acc + len, 0);
console.log("Part 1", p1);

const p2 = lines
  .map((original) => JSON.stringify(original).length - original.length)
  .reduce((acc, len) => acc + len, 0);
console.log("Part 1", p2);
