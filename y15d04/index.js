import { readFileSync } from "fs";
import { createHash } from "crypto";

const input = readFileSync(process.argv[2], { encoding: "utf8" }).trim();

const findNum = (prefix) => {
  for (let i = 0; ; i++) {
    const hash = createHash("md5");
    hash.update(input + i);
    if (hash.digest("hex").startsWith(prefix)) {
      console.log(i);
      break;
    }
  }
};

findNum("00000"); // part 1
findNum("000000"); // part 2
