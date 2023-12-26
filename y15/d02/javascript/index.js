import fs from "fs";
const file = fs.readFileSync(process.argv[2], { encoding: "utf8" }).trim();
const lines = file.split("\n");
let paper = 0,
  ribbon = 0;
lines.forEach((line) => {
  const [l, w, h] = line.split("x").map((d) => Number.parseInt(d));
  // part 1
  paper += 2 * (l * w + w * h + h * l) + Math.min(l * w, w * h, h * l);
  // part 2
  ribbon += l * w * h + 2 * Math.min(l + w, w + h, h + l);
});
console.log("Total Wrapping Paper", paper);
console.log("Total Ribbon", ribbon);
