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
	patterns := [][][]string{}
	currentPattern := [][]string{}
	for _, line := range lines {
		if line == "" {
			patterns = append(patterns, currentPattern)
			currentPattern = [][]string{}
			continue
		}
		currentPattern = append(currentPattern, strings.Split(line, ""))
	}
	patterns = append(patterns, currentPattern)
	sum1 := 0
	sum2 := 0
	for _, p := range patterns {
		pt := gridTranspose(p)
		r1 := getReflection(p, true)
		// column
		c1 := getReflection(pt, true)

		// row
		r2 := getReflection(p, false)
		// column
		c2 := getReflection(pt, false)
		sum1 += r1*100 + c1
		sum2 += r2*100 + c2
	}

	fmt.Println("Part 1", sum1)
	fmt.Println("Part 2", sum2)
}

func getReflection(grid [][]string, part1 bool) int {
	R := len(grid)
	for r := 1; r < R; r++ {
		topStart, topEnd, bottomStart, bottomEnd := getRelfectionIndices(r, R)
		top := grid[topStart:topEnd]
		bottom := grid[bottomStart:bottomEnd]
		top = reverseSlice(top)

		isSymmetric := true
		smudgeCount := 0 // FOR PART 2
		for i, r1 := range top {
			r2 := bottom[i]
			if rowEqual(r1, r2) == false && part1 {
				isSymmetric = false
				break
			}
			// FOR PART 2
			smudgeCount += getSmudgeCount(r1, r2)
		}

		if isSymmetric && part1 {
			return r
		}

		// FOR PART 2
		if smudgeCount == 1 && !part1 {
			return r
		}
	}
	return 0
}

// returns topStart, topEnd, bottomStart, bottomEnd
func getRelfectionIndices(r, R int) (int, int, int, int) {
	overlap := min(r, R-r)
	return max(0, r-overlap), r, r, r + overlap
}

func rowEqual(r1, r2 []string) bool {
	for i, s := range r1 {
		if s != r2[i] {
			return false
		}
	}

	return true
}

func getSmudgeCount(r1, r2 []string) int {
	smudgeCount := 0
	for i, s := range r1 {
		if s != r2[i] {
			smudgeCount++
		}
	}

	return smudgeCount
}
