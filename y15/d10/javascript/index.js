import fs from "fs";
const file = fs.readFileSync(process.argv[2], { encoding: "utf8" }).trim();
const inputSequence = file.split("\n")[0].split("");

const lookAndSay = (sequence) => {
  let result = [];
  let currentDigit = sequence[0];
  let count = 0;

  for (const d of sequence) {
    if (currentDigit == d) {
      count++;
    } else {
      result.push(...count.toString().split(""));
      result.push(currentDigit);

      currentDigit = d;
      count = 1;
    }
  }

  // add the last
  result.push(...count.toString().split(""));
  result.push(currentDigit);

  return result;
}

const evolveSequence = (sequence, generation) => {
  for (let i = 0; i < generation; i++) {
    sequence = lookAndSay(sequence);
  }

  return sequence;
}

console.log("Array");

let part1 = evolveSequence(inputSequence, 40);
console.log("Part 1:", part1.length);


let part2 = evolveSequence(inputSequence, 50);
console.log("Part 2:", part2.length);
