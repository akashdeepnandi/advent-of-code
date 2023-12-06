// node version v18.12.1
// npm version v9.8.1
import fs from "fs"

// SETUP
const TEST = false
const INPUT = TEST ? "./sinput" : "./input"
const file = fs.readFileSync(INPUT, { encoding: "utf8" });

const lines = file.split("\n");
// Popping the blank line in end
lines.pop();
// SETUP

// SOLUTION -- START
let sum1 = 0;
// Fill the values of card copy tracker, default 1 - each card has one copy
const cardCopyTracker = lines.map(() => 1);

lines.forEach((card, cardIdx) => {
  const nums = card.split(": ")[1];
  // split the nums, and parse to int
  const [winNums, myNums] = nums.split(" | ").map((numStr) =>
    numStr.split(" ").filter(s => s !== "").map((n) => Number.parseInt(n))
  );

  let matchCount = 0;
  myNums.forEach((n) => {
    // check if my number is in winning number list
    if (winNums.includes(n)) matchCount++;
  });

  if (matchCount > 0) {
    // calculate the points and add the sum
    sum1 += Math.pow(2, matchCount - 1)
    // FOR PART 2
    // get current card's copy count
    const currentCardCopies = cardCopyTracker[cardIdx];
    // get the card indexes for which we need to increase the count
    const startIdx = cardIdx + 1;
    const endIdx = Math.min(lines.length - 1, cardIdx + matchCount)
    // run the loop from start to end and add the counts
    for (let i = startIdx; i <= endIdx; i++) {
      cardCopyTracker[i] += currentCardCopies;
    }
  }
})

console.log("Part 1 Sum", sum1)

let sum2 = cardCopyTracker.reduce((acc, count) => (acc += count), 0)
console.log("Part 1 Sum", sum2)
// SOLUTION -- END
