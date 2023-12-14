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

var R, C int

func solve(lines []string) {
	R = len(lines)
	C = len(lines[0])
	grid := make([][]string, 0, R)

	for _, line := range lines {
		row := make([]string, 0, C)
		for _, c := range strings.Split(line, "") {
			row = append(row, c)
		}
		grid = append(grid, row)
	}

	part1(grid)
	part2(grid)
}

func part1(grid [][]string) {
	grid = gridRotate(grid)
	grid = gridTilt(grid)
	fmt.Println("Part 1", gridLoad(grid))
}

func part2(grid [][]string) {
	CYCLES_LENGTH := 1000000000
	gridSeen, key, loopEnd := make(map[string]int), "", 0
	for i := 0; i < CYCLES_LENGTH; i++ {
		key = gridString(grid)
		if gridSeen[key] != 0 {
			loopEnd = i
			break
		} else {
			gridSeen[key] = i
		}
		grid = gridCyle(grid)
	}

	loopStart := gridSeen[key]

	cycleEnd := (CYCLES_LENGTH-loopStart)%(loopStart-loopEnd) + loopStart
	for k, i := range gridSeen {
		if cycleEnd == i {
			key = k
			break
		}
	}

	grid = stringToGrid(key)

	grid = gridRotate(grid)
	fmt.Println("Part 2", gridLoad(grid))

}

func stringToGrid(gridStr string) [][]string {
	g := strings.Split(gridStr, "")
	grid := make([][]string, R)
	gindex := 0
	for r := 0; r < R; r++ {
		row := make([]string, C)
		for c := 0; c < C; c++ {
			row[c] = g[gindex]
			gindex++
		}

		grid[r] = row
	}

	return grid
}

func gridString(grid [][]string) string {
	s := ""
	for r := 0; r < R; r++ {
		s += strings.Join(grid[r], "")
	}

	return s
}

func gridCyle(grid [][]string) [][]string {
	for i := 0; i < 4; i++ {
		grid = gridRotate(grid)
		grid = gridTilt(grid)
	}
	return grid
}

func gridRotate(grid [][]string) [][]string {
	grid = gridTranspose(grid)
	for i := 0; i < R; i++ {
		start := 0
		end := C - 1
		// reverses the rows
		for start < end {
			grid[i][start], grid[i][end] = grid[i][end], grid[i][start]
			start++
			end--
		}
	}
	return grid
}

func gridLoad(grid [][]string) int {
	totalLoad := 0
	for r := 0; r < R; r++ {
		rowLoad := 0
		for c := 0; c < C; c++ {
			if isRound(grid[r][c]) {
				rowLoad += c + 1
			}
		}

		totalLoad += rowLoad
	}

	return totalLoad
}

func gridTilt(grid [][]string) [][]string {
	// pushes all the O to the right side
	for r := 0; r < R; r++ {
		row := grid[r]
		rowS := strings.Join(row, "")
		parts := strings.Split(rowS, "#")
		for i, part := range parts {
			parts[i] = sortString(part, false)
		}
		rowS = strings.Join(parts, "#")
		grid[r] = strings.Split(rowS, "")
	}

	return grid
}

func isRound(s string) bool {
	return s == "O"
}

func isEmpty(s string) bool {
	return s == "."
}
