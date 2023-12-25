import { readFileSync } from "fs"
import Vector from "./vector.js";
const moves = [...readFileSync(process.argv[2], "utf8").trim()];

let moveDir = { "^": new Vector(-1, 0), "v": new Vector(1, 0), "<": new Vector(0, -1), ">": new Vector(0, 1) }
const nextPos = (v, move) => v.add(moveDir[move])

let santaPos = new Vector(0,0), bothPos = [new Vector(0,0), new Vector(0,0)]
const visited1 = new Set(), visited2 = new Set()

moves.forEach((move, i) => {
  santaPos = nextPos(santaPos, move)
  visited1.add(santaPos.toString()) // part 1

  bothPos[i % 2] = nextPos(bothPos[i % 2], move)
  visited2.add(bothPos[i % 2].toString()) // part 2
})

console.log("Part 1", visited1.size)
console.log("Part 2", visited2.size)
