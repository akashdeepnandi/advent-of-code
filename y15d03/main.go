package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	f := flag.String("f", "", "Input File Name")
	flag.Parse()
	lines := readLines(*f)
	solve(lines)
}

var moveDir = map[string]Vector{"^": {-1, 0}, "v": {1, 0}, "<": {0, -1}, ">": {0, 1}}

func (v Vector) move(m string) Vector {
	o := moveDir[m]
	return v.Add(o)
}

func solve(lines []string) {
	moves := strings.Split(lines[0], "")
	santaPos, bothPos := Vector{0, 0}, []Vector{{0, 0}, {0, 0}}
	visited1, visited2 := make(map[Vector]int), make(map[Vector]int)
	for i, move := range moves {
		santaPos = santaPos.move(move)
		visited1[santaPos] = 1 // for part 1

		bothPos[i%2] = bothPos[i%2].move(move)
		visited2[bothPos[i%2]] = 1 // for part 2
	}
	fmt.Println("Part 1", len(visited1))
	fmt.Println("Part 2", len(visited2))
}
