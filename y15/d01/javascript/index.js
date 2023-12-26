import fs from "fs";
const file = fs.readFileSync(process.argv[2], { encoding: "utf8" });
// parse the input
const lines = file.trim().split("\n");
// convert into character list
const input = lines[0].split("");
let pos = 0;
const floor = input.reduce((acc, p, i) => {
  // check if the current acc = -1, then in last index, we reached the basement
  // capture the position
  if (acc == -1 && pos == 0) pos = i;
  // add the value based on the symbol. ( -> 1 ) -> -1
  return acc + (p == "(" ? 1 : -1);
}, 0);

console.log("Part 1", floor);
console.log("Part 2", pos);
