import fs from "fs"
const file = fs.readFileSync(process.argv[2], { encoding: "utf8" }).trim();
const lines = file.split("\n");
// Part 1
let sCount = 0;
for (let i = 0; i < lines.length; i++) {
  const line = lines[i];
  const vowels = [...line.matchAll(/(a|e|i|o|u)/g)];
  if (vowels.length < 3) continue;
  const doubles = [...line.matchAll(/(\w)\1/g)];
  if (doubles.length <= 0) continue;
  const bad = [...line.matchAll(/(ab|cd|pq|xy)/g)];
  if (bad.length > 0) continue;
  sCount++;
}

console.log(sCount)

// Part 2
sCount = 0
for (let i = 0; i < lines.length; i++) {
  const line = lines[i];
  const pairs = [...line.matchAll(/(\w{2}).*\1/g)]
  if (pairs.length <= 0) continue
  const onebetween = [...line.matchAll(/(\w)\w\1/g)]
  if (onebetween.length <= 0) continue
  sCount++;
}
console.log(sCount)
