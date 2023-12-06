// node version v18.12.1
// npm version v9.8.1
import fs from "fs"

// SETUP
const INPUT = true ? "./sinput" : "./input"
const file = fs.readFileSync(INPUT, { encoding: "utf8" });

const lines = file.split("\n");
// Popping the blank line in end
lines.pop();
// SETUP

// SOLUTION -- START
const getWinCountUsingRoots = (time, distance) => {
  // here from ax2 + bx + c, a = 1, b = -time, c = distance
  let sqrt = Math.sqrt(time * time - 4 * distance);
  let r1 = (time - sqrt) / 2
  let r2 = (time + sqrt) / 2
  if (Number.isInteger(r1)) r1++;
  if (Number.isInteger(r2)) r2--
  r1 = Math.ceil(r1)
  r2 = Math.floor(r2)
  return (r2 - r1) + 1;
}

let part1Input = []
let part2Input = ["", ""];
lines.forEach((line, lineIdx) => {
  // logic
  part1Input.push(
    [...line.matchAll(/\d+/g)].map((l) => {
      part2Input[lineIdx] += l
      return Number.parseInt(l[0])
    })

  )
})

// FOR PART 1

let [times, distances] = part1Input;
let part1Margin = 1
times.forEach((t, i) => {
  let d = distances[i]
  part1Margin *= getWinCountUsingRoots(t, d)
});

// FOR PART 1

// FOR PART 2
const [bigTime, bigDistance] = part2Input.map((n) => Number.parseInt(n))
const part2Ans = getWinCountUsingRoots(bigTime, bigDistance)

console.log("Part 1", part1Margin)
console.log("Part 2", part2Ans)
// SOLUTION -- END
