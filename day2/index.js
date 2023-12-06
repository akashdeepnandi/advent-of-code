// node version v18.12.1
// npm version v9.8.1
import fs from "fs"

// SETUP
const INPUT = false ? "./sinput" : "./input"
const file = fs.readFileSync(INPUT, { encoding: "utf8" });

const lines = file.split("\n");
// Popping the blank line in end
lines.pop();
let sum1 = 0;
let sum2 = 0;
// SETUP

// SOLUTION -- START
lines.forEach((line, lineIdx) => {
  const redCounts = [...line.matchAll(/(\d+) red/g)].map((m) => Number.parseInt(m[1]))
  const greenCounts = [...line.matchAll(/(\d+) green/g)].map((m) => Number.parseInt(m[1]))
  const blueCounts = [...line.matchAll(/(\d+) blue/g)].map((m) => Number.parseInt(m[1]))
  Math.max
  // get the max count
  const redMax = Math.max(...redCounts)
  const greenMax = Math.max(...greenCounts)
  const blueMax = Math.max(...blueCounts)
  const isGameValid = redMax <= 12 && greenMax <= 13 && blueMax <=14

  if(isGameValid) sum1 += lineIdx + 1 // gameIndex = lineIdx + 1
  sum2 += redMax * greenMax * blueMax // calculate power and add it
})

console.log("Part 1 Sum", sum1)
console.log("Part 2 Sum", sum2)
// SOLUTION -- END
