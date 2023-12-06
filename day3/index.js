// node version v18.12.1
// npm version v9.8.1
// This Code does not focus on conciseness but readibility. 
// I know there are 100 ways to make this snappy and short
// but target is to make this beginner friend
import fs from "fs"


const TEST = false
const INPUT = TEST ? "./sinput" : "./input"
const file = fs.readFileSync(INPUT, {
  encoding: "utf8"
});

const lines = file.split("\n");
lines.pop();
// regex for matching all the numbers
const reNum = /\d+/g

// regex for matching anything which is not a number or dot
const reSymbol = /[^\d\.]/
let sum1 = 0;

// FOR PART 2
// will be used for tracking the pair of part numbers adjacent to '*'
const gears = {}

// think of lines as a grid, it will help
lines.forEach((line, lineIndex) => {
  const numberMatches = [...line.matchAll(reNum)]
  // Uncomment below to check the matches
  // here each num match is a regex match object
  // containing match at 0 and startIndex with index
  numberMatches.forEach((numMatch) => {
    // get the number and starting, ending indices of the match
    // for example: for string "...431.." ->
    // num = 43, matchStart = 3, matchEnd = 5
    const num = Number.parseInt(numMatch[0])
    const matchStart = numMatch.index
    const matchEnd = matchStart + numMatch[0].length - 1
    let isPart = false
    // FOR PART 2 
    let gearIndex = "" // to track

    // for PART 1
    // we need to check if their is any symbol adjacent to the number
    // if there is, then it is a part number
    // positions we will need to check are denoted by x. Basically all surrounding positions
    // [x] [x] [x] [ ] - to check: entire top row and top diagonals
    // [x] [m] [x] [ ] - to check: left and right neighbour
    // [x] [x] [x] [ ] - to check: entire bottom row and bottom diagonals

    // check left: matchStart - 1
    if (matchStart - 1 >= 0) {// just to make sure, this is positive
      const char = line[matchStart - 1] // get the character from left
      const isSymbol = !!reSymbol.exec(char) // convert to bool
      if (isSymbol) {
        isPart = true
        if (char === "*") {
          // get the gear coordinates
          const x = matchStart - 1
          const y = lineIndex
          const coordinate = `x:${x},y:${y}`
          const currentValue = gears[coordinate] || [];
          gears[coordinate] = [...currentValue, num]
        }
      }
    }

    // check right: matchEnd + 1
    if (matchEnd + 1 < line.length) {// just to make sure, we are in bounds
      const char = line[matchEnd + 1] // get the character from right
      const isSymbol = !!reSymbol.exec(char) // convert to bool
      if (isSymbol) {
        isPart = true
        if (char === "*") {
          // get the gear coordinates
          const x = matchEnd + 1
          const y = lineIndex
          const coordinate = `x:${x},y:${y}`
          const currentValue = gears[coordinate] || [];
          gears[coordinate] = [...currentValue, num]
        }
      }
    }

    // check top row and top diagonals: matchStart - 1 to matchEnd + 1 
    if (lineIndex != 0) { // there is no top row for first row
      const topLine = lines[lineIndex - 1] // get the top line
      const topStart = Math.max(0, matchStart - 1) // to stay in bounds
      const topEnd = Math.min(topLine.length - 1, matchEnd + 1) // to stay in bounds

      for (let topIndex = topStart; topIndex <= topEnd; topIndex++) {
        const topChar = topLine[topIndex];
        const isSymbol = !!reSymbol.exec(topChar) // convert to bool
        if (isSymbol) {
          isPart = true
          if (topChar === "*") {
            // get the gear coordinates
            const x = topIndex
            const y = lineIndex - 1
            const coordinate = `x:${x},y:${y}`
            const currentValue = gears[coordinate] || [];
            gears[coordinate] = [...currentValue, num]
          }
        }

      }
    }

    // check bottom row and bottom diagonals: matchStart - 1 to matchEnd + 1 
    if (lineIndex != lines.length - 1) { // there is no bottom row for last row
      const bottomLine = lines[lineIndex + 1] // get the bottom line
      const bottomStart = Math.max(0, matchStart - 1) // to stay in bounds
      const bottomEnd = Math.min(bottomLine.length - 1, matchEnd + 1) // to stay in bounds

      for (let bottomIndex = bottomStart; bottomIndex <= bottomEnd; bottomIndex++) {
        const bottomChar = bottomLine[bottomIndex];
        const isSymbol = !!reSymbol.exec(bottomChar) // convert to bool
        if (isSymbol) {
          isPart = true
          if (bottomChar === "*") {
            // get the gear coordinates
            const x = bottomIndex
            const y = lineIndex + 1
            const coordinate = `x:${x},y:${y}`
            const currentValue = gears[coordinate] || [];
            gears[coordinate] = [...currentValue, num]
          }
        }
      }
    }

    if (isPart) {
      sum1 += num
    }
  })
})

console.log("PART 1, Sum is", sum1)
// FOR PART 2
// calculate the part numbers which are adjacent to gear
// filter it to keep only 2 items as that makes it a pair
// OPTION 1
let sum2 = 0;
let gearPairs = [];
for (let coordinate in gears) {
  const numbers = gears[coordinate];
  if (numbers.length === 2) {
    gearPairs.push(numbers)
  }
}

for (let i = 0; i < gearPairs.length; i++) {
  const pair = gearPairs[i];
  const p0 = pair[0]
  const p1 = pair[1]

  const gearRatio = p0 * p1;
  sum2 += gearRatio
}
// OPTION 1


// alternate way to do the above with higher order functions
// comment the OPTION 1, uncomment code between OPTION 2
// check what Object.values does to a object
// OPTION 2
// let sum2 = Object.values(gears)
//   .filter((numbers) => numbers.length === 2)
//   .reduce((acc, [p0, p1]) => (acc += p0 * p1), 0)
// OPTION 2

console.log("PART 2, Sum is", sum2)
