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

func solve(lines []string) {
	operateLight := func(v, mode int) int {
		if mode == 0 {
			if v == 1 {
				v = 0
			} else {
				v = 1
			}
		} else if mode == 1 {
			v = 1
		} else {
			v = 0
		}
		return v

	}
	fmt.Println("Part 1: ", operateDecoration(lines, operateLight))
	operateBrightness := func(v, mode int) int {
		if mode == 0 {
			v += 2
		} else if mode == 1 {
			v++
		} else {
			v = max(0, v-1)
		}
		return v

	}
	fmt.Println("Part 2: ", operateDecoration(lines, operateBrightness))
}

func parseInstruction(operation []string) (int, []int, []int) {
	mode, si, ei := -1, 2, 4 // mode -> operation on lights
	if len(operation) == 4 {
		mode = 0 // toggle
		si, ei = 1, 3
	} else if operation[1] == "on" {
		mode = 1 // turn on
	} else {
		mode = 2 // turn off
	}
	return mode,
		stringToIntSlice(strings.Split(operation[si], ",")),
		stringToIntSlice(strings.Split(operation[ei], ","))
}

func operateDecoration(lines []string, operate func(int, int) int) int {
	S := 1000

	grid := make([][]int, S)
	for i := 0; i < S; i++ {
		grid[i] = make([]int, S)
	}
	count := 0

	for _, line := range lines {
		operation := strings.Split(line, " ")
		mode, start, end := parseInstruction(operation)
		for r := start[0]; r <= end[0]; r++ {
			for c := start[1]; c <= end[1]; c++ {
				grid[r][c] = operate(grid[r][c], mode)
			}
		}
	}

	for r := 0; r < S; r++ {
		for c := 0; c < S; c++ {
			count += grid[r][c]
		}
	}
	return count
}
