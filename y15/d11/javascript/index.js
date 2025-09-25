import fs from "fs";
const file = fs.readFileSync(process.argv[2], { encoding: "utf8" }).trim();

let password = file.split("\n")[0];

const letters = "abcdefghijklmnopqrstuvwxyz"
const itol = letters.split("");
let ltoi = {}
itol.forEach((l, i) => {
  ltoi[l] = i;
});
const invalidLetters = ["i", "l", "o"];

const validatePassword = (password) => {
  let lastLetter = password[0];
  let lastLetterIndex = ltoi[lastLetter];
  let threeCount = 1;
  let hasThreeStraight = false;
  let pairCount = { [lastLetter]: 1 };

  if (invalidLetters.includes(lastLetter)) return false;

  for (let i = 1; i < password.length; i++) {
    let currentLetter = password[i];
    let currentLetterIndex = ltoi[currentLetter];
    if (invalidLetters.includes(currentLetter)) return false;

    if (threeCount == 3) {
      hasThreeStraight = true;
    } else if (lastLetterIndex + 1 === currentLetterIndex) {
      threeCount++;
    } else {
      threeCount = 1;
    }

    if (lastLetter === currentLetter) {
      pairCount[lastLetter]++;
    } else {
      pairCount[currentLetter] = 1;
    }

    lastLetter = currentLetter;
    lastLetterIndex = currentLetterIndex;
  }

  let hasTwoPairs = Object.values(pairCount).filter((v) => v >= 2).length >= 2;

  return hasTwoPairs && hasThreeStraight;
}

const incrementPassword = (password) => {
  let result = password.split("");

  // fix invalid passwords
  for (let i = 0; i < result.length; i++) {
    let currentLetter = result[i];
    if (invalidLetters.includes(currentLetter)) {
      let currentLetterIndex = ltoi[currentLetter];
      let newCurrentLetter = itol[(currentLetterIndex + 1) % itol.length];
      result[i] = newCurrentLetter;

      for (let j = i + 1; j < result.length; j++) {
        result[j] = "a"
      }
      return result.join("");
    }
  }

  // increment
  for (let i = result.length - 1; i >= 0; i--) {
    let currentLetter = result[i];
    let currentLetterIndex = ltoi[currentLetter];
    let nextLetter = itol[(currentLetterIndex + 1) % itol.length];
    result[i] = nextLetter;

    if (nextLetter !== "a") break;
  }

  return result.join("");
}

while (true) {
  password = incrementPassword(password)
  if (validatePassword(password)) break;
}

console.log("Part 1:", password);

while (true) {
  password = incrementPassword(password)
  if (validatePassword(password)) break;
}

console.log("Part 2:", password);
