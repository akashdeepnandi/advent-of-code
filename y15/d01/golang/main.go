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
	input := strings.Split(lines[0], "")
	floor := 0
	pos := 0
	for i, p := range input {
		// part 1
		if p == "(" {
			floor++
		} else {
			floor--
		}
		// part 2
		if floor == -1 && pos == 0 {
			pos = i + 1
		}

	}

	fmt.Println("Part 1", floor)
	fmt.Println("Part 2", pos)
}
