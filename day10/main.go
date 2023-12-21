package main

import (
	"bufio"
	"fmt"
)

// | is a vertical pipe connecting north and south.
// - is a horizontal pipe connecting east and west.
// L is a 90-degree bend connecting north and east.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting south and west.
// F is a 90-degree bend connecting south and east.
// . is ground; there is no pipe in this tile.
// S is the starting position of the animal; there is a pipe on this
func main() {
	INPUT := "sinput"
	lines := readLines(INPUT)
	solution(lines)
}

func solution(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}

func part2(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
