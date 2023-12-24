import fs from "fs"
const file = fs.readFileSync(process.argv[2], { encoding: "utf8" }).trim();
const lines = file.split("\n");

lines.forEach((line, lineIdx) => {
  // logic
  console.log(line)
})

