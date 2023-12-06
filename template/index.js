// node version v18.12.1
// npm version v9.8.1
import fs from "fs"

// SETUP
const INPUT = true ? "./sinput" : "./input"
const file = fs.readFileSync(INPUT, { encoding: "utf8" });

const lines = file.split("\n");
// Popping the blank line in end
lines.pop();
let sum1 = 0;
let sum2 = 0;
// SETUP

// SOLUTION -- START
lines.forEach((line, lineIdx) => {
  // logic
})

console.log("Part 1 Sum", sum1)
console.log("Part 2 Sum", sum2)
// SOLUTION -- END
